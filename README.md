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
  ```bash
  git clone https://github.com/visheshc14/report-service.git
  cd report-service
- Install dependencies:
  ```bash
  go mod tidy 
- Generate protobuf code:
  ```bash
  ./generate.sh
- Build and run:
  ```bash
  make build
  make run
- Docker
  ```bash
  make docker-build
  make docker-run
    
### GenerateReport
```protobuf
rpc GenerateReport(GenerateReportRequest) returns (GenerateReportResponse)





