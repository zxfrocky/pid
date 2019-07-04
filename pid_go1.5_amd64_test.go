// +build amd64 amd64p32
// +build gc,go1.5

package goid

import (
	"testing"
)

func TestPid(t *testing.T) {
	if GetPid() != getTID() {
		t.Fatalf("The result of GetPid is not equal to procPin!")
	}
}
