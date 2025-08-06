package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Errorf(
			"Response failed with status code: %d and\nbody: %s\n",
			res.StatusCode,
		)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}

