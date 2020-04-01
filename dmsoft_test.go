// +build windows

// export GOARCH=386

package dmsoft

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		wantDm  *dmsoft
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDm, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDm, tt.wantDm) {
				t.Errorf("New() = %v, want %v", gotDm, tt.wantDm)
			}
		})
	}
}

func Test_dmsoft_Release(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.com.Release()
		})
	}
}

func TestSetDllPathA(t *testing.T) {
	type args struct {
		path string
		mode int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetDllPathA(tt.args.path, tt.args.mode); got != tt.want {
				t.Errorf("SetDllPathA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetDllPathW(t *testing.T) {
	type args struct {
		path string
		mode int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetDllPathW(tt.args.path, tt.args.mode); got != tt.want {
				t.Errorf("SetDllPathW() = %v, want %v", got, tt.want)
			}
		})
	}
}
