package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/visheshc14/report-service/internal/server"
	"github.com/visheshc14/report-service/internal/service"

	gen "github.com/visheshc14/report-service/gen"
)

func main() {
	// Initialize services
	reportService := service.NewReportService()
	grpcServer := server.NewGRPCServer(reportService)

	// Start cron job for periodic report generation
	c := cron.New()
	predefinedUserIDs := []string{"Vishesh", "Neeti", "Yashraj", "Yuvraj", "Shibin"}

	// Add cron job that runs every 10 seconds
	_, err := c.AddFunc("@every 10s", func() {
		log.Printf("[%s] Cron job triggered", time.Now().Format(time.RFC3339))

		for _, userID := range predefinedUserIDs {
			ctx := context.Background()
			resp, err := reportService.GenerateReport(ctx, &gen.GenerateReportRequest{UserId: userID})
			if err != nil {
				log.Printf("[%s] Error generating report for user %s: %v",
					time.Now().Format(time.RFC3339), userID, err)
			} else if resp.Error != "" {
				log.Printf("[%s] Report generation failed for user %s: %s",
					time.Now().Format(time.RFC3339), userID, resp.Error)
			} else {
				log.Printf("[%s] Cron-generated report: %s for user: %s",
					time.Now().Format(time.RFC3339), resp.ReportId, userID)
			}

			// Small delay between user reports
			time.Sleep(100 * time.Millisecond)
		}

		// Log current report count
		log.Printf("[%s] Total reports generated: %d",
			time.Now().Format(time.RFC3339), reportService.GetReportCount())
	})

	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}

	// Start cron
	c.Start()
	log.Printf("Cron job started - will run every 10 seconds")

	// Start gRPC server in a goroutine
	go func() {
		if err := grpcServer.Start(":50051"); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Printf("Shutdown signal received, stopping services...")

	// Stop cron
	c.Stop()
	log.Printf("Cron job stopped")

	// Stop gRPC server
	grpcServer.Stop()

	log.Printf("Service shutdown completed")
}
