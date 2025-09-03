# Report Service

A Go service with gRPC API and cron jobs for report generation.

## Features

- gRPC API with GenerateReport and HealthCheck endpoints
- Cron job that runs every 10 seconds to generate reports for predefined users
- In-memory report storage
- Comprehensive logging with timestamps
- Docker containerization


### Installation

- Clone the repository: 
 ```git clone https://github.com/visheshc14/report-service.git```
 ```cd report-service```
- Install dependencies:
  ```go mod tidy``` 
- Generate protobuf code:
  ```./generate.sh```
- Build and run:
  ```make build``
  make run```
- Docker
  ```bash
  make docker-build
  make docker-run```

### TESTING
Use ``grpcurl`` to test the gRPC endpoints:
# Health check
```grpcurl -plaintext -d '{}' localhost:50051 report.ReportService/HealthCheck
# Generate report
```grpcurl -plaintext -d '{"user_id": "test_user"}' localhost:50051 report.ReportService/GenerateReport
  
### GenerateReport
```protobuf
rpc GenerateReport(GenerateReportRequest) returns (GenerateReportResponse)





