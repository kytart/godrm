package drm

import (
	"fmt"
)

func ExampleHasDumbBuffer() {
	// This example shows how to test if your graphics card
	// supports 'dumb buffers' capability. With this capability
	// you can create simple memory-mapped buffers without any
	// driver-dependent code.

	file, err := OpenCard(0)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}
	defer file.Close()
	if !HasDumbBuffer(file) {
		fmt.Printf("drm device does not support dumb buffers")
		return
	}
	fmt.Printf("ok")

	// Output: ok
}

func ExampleListDevices() {
	// Shows how to enumerate the available dri devices
	for _, dev := range ListDevices() {
		fmt.Printf("Driver name: %s\n", dev.Name)
	}

	// Output: Driver name: i915
}
