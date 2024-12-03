package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	var result string
	for i := 0; i < length; i++ {
		result += strconv.Itoa(rand.Intn(10))
	}
	return result

}

func GenerateString(length int) string {
	rand.Seed(time.Now().UnixNano())
	RuneArr := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var sb strings.Builder
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteRune(RuneArr[rand.Intn(len(RuneArr))])
	}
	return sb.String()
}

func GenerateGUID() string {
	guid := uuid.New().String()

	return guid
}

func GenerateAlphanumeric(length int) string {
	RuneArr := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteRune(RuneArr[rand.Intn(len(RuneArr))])
	}
	return sb.String()
}

func GenerateEnum(values []string) string {
	randomIndex := rand.Intn(len(values))
	return values[randomIndex]
}