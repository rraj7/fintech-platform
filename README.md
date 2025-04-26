# Fintech Platform API (Go)
=========================

A simple, enterprise-grade financial platform API built with Go (Gin Framework).
This project simulates core banking operations like user registration, authentication, transaction processing, and dashboard views — designed with clean architecture principles and cloud deployment in mind.

## Features
--------
- User Registration with full validation (email, password, DOB, etc.)
- Secure Authentication using JWT (WIP)
- Transaction APIs: Transfer funds & view history
- User Dashboard endpoint
- Modular, scalable project structure (Clean Architecture)
- Ready for PostgreSQL integration
- Containerization & Cloud-ready (Azure focus)
- Future-proofed for CI/CD pipelines

## 🏗️ Tech Stack
- **Language:** Go (Golang)
- **Framework:** Gin
- **Auth:** JWT (Coming Soon)
- **Database:** PostgreSQL (Planned)
- **Deployment:** Docker + Azure (Planned)
- **Infra-as-Code:** Terraform (Coming Soon)

## Getting Started

### Prerequisites
- Go 1.20+
- Git
- (Optional) Docker & PostgreSQL for DB setup

### Clone the Repo
```bash
git clone https://github.com/rraj7/fintech-platform.git
cd fintech-platform
```
Run the API Locally:
```bash
go run cmd/main.go
```
**By default, server runs on: http://localhost:8080**

## API Endpoints
| Method | Endpoint                     | Description         |
|--------|------------------------------|---------------------|
| POST   | /api/auth/register           | User Registration   |
| POST   | /api/auth/login              | User Login          |
| GET    | /api/user/dashboard          | View User Dashboard |
| POST   | /api/transactions/transfer   | Transfer Funds      |
| GET    | /api/transactions/history    | Transaction History |

## Project Structure

``` 
fintech-platform/
├── cmd/                - Entry point
├── internal/
│   ├── adapters/       - HTTP Handlers, Middleware
│   ├── service/        - Business Logic (Validations, Auth)
│   ├── domain/         - Core Entities (User, Transaction)
│   └── ports/          - Interfaces for DB, Services
└── go.mod
```

## Roadmap

- [x] Basic API Routing
- [x] User Input Validation
- [ ] JWT Authentication
- [ ] PostgreSQL Integration
- [ ] Dockerize Application
- [ ] Terraform Azure Deployment
- [ ] GitHub Actions CI/CD
- [ ] API Documentation (Swagger)

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License

MIT License

## Author
------
Rishi Raj
Cloud & DevOps Engineer | Go Enthusiast | Azure Specialist
LinkedIn: https://linkedin.com/in/rishiraj31
Portfolio: https://rishi-raj.me
