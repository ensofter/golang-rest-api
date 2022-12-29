package db

import (
	"rest-api-tutorial/internal/user"
	"rest-api-tutorial/pkg/logging"
)

type db struct {
}

func NerStorage(collection string, logger *logging.Logger) user.Storage {
	return &db{}
}
