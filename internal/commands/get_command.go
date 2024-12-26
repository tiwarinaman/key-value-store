package commands

import (
	"errors"
	"own-redis/internal/storage"
	"own-redis/pkg/constants"
)

type GetCommand struct {
	BaseCommand
}

func (g *GetCommand) PreExecute(args []string, storage *storage.Storage) error {
	if len(args) != 1 {
		return errors.New("GET requires a key")
	}
	return nil
}

func (g *GetCommand) ExecuteCore(args []string, store *storage.Storage) string {
	key := args[0]
	value := store.Get(key)
	if value == nil {
		return constants.Nil
	}
	return value.(string)
}

func (g *GetCommand) GetType() string {
	return constants.Get
}
