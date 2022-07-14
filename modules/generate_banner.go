package modules

import (
	"github.com/disintegration/imaging"
	"image"
	"strconv"
)

func GenerateBanner(profilePhotos []string, bgFile string) error {
	var i = 0
	var tmp *image.NRGBA

	for _, profilePhoto := range profilePhotos {
		i++
		var loopImage = "img/tmp/" + strconv.Itoa(i) + ".jpg"

		var _, err = DownloadFile(loopImage, profilePhoto)
		if err != nil {
			return err
		}

		src, err := imaging.Open(bgFile)
		if err != nil {
			return err
		}

		photo, err := imaging.Open(loopImage)
		if err != nil {
			return err
		}

		photo = imaging.Resize(photo, 140, 140, imaging.Lanczos)

		switch i {
		case 1:
			tmp = imaging.Paste(src, photo, image.Pt(623, 180))
		case 2:
			tmp = imaging.Paste(src, photo, image.Pt(785, 180))
		case 3:
			tmp = imaging.Paste(src, photo, image.Pt(947, 180))
		case 4:
			tmp = imaging.Paste(src, photo, image.Pt(1109, 180))
		case 5:
			tmp = imaging.Paste(src, photo, image.Pt(1271, 180))
		}

		err = imaging.Save(tmp, bgFile, imaging.JPEGQuality(100))
		if err != nil {
			return err
		}
	}
	return nil
}
