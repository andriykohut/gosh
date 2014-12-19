package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"
)

const ShortenerApiUrl = "https://www.googleapis.com/urlshortener/v1/url"

func Shorten(longUrl string) string {
	jsonBytes, _ := json.Marshal(map[string]string{"longUrl": longUrl})
	req, _ := http.NewRequest("POST", ShortenerApiUrl, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	response := make(map[string]string)
	json.Unmarshal(body, &response)
	return response["id"]
}

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	for _, arg := range os.Args[1:] {
		fmt.Fprintf(w, "%s\t => %s\n", arg, Shorten(arg))
	}
	w.Flush()
}
