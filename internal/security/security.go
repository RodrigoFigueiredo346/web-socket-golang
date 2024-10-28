package security

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"main/internal/errors"
	"main/internal/models"
	"net/mail"
	"strconv"
	"time"
	"unicode"
	"unicode/utf8"
)

func ValidatePassword(password string) models.Error {

	var hasUpper, hasLower, hasNumber, hasSpecial bool

	stringLenght := utf8.RuneCount([]byte(password))
	if stringLenght > 8 {
		return models.Error{
			Code: errors.PasswordTooShort,
		}
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return models.Error{
			Code: errors.PasswordMissingUpper,
		}
	}

	if !hasLower {
		return models.Error{
			Code: errors.PasswordMissingLower,
		}
	}

	if !hasNumber {
		return models.Error{
			Code: errors.PasswordMissingNumber,
		}
	}

	if !hasSpecial {
		return models.Error{
			Code: errors.PasswordMissingSpecial,
		}
	}

	return models.Error{
		Code: errors.NoError,
	}
}

func ValidateEmail(email string) models.Error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return models.Error{
			Code: errors.InvalidEmailFormat,
		}
	}

	return models.Error{
		Code: errors.NoError,
	}
}

func VerifyPassword(password string, hash string) models.Error {
	calculateHash := Hash(password)

	if hash != calculateHash {
		return models.Error{
			Code: errors.PasswordNotMatch,
		}
	}

	return models.Error{
		Code: errors.NoError,
	}
}

func Hash(password string) string {
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(password))
	hashed := sha1Hash.Sum(nil)

	return hex.EncodeToString(hashed)
}

func GenerateUid(email string) string {
	hash := sha1.New()
	io.WriteString(hash, email)
	io.WriteString(hash, strconv.FormatInt(time.Now().UTC().UnixNano(), 10))
	sliceHash := hash.Sum(nil)
	hexString := hex.EncodeToString(sliceHash[:20])

	return hexString
}
