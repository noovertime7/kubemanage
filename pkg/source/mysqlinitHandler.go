package source

import (
	"context"
	"gorm.io/gorm"
)

type mysqlInitHandler struct {
	db *gorm.DB
}

var _ InitHandler = &mysqlInitHandler{}

func newMysqlInitHandler(db *gorm.DB) InitHandler {
	return &mysqlInitHandler{db: db}
}

func (m *mysqlInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return m.createTables(ctx, inits)
}

func (m *mysqlInitHandler) createTables(ctx context.Context, inits initSlice) error {
	for _, init := range inits {
		if init.TableCreated(ctx, m.db) {
			continue
		}
		if err := init.MigrateTable(ctx, m.db); err != nil {
			return err
		}
	}
	return nil
}
