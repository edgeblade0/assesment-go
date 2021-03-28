package repository

import (
	"io/ioutil"
	"net/http"
)

func (r *repo) GetTracking() (string, int, error) {
	var htmlString string

	url := "https://gist.githubusercontent.com/nubors/eecf5b8dc838d4e6cc9de9f7b5db236f/raw/d34e1823906d3ab36ccc2e687fcafedf3eacfac9/jne-awb.html"

	response, err := http.Get(url)
	if err != nil {
		return htmlString, 0, err
	}

	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return htmlString, response.StatusCode, err
		}
		htmlString = string(bodyBytes)

	}

	return htmlString, response.StatusCode, err
}
