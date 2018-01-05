package av

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	if NewClient(testApiKey) == nil {
		t.Error("got nil client")
	}
}
