package gorge

import (
	"io"
	"io/ioutil"
	"net/http"
	"fmt"
	"log"
	"strings"
)

func gorge.post(url string, contentType string, body io.Reader) {

    resp, err := http.Post(url, contentType, body)
    if err != nil {
        panic(err)
	}
	
	defer resp.Body.Close() 
}

func gorge.put(method string, value string, body io.Reader) {
	resp, err := http.NewRequest("PUT", value, body)
    if err != nil {
        panic(err)
	}

	defer resp.Body.Close() 

}


func gorge.get(url string) {
	resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	g := func(w http.ResponseWriter, _ *http.Request) {
        fmt.Fprintf(w, string(body))
	}

	http.HandleFunc("/get", g)
}

func gorge.end() {
	log.Fatal(http.ListenAndServe(":6789", nil))
}

func gorge.say(value string) {

	t := func(w http.ResponseWriter, _ *http.Request) {
        fmt.Fprintf(w, value)
	}

	http.HandleFunc("/", t)
}

func gorge.do(value string) {

	r := strings.NewReader(value)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

}

const (
    MethodGet     = "GET"
    MethodHead    = "HEAD"
    MethodPost    = "POST"
    MethodPut     = "PUT"
    MethodPatch   = "PATCH" // RFC 5789
    MethodDelete  = "DELETE"
    MethodConnect = "CONNECT"
    MethodOptions = "OPTIONS"
    MethodTrace   = "TRACE"
)

func gorge.method(method string, value string) {
	http.NewRequest(method, value, nil)
}
