package internal

import "context"

type Auditor interface {
	Audit(ctx context.Context) error
}

type Printer interface {
	PrintLog(ctx context.Context)
}
