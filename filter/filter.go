package filter

import (
	"github.com/disintegration/imaging"
)

type GrayscaleFilter struct{}

func (f *GrayscaleFilter) Process(srcPath, dstPath string) error {
	srcImg, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	grayImg := imaging.Grayscale(srcImg)

	err = imaging.Save(grayImg, dstPath)
	if err != nil {
		return err
	}

	return nil
}

type BlurFilter struct{}

func (f *BlurFilter) Process(srcPath, dstPath string) error {
	srcImg, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	blurImg := imaging.Blur(srcImg, 10.0)

	err = imaging.Save(blurImg, dstPath)
	if err != nil {
		return err
	}

	return nil
}
