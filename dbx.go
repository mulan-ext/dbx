package dbx

import (
	"errors"
	"net/url"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/virzz/mulan/db"
)

func New(cfg *db.Config) (*gorm.DB, error) {
	_dsn, err := url.Parse(cfg.DSN)
	if err != nil {
		return nil, err
	}
	gormCfg := &gorm.Config{
		Logger: db.NewLogger(cfg.Debug),
	}
	var dialector gorm.Dialector
	switch _dsn.Scheme {
	case "mysql":
		dialector = mysql.New(mysql.Config{
			DSN:                    cfg.String(),
			DefaultStringSize:      255,
			DontSupportRenameIndex: true,
		})
		return db.New(dialector, cfg, gormCfg)
	case "postgres":
		dialector = postgres.New(postgres.Config{
			DSN:                  cfg.String(),
			PreferSimpleProtocol: true,
		})
	case "sqlite", "sqlite3":
		dialector = sqlite.Open(cfg.String())
	default:
		return nil, errors.New("unsupported scheme: " + _dsn.Scheme)
	}
	return db.New(dialector, cfg, gormCfg)
}
