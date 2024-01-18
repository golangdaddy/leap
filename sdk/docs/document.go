package docs

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/golangdaddy/leap/sdk/common"
	"golang.org/x/crypto/sha3"
)

var app *common.App

func init() {
	app = &common.App{}
	app.UseCBOR()
}

func New(salt, bucketName string) *Document {

	bucketID := "z" + hex.EncodeToString(app.SHA256(append([]byte(bucketName), []byte(salt)...))[:30])

	doc := &Document{
		Bucket: bucketID,
	}
	return doc
}

func EmptyDocument() *Document {
	doc := &Document{}
	return doc
}

type Document struct {
	Bucket   string
	Lat, Lng float64
	Place    []string
	Parent   string
	Class    string
	Data     interface{}
}

func (self *Document) ID() string {
	serial, err := app.MarshalCBOR(self)
	if err != nil {
		panic(err)
	}
	h := sha3.New224()
	n, _ := h.Write([]byte(serial))
	// encode the length of the document into its id
	return hex.EncodeToString(h.Sum(nil)[:16]) + "-" + strconv.Itoa(n)
}

func (self *Document) TimePrefix(t time.Time) string {

	return fmt.Sprintf("%d/%02d/%02d/%02d/%02d/%02d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
}
