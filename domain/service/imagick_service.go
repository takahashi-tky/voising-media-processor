package service

import (
	"bytes"
	"cloud.google.com/go/storage"
	"fmt"
	"os/exec"
	"strings"
)

type ImagickService interface {
	GetFileFormat(*storage.Reader) (string, error)
	ConvertResize(reader *storage.Reader, width uint, height uint) (buf bytes.Buffer, err error)
}

type imagickService struct {
}

func (i *imagickService) ConvertResize(reader *storage.Reader, width uint, height uint) (buf bytes.Buffer, err error) {
	cmd := exec.Command("convert", "-", "-resize", fmt.Sprintf("%dx%d", width, height), "-")
	cmd.Stdin = reader
	cmd.Stdout = &buf
	err = cmd.Run()
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("cmd.Run: %w", err)
	}
	return buf, err
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
