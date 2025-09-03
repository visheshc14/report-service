# Report Service

A Go service with gRPC API and cron jobs for report generation.

## Features

- gRPC API with GenerateReport and HealthCheck endpoints
- Cron job that runs every 10 seconds to generate reports for predefined users
- In-memory report storage
- Comprehensive logging with timestamps
- Docker containerization

## Installation 
- Clone the repository:
```bash
git clone https://github.com/visheshc14/report-service.git 
cd report-service
```
- Install dependencies:
```bash
go mod tidy
```
- Generate protobuf code:
```bash
./generate.sh
```  
- Build and run:
```bash
make build
make run
```
## Docker
```bash
make build
make run
```
## Testing

Use **grpcurl** to test the gRPC endpoints:

- Health check
```bash
grpcurl -plaintext -d '{}' localhost:50051 report.ReportService/HealthCheck
```
- Generate report
```bash
grpcurl -plaintext -d '{"user_id": "test_user"}' localhost:50051 report.ReportService/GenerateReport
```
### GenerateReport
```protobuf
rpc GenerateReport(GenerateReportRequest) returns (GenerateReportResponse)
```
  
<img width="1051" height="394" alt="Screenshot 2025-09-03 132408" src="https://github.com/user-attachments/assets/db2d1be0-e801-44fd-9b87-930fc6e0c316" />

<img width="769" height="212" alt="Screenshot 2025-09-03 132548" src="https://github.com/user-attachments/assets/3def989b-0b20-40ce-b195-00d02537f551" />

## Installation Scaling Report Service for High Concurrency

To handle 10,000 concurrent gRPC requests per second across multiple data centers, I would horizontally scale the service by deploying multiple containerized instances using Kubernetes with auto-scaling based on CPU, memory, and request load, distributed across regions for low latency and high availability. I would expose the gRPC service through **Traefik** as an ingress/load balancer, enabling TLS termination, routing, and observability, and also enable gRPC client-side load balancing and reflection. Reports would be stored in a distributed persistent database such as PostgreSQL, Cassandra, or Redis, with caching for frequently accessed data to ensure durability and fast access. Heavy report generation could be offloaded to a message queue like Kafka for asynchronous processing, and server streaming or batch processing could optimize throughput and reduce memory pressure. Reliability would be enhanced with health checks, retries with exponential backoff, circuit breakers, and multi-region failover. Comprehensive monitoring and tracing using Prometheus, Grafana, and OpenTelemetry would provide observability and allow proactive performance tuning. Logging, metrics, and alerting would ensure SLAs are met while handling high-volume, low-latency gRPC traffic efficiently.
