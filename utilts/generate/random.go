package generate

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateRandomPassword(lenNumber int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	password := make([]byte, lenNumber)
	for i := range password {
		password[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(password), nil
}

func GenerateOTP(length int) int {
	rand.Seed(time.Now().UnixNano())
	characters := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	otp := 0
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(characters))
		otp = otp*10 + characters[randomIndex]
	}

	// ตรวจสอบความยาวของ OTP
	for len(strconv.Itoa(otp)) < length {
		randomIndex := rand.Intn(len(characters))
		otp = otp*10 + characters[randomIndex]
	}

	return otp
}

func GenerateReferenceID(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
	}
	return sb.String()
}
