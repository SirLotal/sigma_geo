package controllers

import (
	"context"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (app *App) Startup(ctx context.Context) {
	app.ctx = ctx
}

func (app App) DomReady(ctx context.Context) {

}

func (app *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

func (app *App) Shutdown(ctx context.Context) {

}
