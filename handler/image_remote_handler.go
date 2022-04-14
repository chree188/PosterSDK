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

	"github.com/chree188/PosterSDK/core"
	"github.com/nfnt/resize"
)

// ImageRemoteHandler 根据URL地址设置图片
type ImageRemoteHandler struct {
	// 合成复用Next
	Next
	X   int
	Y   int
	Width uint
	Hight uint
	URL string //http://xxx.png
}

// Do 地址逻辑
func (h *ImageRemoteHandler) Do(c *Context) (err error) {
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
	srcPoint := image.Point{
		X: h.X,
		Y: h.Y,
	}
	core.MergeImage(c.PngCarrier, srcImage, srcImage.Bounds().Min.Sub(srcPoint))
	return
}
