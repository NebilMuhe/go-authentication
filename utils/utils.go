package utils

import "github.com/segmentio/ksuid"

func GenerateID() string {
	return ksuid.New().String()
}

func RemoveIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}
