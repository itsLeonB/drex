package provider

import (
	"errors"

	"github.com/itsLeonB/drex/internal/config"
	"github.com/itsLeonB/ezutil/v2"
)

type Provider struct {
	Logger ezutil.Logger
	*DBs
	*Repositories
	*Services
}

func All(configs config.Config) *Provider {
	dbs := ProvideDBs(configs.DB)
	repos := ProvideRepositories(dbs.GormDB)

	return &Provider{
		Logger:       ProvideLogger(configs.App),
		DBs:          dbs,
		Repositories: repos,
		Services:     ProvideServices(repos),
	}
}

func (p *Provider) Shutdown() error {
	var err error
	if p.DBs != nil {
		if e := p.DBs.Shutdown(); e != nil {
			err = errors.Join(err, e)
		}
	}
	return err
}
