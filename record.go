package senml

import (
	//"encoding/json"
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
