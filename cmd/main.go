package main

import (
	_ "go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
)

func main() {
	app, err := NewApp()
	if err != nil {
		app.Logger.Fatal("NewApp", zap.Error(err))
	}
	if err := app.Start(); err != nil {
		app.Logger.Fatal("Start", zap.Error(err))
	}
}
