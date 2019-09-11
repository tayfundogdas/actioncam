package libipcamera

import (
	"fmt"
	"net"
)

func ExampleCreateCamera() {
	cameraIP := net.ParseIP("192.168.0.1")

	// Create a camera
	camera := CreateCamera(cameraIP, 6666, "admin", "12345")
	defer camera.Disconnect()

	// Enable verbose output for debugging
	camera.SetVerbose(true)

	// Connect to the camera and start responding to keep-alive messages
	camera.Connect()

	// Send a login packet to enable camera control
	err := camera.Login()
	if err != nil {
		fmt.Printf("Failed to Login: %s\n", err)
	}

	// Make the camera take a still image
	err = camera.TakePicture()
	if err != nil {
		fmt.Printf("Failed to take a picture: %s\n", err)
	}
}
