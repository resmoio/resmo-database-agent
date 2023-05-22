package pkg

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"reflect"
	"resmo-db-mapper/pkg/config"
	"resmo-db-mapper/pkg/resmo"
)

func RunSQLDatabaseQueries(ctx context.Context, config config.Config, dbType string) error {
	var queries []Data

	switch dbType {
	case "mysql":
		queries = mysqlQueries
	case "postgres":
		queries = postgresQueries
	case "clickhouse":
		queries = clickhouseQueries
	default:
		return fmt.Errorf("unsupported database type: %s", dbType)
	}

	log.Printf("connecting to database: %s\n", config.DSN)
	db, err := sqlx.Open(dbType, config.DSN)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	log.Printf("connected to database successfully\n")
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			fmt.Errorf("error closing database: %w", err)
		}
	}(db)

	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	for _, query := range queries {
		data, err := QueryDBData(ctx, db, query.values, query.query)
		if err != nil {
			// todo think about fail situation
			log.Printf("failed to map data for %s: %s", query.name, err)
			continue
		}

		err = resmo.Ingest(ctx, config, dbType, query.name, data)
		if err != nil {
			return fmt.Errorf("failed to send data for %s: %w", query.name, err)
		}
	}

	return nil
}

func QueryDBData(ctx context.Context, db *sqlx.DB, dest interface{}, query string) (interface{}, error) {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr || reflect.TypeOf(dest).Elem().Kind() != reflect.Slice {
		return nil, fmt.Errorf("database conversation has failed")
	}

	destValue := reflect.ValueOf(dest).Elem()

	rows, err := db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Errorf("error closing rows: %w", err)
		}
	}(rows)

	for rows.Next() {
		item := reflect.New(reflect.TypeOf(dest).Elem().Elem()).Interface()
		err = rows.StructScan(item)
		if err != nil {
			log.Printf("failed to scan row: %s", err)
			continue
		}
		destValue.Set(reflect.Append(destValue, reflect.ValueOf(item).Elem()))
	}

	if rows.Err() != nil {
		fmt.Errorf("rows returned with error: %w", err)
	}

	return dest, nil
}
