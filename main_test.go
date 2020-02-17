package main

import (
	"reflect"
	"testing"
)

func TestDial_convertDigits(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name     string
		args     args
		wantList []string
	}{
		{name: "simple", args: args{[]int{2, 3}}, wantList: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{name: "zero simple", args: args{[]int{0, 3}}, wantList: []string{"d", "e", "f"}},
		{name: "another simple", args: args{[]int{6, 3}}, wantList: []string{"md", "me", "mf", "nd", "ne", "nf", "od", "oe", "of"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Dial{
				m: map[int]string{2: "abc", 3: "def", 4: "ghi", 5: "jkl", 6: "mno", 7: "pqrs", 8: "tuv", 9: "wxyz"},
			}
			if gotList := d.convertDigits(tt.args.arr...); !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("Dial.convertDigits() = %v, want %v", gotList, tt.wantList)
			}
		})
	}
}
