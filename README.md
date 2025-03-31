# EventsIO

EventsIO is a simple and flexible library designed to handle events within G3deon. Its main goal is to centralize event logic in the application, providing a unified interface for different messaging implementations.

## Features

- **Simple Interface**: Intuitive API for publishing and **subscribing** to events
- **Multiple Implementations**: Support for different messaging backends
  - RabbitMQ (implemented)
  - Kafka (coming soon)
  - Memory (for testing)
- **Flexibility**: Allows advanced options when needed
- **Thread-Safe**: Concurrency-safe implementation
- **JSON Serialization**: Automatic event serialization/deserialization

## Installation

```bash
go get github.com/g3deon/aventsio
```

## Getting Started

### Pre-commit Hook Setup

To ensure code quality, the project includes a pre-commit hook that runs:
- Go code formatting (`go fmt`)
- Linting with `golangci-lint`
- Unit tests (`go test`)

To set up the hook:

```bash
# Configure hooks directory
git config core.hooksPath githooks

# Give execution permissions to the script
git add githooks/pre-commit.sh
git update-index --chmod=+x githooks/pre-commit.sh
```

Now, every time you try to commit, these checks will run automatically.

### Basic Usage

```go
import (
    "github.com/g3deon/aventsio"
    "github.com/g3deon/aventsio/adapters/memory"
)

// Create in-memory event bus
bus := memory.NewEventBus()

// Publish event
event := eventsio.NewEvent("event-id")
bus.Publish("users", event)

// Subscribe to events
bus.Subscribe("users", func(event eventsio.Event) {
    fmt.Printf("Received event: %s\n", event.GetID())
})
```

## Advanced Options

### With Retries

```go
// Publish with automatic retry
bus.Publish("users", event, eventsio.WithRetry())
```

```go
// Retry count its increment every time the event its re-delivered
bus.Subscribe("users", func(event eventsio.Event) {
    fmt.Printf(
      "Received event: %s\n, with re-try: %d", 
      event.GetID(), 
      event.GetRetryCount(),
    )
})
```

More comming soon...

## Event Structure

Each event contains:
- Unique ID
- Send timestamp
- Customizable headers
- Event-specific data

## Contributing

Contributions are welcome. Please ensure to:
1. Follow existing code conventions
2. Add tests for new functionality
3. Update documentation as needed

## License

MIT
