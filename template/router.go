package template

var RouterTmpl = `package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"log"
	"net/url"
)

// example for init the database:
//
//  DB, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/employees?charset=utf8&parseTime=true")
//  if err != nil {
//  	panic("failed to connect database: " + err.Error())
//  }
//  defer db.Close()

var DB *gorm.DB

func ConfigRouter() http.Handler {
	router := httprouter.New()
    {{range .}}Config{{pluralize .}}Router(router)
    {{end}}
	
	return router
}

func readInt(r *http.Request, paramName string, defaultNumber int) (int, error) {
	queryString := r.URL.RawQuery
	if queryString == "" {
		return defaultNumber, nil
	}
	m, err := url.ParseQuery(queryString)
	if err != nil {
		log.Fatal(err)
		return defaultNumber, nil
	}
	if m[paramName] != nil {
		integerValues, err := strconv.Atoi(m[paramName][0])
		if err != nil {
			return defaultNumber, err
		}
		return integerValues, nil
	}
	return defaultNumber, nil
}


func writeJSON(w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, v)
}
`
