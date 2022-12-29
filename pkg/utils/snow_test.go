package utils

import (
	"fmt"
	"testing"
)

func TestGetSnowflakeID(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(GetSnowflakeID())
	}
}
