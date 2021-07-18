package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	gohttp "net/http"

	"github.com/micheam/wiseman/internal/http"
	"github.com/micheam/wiseman/internal/localdata"
	"github.com/micheam/wiseman/internal/usecase"
	"github.com/micheam/wiseman/scrumwise"
)

var (
	projects = flag.String("projects", "",
		"A comma-separated list of the ids of the scrumwise projects")

	interval = flag.Int("interval", 5,
		"The interval to check for updates in seconds")

	sockfile = flag.String("sockfile", filepath.Join(os.Getenv("HOME"),
		".wiseman", "server.sock"), "path of unix socket file")
)

func main() {
	err := run(context.Background())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

var (
	getDataVersion = scrumwise.GetDataVersion
	getData        = func(ctx context.Context) (*scrumwise.Data, scrumwise.DataVersion, error) {
		ps := strings.Split(*projects, ",")
		param := scrumwise.NewGetDataParam(ps...)
		r, err := scrumwise.GetData(ctx, *param)
		if err != nil {
			return nil, 0, err
		}
		return r.Data, r.DataVersion, nil
	}
)

func run(ctx context.Context) error {
	flag.Parse()

	if len(*projects) == 0 {
		flag.Usage()
		return fmt.Errorf("projects is empty or not set")
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	_ = os.Remove(*sockfile)
	uaddr, err := net.ResolveUnixAddr("unix", *sockfile)
	if err != nil {
		return err
	}
	log.Printf("server start at addr %v", uaddr)
	listener, err := net.ListenUnix("unix", uaddr)
	if err != nil {
		return fmt.Errorf("can't listen via %q: %w", *sockfile, err)
	}
	defer func() { _ = listener.Close() }()

	personGateway := localdata.NewPersonRepository(localdata.DataCache{
	contaner := http.Container{
		UseCaseListPersons: usecase.NewUseCaseListPersons(personGateway),
	}
	srv := &gohttp.Server{Handler: http.Router(ctx, contaner)}
	go func() {
		if err := srv.Serve(listener); err != nil && err != gohttp.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")
	<-done

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}
	log.Print("Server Exited Properly")
	return nil
}
