// +build amd64 amd64p32
// +build gc,go1.5

package goid

import (
	"testing"
)

func TestPid(t *testing.T) {
	p1 := GetPid()
	p2 := getTID()
	if p1 != p2 {
		t.Fatalf("The result of GetPid %d procPin %d are not equal!", p1, p2)
	}
}
