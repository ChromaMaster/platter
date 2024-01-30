package test

import "testing"

func SkipIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test...")
	}
}

func SkipUnit(t *testing.T) {
	if !testing.Short() {
		t.Skip("Skipping unit test...")
	}
}
