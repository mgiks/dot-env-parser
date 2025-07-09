package parser

import "testing"

func TestReadDotEnvFile(t *testing.T) {
	data, err := ReadDotEnvFile("/home/mgik/Projects/dotenvparser/parser/fixtures/.env")

	if err != nil {
		t.Error("should have run without errors")
	}

	if data == nil {
		t.Error("should have returned non-nil data")
	}
}
