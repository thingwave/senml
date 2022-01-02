package senml

import (
	//"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type SenMLRecord struct {
	Bn   string      `json:"bn,omitempty"`
	Bt   interface{} `json:"bt,omitempty"`
	Bu   string      `json:"bu,omitempty"`
	Bv   interface{} `json:"bv,omitempty"`
	Bs   interface{} `json:"bs,omitempty"`
	Bver *int16      `json:"bver,omitempty"`

	N  string      `json:"n,omitempty"`
	U  string      `json:"u,omitempty"`
	V  interface{} `json:"v,omitempty"`
	Vs string      `json:"vs,omitempty"`
	Vb *bool       `json:"vb,omitempty"`
	Vd string      `json:"vd,omitempty"`
	S  interface{} `json:"s,omitempty"`
	T  interface{} `json:"t,omitempty"`
	Ut interface{} `json:"ut,omitempty"`
}

type SenMLMessage []SenMLRecord

type SenML interface {
	Validate(records SenMLMessage) error
}

func Validate(records SenMLMessage, ipsoMode bool) error {
	var bn string = ""

	if len(records) == 0 {
		return errors.New("Empty array")
	}

	for i, record := range records {
		if record.Bn != "" {
			bn = record.Bn
		} else {
			record.Bn = bn
		}

		if record.Bt == nil { // record.T something.. check RFC
			record.Bt = time.Now().Unix()
		}

		if record.N == "" {
			record.N = bn
		} else {
			record.N = bn + record.N
		}

		if strings.HasPrefix(record.N, "urn:dev:mac:") {
			fmt.Printf("bn(MAC): %s, n: %s\n", record.Bn, record.N)
			parts := strings.Split(record.Bn, ":")
			fmt.Printf("\tMAC: %s\n", parts[3])
		} else if strings.HasPrefix(record.N, "urn:dev:sn:") {

		} else if strings.HasPrefix(record.N, "urn:sys:name:") {
			fmt.Printf("bn(SYS): %s, n: %s\n", record.Bn, record.N)
		}

		if ipsoMode {
			parts := strings.Split(record.N, "/")
			major, _ := strconv.Atoi(parts[0])
			channel, _ := strconv.Atoi(parts[1])
			minor, _ := strconv.Atoi(parts[2])
			fmt.Printf("MAJOR: %d, MIDDLE: %d, MINOR: %d\n", major, channel, minor)

			fmt.Printf("%v, %+v\n", i, record)
			/*if record.Bn == nil {
			        return errors.New("Missing tags")
			}*/
		}
	}

	return nil
}
