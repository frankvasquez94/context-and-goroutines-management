package detach

import (
	"context"
	"context-and-goroutines-management/internal"
	"context-and-goroutines-management/internal/platform/kit"
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

	// Creating my detached context
	dContextAuditor := kit.DetachedContext(ctx)
	dContextPrinter := kit.DetachedContext(ctx)

	go s.auditor.Audit(dContextAuditor)
	go s.printer.PrintLog(dContextPrinter)

	fmt.Println("End the proccess, but go routines are always working...")

	return nil
}
