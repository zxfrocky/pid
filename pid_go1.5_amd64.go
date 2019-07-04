// Copyright 2019 Cholerae Hu.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.

// +build amd64 amd64p32
// +build gc,go1.5,!go1.13

package goid

import "unsafe"

//go:nosplit
func GetM() uintptr

//go:nosplit
func GetPid() int {
	return int((*m)(unsafe.Pointer(GetM())).p.id)
}

//go:linkname procPin runtime.procPin
//go:nosplit
func procPin() int

//go:linkname procUnpin runtime.procUnpin
//go:nosplit
func procUnpin()

func getTID() int {
	tid := procPin()
	procUnpin()
	return tid
}
