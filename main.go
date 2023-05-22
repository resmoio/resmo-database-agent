package main

import (
	"context"
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"resmo-db-mapper/pkg"
	"resmo-db-mapper/pkg/config"
	"strings"
	"time"
)

var (
	version string
)

func main() {
	if err := realMain(); err != nil {
		log.Fatalln("failed to run resmo-database-agent:", err)
	}
}

func realMain() error {
	var config config.Config
	err := config.ReadConfig(version)
	if err != nil {
		return fmt.Errorf("error while reading config : %w", err)
	}

	ctxDur, err := time.ParseDuration(config.Timeout)
	if err != nil {
		log.Printf("failed to parse the context duration from the configuration: %s. The context timeout duration is set to 10 seconds\n", err)
		ctxDur = 10 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), ctxDur)
	defer cancel()

	log.Printf("config: %s read successfully, will start to run queries\n", config)

	dbType, err := getDatabaseType(config.DSN)
	if err != nil {
		return fmt.Errorf("error while getting database type from DSN: %w", err)
	}
	if config.Schedule == "" {
		err := runQueries(ctx, config, dbType)
		if err != nil {
			return fmt.Errorf("error while running queries: %w", err)
		}
		log.Println("database resources ingested successfully")
		return nil
	}
	dur, err := time.ParseDuration(config.Schedule)
	if err != nil {
		return fmt.Errorf("could not parse schedule from config: %w", err)
	}

	ticker := time.NewTicker(dur)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("running queries again with schedule: ", config.Schedule)
		ctx, cancel := context.WithTimeout(context.Background(), ctxDur)
		err := runQueries(ctx, config, dbType)
		if err != nil {
			return fmt.Errorf("error while running queries: %w", err)
		}
		cancel()
	}

	log.Println("database resources ingested successfully")
	return nil
}

func runQueries(ctx context.Context, config config.Config, dbType string) error {
	switch dbType {
	case "mongo":
		err := pkg.RunMongoQueries(ctx, config, dbType)
		if err != nil {
			return fmt.Errorf("mongo runner error: %w", err)
		}
	default:
		err := pkg.RunSQLDatabaseQueries(ctx, config, dbType)
		if err != nil {
			return fmt.Errorf("sql runner error: %w", err)
		}
	}

	return nil
}

func getDatabaseType(connectionString string) (string, error) {
	var dbType string

	if strings.HasPrefix(connectionString, "postgres://") || strings.HasPrefix(connectionString, "postgresql://") {
		dbType = "postgres"
	} else if strings.HasPrefix(connectionString, "mongodb://") {
		dbType = "mongo"
	} else if strings.HasPrefix(connectionString, "clickhouse://") {
		dbType = "clickhouse"
	} else if strings.HasPrefix(connectionString, "mysql://") {
		dbType = "mysql"
	} else {
		return "", fmt.Errorf("unsupported database type for connection string: %s", connectionString)
	}

	return dbType, nil
}
