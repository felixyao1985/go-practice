package v2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const APIVersion = "/api/v2"

type CayleyClinet struct {
	address string // cayley server address
}

type Resq struct {
	Err    string `json:"error,omitempty"`
	Result string `json:"result,omitempty"`
	Count  int    `json:"count,omitempty"`
}

type Quad struct {
	Subject   string `json:"subject"`
	Predicate string `json:"predicate"`
	Object    string `json:"object"`
	Label     string `json:"label"`
}

type Triads []Quad

func NewAPIClient(address string) CayleyClinet {
	return CayleyClinet{
		address: address + APIVersion,
	}
}

func (c *CayleyClinet) StringList(preds []string) string {
	var items string
	items = "["
	for i, _ := range preds {
		items = items + `"` + preds[i] + `",`
	}
	items = items + "]"

	return items
}

func (c *CayleyClinet) MakeQuad(q Triads) string {
	var items string
	for _, item := range q {
		if item.Label == "" {
			item.Label = "."
		}

		items += fmt.Sprintln(item.Subject, item.Predicate, item.Object, item.Label)
	}

	return items
}

func (c *CayleyClinet) Write(q Triads) error {
	address := c.address + "/write"

	triad := c.MakeQuad(q)

	var x bytes.Buffer
	x.Write([]byte(triad))
	resp, err := http.Post(address, "text/json", &x)
	if err != nil {
		return err
	}
	r, readerr := ioutil.ReadAll(resp.Body)
	ret := Resq{}

	defer resp.Body.Close()
	if readerr != nil {
		return readerr
	}

	err = json.Unmarshal(r, &ret)
	if err != nil {
		return err
	}

	if ret.Err != "" {
		return errors.New(ret.Err)
	}
	return nil

}

func (c *CayleyClinet) Delete(q Triads) error {
	address := c.address + "/delete"
	triad := c.makeQuad(q)

	var x bytes.Buffer
	x.Write([]byte(triad))
	resp, err := http.Post(address, "text/json", &x)
	if err != nil {
		return err
	}
	_, readerr := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if readerr != nil {
		return readerr
	}
	return nil

}

func (c *CayleyClinet) Gremlin(q string) ([]byte, error) {
	address := c.address + "/query?lang=gizmo"
	var x bytes.Buffer
	x.Write([]byte(q))
	resp, err := http.Post(address, "text/plain", &x)
	if err != nil {
		return nil, err
	}
	data, readerr := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return data, readerr

}
