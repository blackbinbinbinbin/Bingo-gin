package data

import (
	"github.com/blackbinbinbinbin/Bingo-gin/internal/log"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/data/conf"
)

// Data .
type Data struct {
	// TODO warpped database client
	// ep:
	// *sql.DB
	// *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Log("msg", "closing the data resources")
	}
	return &Data{}, cleanup, nil
}
