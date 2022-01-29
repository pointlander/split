// Copyright 2022 The Split Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math/rand"

	"github.com/itsubaki/q"
)

func main() {
	rand.Seed(1)

	qsim := q.New()
	q0 := qsim.Zero()
	q1 := qsim.Zero()
	q2 := qsim.Zero()
	q3 := qsim.Zero()
	q4 := qsim.Zero()
	q5 := qsim.Zero()
	q6 := qsim.Zero()

	qsim.H(q2)
	qsim.H(q3)
	qsim.H(q4)
	qsim.H(q5)
	qsim.H(q6)

	qsim.ControlledZ([]q.Qubit{q2}, q3)
	qsim.ControlledZ([]q.Qubit{q3}, q4)
	qsim.ControlledZ([]q.Qubit{q4}, q5)
	qsim.ControlledZ([]q.Qubit{q5}, q6)

	qsim.H(q2)
	qsim.H(q4)
	qsim.H(q6)

	qsim.Swap(q2, q4)
	qsim.Swap(q4, q6)
	qsim.Swap(q1, q3)

	qsim.H(q0)
	qsim.H(q2)
	qsim.H(q4)

	qsim.CNOT(q1, q5)
	qsim.CNOT(q3, q5)
	qsim.CNOT(q1, q6)
	qsim.CNOT(q5, q6)
	qsim.CNOT(q6, q5)

	m := qsim.Measure(q0, q1, q2, q3, q4, q5, q6)
	fmt.Println(m)
	fmt.Println(m.Int())
	fmt.Println(m.BinaryString())
}
