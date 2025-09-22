package auditor

import (
	"context"
	"fmt"
	"time"
)

type repo struct {
}

func New() *repo {
	return &repo{}
}

func (a *repo) Audit(ctx context.Context) error {

	traceID := fmt.Sprintf("%v", ctx.Value("traceID"))
	fmt.Println("I am auditing the request and response with traceID: ", traceID)
	time.Sleep(15 * time.Second)
	fmt.Println("Auditing the request and response successfully, with traceID: ", traceID)
	return nil
}
