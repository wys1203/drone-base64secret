package main

import (
	"reflect"
	"testing"
)

func TestPlugin_Exec(t *testing.T) {
	tests := []struct {
		name    string
		p       *Plugin
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Exec(); (err != nil) != tt.wantErr {
				t.Errorf("Plugin.Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDisplay(t *testing.T) {
	type args struct {
		name string
		x    interface{}
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Display(tt.args.name, tt.args.x)
		})
	}
}

func Test_display(t *testing.T) {
	type args struct {
		path string
		v    reflect.Value
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			display(tt.args.path, tt.args.v)
		})
	}
}

func Test_formatAtom(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatAtom(tt.args.v); got != tt.want {
				t.Errorf("formatAtom() = %v, want %v", got, tt.want)
			}
		})
	}
}
