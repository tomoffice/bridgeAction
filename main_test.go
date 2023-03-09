package main

import (
	"reflect"
	"testing"
)

func TestNewComputer(t *testing.T) {
	hd := &HD{}
	mem := &DDR4{}
	graph := &BuiltIn{}
	want := &Computer{hd, mem, graph}
	got := NewComputer(hd, mem, graph)
	if !reflect.DeepEqual(got, want) {
		t.Error(want, got)
	}
}
func TestON(t *testing.T) {
	type args struct {
		hd    Drive
		mem   Memory
		graph Graphic
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "A",
			args: args{
				hd:    &HD{},
				mem:   &DDR4{},
				graph: &BuiltIn{},
			},
		}, {
			name: "B",
			args: args{
				hd:    &SSD{},
				mem:   &DDR5{},
				graph: &Discrete{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Computer{
				tt.args.hd,
				tt.args.mem,
				tt.args.graph}
			c.On()
		})
	}
}
func TestOff(t *testing.T) {
	type args struct {
		hd    Drive
		mem   Memory
		graph Graphic
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "A",
			args: args{
				hd:    &HD{},
				mem:   &DDR4{},
				graph: &BuiltIn{},
			},
		}, {
			name: "B",
			args: args{
				hd:    &SSD{},
				mem:   &DDR5{},
				graph: &Discrete{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Computer{
				tt.args.hd,
				tt.args.mem,
				tt.args.graph}
			c.Off()
		})
	}
}
