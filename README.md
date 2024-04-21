# ThingWave SenML module for Golang
This is a Golang module for the JSON-based SenML sensor data format. This module is compatible with IPSO Smart Objects, OMA LwM2M, and Eclipse Arrowhead usage.

## Introduction
SenML was developed as an XML and JSON based data format for sensor data. Later CBOR was added as well.

## Features
Current features include  Golang struct and data type that can capture all features in SenML.

## Installation
Install using
```
go get github.com/thingwave/senml-go
```
Then import with
```
import (
  "fmt"
  "time"
  "encoding/json"
  senml "github.com/thingwave/senml-go"
)
```


To, use, create a SenMLMessage as an array of SenMLRecords (example has OMA LwM2M records for temperature of 23.1 degrees and humidity of 42%)
```
smlMsg := make([]senml.SenMLRecord, 0)
var bver int16 = 5
head := senml.SenMLRecord{Bn: "urn:dev:mac:abcd1234:", Bver: &bver, Bt: float64(time.Now().UnixMilli())/1000.0, N: "3303/0/5700", V: 23.1}
hum := senml.SenMLRecord{N: "3304/0/5700", V: 42}
smlMsg = append(smlMsg, head)
smlMsg = append(smlMsg, hum)
jsonData, err := json.Marshal(smlMsg)
if err == nil {
  
  // do something with the message, like send over MQTT, save to file or database, etc.
  fmt.Printf("%s\n", string(jsonData))
}
```

## Future work
Future work includes:
 * JSON to Go validation
 * SenML data validation
 * CBOR support
 * XML support

## Contact
Contact support@thingwave.com for feature requests and bug reports.
