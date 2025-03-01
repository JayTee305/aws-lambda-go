// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
package events

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events/test"
	"github.com/stretchr/testify/assert"
)

func TestKinesisEventMarshaling(t *testing.T) {

	// 1. read JSON from file
	inputJSON := test.ReadJSONFromFile(t, "./testdata/kinesis-event.json")

	// 2. de-serialize into Go object
	var inputEvent KinesisEvent
	if err := json.Unmarshal(inputJSON, &inputEvent); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	// 3. serialize to JSON
	outputJSON, err := json.Marshal(inputEvent)
	if err != nil {
		t.Errorf("could not marshal event. details: %v", err)
	}

	// 4. check result
	assert.JSONEq(t, string(inputJSON), string(outputJSON))
}

func TestKinesisMarshalingMalformedJson(t *testing.T) {
	test.TestMalformedJson(t, KinesisEvent{})
}
