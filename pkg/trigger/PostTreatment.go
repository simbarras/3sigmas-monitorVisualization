package trigger

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func Trigger(api, bucket string) error {
	resp, err := http.PostForm(api+"/"+bucket, url.Values{})
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error while closing response body: %s\n", err)
		}
	}(resp.Body)
	if resp.StatusCode != 200 {
		return err
	}
	return nil
}
