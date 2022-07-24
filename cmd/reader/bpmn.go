package reader

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/hendriknielaender/go-bpmn/pkg/bpmn"
)

func LoadFromFile(filename string) (*bpmn.Definitions, error) {
	xmlData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return LoadFromBytes(xmlData)
}

func LoadFromBytes(xmlData []byte) (*bpmn.Definitions, error) {
	var definitions bpmn.Definitions
	err := xml.Unmarshal(xmlData, &definitions)
	if err != nil {
		return nil, err
	}

	return &definitions, nil
}
