# Beaver - Logging Service

## Overview

Beaver is a lightweight and efficient logging service written in Go. It provides structured logging capabilities, including middleware for HTTP routes, to enhance observability and debugging for your applications.

## Features

- Structured logging with JSON output.
- Middleware for logging HTTP requests and responses.
- Configurable log levels (info, warn, error).
- Support for logging to different outputs (console, file, remote services).
- Easy integration with existing Go applications.

## Installation

```sh
# Clone the repository
git clone https://github.com/yourusername/beaver.git
cd beaver

# Build the service
go build -o beaver

# Run the service
./beaver
```

## Usage

### Importing Beaver in Your Go Application

```go
import (
    "github.com/yourusername/beaver"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })
    
    loggedMux := beaver.LoggingMiddleware(mux)
    
    http.ListenAndServe(":8080", loggedMux)
}
```

### Logging Example

```go
package main

import (
    "github.com/yourusername/beaver"
    "log"
)

func main() {
    logger := beaver.NewLogger(beaver.Config{Level: "info"})
    
    logger.Info("Beaver logging initialized successfully")
    logger.Warn("This is a warning message")
    logger.Error("An error occurred")
}
```

## Configuration

Beaver can be configured using environment variables or a config file:

### Using Environment Variables

Set the following environment variables before running your application:

```sh
export BEAVER_LOG_LEVEL=info
export BEAVER_LOG_OUTPUT=console
export BEAVER_LOG_FILE=logs/app.log
```

### Using a Config File

Beaver also supports configuration via a JSON or YAML file. Example JSON configuration:

```json
{
    "log_level": "info",
    "log_output": "file",
    "log_file": "logs/app.log"
}
```

Example YAML configuration:

```yaml
log_level: info
log_output: file
log_file: logs/app.log
```

To use a config file, specify the file path when initializing the logger:

```go
logger := beaver.NewLoggerFromFile("config.json")
```

| Environment Variable | Description | Default |
|----------------------|-------------|---------|
| `BEAVER_LOG_LEVEL` | Log level (`info`, `warn`, `error`) | `info` |
| `BEAVER_LOG_OUTPUT` | Log output (`console`, `file`, `remote`) | `console` |
| `BEAVER_LOG_FILE` | File path if `file` output is selected | `logs/app.log` |

## Authors

- [Ayden](https://github.com/ayden-boyko)
