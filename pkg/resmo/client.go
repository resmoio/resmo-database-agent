package resmo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
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

	dbUrl, err := extractDomainAndPort(config.DSN)
	if err != nil {
		return fmt.Errorf("error while extracting domain and port from DSN: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Ingest-Key", config.IngestKey)
	req.Header.Set("Resmo-Database-Agent", config.Version)
	req.Header.Set("DB-URL", dbUrl)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	defer resp.Body.Close()

	return nil
}
func extractDomainAndPort(urlStr string) (string, error) {
	uri, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("invalid URL format for given URL: %s", urlStr)
	}

	host, port, err := net.SplitHostPort(uri.Host)
	if err != nil {
		return "", fmt.Errorf("failed to split host and port: %w", err)
	}

	if port != "" {
		return fmt.Sprintf("%s:%s", host, port), nil
	}

	return host, nil
}
