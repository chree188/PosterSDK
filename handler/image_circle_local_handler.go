/**
 * @Author: admin@chree.cn
 * @Description:
 * @File:  image_local_handler
 * @Version: 1.1.0
 * @Date: 2020/5/22 08:51
 */

package handler

import (
	"fmt"
	"image"
	"os"

	"github.com/chree188/PosterSDK/circlemask"
	"github.com/chree188/PosterSDK/core"
	"github.com/nfnt/resize"
)

// ImageCircleLocalHandler 根据Path路径设置圆形图片
type ImageCircleLocalHandler struct {
	// 合成复用Next
	Next
	X     int
	Y     int
	Width uint
	Hight uint
	Path  string //./images/xx.png
}

// Do 地址逻辑
func (h *ImageCircleLocalHandler) Do(c *Context) (err error) {

	imageFile, err := os.Open(h.Path)
	if err != nil {
		fmt.Errorf("os.Open err：%v", err)
	}

	srcImage, _, err := image.Decode(imageFile)

	if err != nil {
		fmt.Errorf("SetRemoteImage image.Decode err：%v", err)
	}
	
	if h.Width > 0 && h.Hight > 0 {
		srcImage = resize.Resize(h.Width, h.Hight, srcImage, resize.Lanczos3)
	}
	
	// 算出图片的宽度和高试
	width := srcImage.Bounds().Max.X - srcImage.Bounds().Min.X
	hight := srcImage.Bounds().Max.Y - srcImage.Bounds().Min.Y

	//把头像转成Png,否则会有白底
	srcPng := core.NewPNG(0, 0, width, hight)
	core.MergeImage(srcPng, srcImage, srcImage.Bounds().Min)

	// 圆的直径以长边为准
	diameter := width
	if width > hight {
		diameter = hight
	}
	// 遮罩
	srcMask := circlemask.NewCircleMask(srcPng, image.Point{0, 0}, diameter)

	srcPoint := image.Point{
		X: h.X,
		Y: h.Y,
	}
	core.MergeImage(c.PngCarrier, srcMask, srcImage.Bounds().Min.Sub(srcPoint))
	return
}
