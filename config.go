package main

import (
	"errors"

	"github.com/caarlos0/env/v6"
)

type config struct {
	Port                   int      `env:"PORT" envDefault:"8080"`
	HoneyPots              []string `env:"HONEYPOTS" envDefault:"_t_email" envSeparator:","`
	DefaultRecipient       string   `env:"EMAIL_TO"`
	AllowedRecipients      []string `env:"ALLOWED_TO" envSeparator:","`
	Sender                 string   `env:"EMAIL_FROM"`
	SmtpUser               string   `env:"SMTP_USER"`
	SmtpPassword           string   `env:"SMTP_PASS"`
	SmtpHost               string   `env:"SMTP_HOST"`
	SmtpPort               int      `env:"SMTP_PORT" envDefault:"587"`
	GoogleApiKey           string   `env:"GOOGLE_API_KEY"`
	Spamlist               []string `env:"SPAMLIST" envSeparator:"," envDefault:"gambling,casino"`
	MessageHeader          string   `env:"MESSAGE_HEADER"`
	MessageFooter          string   `env:"MESSAGE_FOOTER"`
	MessageSubmitter       string   `env:"MESSAGE_SUBMITTER" envDefault:"false"`
	MessageSubmitterHeader string   `env:"MESSAGE_SUBMITTER_HEADER"`
	MessageSubmitterFooter string   `env:"MESSAGE_SUBMITTER_FOOTER"`
}

func parseConfig() (*config, error) {
	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		return cfg, errors.New("failed to parse config")
	}
	return cfg, nil
}

func checkRequiredConfig(cfg *config) bool {
	if cfg.DefaultRecipient == "" {
		return false
	}
	if len(cfg.AllowedRecipients) < 1 {
		return false
	}
	if cfg.Sender == "" {
		return false
	}
	if cfg.SmtpUser == "" {
		return false
	}
	if cfg.SmtpPassword == "" {
		return false
	}
	if cfg.SmtpHost == "" {
		return false
	}
	return true
}
