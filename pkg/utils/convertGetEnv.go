package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetenvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		errMsg := fmt.Sprintf("environment variable '%s' not found", key)
		return v, errors.New(errMsg)
	}
	return v, nil
}

func GetenvInt(key string) (int, error) {
	s, err := GetenvStr(key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func GetenvBool(key string) (bool, error) {
	s, err := GetenvStr(key)
	if err != nil {
		return false, err
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return v, nil
}
