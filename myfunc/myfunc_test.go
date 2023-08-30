package util

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	fmt.Println("max:", Max(1, 2))
}

func TestMin(t *testing.T) {
	fmt.Println("min:", Min(1, 2))
}
