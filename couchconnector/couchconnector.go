package couchconnector

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"errors"

	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
)

var couchDB CouchDB

func InitDatabase(dbserver string, dbname string, dbusername string, dbpassword string) error {
	couchDB = CouchDB{dbserver, dbname, dbusername, dbpassword}
	path := fmt.Sprintf("%s/%s", couchDB.DatabaseServer, couchDB.DatabaseName)
	if _, err := getData(path); err != nil {
		return err
	}
	return nil
}

//  GetDocumentByID return a document with a specific id
func GetDocumentByID(docID string) (gjson.Result, error) {
	path := fmt.Sprintf("%s/%s/%s", couchDB.DatabaseServer, couchDB.DatabaseName, docID)
	return getData(path)

}

func CheckIfDocumentExistByID(docID string) bool {
	path := fmt.Sprintf("%s/%s/%s", couchDB.DatabaseServer, couchDB.DatabaseName, docID)
	res, _ := getData(path)
	if res.Get("error").Exists() {
		return false
	}
	return true
}

func CheckIfDocumentExistWithSingleKey(viewObject string, viewname string, key string) (bool, error) {
	path := fmt.Sprintf("%s/%s/_design/%s/_view/%s?key=\"%s\"&include_docs=true", couchDB.DatabaseServer, couchDB.DatabaseName, viewObject, viewname, key)
	res, _ := getData(path)
	if res.Get("error").Exists() {
		return false, errors.New("Wrong format request")
	}
	arr := res.Get("rows").Array()
	if len(arr) > 0 {
		return true, nil
	}
	return false, nil
}

//  GetViewBySingleKey return a view using a single key
func GetViewBySingleKey(viewObject string, viewname string, key string, reduce bool) (gjson.Result, error) {
	path := ""
	if !reduce {
		path = fmt.Sprintf("%s/%s/_design/%s/_view/%s?key=\"%s\"&include_docs=true", couchDB.DatabaseServer, couchDB.DatabaseName, viewObject, viewname, key)
	} else {
		path = fmt.Sprintf("%s/%s/_design/%s/_view/%s?key=\"%s\"", couchDB.DatabaseServer, couchDB.DatabaseName, viewObject, viewname, key)
	}
	//path := fmt.Sprintf("%s/%s/_design/%s/_view/%s?key=\"%s\"&include_docs=true", couchDB.DatabaseServer, couchDB.DatabaseName, viewObject, viewname, key)
	return getData(path)
}

//  GetViewByMultipleKeys return a view using a multiple keys
func GetViewByMultipleKeys(viewObject string, viewname string, keys []string, reduce bool) (gjson.Result, error) {
	keyString := "["
	for _, key := range keys {
		keyString += "\"" + key + "\","
	}
	keyString = keyString[:len(keyString)-1]
	keyString += "]"
	path := ""
	if !reduce {
		path = fmt.Sprintf("%s/%s/_design/%s/_view/%s?key=%s&include_docs=true", couchDB.DatabaseServer, couchDB.DatabaseName, viewObject, viewname, keyString)
	} else {
		path = fmt.Sprintf("%s/%s/_design/%s/_view/%s?key=%s", couchDB.DatabaseServer, couchDB.DatabaseName, viewObject, viewname, keyString)
	}
	return getData(path)
}

func PostDocument(document interface{}) string {
	path := fmt.Sprintf("%s/%s", couchDB.DatabaseServer, couchDB.DatabaseName)
	id := postData(path, document)
	return id
}

func PutDocument(document interface{}) string {
	path := fmt.Sprintf("%s/%s", couchDB.DatabaseServer, couchDB.DatabaseName)
	rev := putData(path, document)
	return rev
}

func getData(path string) (gjson.Result, error) {
	req := fasthttp.AcquireRequest()
	autheticationData := fmt.Sprintf("%s:%s", couchDB.DatabaseUsername, couchDB.DatabasePassword)
	authenticationHash := fmt.Sprintf("Basic %s", b64.StdEncoding.EncodeToString([]byte(autheticationData)))
	req.SetRequestURI(path)
	req.Header.Add("Authorization", authenticationHash)
	res := fasthttp.AcquireResponse()
	err := fasthttp.DoTimeout(req, res, time.Second*2)
	if err != nil {
		return gjson.Parse("[]"), err
	}
	return gjson.ParseBytes(res.Body()), nil
}

func postData(path string, object interface{}) string {
	client := &http.Client{}
	json, _ := json.Marshal(object)
	request, _ := http.NewRequest("POST", path, strings.NewReader(string(json)))
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(couchDB.DatabaseUsername, couchDB.DatabasePassword)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		obj := gjson.ParseBytes(body)
		return obj.Get("id").String()
	}
	return ""
}

func putData(path string, object interface{}) string {
	client := &http.Client{}
	json, _ := json.Marshal(object)
	request, _ := http.NewRequest("PUT", path, strings.NewReader(string(json)))
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(couchDB.DatabaseUsername, couchDB.DatabasePassword)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		obj := gjson.ParseBytes(body)
		return obj.Get("rev").String()
	}
	return ""
}
