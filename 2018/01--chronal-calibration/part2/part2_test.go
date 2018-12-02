package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	if os.Getenv("GO_TEST_RUN_MAIN") == "1" {
		main()
	} else {
		os.Exit(m.Run())
	}
}

func TestPart2(t *testing.T) {
	cmd := exec.Command(os.Args[0])
	cmd.Dir = "../"
	cmd.Env = []string{"GO_TEST_RUN_MAIN=1"}
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("got error: %v", err)
	}
	output := string(outputBytes)
	wantOutput := "Solution: 69074\n"
	if output == wantOutput {
		t.Logf("got %q", output)
	} else {
		t.Errorf("got %q; want %q", output, wantOutput)
	}
}
