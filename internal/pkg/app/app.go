package app

import (
	"awesomeProject/internal/app/config"
	"awesomeProject/internal/app/dsn"
	"awesomeProject/internal/app/repository"
	"context"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Conf   *config.Config
	repo   *repository.Repository
	Router gin.IRouter
}

func New() (*Application, error) {
	var ctx context.Context
	cnf, err := config.NewConfig(ctx)
	if err != nil {
		return nil, err
	}

	dsnStr := dsn.FromEnv()
	rep, err := repository.New(dsnStr)
	if err != nil {
		return nil, err
	}
	a := &Application{Conf: cnf, repo: rep}

	return a, nil
}
