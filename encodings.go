package gorge

import (
	"crypto/x509"
	"fmt"
	"log"
	"io"
	"strings"
	"encoding/json"
	"encoding/pem"
	"encoding/xml"
	"net/http"
)

func decodepem(file string) {
	var pubPEMData = []byte(file)

	block, rest := pem.Decode(pubPEMData)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	pem := func(w http.ResponseWriter, _ *http.Request) {
        log.Println(rest, pub)
	}

	http.HandleFunc("/pem", pem)
}

func decodejson(file string) {
	const jsonStream = `file`

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		if err := dec.Decode(file); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		json := func(w http.ResponseWriter, _ *http.Request) {
        fmt.Fprintf(w, file)
	}

	http.HandleFunc("/json", json)
		
	}
}

func decodexml(file string) {
	data := file
	err := xml.Unmarshal([]byte(data), file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	xml := func(w http.ResponseWriter, _ *http.Request) {
        fmt.Fprintf(w, data)
	}

	http.HandleFunc(":6789/xml", xml)
}