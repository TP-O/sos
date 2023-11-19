// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package pggenerated

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreatePlayer(ctx context.Context, arg CreatePlayerParams) (Player, error)
	PlayerByID(ctx context.Context, userID pgtype.UUID) (Player, error)
	PlayersByUsername(ctx context.Context, username string) ([]PlayersByUsernameRow, error)
}

var _ Querier = (*Queries)(nil)
