package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

)

func Get(reqtype, name, url string) {
    fmt.Printf("Type: %s\nName: %s, URL: %s\n", reqtype, name, url)
  
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Error making GET request: %v\n", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error reading response body: %v\n", err)
        return
    }

    var prettyBody bytes.Buffer
    err = json.Indent(&prettyBody, body, "", "  ")
    if err != nil {
        fmt.Printf("Error indenting JSON: %v\n", err)
        return
    }

    fmt.Println(prettyBody.String())
}
