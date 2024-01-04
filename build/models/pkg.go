package models

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/richardboase/npgpublic/sdk/cloudfunc"
	"github.com/richardboase/npgpublic/sdk/common"
)

func getTime() int64 {
	return time.Now().UTC().Unix()
}

type Generic struct {
	Meta Internals
}

func GetMetadata(app *common.App, id string) (*Internals, error) {

	dst := &Generic{}

	i := Internal(id)
	path := i.DocPath()

	println("GET DOCUMENT", path)

	doc, err := app.Firestore().Doc(path).Get(app.Context())
	if err != nil {
		return nil, err
	}
	return &dst.Meta, doc.DataTo(dst)
}

func AssertRange(w http.ResponseWriter, min, max float64, value interface{}) bool {

	var val float64
	switch v := value.(type) {
	case int:
		val = float64(v)
	case float64:
		val = v
	case string:
		val = float64(len(v))
	default:
		log.Println("ignoring range assertion for unknown type")
	}

	err := fmt.Errorf("value %v exceeded value of range min: %v max: %v ", value, min, max)
	if val < min {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return false
	}
	if val > max {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return false
	}
	return true
}

func AssertSTRING(w http.ResponseWriter, m map[string]interface{}, key string) (string, bool) {
	s, ok := m[key].(string)
	if !ok {
		err := fmt.Errorf("'%s' is required for this request", key)
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return s, false
	}
	return s, true
}

func AssertSTRINGS(w http.ResponseWriter, m map[string]interface{}, key string) ([]string, bool) {
	a, ok := m[key].([]interface{})
	if !ok {
		err := fmt.Errorf("'%s' is required for this request", key)
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return nil, false
	}
	b := []string{}
	for _, v := range a {
		s, ok := v.(string)
		if !ok {
			err := fmt.Errorf("strings are required for this request: %s", key)
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return nil, false
		}
		b = append(b, s)
	}
	return b, true
}

func AssertFLOAT64(w http.ResponseWriter, m map[string]interface{}, key string) (float64, bool) {
	f, ok := m[key].(float64)
	if !ok {
		err := fmt.Errorf("'%s' is required for this request", key)
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return 0, false
	}
	return f, true
}

func AssertBOOL(w http.ResponseWriter, m map[string]interface{}, key string) (bool, bool) {
	v, ok := m[key].(bool)
	if !ok {
		err := fmt.Errorf("'%s' is required for this request", key)
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return false, false
	}
	return v, true
}

func AssertINT(w http.ResponseWriter, m map[string]interface{}, key string) (int, bool) {
	v, ok := m[key].(float64)
	if !ok {
		err := fmt.Errorf("'%s' is required for this request", key)
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return 0, false
	}
	return int(v), true
}
