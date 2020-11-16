package models

import (
	"fmt"
	"testing"
)

func TestGetPointsByDistance(t *testing.T) {
	r := CreateRequest("1", "2", 42.021, 312.4)
	fmt.Println(r)
}
