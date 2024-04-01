package config

import "github.com/joho/godotenv"

type PGConfig interface {
	DSN() string
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}
