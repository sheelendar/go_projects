package main

import "testing"

// This is unit test case file we write unit test case for each function to cover every step.
func Test_updateOutputStream(t *testing.T) {
	oldstreamOutput := streamOutput
	streamOutput = make(map[string]Output)
	defer func() {
		streamOutput = oldstreamOutput
	}()

	testCases := []struct {
		name     string
		input    Input
		expected map[string]Output
	}{
		{
			name:     "successCase",
			input:    Input{UserID: 1, EventType: "post", Timestamp: 1672444800},
			expected: map[string]Output{"1_2023-01-01": Output{}},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			updateOutputStream(&test.input)
			assert.Equal(t, test.expected, streamOutput)
		})

	}
}

func Test_validateInputParse(t *testing.T) {

}

func Test_writeDataInOutputFile(t *testing.T) {

}

func Test_getUserKey(t *testing.T) {

}
