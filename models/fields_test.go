package models

import (
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

func TestFields(t *testing.T) {
	assert := assert.New(t)

	// string
	{
		f := Get("string", "250")
		pretty.Println(f)
		assert.Equal(float64(1), f.Range.Min)
		assert.Equal(float64(250), f.Range.Max)
	}
	{
		f := Get("string", "250", "750")
		pretty.Println(f)
		assert.Equal(float64(250), f.Range.Min)
		assert.Equal(float64(750), f.Range.Max)
	}

	// uint
	{
		f := Get("uint")
		pretty.Println(f)
		assert.NotNil(f.Range)
	}

	// int
	{
		f := Get("int")
		pretty.Println(f)
		assert.Nil(f.Range)
	}

	// float 64
	{
		f := Get("float64", "250", "750")
		pretty.Println(f)
		if assert.NotNil(f.Range) {
			assert.Equal(float64(250), f.Range.Min)
			assert.Equal(float64(750), f.Range.Max)
		}
	}
	{
		f := Get("float64", "1", "750")
		pretty.Println(f)
		if assert.NotNil(f.Range) {
			assert.Equal(float64(1), f.Range.Min)
			assert.Equal(float64(750), f.Range.Max)
		}
	}
	{
		f := Get("float64", "750")
		pretty.Println(f)
		if assert.NotNil(f.Range) {
			assert.Equal(float64(0), f.Range.Min)
			assert.Equal(float64(750), f.Range.Max)
		}
	}
	{
		f := Get("float64")
		pretty.Println(f)
		assert.Nil(f.Range)
	}

	{
		f := Get("address")
		pretty.Println(f)
		assert.Nil(f.Range)
		assert.Equal(5, len(f.Inputs))
		assert.True(f.Inputs[0].Required)
		assert.False(f.Inputs[1].Required)
	}

	{
		f := Get("name.person")
		pretty.Println(f)
		assert.Equal(3, len(f.Inputs))
		assert.True(f.Inputs[0].Required)
		assert.False(f.Inputs[1].Required)
		assert.True(f.Inputs[2].Required)
	}

	{
		f := Get("social.account", "twitter", "facebook", "google")
		pretty.Println(f)
		assert.NotNil(f.Inputs)
		assert.Equal(2, len(f.Inputs))
		assert.Equal(3, len(f.Inputs[0].InputOptions))
		assert.True(f.Inputs[0].Required)
		assert.True(f.Inputs[1].Required)
		assert.Equal("google", f.Inputs[0].InputOptions[2])
	}

	{
		f := Get("name.company")
		pretty.Println(f)
		if assert.NotNil(f.Inputs[0].Range) {
			assert.Equal(float64(1), f.Inputs[0].Range.Min)
			assert.Equal(float64(160), f.Inputs[0].Range.Max)
		}
	}
}
