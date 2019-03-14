package auto

import (
	"bytes"
	"log"
	"math/rand"
	"os/exec"
	"strings"
)

func runAdb(args ...string) {
	var b bytes.Buffer
	cmd := exec.Command("adb", args...)
	cmd.Stdout = &b
	cmd.Stderr = &b
	log.Printf("adb %s", strings.Join(args, " "))
	err := cmd.Run()
	if cmd.Process != nil {
		cmd.Process.Kill()
	}
	if err != nil {
		log.Fatalf("adb %s: %v", strings.Join(args, " "), err.Error())
	}
}

func RandInt64(min, max int) int {
	return rand.Intn(max-min) + min
}
