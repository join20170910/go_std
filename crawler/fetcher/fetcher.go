package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(any(err))
		}
		return all, err
	} else {
		return nil, fmt.Errorf("error status code: %d", resp.StatusCode)
	}
	return nil, nil
}
