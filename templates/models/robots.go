package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/go-resty/resty/v2"
)

func getRobotURL(target string, query *url.Values) (string, error) {

	host := "https://europe-west2-ninja-punk-girls.cloudfunctions.net/"

	switch target {

	case "generate":
		host += "robot-generate"

	case "prepare":
		host += "robot-prepare"

	default:
		return "", fmt.Errorf("no switch case found for %s", target)
	}

	return host + "?" + query.Encode(), nil
}

func CallRobot(method, target string, query *url.Values, src, dst interface{}) (interface{}, error) {

	host, err := getRobotURL(target, query)
	if err != nil {
		return nil, err
	}

	log.Println("calling robot:", method, host)

	client := resty.New()

	var response *resty.Response

	switch method {
	case "GET":
		response, err = client.R().SetResult(dst).Get(host)
		if err != nil {
			return nil, err
		}
	case "POST":
		response, err = client.R().SetResult(dst).SetBody(src).Post(host)
		if err != nil {
			return nil, err
		}
	}

	return response.Result(), nil
}
