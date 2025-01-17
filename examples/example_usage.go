package main

import (
	"Go360/pkg/go360"
	"fmt"
	"log"
)

func main() {
	// Replace with your Xbox console's IP and port
	xboxIP := "192.168.4.39"
	xboxPort := 730 // XBDM's default port

	// Create a new Xbox360 instance with default configuration
	xbox := go360.NewXbox360(go360.DefaultConfig())

	// If you want you can play around with your own config here:
	/*
	   config := &go360.Config{
	       Timeout: 10 * time.Second,
	   }
	   xbox := go360.NewXbox360(config)
	*/

	// Connect to the console
	err := xbox.Connect(xboxIP, xboxPort)
	if err != nil {
		log.Fatalf("Failed to connect to Xbox 360: %v", err)
	}
	defer func() {
		// Disconnect from the console when the program ends
		err := xbox.Disconnect()
		if err != nil {
			log.Printf("Failed to disconnect from Xbox 360: %v", err)
		}
	}()

	fmt.Println("Connected to:", xboxIP)

	// Example 1: Send an XNotify notification
	err = xbox.XNotify("Hello from go360!")
	if err != nil {
		log.Printf("Failed to send XNotify: %v", err)
	} else {
		fmt.Println("XNotify sent successfully!")
	}

	// Example 2: Launch a .XEX file
	result, err := xbox.LaunchXeX("Hdd:\\Applications\\DashLaunch\\Installer\\default.xex")
	if err != nil {
		log.Fatalf("Error launching XEX: %v", err)
	}
	fmt.Println(result)
}
