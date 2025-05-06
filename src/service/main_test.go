//go:build database

package service

import (
	"context"
	"os"
	"testing"

	"catalog/src"
	"catalog/src/datastore"
	"catalog/src/datastore/ent"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var DbClient *ent.Client

func TestMain(m *testing.M) {
	src.SetLogLevel(zapcore.DebugLevel)
	logger := src.DefaultLogger

	dsn := src.Dsn{
		Host:     src.GetEnvOrDefault(src.EnvDbHost, "localhost"),
		Port:     src.GetEnvOrDefault(src.EnvDbPort, "3306"),
		Username: src.GetEnvOrDefault(src.EnvDbRootUser, "root"),
		Password: src.GetEnvOrDefault(src.EnvDbRootPassword, "catalog_root_password"),
		Database: src.GetEnvOrDefault(src.EnvDbSchema, "test_schema"),
	}

	dbClient, err := datastore.Onboarding(
		context.Background(),
		logger,
		dsn,
		src.GetEnvOrDefault(src.EnvDbSupportUser, "catalog_support"),
		src.GetEnvOrDefault(src.EnvDbSupportPassword, "catalog_support_password"),
		src.GetEnvOrDefault(src.EnvDbUser, "catalog_user"),
		src.GetEnvOrDefault(src.EnvDbPassword, "catalog_user_password"),
		src.GetEnvOrDefault(src.EnvDbSchema, "test_schema"),
	)
	if err != nil {
		logger.Error("Failed to deploy test database", zap.Error(err))
		os.Exit(1)
	}
	DbClient = dbClient

	code := m.Run()

	logger.Info("Tests done")
	_ = logger.Sync()

	os.Exit(code)
}
