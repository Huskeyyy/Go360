
# Go360

<p>
  <a href="https://pkg.go.dev/github.com/Huskeyyy/Go360">
    <img src="https://pkg.go.dev/badge/github.com/Huskeyyy/Go360.svg" alt="Go Reference">
  </a>
  <a href="https://goreportcard.com/report/github.com/Huskeyyy/Go360">
    <img src="https://goreportcard.com/badge/github.com/Huskeyyy/Go360" alt="Go Report Card">
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License: MIT">
  </a>
  <a href="https://github.com/Huskeyyy/Go360">
    <img src="https://img.shields.io/github/go-mod/go-version/Huskeyyy/Go360" alt="Go Version">
  </a>
</p>


Go360 is an open-source Go library for interacting with RGH/XDK Xbox 360 consoles over TCP.

## Features

- TCP-based communication
- Memory reading and writing capabilities
- XEX file launching
- System control (shutdown, reboot, pause)
- XNotify message support
- Configurable timeouts and connection settings

## Installation

```bash
go get github.com/huskeyyy/go360
```

## Quick Start

```go
package main

import (
    "github.com/huskeyyy/go360/pkg/go360"
    "fmt"
    "log"
)

func main() {
    xbox := go360.NewXbox360(go360.DefaultConfig())
    
    err := xbox.Connect("192.168.1.100", 730)
    if err != nil {
        log.Fatal(err)
    }
    defer xbox.Disconnect()

    // Send a notification to the console
    err = xbox.XNotify("Hello from Go360!")
    if err != nil {
        log.Printf("Failed to send notification: %v", err)
    }
}
```

## Documentation

Full documentation is available in the [docs](./docs) directory.


## Examples

Check the [examples](./examples) directory for more detailed usage examples.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

This project is not affiliated with Microsoft or Xbox. Use at your own risk.
