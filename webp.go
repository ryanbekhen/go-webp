package webp

/*
#cgo LDFLAGS: -lwebp
#include <stdlib.h>
#include <webp/encode.h>

int encodeWebP(uint8_t* img, int width, int height, int stride, float quality_factor, uint8_t** output, size_t* output_size) {
    WebPConfig config;
    if (!WebPConfigInit(&config)) {
        return 0;
    }
    config.quality = quality_factor;
    if (!WebPValidateConfig(&config)) {
        return 0;
    }
    *output_size = WebPEncodeRGBA(img, width, height, stride, quality_factor, output);
    return 1;
}
*/
import "C"
import (
	"errors"
	"image"
	"io"
	"unsafe"
)

// Encode encodes the input image in the WebP format with the specified quality and writes it to the given destination writer.
// The function converts the image into a byte slice and calls a C function to perform the encoding.
// The resulting encoded data is written to the destination writer.
// If there are any errors during the encoding or writing process, an error is returned.
// Note: The C function is not included in this code snippet, it must be implemented separately.
// Example usage:
//
//	err := webp.Encode(img, 75, w)
//	if err != nil {
//	    log.Fatal(err)
//	}
func Encode(img image.Image, quality float32, dest io.Writer) error {
	if img == nil {
		return errors.New("img must not be nil")
	}

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	output := make([]byte, 4*width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			index := y*width*4 + x*4
			output[index+0] = byte(r >> 8)
			output[index+1] = byte(g >> 8)
			output[index+2] = byte(b >> 8)
			output[index+3] = byte(a >> 8)
		}
	}

	// Convert the go byte slice to c array.
	cArray := (*C.uint8_t)(C.CBytes(output))
	defer C.free(unsafe.Pointer(cArray))

	var outputSize C.size_t
	result := C.encodeWebP(cArray, C.int(width), C.int(height), C.int(width*4), C.float(quality), &cArray, &outputSize)

	if result == 0 {
		return errors.New("failed to encode image to webp")
	}

	// Get the bytes, don't forget to free it after getting the bytes.
	encodedData := C.GoBytes(unsafe.Pointer(cArray), C.int(outputSize))
	C.free(unsafe.Pointer(cArray))

	_, err := dest.Write(encodedData)
	if err != nil {
		return errors.New("failed to write encoded data")
	}

	return nil
}
