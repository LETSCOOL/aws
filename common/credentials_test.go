package common

import (
	"testing"
)

func TestCredentialsCSV(t *testing.T) {
	_, err := NewCredentialsFromCSV(TEST_CREDENTIALS_FILE)
	if err != nil {
		t.Errorf("Create credentials fail (%s)\n", err)
		return
	}
}







