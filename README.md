# Go Fetch Receipt

This repository contains the Fetch Receipt service implementation. Follow the instructions below to set up and run the project.

## Quick Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd go-fetch-receipt
```

2. Set up environment variables:
```bash
cp .env_template .env
```
Edit the `.env` file with your specific configuration values.

3. Start the services:
```bash
docker compose up
```

4. Run database migrations:
```bash
make migrate-up
```

## Project Status

### Tasks Left
- [ ] Implement comprehensive error handling
- [ ] Complete project documentation
    - API documentation
    - Architecture overview
    - Development guidelines
    - Testing procedures

## Prerequisites

- Docker and Docker Compose
- Make
- Go 

## Development

### Environment Configuration
The `.env` file controls various aspects of the application. Make sure to configure the following:
- Database credentials
- Service ports

### Database Migrations
Migrations are managed through the `make` command:
```bash
# Apply migrations
make migrate-up

# Rollback migrations
make migrate-down

# Drop database
make drop
```


## Query
Email Author

## Author
Sagar Shrestha (https://github.com/znerf)

