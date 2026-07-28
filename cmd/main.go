package main

import (
	"context"
	"os"
	"os/signal"
	"stations/internal"
	"sync"
	"syscall"
	"time"
)

func main() {
	timeCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	ctx, stop := signal.NotifyContext(timeCtx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	wg.Go(func() {

		mapfile, start, end, trainCount, ok := internal.ParseArgs()
		if !ok {
			ExitWithErrors()
		}

		contents := internal.ReadMapFile(mapfile)

		graphList, ok := internal.ParseMap(string(contents), start, end)
		if !ok {
			ExitWithErrors()
		}

		_, paths := internal.MaxFlow(graphList, start, end)
		internal.CreateSchedule(graphList, paths, start, end, trainCount, true)
	})

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		internal.Log("Program exited successfully")
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			internal.PrintWarn("Program timeout exceeded, program exited forcefully")
			os.Exit(1)
		}
		internal.Log("Shutdown initiated")
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		select {
		case <-done:
			internal.Log("Program exited successfully")
		case <-shutdownCtx.Done():
			internal.PrintWarn("Shutdown timeout exceeded, program exited forcefully")
			os.Exit(1)
		}
	}

}

func ExitWithErrors() {
	internal.PrintWarn("Program exited with errors")
	os.Exit(1)
}
