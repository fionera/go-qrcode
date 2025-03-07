package qrcode

import (
	"fmt"
	"image"
	"image/color"

	"github.com/yeqown/go-qrcode/matrix"
)

// defaultOutputOption default output image background color and etc options
func defaultOutputOption() *outputImageOptions {
	return &outputImageOptions{
		bgColor:      color.White,     // white
		qrColor:      color.Black,     // black
		logo:         nil,             //
		qrWidth:      20,              //
		shape:        _shapeRectangle, //
		imageEncoder: jpegEncoder{},
		encMode:      EncModeByte,
		ecLevel:      Quart,
	}
}

// outputImageOptions to output QR code image
type outputImageOptions struct {
	// bgColors
	bgColor color.Color

	// qrColor
	qrColor color.Color

	// logo this icon image would be put the center of QR Code image
	// NOTE: logo only should has 1/5 size of QRCode image
	logo image.Image

	// qrWidth width of each qr block
	qrWidth int

	// shape means how to draw the shape of each cell.
	shape IShape

	// imageEncoder specify which file format would be encoded the QR image.
	imageEncoder ImageEncoder

	// encMode specifies which encMode to use
	encMode encMode

	// ecLevel specifies which ecLevel to use
	ecLevel ecLevel
}

func (oo *outputImageOptions) backgroundColor() color.Color {
	if oo == nil || oo.bgColor == nil {
		return color.White
	}

	return oo.bgColor
}

// DEPRECATED
// qrColor would be save into `_stateToRGBA`
//func (oo *outputImageOptions) foregroundColor() color.Color {
//	if oo == nil || oo.qrColor == nil {
//		return color.Black
//	}
//
//	return oo.qrColor
//}

func (oo *outputImageOptions) logoImage() image.Image {
	if oo == nil || oo.logo == nil {
		return nil
	}

	return oo.logo
}

func (oo *outputImageOptions) qrBlockWidth() int {
	if oo == nil || (oo.qrWidth <= 0 || oo.qrWidth > 255) {
		return 20
	}

	return oo.qrWidth
}

func (oo *outputImageOptions) getShape() IShape {
	if oo == nil || oo.shape == nil {
		return _shapeRectangle
	}

	return oo.shape
}

var (
	// _stateToRGBA state map tp color.Gray16
	_stateToRGBA = map[matrix.State]color.Color{
		matrix.StateFalse: hexToRGBA("#ffffff"),
		matrix.StateTrue:  hexToRGBA("#000000"),
		matrix.StateInit:  hexToRGBA("#cdc9c3"),
		//matrix.StateVersion: hexToRGBA("#444444"),
		//matrix.StateFormat: hexToRGBA("#555555"),
		//matrix.StateFinder: hexToRGBA("#2BA859"),
		matrix.StateFinder: hexToRGBA("#000000"),
	}

	// _defaultStateColor default color of undefined matrix.State
	// it shouldn't be used.
	_defaultStateColor = hexToRGBA("#ff414d")
)

// stateRGBA get color.Color by value State
func (oo *outputImageOptions) stateRGBA(v matrix.State) color.Color {
	if v, ok := _stateToRGBA[v]; ok {
		return v
	}

	return _defaultStateColor
}

// hexToRGBA convert hex string into color.RGBA
func hexToRGBA(s string) color.RGBA {
	c := color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 0xff,
	}

	var err error
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")
	}
	if err != nil {
		panic(err)
	}

	return c
}
