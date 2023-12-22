package models

import (
	"fmt"
	"net/http"
	"time"

	"github.com/richardboase/npgpublic/sdk/cloudfunc"
)

func getTime() int64 {
	return time.Now().UTC().Unix()
}

func AssertKeyValueSTRING(w http.ResponseWriter, m map[string]interface{}, key string) (string, bool) {
	s, ok := m[key].(string)
	if !ok {
		err := fmt.Errorf("'%s' is required for this request", key)
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return s, false
	}
	return s, true
}

func AssertKeyValueSTRINGS(w http.ResponseWriter, m map[string]interface{}, key string) ([]string, bool) {
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

func AssertKeyValueFLOAT64(w http.ResponseWriter, m map[string]interface{}, key string) (float64, bool) {
	f, ok := m[key].(float64)
	if !ok {
		err := fmt.Errorf("'%s' is required for this request", key)
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return 0, false
	}
	return f, true
}

func AssertKeyValueBOOL(w http.ResponseWriter, m map[string]interface{}, key string) (bool, bool) {
	v, ok := m[key].(bool)
	if !ok {
		err := fmt.Errorf("'%s' is required for this request", key)
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return false, false
	}
	return v, true
}

func AssertKeyValueINT(w http.ResponseWriter, m map[string]interface{}, key string) (int, bool) {
	v, ok := m[key].(float64)
	if !ok {
		err := fmt.Errorf("'%s' is required for this request", key)
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return 0, false
	}
	return int(v), true
}
