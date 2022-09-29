package token

import (
	"crud.com/packages/token"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const testEmail string = "test@mail.com"

var secretKey = []byte("secretekey")

func TestJwt(t *testing.T) {
	var duration, err = time.ParseDuration("12h")
	if err != nil {
		t.Errorf(err.Error())
	}
	var maker = token.JWTMaker{}
	var jwt, _ = maker.NewJWTMaker(secretKey)

	tkn, err := jwt.CreateToken(testEmail, duration)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.NotNilf(t, tkn, "this shit is broken")
}
