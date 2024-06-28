# GO WebP

This is a simple library Go to convert images to WebP format using the [libwebp](https://developers.google.com/speed/webp/docs/api) library.

## Development Requirements

- Go 1.22
- libwebp-dev

Before running the project, you need to install the libwebp-dev package. You can install it by running the following command:

```bash
# Ubuntu
sudo apt-get install libwebp-dev
```

## Installation

To install the library, you need to run the following command:

```bash
go get github.com/ryanbekhen/go-webp
```

## Usage

To use the library, you need to run the following command:

```go
package main

import (
	"github.com/ryanbekhen/go-webp"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	input := "img.png"
	output := "img.webp"

	// Open the input image file.
	imgFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()

	// Decode the input image file.
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	// Create the output image file.
	outFile, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Encode the input image file.
	err = webp.Encode(img, 75, outFile)
	if err != nil {
		panic(err)
	}
}
```

> Note: The quality parameter is a number between 0 and 100.
> import _ "image/jpeg" and import _ "image/png" are required to decode the input image file.

## Contributing

Contributions are welcome! For feature requests and bug reports please [submit an issue](https://github.com/ryanbekhen/go-webp/issues).

## License

This project is licensed under the [MIT License](LICENSE).