package server

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"os"
)

func writeToFile(fileName string, data []byte) error {
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write QR code to file: %w", err)
	}
	return nil
}

func generateQR(url string) ([]byte, error) {
	qrCode, err := qrcode.New(url, qrcode.Highest)
	if err != nil {
		return nil, fmt.Errorf("failed to create QR code: %w", err)
	}

	// Generate the PNG data of the QR code
	qrBytes, err := qrCode.PNG(256)
	if err != nil {
		return nil, fmt.Errorf("failed to generate QR PNG data: %w", err)
	}

	return qrBytes, nil
}
