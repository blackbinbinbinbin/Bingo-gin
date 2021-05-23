package ent

import (
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/ent"
)

func Open(db *sql.DB) (*ent.Client, error) {
	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv)), nil
}