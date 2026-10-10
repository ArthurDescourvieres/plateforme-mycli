package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func runMain(t *testing.T, args ...string) (string, error) {
	t.Helper()

	cmd := exec.Command(os.Args[0], "-test.run=^TestRunMainSubprocess$")
	cmd.Env = append(os.Environ(),
		"MYCLI_RUN_MAIN=1",
		"MYCLI_ARGS="+strings.Join(args, " "),
		"MYCLI_CONFIG="+filepath.Join(t.TempDir(), "missing.json"),
		"MYCLI_ACCESS_KEY=",
		"MYCLI_SECRET_KEY=",
	)

	output, err := cmd.CombinedOutput()
	return string(output), err
}

func TestRunMainSubprocess(t *testing.T) {
	if os.Getenv("MYCLI_RUN_MAIN") != "1" {
		t.Skip("only runs as a subprocess of TestHelpWithoutCredentials")
	}

	os.Args = append([]string{"mycli"}, strings.Fields(os.Getenv("MYCLI_ARGS"))...)
	main()
}

func TestHelpWithoutCredentials(t *testing.T) {
	for _, args := range [][]string{
		{"--help"},
		{"--version"},
		{"bucket", "--help"},
		{"object", "--help"},
	} {
		if output, err := runMain(t, args...); err != nil {
			t.Errorf("mycli %s: %v\n%s", strings.Join(args, " "), err, output)
		}
	}

	if _, err := runMain(t, "bucket", "list"); err == nil {
		t.Error("mycli bucket list: expected an error without credentials")
	}
}
