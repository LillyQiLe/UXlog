package controllers

import (
	"testing"
	"crypto"
	"encoding/hex"
	"strconv"
	"time"
	"os"
)

func TestRegisterController_Post(t *testing.T) {
	md5Ctx := crypto.SHA256.New()
	md5Ctx.Write([]byte("greatpassword&$^.." + strconv.FormatInt(time.Now().Unix(), 10)))
	md5PasswordHex := md5Ctx.Sum(nil)
	println(hex.EncodeToString(md5PasswordHex))
}

func TestPicMaker_OutputFile(t *testing.T) {
	const (
		dx	= 100
		dy	= 40
		fontFile = "SentyZHAO.ttf"
		fontSize = 23
		fontDPI = 72
	)

	mk := PicMaker{}
	mk.SetFormate(dx, dy, fontFile, fontSize, fontDPI)
	file,err := os.Create("qwe.png")
	if err != nil{
		println(err.Error())
	}
	defer file.Close()
	mk.OutputFile("hello")
	_,err = mk.WriteTo(file)
	if err != nil{
		println(err.Error())
	}

}