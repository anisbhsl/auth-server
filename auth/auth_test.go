package auth

import (
	"github.com/anisbhsl/auth-server/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordValidation(t *testing.T) {
	authService := mock.AuthService{}

	passwordValidationTests := []struct {
		rawPass    string
		hashedPass string
		expected   bool
	}{
		{
		"letmein",
		authService.EncryptPassword("letmein"),
		true,
		},
		{
			"adm123",
			authService.EncryptPassword("admin123"),
			false,
		},
	}

	for _,tt:=range passwordValidationTests{
		assert.Equal(t,tt.expected,authService.ValidatePasswordHash(tt.rawPass,tt.hashedPass))
	}
}
