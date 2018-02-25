package controllers

import (
	"image"
	"github.com/golang/freetype"
	"io/ioutil"
	"image/color"
	"io"
	"image/png"
)

type setting struct {
	Dx    int
	Dy    int
	FontFile string
	FontSize float64
	FontDPI float64
}

type PicMaker struct{
	img *image.NRGBA
	c	*freetype.Context
	code_setting setting
}


func (pm *PicMaker)SetFormate(dx int, dy int, fontFile string, fontSize float64, fontDPI float64)(error){
	pm.code_setting = setting{dx, dy, fontFile, fontSize, fontDPI}
	pm.img = image.NewNRGBA(image.Rect(0,0,dx,dy))
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err!=nil{
		return err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err!=nil{
		return err
	}
	pm.c = freetype.NewContext()
	pm.c.SetDPI(fontDPI)
	pm.c.SetFont(font)
	pm.c.SetFontSize(fontSize)
	pm.c.SetClip(pm.img.Bounds())
	pm.c.SetDst(pm.img)
	pm.c.SetSrc(image.White)
	return nil
}

func (pm *PicMaker)drawBackground(){
	for y := 0; y < pm.code_setting.Dy; y++{
		for x := 0; x < pm.code_setting.Dx; x++{
			pm.img.Set(x,y,color.RGBA{uint8(x), uint8(y), 0, 255})
		}
	}
}

func (pm *PicMaker)setText (text string) (error){
	pt := freetype.Pt(10, 10 + 15) // 字出现的位置
	_, err := pm.c.DrawString(text, pt)
	if err!=nil{
		return err
	}else{
		return nil
	}
}

func (pm *PicMaker)OutputFile(str string){
	pm.drawBackground()
	pm.setText(str)
}

func (pm *PicMaker) WriteTo(w io.Writer) (int64, error) {
	return 0, png.Encode(w, pm.img)
}
