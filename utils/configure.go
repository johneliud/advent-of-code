package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Configure() (string, string, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return "", "", fmt.Errorf("failed loading .env file: %v", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", "", fmt.Errorf("failed getting working directory: %v", err)
	}

	splitWd := strings.Split(wd, "/")
	if len(splitWd) < 2 {
		fmt.Println("Invalid working directory.")
	}
	
	year, day := splitWd[len(splitWd)-2], splitWd[len(splitWd)-1][3:]
	return year, day, nil
}
