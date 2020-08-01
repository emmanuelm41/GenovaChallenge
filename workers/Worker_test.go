package workers

import (
	"testing"
)

func TestGetLocation(t *testing.T) {
	_, _, _, err1 := getLocation(5.2, 8.6, 12.5)

	if err1 != nil {
		t.Errorf("Error found trying to get position: %v \n", err1)
	}
}

func TestGetMessage(t *testing.T) {
	_, err2 := getMessage([]string{"This", "", "", "message"}, []string{"", "a", "new", "message"}, []string{"This", "", "", ""})

	if err2 != nil {
		t.Errorf("Error found trying to get message: %v \n", err2)
	}
}
