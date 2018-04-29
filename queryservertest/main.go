package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
)

var (
	stdin              = bufio.NewReader(os.Stdin)
	stdout             = bufio.NewWriter(os.Stdout)
	stderr             = bufio.NewWriter(os.Stderr)
	regexDocumentID    = regexp.MustCompile("^[A-Za-z0-9]{16}$")
	base62charSet      = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	timestampTolerance = int64(1000 * 60 * 5)
)

func main() {
	initCustomValidation()
	for {
		data, err := stdin.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log("Error reading from stdin: %s", err.Error())
			continue
		}
		if !gjson.Valid(string(data)) {
			log("Input is not valid JSON: %s", string(data))
			continue
		}
		stdout.Write(handleData(data))
		stdout.WriteString("\n")
		stdout.Flush()
	}
}

func handleData(data []byte) []byte {
	result := gjson.GetBytes(data, "0")
	if !result.Exists() {
		response, _ := json.Marshal([]string{"error", "value_error", "Could not find command"})
		stdout.Write(response)
	}
	log("Received valid JSON document: %s", string(data))
	command := result.String()
	switch command {
	case "ddoc":
		docName := gjson.GetBytes(data, "1")
		if docName.String() == "new" {
			tolerance := gjson.GetBytes(data, "3.options.timestamp_tolerance")
			if tolerance.Type == gjson.Number && tolerance.Num > 0 {
				log("Setting timestamp tolerance to %d", tolerance.Int())
				timestampTolerance = tolerance.Int()
			}
			return statusOk()
		}
		subcommand := gjson.GetBytes(data, "2.0")
		if subcommand.String() == "validate_doc_update" {
			parts := gjson.GetBytes(data, "3").Array()
			if len(parts) == 4 {
				return validate(parts[0], parts[1], parts[2], parts[3])
			}
			return statusErr("value_error", "No data to validate")
		}
		return statusErr("value_error", "Unknown subcommand")
	default:
		return statusOk()
	}
}

func log(message string, args ...interface{}) {
	stderr.WriteString(fmt.Sprintf(message, args...))
	stderr.WriteByte('\n')
	stderr.Flush()
}

func statusErr(errType string, errMessage string) []byte {
	result, _ := json.Marshal([]string{"error", errType, errMessage})
	return result
}

func statusForbidden(reason string) []byte {
	data, _ := json.Marshal(map[string]string{"forbidden": reason})
	return data
}

func statusOk() []byte {
	return []byte("true")
}

func statusValid() []byte {
	return []byte("1")
}

func timeFromID(id string) int64 {
	var result int64
	for i := 0; i < 8; i++ {
		result += int64(strings.Index(base62charSet, string(id[i]))) * int64(math.Pow(62, float64(7-i)))
	}
	return result
}

func checkTime(timestamp int64) bool {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	return timestamp >= now-timestampTolerance && timestamp <= now+timestampTolerance
}

func couchRequest(path string) gjson.Result {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(fmt.Sprintf("http://127.0.0.1:5984%s", path))
	req.Header.Add("Authorization", "Basic YWRtaW46SWd1bUNhdDU=")
	res := fasthttp.AcquireResponse()
	err := fasthttp.DoTimeout(req, res, time.Second*2)
	if err != nil {
		log("Error getting player: %s", err.Error())
		return gjson.Parse("[]")
	}
	return gjson.ParseBytes(res.Body())
}

func getObjectByID(objectID string) gjson.Result {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(fmt.Sprintf("https://adh.rapidnet.de:6984/all_day_hero/%s", objectID))
	req.Header.Add("Authorization", "Basic YWRtaW46SWd1bUNhdDU=")
	res := fasthttp.AcquireResponse()
	err := fasthttp.DoTimeout(req, res, time.Second*2)
	if err != nil {
		log("Error getting object: %s", err.Error())
		return gjson.Parse("[]")
	}
	object := gjson.ParseBytes(res.Body())
	if object.Get("error").String() == "not_found" {
		return gjson.Parse("[]")
	}
	return object
}

func validate(newDoc gjson.Result, oldDoc gjson.Result, userCtx gjson.Result, security gjson.Result) []byte {
	if !regexDocumentID.MatchString(newDoc.Get("_id").String()) {
		return statusForbidden("Document id has wrong format")
	}

	if !oldDoc.Get("_id").Exists() && !checkTime(timeFromID(newDoc.Get("_id").String())) {
		return statusForbidden("Document timestamp is not within tolerance")
	}
	newDocType := newDoc.Get("type")
	if newDocType.Type != gjson.String {
		return statusForbidden("Document must have a type")
	}
	oldDocType := oldDoc.Get("type")
	if oldDocType.Exists() && newDocType.String() != oldDocType.String() {
		return statusForbidden("Document type cannot be changed")
	}

	update := oldDoc.Get("_id").Exists()

	player := couchRequest(fmt.Sprintf("/_users/org.couchdb.user:%s", userCtx.Get("name").String()))

	if !regexDocumentID.MatchString(player.Get("player_id").String()) {
		return statusForbidden("Authorized user does not have a PLAYERID")
	}

	switch newDocType.String() {
	case "player":
		return validatePlayer(newDoc, update)
	case "message":
		return validateMessage(newDoc, update)
	case "challenge":
		return validateChallenge(newDoc, update)
	case "friendship":
		return validateFriendship(newDoc, update)
	case "join":
		return validateJoin(newDoc, update)
	case "post":
		return validatePost(newDoc, update)
	case "comment":
		return validateComment(newDoc, update)
	case "rating":
	default:
		return statusForbidden("Unknown document type")
	}
	return statusValid()
}
