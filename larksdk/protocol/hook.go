package protocol

import "context"

type Hook interface {
	BeforeProcess(ctx context.Context) (context.Context, error)
	AfterProcess(ctx context.Context) error
}
