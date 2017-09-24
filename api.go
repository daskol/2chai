package main

import (
	"net/http"
	"strconv"

	"github.com/pquerna/ffjson/ffjson"
)

func ListThreads(board string, page int) (*Threads, error) {
	endpoint := strconv.Itoa(page)

	if page == 0 {
		endpoint = "index"
	}

	url := "https://2ch.hk/" + board + "/" + endpoint + ".json"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	cli := &http.Client{}
	res, err := cli.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	dec := ffjson.NewDecoder()
	threads := &Threads{}

	if err := dec.DecodeReader(res.Body, threads); err != nil {
		return nil, err
	}

	return threads, nil
}

func ListThreadCatalog(board string) (*Threads, error) {
	url := "https://2ch.hk/" + board + "/catalog.json"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	cli := &http.Client{}
	res, err := cli.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	dec := ffjson.NewDecoder()
	threads := &Threads{}

	if err := dec.DecodeReader(res.Body, &threads); err != nil {
		return nil, err
	}

	return threads, nil
}
