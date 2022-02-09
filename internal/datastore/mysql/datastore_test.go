//go:build ci
// +build ci

package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"

	"github.com/authzed/spicedb/internal/datastore/mysql/migrations"
	"github.com/authzed/spicedb/pkg/migrate"
	"github.com/authzed/spicedb/pkg/secrets"
)

const (
	mysqlPort    = 3306
	testDBPrefix = "spicedb_test"
	creds        = "root:secret"
)

var containerPort string

func createMigrationDriver(connectStr string) (*migrations.MysqlDriver, error) {
	migrationDriver, err := migrations.NewMysqlDriver(connectStr)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize migration engine: %w", err)
	}

	return migrationDriver, nil
}

func TestMySQLMigrations(t *testing.T) {
	req := require.New(t)

	connectStr := setupDatabase()

	migrationDriver, err := createMigrationDriver(connectStr)
	req.NoError(err)

	version, err := migrationDriver.Version()
	req.NoError(err)
	req.Equal("", version)

	err = migrations.Manager.Run(migrationDriver, migrate.Head, migrate.LiveRun)
	req.NoError(err)

	version, err = migrationDriver.Version()
	req.NoError(err)

	headVersion, err := migrations.Manager.HeadRevision()
	req.NoError(err)
	req.Equal(headVersion, version)
}

func setupDatabase() string {
	var db *sql.DB
	connectStr := fmt.Sprintf("%s@(localhost:%s)/mysql", creds, containerPort)
	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		log.Fatalf("couldn't open DB: %s", err)
	}
	defer func() {
		err := db.Close() // we do not want this connection to stay open
		if err != nil {
			log.Fatalf("failed to close db: %s", err)
		}
	}()

	uniquePortion, err := secrets.TokenHex(4)
	if err != nil {
		log.Fatalf("Could not generate unique portion of db name: %s", err)
	}
	dbName := testDBPrefix + uniquePortion

	tx, err := db.Begin()
	_, err = tx.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
	if err != nil {
		log.Fatalf("failed to create database: %s: %s", dbName, err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalf("failed to commit: %s", err)
	}

	return fmt.Sprintf("%s@(localhost:%s)/%s?parseTime=true", creds, containerPort, dbName)
}

func migrateDatabase(connectStr string) {
	migrationDriver, err := createMigrationDriver(connectStr)
	if err != nil {
		log.Fatalf("failed to run migration: %s", err)
	}

	err = migrations.Manager.Run(migrationDriver, migrate.Head, migrate.LiveRun)
	if err != nil {
		log.Fatalf("failed to run migration: %s", err)
	}
}

func TestMain(m *testing.M) {
	mysqlContainerRunOpts := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "5",
		Env:        []string{"MYSQL_ROOT_PASSWORD=secret"},
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		fmt.Printf("could not connect to docker: %s\n", err)
		os.Exit(1)
	}

	// only bring up the container once
	containerResource, err := pool.RunWithOptions(mysqlContainerRunOpts)
	if err != nil {
		fmt.Printf("could not start resource: %s\n", err)
		os.Exit(1)
	}

	containerCleanup := func() {
		// When you're done, kill and remove the container
		if err := pool.Purge(containerResource); err != nil {
			fmt.Printf("could not purge resource: %s\n", err)
			os.Exit(1)
		}
	}

	containerPort = containerResource.GetPort(fmt.Sprintf("%d/tcp", mysqlPort))
	connectStr := fmt.Sprintf("%s@(localhost:%s)/mysql?parseTime=true", creds, containerPort)

	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		fmt.Printf("failed to open db: %s\n", err)
		containerCleanup()
		os.Exit(1)
	}

	defer func() {
		err := db.Close() // we do not want this connection to stay open
		if err != nil {
			fmt.Printf("failed to close db: %s\n", err)
			containerCleanup()
			os.Exit(1)
		}
	}()

	err = pool.Retry(func() error {
		var err error
		err = db.Ping()
		if err != nil {
			return fmt.Errorf("couldn't validate docker/mysql readiness: %w", err)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("mysql database error: %s\n", err)
		containerCleanup()
		os.Exit(1)
	}

	exitStatus := m.Run()
	containerCleanup()
	os.Exit(exitStatus)
}
