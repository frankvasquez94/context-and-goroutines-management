package main

import (
	"context"
	"context-and-goroutines-management/internal/operations/background"
	"context-and-goroutines-management/internal/operations/detach"
	"context-and-goroutines-management/internal/platform/repositories/auditor"
	"context-and-goroutines-management/internal/platform/repositories/printer"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// proccessWorkWitDetachedContext()
	proccessWorkWithBackgroundContext()
	// Wait to terminate main go routine
	time.Sleep(1 * time.Minute)

	fmt.Println("All the proccess took: ", time.Since(start))

}

func proccessWorkWitDetachedContext() {
	start := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	auditorRepo := auditor.New()
	printerRepo := printer.New()

	svc := detach.New(auditorRepo, printerRepo)

	err := svc.Proccess(ctx)

	if err != nil {
		fmt.Println("Error proccessing request", err.Error())
	}

	select {
	case <-ctx.Done():
		fmt.Println("The request has been complete, however audit and printer goroutines may still be proccessing", "the request took: ", time.Since(start))
		return
	case <-time.After(5 * time.Second):
		fmt.Println("All the proccess has been completed ...")
	}

}

func proccessWorkWithBackgroundContext() {
	start := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	auditorRepo := auditor.New()
	printerRepo := printer.New()

	svc := background.New(auditorRepo, printerRepo)

	err := svc.Proccess(ctx)

	if err != nil {
		fmt.Println("Error proccessing request", err.Error())
	}

	select {
	case <-ctx.Done():
		fmt.Println("The request has been complete, however audit and printer goroutines may still be proccessing", "the request took: ", time.Since(start))
		return
	case <-time.After(5 * time.Second):
		fmt.Println("All the proccess has been completed ...")
	}

}
