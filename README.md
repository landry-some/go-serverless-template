# go-serverless-template

go-serverless-template is a production-ready template for building serverless applications in Go using AWS Lambda and the Serverless Framework.

It provides a clean project structure, local development setup, Docker support, and testing with mocks.

---

## Features

- AWS Lambda functions written in Go
- Serverless Framework configuration
- Docker-based local development
- Structured project layout (`functions/`, `pkg/`)
- Mock generation using Mockery
- Environment-based configuration
- CI/CD-friendly setup

---

## Project Structure

- `functions/` – Lambda function entry points  
- `pkg/` – Reusable business logic  
- `mocks/` – Generated mocks for testing  
- `serverless.yml` – Serverless configuration  
- `Dockerfile.golang` – Build container  
- `docker-compose.yml` – Local development setup  
- `.mockery.yml` – Mock generation config  
- `Makefile` – Development commands  

---

## Requirements

- Go 1.20+
- Node.js (for Serverless Framework)
- Serverless Framework
- Docker (optional for local builds)

Install Serverless globally:

```bash
npm install -g serverless
```

---

## Setup

Clone the repository:

```bash
git clone https://github.com/<your-username>/go-serverless-template.git
cd go-serverless-template
```

Copy environment variables:

```bash
cp .env.example .env
```

Edit `.env` with your AWS credentials and configuration.

---

## Local Development

Install Go dependencies:

```bash
go mod download
```

Run tests:

```bash
go test ./...
```

Generate mocks:

```bash
make mocks
```

---

## Deploy to AWS

Deploy using Serverless:

```bash
serverless deploy
```

Or specify stage:

```bash
serverless deploy --stage dev
```

---

## Local Testing (Optional)

If configured with Docker:

```bash
docker-compose up --build
```

---

## Adding a New Function

1. Create a new folder under `functions/`
2. Add the handler implementation
3. Register it in `serverless.yml`
4. Deploy

---
