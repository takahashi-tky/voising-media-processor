package service

import (
	"cloud.google.com/go/storage"
	"fmt"
	"os/exec"
	"strings"
)

type ImagickService interface {
	GetFileFormat(*storage.Reader) (string, error)
	ConvertResize(reader *storage.Reader, width uint, height uint) error
}

type imagickService struct {
}

func (i *imagickService) ConvertResize(reader *storage.Reader, width uint, height uint) error {
	cmd := exec.Command("convert", "-", "-resize", fmt.Sprintf("%dx%d", width, height), "-")
	cmd.Stdin = reader

	err := cmd.Run()
	if err != nil {
		return err
	}
	//TODO implement me
	panic("implement me")
}

func (i *imagickService) GetFileFormat(reader *storage.Reader) (string, error) {
	cmd := exec.Command("identify", "-")
	cmd.Stdin = reader
	result, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.Split(string(result), " ")[1], err
}

func NewImagickService() ImagickService {
	return &imagickService{}
}
