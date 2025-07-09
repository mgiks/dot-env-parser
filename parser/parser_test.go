package parser

import "testing"

const pathOfEmptyEnvFile = "/home/mgik/Projects/dotenvparser/parser/fixtures/.env"

func TestReadDotEnvFile(t *testing.T) {
	data, err := ReadDotEnvFile(pathOfEmptyEnvFile)

	if err != nil {
		t.Error("should have run without errors")
	}

	if data == nil {
		t.Error("should have returned non-nil data")
	}
}
