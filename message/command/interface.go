package command

import "context"

type Handler interface {
	CommandName() string

	Exec(ctx context.Context, caller string, args []string) error
}
