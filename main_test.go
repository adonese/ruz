package main

import (
	"testing"
)

func Test_parseFortune(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"testing parsing", args{"fortune.txt"}, []string{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFortune(tt.args.fname)
			if err != nil {
				t.Errorf("Error is: %v", err)
			}
			if len(got) != 210 {
				t.Errorf("parseFortune() small length: have: %v, want: %v", len(got), 209)
			}
		})
	}
}
