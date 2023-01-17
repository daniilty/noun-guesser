package main

import (
	"bufio"
	"context"
	"log"
	"noun-guesser/internal/server"
	"noun-guesser/internal/tree"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func run() error {
	addr := os.Getenv("HTTP_ADDR")
	source := os.Getenv("SOURCE")

	f, err := os.Open(source)
	if err != nil {
		return err
	}

	t := tree.NewWord()

	s := bufio.NewScanner(f)
	for s.Scan() {
		t.Insert(s.Text())
	}

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	server := server.NewServer(addr, t)

	wg.Add(1)
	go func() {
		server.Run(ctx)
		wg.Done()
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-termChan
	cancel()
	wg.Wait()

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
