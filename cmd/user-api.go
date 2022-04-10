package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	delivery "users-api-server/internal/user/delivery/http"
	"users-api-server/internal/user/repository/sqllite"
	"users-api-server/internal/user/service"
)

func Run(port int) error {

	router := gin.Default()

	routerGroup := router.Group("/api/v1")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		FullSaveAssociations: true,
		QueryFields:          true,
	})
	if err != nil {
		return errors.New("failed to connect to sqllite database: " + err.Error())
	}

	repo := sqllite.NewSqlLiteRepo(db)
	if err = repo.Migrate(); err != nil {
		return err
	}

	srv := service.New(repo)

	delivery.RegisterHttpEndPoints(routerGroup, srv)

	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", port), //specify port here
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err = httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return httpServer.Shutdown(ctx)
}
