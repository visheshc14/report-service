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

<img width="1051" height="394" alt="Screenshot 2025-09-03 132408" src="https://github.com/user-attachments/assets/db2d1be0-e801-44fd-9b87-930fc6e0c316" />

<img width="769" height="212" alt="Screenshot 2025-09-03 132548" src="https://github.com/user-attachments/assets/3def989b-0b20-40ce-b195-00d02537f551" />




