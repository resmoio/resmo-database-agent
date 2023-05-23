package resmo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/kos-v/dsnparser"
	"net/http"
	"resmo-db-mapper/pkg/config"
)

func Ingest(ctx context.Context, config config.Config, driverType string, resourceKey, results interface{}) error {
	ingestUrl := "https://id.resmo.app:443/integration/%s/ingest/%s"
	if config.DomainOverride != "" {
		ingestUrl = "https://" + config.DomainOverride + "/integration/%s/ingest/%s"
	}

	data, err := json.Marshal(results)
	if err != nil {
		return fmt.Errorf("error marshaling %s results: %w", resourceKey, err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf(ingestUrl, driverType, resourceKey), bytes.NewBufferString(string(data)))

	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	dbIdentifier := config.DbIdentifier
	dbIdentifier, err = extractDomainAndPort(dbIdentifier, config)
	if err != nil {
		return fmt.Errorf("error while extracting domain and port from DSN: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Ingest-Key", config.IngestKey)
	req.Header.Set("Resmo-Database-Agent", config.Version)
	req.Header.Set("DB-Identifier", dbIdentifier)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	defer resp.Body.Close()

	return nil
}

func extractDomainAndPort(dbIdentifier string, config config.Config) (string, error) {
	if dbIdentifier == "" {
		dsn := dsnparser.Parse(config.DSN)
		if dsn.GetHost() == "" {
			return "", fmt.Errorf("invalid format for dsn: %s", dsn.GetRaw())
		}

		if dsn.GetScheme() == "" {
			dbIdentifier = dsn.GetHost()
		}

		dbIdentifier = fmt.Sprintf("%s:%s", dsn.GetHost(), dsn.GetPort())
	}
	return dbIdentifier, nil
}
