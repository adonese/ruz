package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

var fortune []string

func init() {
	var err error
	fortune, err = parseFortune("fortune.txt")
	if err != nil {
		log.Fatalf("Error in parsing fortune: %v", err)
	}
}

func parseFortune(fname string) ([]string, error) {
	f, err := ioutil.ReadFile(fname)
	if err != nil {
		return []string{}, err
	}

	s := string(f)
	data := strings.Split(s, "\n\n")
	return data, nil

}

func send(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	i := rand.Intn(209)
	f := fortune[i]
	f = strings.ReplaceAll(f, "\n", " ")
	data := map[string]string{"result": f}
	res, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error in marshaling: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/fortune", send)
	http.ListenAndServe(":7010", mux)
}
