// 汇编

package dmsoft

import "testing"

func Test_dmsoft_AsmAdd(t *testing.T) {
	type args struct {
		asmIns string
	}
	tests := []struct {
		name string
		com  *dmsoft
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.AsmAdd(tt.args.asmIns); got != tt.want {
				t.Errorf("dmsoft.AsmAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_AsmCall(t *testing.T) {
	type args struct {
		hwnd int
		mode int
	}
	tests := []struct {
		name string
		com  *dmsoft
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.AsmCall(tt.args.hwnd, tt.args.mode); got != tt.want {
				t.Errorf("dmsoft.AsmCall() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_AsmCallEx(t *testing.T) {
	type args struct {
		hwnd     int
		mode     int
		baseAddr string
	}
	tests := []struct {
		name string
		com  *dmsoft
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.AsmCallEx(tt.args.hwnd, tt.args.mode, tt.args.baseAddr); got != tt.want {
				t.Errorf("dmsoft.AsmCallEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_AsmClear(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.AsmClear(); got != tt.want {
				t.Errorf("dmsoft.AsmClear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_AsmSetTimeout(t *testing.T) {
	type args struct {
		timeOut int
		param   int
	}
	tests := []struct {
		name string
		com  *dmsoft
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.AsmSetTimeout(tt.args.timeOut, tt.args.param); got != tt.want {
				t.Errorf("dmsoft.AsmSetTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_Assemble(t *testing.T) {
	type args struct {
		timeOut int64
		is64bit int
	}
	tests := []struct {
		name string
		com  *dmsoft
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.Assemble(tt.args.timeOut, tt.args.is64bit); got != tt.want {
				t.Errorf("dmsoft.Assemble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_DisAssemble(t *testing.T) {
	type args struct {
		asmCode  string
		baseAddr int64
		is64bit  int
	}
	tests := []struct {
		name string
		com  *dmsoft
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.DisAssemble(tt.args.asmCode, tt.args.baseAddr, tt.args.is64bit); got != tt.want {
				t.Errorf("dmsoft.DisAssemble() = %v, want %v", got, tt.want)
			}
		})
	}
}
