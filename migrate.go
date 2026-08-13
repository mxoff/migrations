package migrations

import (
	"errors"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
)

/* files name
000001_init.up.sql
000001_init.down.sql
*/

func New(dsn string, files string) *Migrate {
	return &Migrate{
		dsn:   dsn,
		files: files,
	}
}

type Migrate struct {
	dsn   string
	files string
}

func (r *Migrate) Migrate() error {
	m, err := migrate.New(r.files, r.dsn)
	if err != nil {
		return fmt.Errorf("migrate.New: %w", err)
	}
	defer func() {
		_, _ = m.Close()
	}()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrate.Up: %w", err)
	}

	return nil
}

func (r *Migrate) Force(version int) error {
	m, err := migrate.New(r.files, r.dsn)
	if err != nil {
		return fmt.Errorf("migrate.New: %w", err)
	}
	defer func() {
		_, _ = m.Close()
	}()

	if err := m.Force(version); err != nil {
		return fmt.Errorf("migrate.Force(%d): %w", version, err)
	}

	return nil
}
