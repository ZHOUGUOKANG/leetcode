package Num_130_solve

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	in := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	expected := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(in)
	//println(fmt.Sprintf("%s,%s",expected,in))
	if !reflect.DeepEqual(expected, in) {
		t.Error(fmt.Sprintf("%s,%s", expected, in))
	}
}
func TestSolve1(t *testing.T) {
	in := [][]byte{{'O', 'X', 'X', 'O', 'X'}, {'X', 'O', 'O', 'X', 'O'}, {'X', 'O', 'X', 'O', 'X'}, {'O', 'X', 'O', 'O', 'O'}, {'X', 'X', 'O', 'X', 'O'}}
	expected := [][]byte{{'O', 'X', 'X', 'O', 'X'}, {'X', 'X', 'X', 'X', 'O'}, {'X', 'X', 'X', 'O', 'X'}, {'O', 'X', 'O', 'O', 'O'}, {'X', 'X', 'O', 'X', 'O'}}
	solve(in)
	//println(fmt.Sprintf("%v,%v",expected,in))
	if !reflect.DeepEqual(expected, in) {
		t.Error(fmt.Sprintf("%s,%s", expected, in))
	}
}

func TestSolve2(t *testing.T) {
	in := [][]byte{{'X', 'X', 'X'}, {'X', 'O', 'X'}, {'X', 'X', 'X'}}
	expected := [][]byte{{'X', 'X', 'X'}, {'X', 'X', 'X'}, {'X', 'X', 'X'}}
	solve(in)
	//println(fmt.Sprintf("%v,%v",expected,in))
	if !reflect.DeepEqual(expected, in) {
		t.Error(fmt.Sprintf("%s,%s", expected, in))
	}
}

func TestSolve3(t *testing.T) {
	in := [][]byte{{'X', 'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'O', 'X'}, {'X', 'X', 'O', 'O', 'X'}, {'X', 'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X', 'X'}}
	expected := [][]byte{{'X', 'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X', 'X'}}
	solve(in)
	//println(fmt.Sprintf("%v,%v",expected,in))
	if !reflect.DeepEqual(expected, in) {
		t.Error(fmt.Sprintf("%s,%s", expected, in))
	}
}
