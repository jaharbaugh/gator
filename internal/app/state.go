package app

import(
	"context"
	"github.com/jaharbaugh/gator/internal/config"
	"github.com/jaharbaugh/gator/internal/database"
)


type State struct {
	CFG *config.Config
	DB *database.Queries
	CTX context.Context
}