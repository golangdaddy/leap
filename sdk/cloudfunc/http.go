package cloudfunc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func QueryParam(r *http.Request, key string) (string, error) {

	v := r.URL.Query().Get(key)
	if len(v) == 0 {
		return "", fmt.Errorf("missing url query param: %s", key)
	}
	return v, nil
}

func HandleCORS(w http.ResponseWriter, r *http.Request, origin string) bool {
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Cache-Control", "no-store")
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "HEAD, POST, GET, OPTIONS, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.WriteHeader(http.StatusNoContent)
		return true
	}
	return false
}

func HttpError(w http.ResponseWriter, err error, status int) {
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), status)
	}
	w.Write([]byte(fmt.Sprintf("REQUEST FAILED: %d %v", status, err)))
}

func ParseJSON(r *http.Request, dst interface{}) error {
	b, err := io.ReadAll(r.Body)
	if r.Body != nil {
		r.Body.Close()
	}
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, dst); err != nil {
		return err
	}
	return nil
}

func ServeJSON(w http.ResponseWriter, src interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}

func ParamValue(r *http.Request, pos int) string {
	return strings.Split(r.URL.Path, "/")[pos]
}
