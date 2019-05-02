package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/log"
)

func installedInPath(name string) bool {
	cmd := exec.Command("which", name)
	outBytes, err := cmd.Output()
	return err == nil && strings.TrimSpace(string(outBytes)) != ""
}

func failf(format string, args ...interface{}) {
	log.Errorf(format, args...)
	os.Exit(1)
}

func main() {

	if !installedInPath("golangci-lint") {
		cmd := command.New("go", "get", "-u", "github.com/golangci/golangci-lint")

		log.Infof("\nInstalling golangci-lint")
		log.Donef("$ %s", cmd.PrintableCommandArgs())

		if out, err := cmd.RunAndReturnTrimmedCombinedOutput(); err != nil {
			failf("Failed to install golangci-lint: %s", out)
		}
	}

	log.Infof("\nRunning golangci-lint...")

	cmd := command.New("golangci-lint", "run")

	log.Printf("$ %s", cmd.PrintableCommandArgs())

	if out, err := cmd.RunAndReturnTrimmedCombinedOutput(); err != nil || strings.TrimSpace(out) != "" {
		log.Errorf("golangci-lint failed")
		log.Printf(out)
		failf("golangci-lint failed: %s", err)
	}
}
