package txmanager

import "context"

type TxManager interface {
	Transactional(ctx context.Context, fn func(context.Context) error) error
}
