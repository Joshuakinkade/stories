package photos

import (
	"image"
	"io"
)

type Processor struct {
	image image.Image
}

func LoadImage(r io.Reader) (Processor, error) {
	return Processor{}, nil
}

func Resize(height, width int) {

}

func Export(w io.Writer, format Format) error {
	return nil
} 
