package dbx

import (
	"context"

	"github.com/virzz/mulan/db"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	cfg *db.Config
	db  *gorm.DB
}

func (s *Service) Serve() (err error) {
	zap.L().Info("Connecting to DB")
	s.db, err = New(s.cfg)
	if err != nil {
		zap.L().Error("Failed to connect to DB", zap.Error(err))
		return err
	}
	return nil
}

func (s *Service) Shutdown(ctx context.Context) error { return s.Close() }

func (s *Service) Close() error {
	sqlDb, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDb.Close()
}

func (s *Service) Raw() any { return s.db }

func NewService(cfg *db.Config) *Service { return &Service{cfg: cfg} }
