package common

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type testStruct struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Other string `json:"other"`
}

func TestCBOR(t *testing.T) {
	assert := assert.New(t)

	app := &App{}
	app.UseCBOR()

	me := &testStruct{
		Name: "John Doe",
		Age:  32323,
	}

	enc, err := app.MarshalCBOR(me)
	if err != nil {
		t.Fatal(err)
		return
	}

	s := hex.EncodeToString(enc)

	log.Println(string(enc))

	d, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
		return
	}

	you := &testStruct{}
	err = app.UnmarshalCBOR(d, you)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(me, you)
}

func TestCBORHelpers(t *testing.T) {
	assert := assert.New(t)

	app := &App{}
	app.UseCBOR()

	me := &testStruct{
		Name: "John Doe",
		Age:  32323,
	}

	enc, err := app.CompactSerial(me)
	if err != nil {
		t.Fatal(err)
		return
	}

	log.Println(enc)

	you := &testStruct{}
	if err = app.ExpandSerial([]byte(enc), you); err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(me, you)
}
