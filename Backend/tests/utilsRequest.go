package tests

import (
	"bytes"
	"encoding/json"
	"github.com/karimabedrabbo/eyo/api/apputils"
	"github.com/karimabedrabbo/eyo/api/managers"
	"log"
	"net/http"
	"net/http/httptest"
)


func IdToRoute(id int64) string {
	return "/" + apputils.IdToString(id)
}

func PerformRequest(router *managers.Router, method, path string, token string, body interface{}) *httptest.ResponseRecorder {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("could not convert body to json: %v", err)
	}
	req, err := http.NewRequest(method, path, bytes.NewReader(bodyBytes))
	if err != nil {
		log.Fatalf("could not process request: %v", err)
	}

	bearer := "Bearer " + token
	req.Header.Add("Authorization", bearer)

	//var bodybyte []byte
	//bodybyte, _ = httputil.DumpRequest(req, true)
	//var permissions os.FileMode = 0644 // or whatever you need
	//err = ioutil.WriteFile("file.txt", bodybyte, permissions)

	w := httptest.NewRecorder()
	router.Engine.ServeHTTP(w, req)
	return w
}
