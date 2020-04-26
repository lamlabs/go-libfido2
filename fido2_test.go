package libfido2_test

import (
	"testing"

	fido2 "github.com/keys-pub/go-libfido2"
)

// See examples package instead.

func TestDetectDevices(t *testing.T) {
	detected, err := fido2.DetectDevices(100)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d devices", len(detected))
}