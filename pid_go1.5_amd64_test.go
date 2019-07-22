// +build amd64 amd64p32
// +build gc,go1.5

package goid

import (
	"testing"
)

func TestPid(t *testing.T) {
	p1 := GetPid()
	p2 := getTID()
	if GetPid() != getTID() {
		t.Fatalf("The result of GetPid %d is not equal to procPin %d!", p1, p2)
	}
}
