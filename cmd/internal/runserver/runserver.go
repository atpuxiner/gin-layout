package runserver

import (
	"context"
	"fmt"
	"github.com/atpuxiner/gin-layout/app/initializer/conf"
	"github.com/atpuxiner/gin-layout/app/initializer/db"
	"github.com/atpuxiner/gin-layout/app/initializer/logger"
	"github.com/atpuxiner/gin-layout/app/router"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func Start(port string) {
	c, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	conf.InitConfig()
	logger.InitLogger()
	defer logger.Logger.Sync()
	db.InitDB(conf.Conf.Db.Driver)
	engine := router.InitRouter()

	if port == "" {
		port = conf.Conf.Server.Port
		if port == "" {
			port = "9000"
		}
	}
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: engine,
	}
	go func() {
		fmt.Println("")
		fmt.Println(time.Now())
		fmt.Printf("Starting server at http://127.0.0.1:%s/\n", port)
		fmt.Printf("Testing server at http://127.0.0.1:%s/ping\n", port)
		fmt.Println("Quit the server with Ctrl+C")
		fmt.Println("")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server start error: %v", err)
		}
	}()

	<-c.Done()

	stop()

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(c); err != nil {
		log.Fatalf("Server stop error: %v", err)
	}
}

func Clean() {
	fmt.Println(time.Now())
	fmt.Println("Exiting the server.....")
	// TODO:
}
