# NovaSearch

### Note this is under developemnt and is not yet suitable for use

NovaSearch is a Go-based application that queries AWS EC2 instances based on tags and caches the results in Redis. The application exposes an HTTP API to retrieve the cached instances.

## Project Structure
awsclient/ awsclient.go 
cache/ cache.go 
config/ config.go 
docker-compose.yaml 
Dockerfile 
go.mod 
go.sum 
main.go


## Prerequisites

- Go 1.21 or later
- Docker
- Docker Compose

## Configuration

Configuration is managed using Viper and should be provided in a `config.yaml` file in the root directory. Example:

```yaml
aws:
  region: "us-east-1"
tags:
  - "Environment"
  - "Project"
```

## Building and Running
Using Docker Compose
1) Build and start the services:
```bash
docker-compose up --build
```
2) Access the API:
```bash
The API will be available at http://127.0.0.1:8080/instances.
```
## API Endpoints
GET /instances: Retrieves the cached EC2 instances.

## License
This project is licensed under the MIT License.
