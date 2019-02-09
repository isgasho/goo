// deal with  pictures.Handling images, such as compression
package imageDealWith

import (
	"fmt"
	"github.com/nfnt/resize"
	"image/gif"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
	"strings"
	"time"
)

//compression pictures. typeValue such as "png" "jpeg" "jpg" "gif".
func Compression(typeValue string, width uint, arg ...string) error {
	typ := strings.ToUpper(typeValue)
	switch typ {
	case "JPG", "JPEG":
		jpegdealWithImage(width, arg...)
	case "PNG":
		pngdealWithImage(width, arg...)
	case "GIF":
		gifdealWithImage(width, arg...)
	default:
		return fmt.Errorf("您输入的类型不正确，仅支持png jpg jpeg gif类型")
	}
	return nil
}
func jpegdealWithImage(width uint, arg ...string) error {
	a := make([]byte, 20)
	for _, v := range arg {
		file, err := os.Open(v)
		if err != nil {
			return err
		}
		img, err := jpeg.Decode(file)
		file.Close()
		if err != nil {
			return err
		}
		img1 := resize.Resize(width, 0, img, resize.Lanczos3)
		source := rand.NewSource(time.Now().UnixNano())
		rand.New(source).Read(a)
		file1, err := os.Create(string(a))
		if err != nil {
			return nil
		}
		jpeg.Encode(file1, img1, nil)
		file1.Close()
	}
	return nil

}

func pngdealWithImage(width uint, arg ...string) error {
	a := make([]byte, 20)
	for _, v := range arg {
		file, err := os.Open(v)
		if err != nil {
			return err
		}
		img, err := png.Decode(file)
		file.Close()
		if err != nil {
			return err
		}
		img1 := resize.Resize(width, 0, img, resize.Lanczos3)
		source := rand.NewSource(time.Now().UnixNano())
		rand.New(source).Read(a)
		file1, err := os.Create(string(a))
		if err != nil {
			return nil
		}
		png.Encode(file1, img1)
		file1.Close()
	}
	return nil

}

func gifdealWithImage(width uint, arg ...string) error {
	a := make([]byte, 20)
	for _, v := range arg {
		file, err := os.Open(v)
		if err != nil {
			return err
		}
		img, err := gif.Decode(file)
		file.Close()
		if err != nil {
			return err
		}
		img1 := resize.Resize(width, 0, img, resize.Lanczos3)
		source := rand.NewSource(time.Now().UnixNano())
		rand.New(source).Read(a)
		file1, err := os.Create(string(a))
		if err != nil {
			return nil
		}
		gif.Encode(file1, img1, nil)
		file1.Close()
	}
	return nil

}
