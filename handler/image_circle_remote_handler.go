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

	"github.com/chree188/PosterSDK/circlemask"
	"github.com/chree188/PosterSDK/core"
	"github.com/nfnt/resize"
)

// ImageCircleRemoteHandler 根据URL地址设置圆形图片
type ImageCircleRemoteHandler struct {
	// 合成复用Next
	Next
	X   int
	Y   int
	Width uint
	Hight uint
	URL string //http://xxx.png
}

// Do 地址逻辑
func (h *ImageCircleRemoteHandler) Do(c *Context) (err error) {
	srcReader, err := core.GetResourceReader(h.URL)
	if err != nil {
		fmt.Errorf("core.GetResourceReader err：%v", err)
	}
	srcImage, imageType, err := image.Decode(srcReader)
	_ = imageType
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
