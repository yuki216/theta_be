package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	usecase "theta-be/application/usecase/record"
	delivery "theta-be/delivery/http"
	"theta-be/delivery/http/healthhandler"
	"theta-be/delivery/http/recordhandler"
	domainrepo "theta-be/domain/repository"
	"theta-be/infrastructure/database"
	"theta-be/infrastructure/repository"
	"theta-be/pkg/config"
)

func main() {
	cfg := config.LoadConfig()

	gin.SetMode(cfg.GinMode)

	// 1. Layer: Infrastructure (Database / Repository)
	var recordRepo domainrepo.RecordRepository
	if cfg.DBDriver == "postgres" {
		db, err := database.NewPostgresDB(cfg)
		if err != nil {
			log.Fatalf("[FATAL] Postgres init failed: %v", err)
		}
		recordRepo = repository.NewRecordGormRepository(db)
		log.Println("[INFO] Layer Infrastructure: PostgreSQL GORM loaded")
	} else {
		recordRepo = repository.NewRecordMemoryRepository()
		log.Println("[INFO] Layer Infrastructure: In-Memory Storage loaded (Zero-DB mode)")
	}

	// 2. Layer: Application (Usecase Orchestration)
	recordUsecase := usecase.NewUsecase(recordRepo)

	// 3. Layer: Delivery (HTTP Handlers & Routing)
	healthHandler := healthhandler.NewHandler()
	recordHttpHandler := recordhandler.NewHandler(recordUsecase)

	router := delivery.SetupRouter(delivery.RouterConfig{
		HealthHandler: healthHandler,
		RecordHandler: recordHttpHandler,
	})

	// 4. Server Execution & Graceful Shutdown
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		log.Printf("[INFO] Server listening on http://localhost:%s\n", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("[FATAL] ListenAndServe failed: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("[INFO] Received shutdown signal. Terminating gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("[FATAL] Forced shutdown: %v\n", err)
	}

	log.Println("[INFO] Service stopped successfully.")
}
