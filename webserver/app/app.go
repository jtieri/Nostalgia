package app

import (
	"github.com/jtieri/Nostalgia/config"
	"gorm.io/gorm"
)

var WebApp *App

type App struct {
	Config *config.Config
	DB     *gorm.DB
}

func New(db *gorm.DB) *App {
	return &App{
		DB: db,
	}
}
