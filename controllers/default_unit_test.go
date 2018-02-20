package controllers

import (
	"testing"
	"crypto"
	"encoding/hex"
	"strconv"
	"time"
)

func TestRegisterController_Post(t *testing.T) {
	md5Ctx := crypto.SHA256.New()
	md5Ctx.Write([]byte("greatpassword&$^.." + strconv.FormatInt(time.Now().Unix(), 10)))
	md5PasswordHex := md5Ctx.Sum(nil)
	println(hex.EncodeToString(md5PasswordHex))
}

