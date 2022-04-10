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
	"users-api-server/config"
	loginDelivery "users-api-server/internal/login/delivery/http"
	onlyadmin "users-api-server/internal/login/service/only-admin"
	userDelivery "users-api-server/internal/user/delivery/http"
	"users-api-server/internal/user/repository/sqllite"
	"users-api-server/internal/user/service"
	"users-api-server/pkg/midleware/auth/midleware"
	"users-api-server/pkg/midleware/auth/service/secret_key"
)

func Run(cfg config.Config) error {

	router := gin.Default()

	db, err := gorm.Open(sqlite.Open(cfg.DatabaseName), &gorm.Config{
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

	userService := service.New(repo)
	authService := secret_key.JWTAuthSecretService(cfg.SecretKey, cfg.Issuer)
	loginService := onlyadmin.StaticLoginService()
	authMidleware := midleware.NewAuthMiddleware(authService)

	routerGroupV1 := router.Group("/api/v1", authMidleware)
	routerGroupV0 := router.Group("/api/v0")

	loginDelivery.RegisterHttpEndPoints(routerGroupV0, loginService, authService)
	userDelivery.RegisterHttpEndPoints(routerGroupV1, userService)

	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%s", cfg.AppPort), //specify port here
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
