# ThingWave SenML module for Golang
This is a Golang module for the JSON-based SenML sensor data format.

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
  "encoding/json"
  senml "github.com/thingwave/senml-go"
)
```


To, use, create a SenMLMessage as an array of SenMLRecords
```
smlMsg := make([]senml.SenMLRecord, 0)
head := senml.SenMLRecord{Bn: "urn:dev:mac:abcd1234:", N: "33303/0/5700", "V": 23.1}
smlMsg = append(smlMsg, head)
jsonData, _ := json.Marshal(smlMsg)
```

## Future work
Future work includes:
 * JSON to Go validation
 * SenML data validation
 * CBOR support
 * XML support

## Contact
Contact support@thingwave.com for feature requests and bug reports.
