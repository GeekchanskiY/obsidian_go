package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("Setting up...")

	exitCode := m.Run()

	println("Tearing down...")

	os.Exit(exitCode)
}

<<<<<<< HEAD
func TestHe(t *testing.T) {
=======
func TestPlayground(t *testing.T) {
>>>>>>> 6331a2e (playing with tests)
	if testing.Short() {
		t.Skip("skipping test in a short mode")
	}
	t.Log("Hello World")

	t.Fail()
}
