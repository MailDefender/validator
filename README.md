# **Validator**

**This microservice acts as a gateway during token validation to avoid directly exposing the engine's APIs.**

> Note: this documentation is currently being drafted and will be completed in a future version.

## 📦 Prerequisites

- **Golang**
- **Docker** (optional)
- **[Environment Variables](#-configuration)**

## 🚀 Installation

### With Docker (Recommended)

```bash
docker build -t maildefender/validator .
docker run -p 8080:8080 --env-file .env maildefender/validator
```

### Without Docker

1. Clone the repository

```bash
# Clone the repoisitory
git clone https://github.com/MailDefender/validator.git
cd validator

# Install dependencies
go mod download

# Build
go build -o validator

# Run
source .env
./validator
```

## 🏃‍♂️ Usage

This app exposes APIs, so please refer to the Swagger to get more details about its usage.

## 🧪 Tests

As this app acts as a gateway, no tests are available.

## 🛠 Configuration

Create a .env file in the project root with the following variables:

```shell
# URL on which the engine can be reached to process token validation request
ENGINE_BASE_ENDPOINT=http://engine:8081
```

> Note: change the host and port according to your local (or docker) configuration

## 📜 License

This project is licensed under MIT.
