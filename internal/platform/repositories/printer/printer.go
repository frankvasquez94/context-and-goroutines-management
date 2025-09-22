package printer

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

func (a *repo) PrintLog(ctx context.Context) {

	traceID := fmt.Sprintf("%v", ctx.Value("traceID"))
	fmt.Println("I am printing the request and response with traceID: ", traceID)
	time.Sleep(25 * time.Second)
	fmt.Println("Printing the request and response successfully with traceID: ", traceID)
}
