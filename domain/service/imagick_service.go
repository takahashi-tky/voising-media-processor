package service

import (
	"bytes"
	"cloud.google.com/go/storage"
	"fmt"
	"log"
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
	var stdout bytes.Buffer
	cmd := exec.Command("identify", "-")
	cmd.Stdin = reader
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	log.Println(stdout.String())
	return strings.Split(stdout.String(), " ")[0], err
}

func NewImagickService() ImagickService {
	log.Println("NewImagickService")
	return &imagickService{}
}
