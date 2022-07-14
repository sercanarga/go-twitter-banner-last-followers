package modules

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(filepath string, url string) (error, error) {
	resp, err := http.Get(url)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err, nil
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err, nil
}
