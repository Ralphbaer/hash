# Cart Service

This repo contains the source code of the Cart service.

## Architecture

![alt text](./hexagonal-macro.png "Title")

## Requirements

| Name | Version | Notes | Mandatory
|------|---------|---------|---------|
| [golang](https://golang.org/dl/) | >= go1.15.14 | Main programming language | true
| [docker](https://www.docker.com/) | n/a | Used to handle core mock images and start local service | true
| [sh/bash] | depending on OS. Anyway, you should be able do execute any .sh file | Used to lint checks, test processes and some console interface customizations | true
| [make](https://www.gnu.org/software/make/) | depending on OS. Anyway, you should be able do execute make commands to run the project, tests and localenvironment | n/a | true

# Usage

### Start Local
Inside /cart, you can run:
```bash
make image           # Builds the container image
make container      # Start the service on port 3000

make discount-image  # Builds the container image
make discount-container # Start the discount service on port 50051
make run           # Start service on port 3000 too (no docker usage[fast-way])
```

# Testing

```bash
make test                 # Run all unit tests
```
## Documentation

To access the docs locally, just change the host in the url to localhost:3000. Something like: http://localhost:3000/cart/docs
