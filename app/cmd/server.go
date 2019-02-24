package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/amebalabs/nocreep/app/store"

	"github.com/amebalabs/nocreep/app/api"
)

// ServerCommand ...
type ServerCommand struct {
	Store StoreGroup
	Port  int
}

// StoreGroup ...
type StoreGroup struct {
	Type string
}

type serverApp struct {
	*ServerCommand
	restSrv     *api.API
	dataService store.Interface
	terminated  chan struct{}
}

// Execute is the entry point for "server" command, called by flag parser
func (s *ServerCommand) Execute() error {
	log.Printf("[INFO] start server on port %d", s.Port)

	ctx, cancel := context.WithCancel(context.Background())
	go func() { // catch signal and invoke graceful termination
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		log.Printf("[WARN] interrupt signal")
		cancel()
	}()

	app, err := s.newServerApp()
	if err != nil {
		log.Printf("[PANIC] failed to setup application, %+v", err)
	}
	if err = app.run(ctx); err != nil {
		log.Printf("[ERROR] remark terminated with error %+v", err)
		return err
	}
	log.Printf("[INFO] remark terminated")
	return nil
}

// newServerApp prepares server application
func (s *ServerCommand) newServerApp() (*serverApp, error) {

	storeEngine, err := s.makeDataStore()
	if err != nil {
		return nil, err
	}

	srv := &api.API{
		Version:   "1",
		Interface: storeEngine}

	return &serverApp{
		ServerCommand: s,
		restSrv:       srv,
		terminated:    make(chan struct{}),
	}, nil
}

// Run all application objects
func (a *serverApp) run(ctx context.Context) error {

	go func() {
		// shutdown on context cancellation
		<-ctx.Done()
		log.Print("[INFO] shutdown initiated")
		a.restSrv.Shutdown()
		if e := a.dataService.Close(); e != nil {
			log.Printf("[WARN] failed to close data store, %s", e)

		}
	}()

	a.restSrv.Run(a.Port)
	close(a.terminated)
	return nil
}

// Wait for application completion (termination)
func (a *serverApp) Wait() {
	<-a.terminated
}

// makeDataStore creates store for all sites
func (s *ServerCommand) makeDataStore() (result store.Interface, err error) {
	log.Printf("[INFO] make data store, type=%s", s.Store.Type)

	switch s.Store.Type {
	case "bolt":
		result, err = store.SetupBoltDB()
	case "mongo":
		fmt.Println("Mongo is not supported yet")
	default:
		return nil, fmt.Errorf("unsupported store type %s", s.Store.Type)
	}
	return result, err
}
