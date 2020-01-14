package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var result []string
	err := json.NewEncoder(os.Stdout).Encode(result)
	if err != nil {
		fmt.Errorf("encode synced PVC list failed: %v", err)
	}
}
