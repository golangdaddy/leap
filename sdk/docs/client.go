package docs

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

const (
	ENDPOINT_INSERT = "https://europe-west2-zoo-dev-01.cloudfunctions.net/func_ingest"
)

type BucketClient struct {
	*Client
	bucketName string
}

type Client struct {
	client *resty.Client
}

func NewClient() *Client {
	return &Client{
		client: resty.New(),
	}
}

func (self *Client) InsertDocument(doc *Document) error {

	resp, err := self.client.R().EnableTrace().SetBody(doc).Post(ENDPOINT_INSERT)
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return errors.New("the request was unsuccessful")
	}

	return nil
}

func (self *Client) Bucket(bucketName string) *BucketClient {
	return &BucketClient{
		Client:     self,
		bucketName: bucketName,
	}
}
