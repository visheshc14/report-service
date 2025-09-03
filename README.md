# Report Service

A Go service with gRPC API and cron jobs for report generation.

## Features

- gRPC API with GenerateReport and HealthCheck endpoints
- Cron job that runs every 10 seconds to generate reports for predefined users
- In-memory report storage
- Comprehensive logging with timestamps
- Docker containerization

### GenerateReport
```protobuf
rpc GenerateReport(GenerateReportRequest) returns (GenerateReportResponse)





