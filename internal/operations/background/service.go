package background

import (
	"context"
	"context-and-goroutines-management/internal"
	"fmt"
	"time"
)

type service struct {
	auditor internal.Auditor
	printer internal.Printer
}

func New(auditor internal.Auditor, printer internal.Printer) *service {
	return &service{
		auditor: auditor,
		printer: printer,
	}
}

func (s *service) Proccess(ctx context.Context) error {

	ctx = context.WithValue(ctx, "traceID", "dfghjgrfecdefrgthyjjuyhtbgvrfcd")

	traceID := fmt.Sprintf("%v", ctx.Value("traceID"))

	fmt.Println("Proccessing my request with traceID: ", traceID)
	time.Sleep(2 * time.Second)

	// Creating a goroutine with context.Background()
	// these goroutine does not cancell, but lose context data from parent
	// such as telemetry data
	go s.auditor.Audit(context.Background())
	go s.printer.PrintLog(context.Background())

	fmt.Println("End the proccess, but go routines are always working...")

	return nil
}
