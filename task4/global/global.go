package global

import (
	"log/slog"

	"gorm.io/gorm"
)

var (
	SDB    *gorm.DB
	Logger *slog.Logger
)
