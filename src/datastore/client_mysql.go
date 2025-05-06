//go:build mysql

package datastore

import (
	"context"
	"database/sql"
	"fmt"

	"catalog/src"
	"catalog/src/datastore/ent"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func GetDbClient(_ context.Context, dsn src.Dsn) (*ent.Client, error) {
	client, err := ent.Open("mysql", dsn.String())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL database: %w", err)
	}
	return client, nil
}

func Onboarding(
	ctx context.Context,
	logger *zap.Logger,
	dsn src.Dsn,
	supportUsername, supportPasswrod,
	username, password,
	database string,
) (*ent.Client, error) {
	db, err := sql.Open("mysql", dsn.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to open MySQL database connection")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to ping MySQL database")
	}

	logger.Warn("Creating database...")
	if _, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", database)); err != nil {
		return nil, errors.Wrapf(err, "failed to create database %s", database)
	}

	logger.Warn("Creating db user...")
	if _, err = db.Exec(fmt.Sprintf("CREATE USER IF NOT EXISTS '%s'@'%%' IDENTIFIED BY '%s'", username, password)); err != nil {
		return nil, errors.Wrapf(err, "failed to create user %s", username)
	}

	// grant SELECT, INSERT, UPDATE to db user
	if _, err = db.Exec(fmt.Sprintf("GRANT CREATE, SELECT, INSERT, UPDATE ON `%s`.* TO '%s'@'%%'", database, username)); err != nil {
		return nil, errors.Wrapf(err, "failed to grant privileges on database %s to user %s", database, username)
	}

	logger.Warn("Creating support user...")
	if _, err = db.Exec(fmt.Sprintf("CREATE USER IF NOT EXISTS '%s'@'%%' IDENTIFIED BY '%s'", supportUsername, supportPasswrod)); err != nil {
		return nil, errors.Wrapf(err, "failed to create user %s", username)
	}

	// grant SELECT to support user
	if _, err = db.Exec(fmt.Sprintf("GRANT SELECT ON `%s`.* TO '%s'@'%%'", database, supportUsername)); err != nil {
		return nil, errors.Wrapf(err, "failed to grant privileges on database %s to user %s", database, username)
	}

	if _, err = db.Exec("FLUSH PRIVILEGES"); err != nil {
		return nil, errors.Wrap(err, "failed to flush privileges")
	}

	dsn.Database = database
	dsn.Username = username
	dsn.Password = password
	entClient, err := GetDbClient(ctx, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ent client")
	}

	logger.Warn("Creating tables...")
	if err := entClient.Schema.Create(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to create schema")
	}

	logger.Warn("Database deployment completed successfully")
	return entClient, nil
}
