package main

import (
	"fmt"
	"os"

	"github.com/think0rcode/machineid"
)

func main() {
	fmt.Println("=== Checking machine id ===")
	fmt.Printf("  /etc/machine_id: %s\n", readFile("/etc/machine-id"))
	fmt.Printf("  /var/lib/dbus/machine_id: %s\n", readFile("/var/lib/dbus/machine-id"))
	fmt.Println()

	fmt.Println("=== Checking product uuid ===")
	fmt.Printf("  /sys/class/dmi/id/product_uuid: %s\n", readFile("/sys/class/dmi/id/product_uuid"))
	fmt.Println()

	fmt.Println("=== Generated ID ===")
	id, err := machineid.ID()
	if err != nil {
		fmt.Printf("Error generating ID: %v\n", err)
	}

	fmt.Printf("Generated ID: %s\n", id)
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "no such file or directory"
		}

		return fmt.Sprintf("error (%v)", err)
	}

	return string(data)
}
