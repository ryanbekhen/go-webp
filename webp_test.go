package webp

import (
	"bytes"
	"image"
	"image/color"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name    string
		img     image.Image
		quality float32
	}{
		{
			name:    "empty image",
			img:     image.NewRGBA(image.Rect(0, 0, 0, 0)),
			quality: float32(75.0),
		},
		{
			name:    "small image",
			img:     image.NewRGBA(image.Rect(0, 0, 2, 2)),
			quality: float32(75.0),
		},
		{
			name:    "large image",
			img:     randomImage(500, 500),
			quality: float32(75.0),
		},
		{
			name:    "high quality",
			img:     randomImage(100, 100),
			quality: float32(100.0),
		},
		{
			name:    "low quality",
			img:     randomImage(100, 100),
			quality: float32(10.0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			err := Encode(tt.img, tt.quality, buf)
			if err != nil {
				t.Fatalf("Encode() error = %v, wantErr = %v", err, false)
			}
		})
	}
}

func TestEncode_EmptyImage(t *testing.T) {
	tests := []struct {
		name    string
		img     image.Image
		quality float32
	}{
		{
			name:    "nil image",
			img:     nil,
			quality: float32(75.0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			err := Encode(tt.img, tt.quality, buf)
			if err == nil {
				t.Fatalf("Encode() error = %v, wantErr = %v", err, true)
			}
		})
	}
}

// This function generates an image with random colors
func randomImage(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{R: uint8(x % 256), G: uint8(y % 256), B: uint8((x + y) % 256), A: 255})
		}
	}
	return img
}
