package assetlayer

import (
	"fmt"

	"github.com/kr/pretty"
)

func (client *Client) Try(method, path string, query map[string]string, body ...interface{}) (interface{}, error) {

	url := client.URL(path)
	println(method, url)

	if query != nil {
		pretty.Println(query)
	}

	response := &Response{}
	r := client.NewRequest().EnableTrace().SetResult(response).SetError(response)
	if query != nil {
		r = r.SetQueryParams(query)
		pretty.Println(query)
	}
	if len(body) > 0 {
		r = r.SetBody(body[0])
	}
	resp, err := r.Execute(method, url)
	if err != nil {
		return nil, err
	}
	code := resp.StatusCode()
	if code < 200 || code >= 300 {
		pretty.Println(response)
		return nil, fmt.Errorf("GetResponse: failed with code: %d", code)
	}
	return response.Body, nil
}
