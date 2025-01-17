// Package go360 provides functionality to interact with Xbox 360 consoles over TCP
package go360

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

// DefaultPort is the default port for XBDM
const DefaultPort = 730

// Xbox360 represents a connection to a console
type Xbox360 struct {
	conn      net.Conn
	connected bool
	timeout   time.Duration
}

// Config holds the configuration options for Xbox360
type Config struct {
	Timeout time.Duration
}

// DefaultConfig returns a Config with default values
func DefaultConfig() *Config {
	return &Config{
		Timeout: 5 * time.Second,
	}
}

// NewXbox360 creates a new Xbox360 instance with the provided configuration
func NewXbox360(cfg *Config) *Xbox360 {
	if cfg == nil {
		cfg = DefaultConfig()
	}
	return &Xbox360{
		timeout: cfg.Timeout,
	}
}

// Connect establishes a connection to the Xbox 360 console
func (x *Xbox360) Connect(ip string, port int) error {
	if port == 0 {
		port = DefaultPort
	}
	address := fmt.Sprintf("%s:%d", ip, port)

	conn, err := net.DialTimeout("tcp", address, x.timeout)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	x.conn = conn
	x.connected = true
	return nil
}

// Disconnect closes the connection to the console
func (x *Xbox360) Disconnect() error {
	if !x.connected {
		return fmt.Errorf("not connected to console")
	}

	if err := x.conn.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	x.connected = false
	return nil
}

// SendCommand sends a command to the console and returns the response
func (x *Xbox360) SendCommand(command string) (string, error) {
	if !x.connected {
		return "", fmt.Errorf("not connected to console")
	}

	if _, err := fmt.Fprintf(x.conn, command+"\r\n"); err != nil {
		return "", fmt.Errorf("failed to send command: %w", err)
	}

	if err := x.conn.SetReadDeadline(time.Now().Add(x.timeout)); err != nil {
		return "", fmt.Errorf("failed to set read deadline: %w", err)
	}
	defer x.conn.SetReadDeadline(time.Time{})

	reader := bufio.NewReader(x.conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return "", nil
		}
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	response = strings.TrimSpace(response)
	if strings.HasPrefix(response, "202") {
		return response, fmt.Errorf("command failed: %s", response)
	}

	return response, nil
}

// SetMemory writes data to the specified address
func (x *Xbox360) SetMemory(address, data []byte) error {
	command := fmt.Sprintf("setmem addr=%s data=%s", address, data)
	_, err := x.SendCommand(command)
	if err != nil {
		return fmt.Errorf("failed to set memory: %w", err)
	}
	return nil
}

// GetMemory reads memory from the specified address
func (x *Xbox360) GetMemory(address string, length int, memoryType string) (string, error) {
	command := fmt.Sprintf("getmem addr=%s length=%d", address, length)

	data, err := x.SendCommand(command)
	if err != nil {
		return "", fmt.Errorf("failed to get memory: %w", err)
	}

	if memoryType == "dec" {
		parsedData, err := ParseHexToDec(data)
		if err != nil {
			return "", fmt.Errorf("failed to parse hex to decimal: %w", err)
		}
		data = fmt.Sprintf("%d", parsedData)
	}

	return data, nil
}

// LaunchXeX launches an XEX file from the specified path
func (x *Xbox360) LaunchXeX(xexPath string) (string, error) {
	directory := xexPath[:strings.LastIndex(xexPath, "\\")+1]
	command := fmt.Sprintf("magicboot title=\"%s\" directory=\"%s\"", xexPath, directory)

	if _, err := x.SendCommand(command); err != nil {
		return "", fmt.Errorf("failed to launch XEX: %w", err)
	}

	return fmt.Sprintf("Launching %s", directory), nil
}

// XNotify sends a notification to the console
func (x *Xbox360) XNotify(message string) error {
	messageLength := len(message)
	messageHex := ToHex(message)

	command := fmt.Sprintf("consolefeatures ver=2 type=12 params=\"A\\0\\A\\2\\%d/%d\\%s\\1\\0\\\"",
		2, messageLength, messageHex)

	_, err := x.SendCommand(command)
	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}
	return nil
}

// System Functions

// PauseSystem pauses the console
func (x *Xbox360) PauseSystem() error {
	_, err := x.SendCommand("debug stop")
	return err
}

// UnPauseSystem unpauses the console
func (x *Xbox360) UnPauseSystem() error {
	_, err := x.SendCommand("debug continue")
	return err
}

// Shutdown turns off the console
func (x *Xbox360) Shutdown() error {
	_, err := x.SendCommand("shutdown")
	return err
}

// ColdReboot performs a cold reboot of the console
func (x *Xbox360) ColdReboot() error {
	_, err := x.SendCommand("magicboot COLD")
	return err
}

// WarmReboot performs a warm reboot of the console
func (x *Xbox360) WarmReboot() error {
	_, err := x.SendCommand("magicboot WARM")
	return err
}
