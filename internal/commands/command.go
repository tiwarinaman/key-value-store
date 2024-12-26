package commands

import "own-redis/internal/storage"

type Command interface {
	Execute(args []string, storage *storage.Storage) string
	GetType() string
}
