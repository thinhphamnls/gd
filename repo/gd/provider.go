package gdrepo

import (
	"github.com/thinhphamnls/gd/config"
	"github.com/thinhphamnls/gd/container"
	"go.uber.org/zap"
)

type GDRepositoriesProvider struct {
	QuickbooksOAuth IQuickBooksOAuthRepo
}

func NewGDRepositoriesProvider(
	cf bootstrap.Config,
	sugar *zap.SugaredLogger,
	databaseProvider container.IDataBaseProvider,
) *GDRepositoriesProvider {
	return &GDRepositoriesProvider{
		QuickbooksOAuth: NewQuickBooksOAuthRepo(databaseProvider, sugar),
	}
}
