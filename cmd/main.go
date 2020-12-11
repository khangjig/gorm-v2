package main

import (
	"context"
	"gorm-v2/client/mysql"
	"gorm-v2/config"
	"gorm-v2/datastores"
	serviceHttp "gorm-v2/delivery/http"
	"gorm-v2/repository"
	"gorm-v2/usecase"
	"log"
	"net"
	_ "net/http/pprof"
	"runtime"
	"time"

	"github.com/soheilhy/cmux"
)

func main() {
	// setup locale
	{
		loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
		if err != nil {
			log.Fatal(err)
			runtime.Goexit()
		}
		time.Local = loc
	}

	var (
		cfg    = config.GetConfig()
		ctx    = context.Background()
		client = mysql.GetClient
		repo   = repository.New(client)
	)

	datastores.Migrate(client(ctx))

	useCase := usecase.New(repo)

	l, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(l)
	httpL := m.Match(cmux.HTTP1Fast())

	errs := make(chan error)

	// http
	{
		h := serviceHttp.NewHTTPHandler(useCase)
		go func() {
			h.Listener = httpL
			errs <- h.Start("")
		}()
	}

	go func() {
		errs <- m.Serve()
	}()

	err = <-errs
	if err != nil {
		log.Fatal(err)
	}

	defer mysql.Disconnect()
	log.Println("exit")
}
