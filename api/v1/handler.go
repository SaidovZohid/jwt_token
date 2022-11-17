package v1

import (
	"github.com/SaidovZohid/jwt_token/config"
	"github.com/SaidovZohid/jwt_token/storage"
)

type handlerV1 struct {
	cfg *config.Config
	Storage storage.StorageI
}

type HandlersV1Options struct {
	Cfg *config.Config
	Storage *storage.StorageI
}

func New(options *HandlersV1Options) *handlerV1 {
	return &handlerV1{
		cfg: options.Cfg,
		Storage: *options.Storage,
	}
}

