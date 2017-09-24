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

func GetThread(board string, thread string) (*ThreadResponse, error) {
	url := "https://2ch.hk/" + board + "/res/" + thread + ".json"
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
	obj := &ThreadResponse{}

	if err := dec.DecodeReader(res.Body, &obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func ListPosts(board string, thread string) ([]*Post, error) {
	if res, err := GetThread(board, thread); err != nil {
		return nil, err
	} else {
		return res.Threads[0].Posts, nil
	}
}
