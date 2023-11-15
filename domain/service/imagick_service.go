package service

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type ImagickService interface {
	DecodeBase64(buffer *bytes.Buffer) (buf bytes.Buffer, err error)
	GetFileFormat(buffer *bytes.Buffer) (string, error)
	ConvertResize(buffer *bytes.Buffer, width uint, height uint) (buf bytes.Buffer, err error)
	ConvertFormat(buffer *bytes.Buffer, format string) (buf bytes.Buffer, err error)
}

type imagickService struct {
}

func (i *imagickService) DecodeBase64(buffer *bytes.Buffer) (buf bytes.Buffer, err error) {
	if strings.Index(buffer.String(), ";base64,") >= 0 {
		var stdout bytes.Buffer
		cmd := exec.Command("identify", "inline:-")
		cmd.Stdin = buffer
		cmd.Stdout = &stdout
		err = cmd.Run()
		if err != nil {
			log.Println("Error identify inline:-")
			return bytes.Buffer{}, fmt.Errorf("cmd.Run: %w", err)
		}
		cmd = exec.Command("convert", "inline:-", strings.Split(stdout.String(), " ")[1]+":-")
		cmd.Stdin = buffer
		cmd.Stdout = &buf
		err = cmd.Run()
		if err != nil {
			return bytes.Buffer{}, fmt.Errorf("cmd.Run: %w", err)
		}
		return buf, err
	} else {
		return *buffer, err
	}
}

func (i *imagickService) ConvertFormat(buffer *bytes.Buffer, format string) (buf bytes.Buffer, err error) {
	cmd := exec.Command("convert", "-", strings.ToUpper(format)+":-")
	cmd.Stdin = buffer
	cmd.Stdout = &buf
	err = cmd.Run()
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("cmd.Run: %w", err)
	}
	return buf, err
}

func (i *imagickService) ConvertResize(buffer *bytes.Buffer, width uint, height uint) (buf bytes.Buffer, err error) {
	cmd := exec.Command("convert", "-", "-resize", fmt.Sprintf("%dx%d", width, height), "-")
	cmd.Stdin = buffer
	cmd.Stdout = &buf
	err = cmd.Run()
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("cmd.Run: %w", err)
	}
	return buf, err
}

func (i *imagickService) GetFileFormat(buffer *bytes.Buffer) (string, error) {
	var stdout bytes.Buffer
	cmd := exec.Command("identify", "-")
	cmd.Stdin = buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.ToLower(strings.Split(stdout.String(), " ")[1]), err
}

func NewImagickService() ImagickService {
	return &imagickService{}
}
