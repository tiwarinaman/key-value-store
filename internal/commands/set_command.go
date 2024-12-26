package commands

import (
	"errors"
	"fmt"
	"own-redis/internal/storage"
	"own-redis/pkg/constants"
)

type SetCommand struct {
	BaseCommand
}

func (s *SetCommand) PreExecute(args []string, storage *storage.Storage) error {
	if len(args) < 2 {
		return errors.New("SET requires a key and value")
	}
	return nil
}

func (s *SetCommand) ExecuteCore(args []string, storage *storage.Storage) string {
	key := args[0]
	value := args[1]
	var expiration int64 = 0

	if len(args) == 3 {
		_, err := fmt.Sscanf(args[2], "%d", &expiration)
		if err != nil {
			return errors.New("invalid expiration").Error()
		}
	}

	storage.Set(key, value, expiration)
	return constants.Ok
}

func (s *SetCommand) GetType() string {
	return constants.Set
}
