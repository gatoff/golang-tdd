package main //golang-tdd

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {

	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "555-2368")

	if buffer.Len() == 0 {
		t.Errorf("Generated QR code lenght is zero")
	}

	_, err := png.Decode(buffer)

	if err != nil {
		t.Errorf("Generated QR Code is not a PNG: %s", err)
	}
}
