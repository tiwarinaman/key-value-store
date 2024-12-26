package commands

import (
	"fmt"
	"own-redis/internal/storage"
	"own-redis/pkg/constants"
)

type BaseCommand struct {
}

func (b *BaseCommand) PreExecute(args []string, storage *storage.Storage) error {
	return nil
}

func (b *BaseCommand) ExecuteCore(args []string, storage *storage.Storage) string {
	return constants.NotImplemented
}

func (b *BaseCommand) PostExecute(args []string, result string, store *storage.Storage) {
	fmt.Printf("Command executed with args: %v, result: %s\n", args, result)
}

func (b *BaseCommand) Execute(args []string, store *storage.Storage) string {
	if err := b.PreExecute(args, store); err != nil {
		return constants.Error + err.Error()
	}

	result := b.ExecuteCore(args, store)
	b.PostExecute(args, result, store)
	return result
}

func (b *BaseCommand) GetType() string {
	return constants.BaseCommand
}
