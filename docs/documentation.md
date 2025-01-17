# Go360 Documentation


## Overview

Go360 provides functionality to interact with Xbox 360 RGH, Jtag or XDK consoles over TCP.

## Constants

```go
const DefaultPort = 730 // Default port for XBDM
```

## Functions

### ParseHexToDec
```go
func ParseHexToDec(hexStr string) (int, error)
```
Converts a hexadecimal string to decimal integer.

### ToHex
```go
func ToHex(str string) string
```
Converts a string to its hexadecimal representation.

## Types

### Config
```go
type Config struct {
    Timeout time.Duration
}
```
Config holds the configuration options for Xbox360.

#### DefaultConfig
```go
func DefaultConfig() *Config
```
DefaultConfig returns a Config with default values.

### Xbox360
```go
type Xbox360 struct {
	conn      net.Conn
    connected bool
    timeout   time.Duration
}
```
Xbox360 represents a connection to a console.

#### NewXbox360
```go
func NewXbox360(cfg *Config) *Xbox360
```
NewXbox360 creates a new Xbox360 instance with the provided configuration.

### Methods

#### Connect
```go
func (x *Xbox360) Connect(ip string, port int) error
```
Establishes a connection to the Xbox 360 console.

#### Disconnect
```go
func (x *Xbox360) Disconnect() error
```
Closes the connection to the console.

#### SendCommand
```go
func (x *Xbox360) SendCommand(command string) (string, error)
```
Sends a command to the console and returns the response.

#### Memory Operations

##### GetMemory
```go
func (x *Xbox360) GetMemory(address string, length int, memoryType string) (string, error)
```
Reads memory from the specified address.

##### SetMemory
```go
func (x *Xbox360) SetMemory(address, data []byte) error
```
Writes data to the specified address.

#### System Control

##### PauseSystem
```go
func (x *Xbox360) PauseSystem() error
```
Pauses the console.

##### UnPauseSystem
```go
func (x *Xbox360) UnPauseSystem() error
```
Unpauses the console.

##### Shutdown
```go
func (x *Xbox360) Shutdown() error
```
Turns off the console.

##### ColdReboot
```go
func (x *Xbox360) ColdReboot() error
```
Performs a cold reboot of the console.

##### WarmReboot
```go
func (x *Xbox360) WarmReboot() error
```
Performs a warm reboot of the console.

#### Other Operations

##### LaunchXeX
```go
func (x *Xbox360) LaunchXeX(xexPath string) (string, error)
```
Launches an XEX file from the specified path.

##### XNotify
```go
func (x *Xbox360) XNotify(message string) error
```
Sends a notification to the console.