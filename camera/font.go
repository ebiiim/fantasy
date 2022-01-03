package camera

import (
	"io/ioutil"

	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

var fontPixelMplusRegular *sfnt.Font
var fontPixelMplusBold *sfnt.Font

func init() {
	var err error
	fontPixelMplusRegular, err = loadFont("assets/font/PixelMplus-20130602/PixelMplus12-Regular.ttf")
	if err != nil {
		panic(err)
	}
	fontPixelMplusBold, err = loadFont("assets/font/PixelMplus-20130602/PixelMplus12-Bold.ttf")
	if err != nil {
		panic(err)
	}
}

func loadFont(src string) (*sfnt.Font, error) {
	f, err := assets.Open(src)
	if err != nil {
		return nil, err
	}
	p, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	tt, err := opentype.Parse(p)
	if err != nil {
		return nil, err
	}
	return tt, nil
}
