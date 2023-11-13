package service

import "gopkg.in/gographics/imagick.v3/imagick"

type ImagickService interface {
	GetBlob() (blob []byte)
	GetFileFormat() string
	ReadBlob(blob []byte) (err error)
	Resize(width uint, height uint) (err error)
	ConvertFormat(format string) (err error)
	Close()
}

type imagickService struct {
	mw         *imagick.MagickWand
	fileFormat string
}

func (i *imagickService) GetBlob() (blob []byte) {
	return i.mw.GetImageBlob()
}

func (i *imagickService) Close() {
	imagick.Terminate()
	i.mw.Destroy()
}

func (i *imagickService) GetFileFormat() string {
	return i.fileFormat
}

func (i *imagickService) ConvertFormat(format string) (err error) {
	return i.mw.SetImageFormat(format)
}

func (i *imagickService) ReadBlob(blob []byte) (err error) {
	err = i.mw.ReadImageBlob(blob)
	if err != nil {
		return err
	}
	i.fileFormat = i.mw.GetFormat()
	return err
}

func (i *imagickService) Resize(width uint, height uint) (err error) {
	return i.mw.ResizeImage(width, height, imagick.FILTER_LANCZOS)
}

func NewImagickService() ImagickService {
	imagick.Initialize()
	return &imagickService{
		mw:         imagick.NewMagickWand(),
		fileFormat: "",
	}
}
