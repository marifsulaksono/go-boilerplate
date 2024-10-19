package contract

import "context"

type UserRepository interface {
	Get(ctx context.Context) string
}
