package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/jtarchie/cl-search/pkg/load"
	"gopkg.in/yaml.v2"
)

func main() {
	var cities load.Cities

	contents, err := ioutil.ReadFile("../../.data/processed.json")
	if err != nil {
		log.Fatalf("cannot load YAML reader: %s", err)
	}

	err = yaml.UnmarshalStrict(contents, &cities)
	if err != nil {
		log.Fatalf("cannot unmarshal YAML reader: %s", err)
	}

	buffer := &bytes.Buffer{}

	fmt.Fprintln(buffer, "package load")
	fmt.Fprintf(
		buffer,
		"var allCities = %s",
		strings.Replace(
			fmt.Sprintf("%#v", cities),
			"load.",
			"",
			-1,
		),
	)

	err = ioutil.WriteFile("all.go", buffer.Bytes(), 0700)
	if err != nil {
		log.Fatalf("cannot write golang file: %s", err)
	}
}
