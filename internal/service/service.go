package service

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/visheshc14/report-service/gen"
)

type ReportService struct {
	gen.UnimplementedReportServiceServer
	reports map[string]*Report
	mu      sync.RWMutex
}

type Report struct {
	ID        string
	UserID    string
	Content   string
	CreatedAt time.Time
}

func NewReportService() *ReportService {
	return &ReportService{
		reports: make(map[string]*Report),
	}
}

func (s *ReportService) GenerateReport(ctx context.Context, req *gen.GenerateReportRequest) (*gen.GenerateReportResponse, error) {
	log.Printf("[%s] Generating report for user: %s", time.Now().Format(time.RFC3339), req.UserId)

	// Personalized report content based on user
	reportContent := s.getPersonalizedReportContent(req.UserId)

	// Simulate report generation
	reportID := fmt.Sprintf("report_%s_%d", req.UserId, time.Now().UnixNano())

	// Store report in memory
	s.mu.Lock()
	s.reports[reportID] = &Report{
		ID:        reportID,
		UserID:    req.UserId,
		Content:   reportContent,
		CreatedAt: time.Now(),
	}
	s.mu.Unlock()

	log.Printf("[%s] Report generated successfully: %s", time.Now().Format(time.RFC3339), reportID)

	return &gen.GenerateReportResponse{
		ReportId: reportID,
		Error:    "",
	}, nil
}

// Helper function for personalized report content
func (s *ReportService) getPersonalizedReportContent(userID string) string {
	// Customize report content based on the user
	switch userID {
	case "Vishesh":
		return fmt.Sprintf("Special performance report for Vishesh - Generated at %s. Excellent work on recent projects!", time.Now().Format(time.RFC3339))
	case "Neeti":
		return fmt.Sprintf("Detailed analytics report for Neeti - Generated at %s. Great progress on data analysis!", time.Now().Format(time.RFC3339))
	case "Yashraj":
		return fmt.Sprintf("Technical report for Yashraj - Generated at %s. Impressive coding contributions!", time.Now().Format(time.RFC3339))
	case "Yuvraj":
		return fmt.Sprintf("Strategic report for Yuvraj - Generated at %s. Outstanding leadership initiatives!", time.Now().Format(time.RFC3339))
	case "Shibin":
		return fmt.Sprintf("Financial report for Shibin - Generated at %s. Impressive financial contributions!", time.Now().Format(time.RFC3339))
	default:
		return fmt.Sprintf("Report for user %s generated at %s", userID, time.Now().Format(time.RFC3339))
	}
}

func (s *ReportService) HealthCheck(ctx context.Context, req *gen.HealthCheckRequest) (*gen.HealthCheckResponse, error) {
	log.Printf("[%s] Health check requested", time.Now().Format(time.RFC3339))

	return &gen.HealthCheckResponse{
		Status: "OK",
	}, nil
}

func (s *ReportService) GetReportCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.reports)
}

func (s *ReportService) GetReports() map[string]*Report {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Return a copy to avoid concurrent access issues
	copy := make(map[string]*Report)
	for k, v := range s.reports {
		copy[k] = v
	}
	return copy
}
