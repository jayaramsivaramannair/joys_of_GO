package utils


import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Parses body of the incoming request from the client
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		//Unmarshal is used to convert the json object back into a struct and save it in x
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return 
		}
	}
}