package main

import (
	"testing"
	"time"
)

func TestSave(t *testing.T) {
	stu := &student{
		Name: "stu01",
		Sex:  "man",
		Age:  10,
	}

	err := stu.Save()
	if err != nil {
		t.Fatalf("save student failed, err:%v", err)
	}
}

func TestLoad(t *testing.T) {
	stu := &student{
		Name: "stu01",
		Sex:  "man",
		Age:  10,
	}

	err := stu.Save()
	if err != nil {
		t.Fatalf("save student failed, err:%v", err)
	}

	stu2 := &student{}
	time.Sleep(10 * time.Second)
	err = stu2.Load()
	if err != nil {
		t.Fatalf("load student failed, err:%v", err)
	}

	if stu.Name != stu2.Name {
		t.Fatalf("load student failed, name not equal")
	}

	if stu.Sex != stu2.Sex {
		t.Fatalf("load student failed, sex not equal")
	}

	if stu.Age != stu2.Age {
		t.Fatalf("load student failed, age not equal")
	}

}

// cmd
// C:\go\src1\go_dev\day7\example\test1>go  test -v
// === RUN   TestAdd
//     calc_test.go:12: test add succ
// --- PASS: TestAdd (0.00s)
// === RUN   TestSub
//     calc_test.go:20: test sub succ
// --- PASS: TestSub (0.00s)
// === RUN   TestSave
// --- PASS: TestSave (0.00s)
// === RUN   TestLoad
// --- PASS: TestLoad (0.00s)
// PASS
// ok      _/C_/go/src1/go_dev/day7/example/test1  0.548s

// C:\go\src1\go_dev\day7\example\test1>go  test
// PASS
// ok      _/C_/go/src1/go_dev/day7/example/test1  0.457s
