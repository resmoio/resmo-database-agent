package config

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

type Config struct {
	URL            string
	IngestKey      string
	Version        string
	DSN            string
	Schedule       string
	Timeout        string
	DomainOverride string
}

func (c *Config) Validate() error {
	if c.IngestKey == "" {
		return errors.New("ingest key is not set")
	}

	if c.DSN == "" {
		return errors.New("datasource name is not set")
	}

	return nil
}

func (c *Config) ReadConfig(ver string) error {
	c.Version = ver

	flag.StringVar(&c.Schedule, "schedule", "", "schedule for running queries")
	flag.StringVar(&c.Timeout, "timeout", "", "timeout duration for connections")
	flag.StringVar(&c.DSN, "dsn", "", "database datasource name")
	flag.StringVar(&c.IngestKey, "ingestKey", "", "ingestKey of the integration")
	flag.StringVar(&c.DomainOverride, "domainOverride", "", "domain url for ingesting")

	flag.Parse()

	if c.DSN != "" && c.IngestKey != "" {
		if err := c.Validate(); err != nil {
			return fmt.Errorf("config is not valid: %w", err)
		}

		return nil
	}

	log.Printf("can not find config from flags, looking for environment variables DSN & RESMO_INGEST_KEY")

	c.DSN = os.Getenv("DSN")
	c.IngestKey = os.Getenv("RESMO_INGEST_KEY")

	if err := c.Validate(); err != nil {
		return fmt.Errorf("config is not valid: %w", err)
	}

	return nil
}
