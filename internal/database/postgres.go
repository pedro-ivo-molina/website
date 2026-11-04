package database

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

func Connect(databaseURL string) (*pgx.Conn, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("Deu ruim")
	}

	log.Println("Deu bom")
	return conn, err
}

func RunMigrations(databaseURL string) error {
	migrations, err := migrate.New(
		"file://db/migrations",
		databaseURL,
	)

	if err != nil {
		return fmt.Errorf("Erro ao inicializar o migrador: %w", err)
	}

	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("Erro ao rodar as migrations: %w", err)
	}

	log.Println("Migrations executadas com sucesso!")
	return nil
}
