package service

import (
	"cloud.google.com/go/storage"
	"os/exec"
	"strings"
)

type ImagickService interface {
	GetFileFormat(*storage.Reader) (string, error)
}

type imagickService struct {
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
