package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// HashPassword uses MD5 which gosec should flag as weak
func HashPassword(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// ExecuteCommand has potential command injection vulnerability
func ExecuteCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)
	return cmd.Run()
}

// ReadFile with potential path traversal issue
func ReadFile(filename string) ([]byte, error) {
	// G304: File path provided as taint input
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	return io.ReadAll(file)
}

// UnsafeFilePermissions sets overly permissive file permissions
func UnsafeFilePermissions(filename string) error {
	// G302: Poor file permissions used when creating a file
	return os.WriteFile(filename, []byte("test"), 0777)
}