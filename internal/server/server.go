package server

import (
	"log"
	"net"

	"github.com/visheshc14/report-service/internal/service"
	"google.golang.org/grpc"

	gen "github.com/visheshc14/report-service/gen"
)

type GRPCServer struct {
	server *grpc.Server
}

func NewGRPCServer(reportService *service.ReportService) *GRPCServer {
	grpcServer := grpc.NewServer()
	gen.RegisterReportServiceServer(grpcServer, reportService)

	return &GRPCServer{
		server: grpcServer,
	}
}

func (s *GRPCServer) Start(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("gRPC server starting on %s", addr)
	return s.server.Serve(lis)
}

func (s *GRPCServer) Stop() {
	log.Printf("Stopping gRPC server")
	s.server.GracefulStop()
}
