package dmsoft

import "testing"

func Test_dmsoft_EnablePicCache(t *testing.T) {
	type args struct {
		enable int
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
			if got := tt.com.EnablePicCache(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.EnablePicCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetBasePath(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetBasePath(); got != tt.want {
				t.Errorf("dmsoft.GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetDmCount(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetDmCount(); got != tt.want {
				t.Errorf("dmsoft.GetDmCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetID(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetID(); got != tt.want {
				t.Errorf("dmsoft.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetLastError(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetLastError(); got != tt.want {
				t.Errorf("dmsoft.GetLastError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetPath(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetPath(); got != tt.want {
				t.Errorf("dmsoft.GetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_Reg(t *testing.T) {
	type args struct {
		regCode string
		verInfo string
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
			if got := tt.com.Reg(tt.args.regCode, tt.args.verInfo); got != tt.want {
				t.Errorf("dmsoft.Reg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_RegEx(t *testing.T) {
	type args struct {
		regCode string
		verInfo string
		ip      string
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
			if got := tt.com.RegEx(tt.args.regCode, tt.args.verInfo, tt.args.ip); got != tt.want {
				t.Errorf("dmsoft.RegEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_RegExNoMac(t *testing.T) {
	type args struct {
		regCode string
		verInfo string
		ip      string
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
			if got := tt.com.RegExNoMac(tt.args.regCode, tt.args.verInfo, tt.args.ip); got != tt.want {
				t.Errorf("dmsoft.RegExNoMac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_RegNoMac(t *testing.T) {
	type args struct {
		regCode string
		verInfo string
		ip      string
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
			if got := tt.com.RegNoMac(tt.args.regCode, tt.args.verInfo, tt.args.ip); got != tt.want {
				t.Errorf("dmsoft.RegNoMac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetDisplayInput(t *testing.T) {
	type args struct {
		mode string
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
			if got := tt.com.SetDisplayInput(tt.args.mode); got != tt.want {
				t.Errorf("dmsoft.SetDisplayInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetEnumWindowDelay(t *testing.T) {
	type args struct {
		delay int
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
			if got := tt.com.SetEnumWindowDelay(tt.args.delay); got != tt.want {
				t.Errorf("dmsoft.SetEnumWindowDelay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetPath(t *testing.T) {
	type args struct {
		path string
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
			if got := tt.com.SetPath(tt.args.path); got != tt.want {
				t.Errorf("dmsoft.SetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetShowErrorMsg(t *testing.T) {
	type args struct {
		show int
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
			if got := tt.com.SetShowErrorMsg(tt.args.show); got != tt.want {
				t.Errorf("dmsoft.SetShowErrorMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SpeedNormalGraphic(t *testing.T) {
	type args struct {
		enable int
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
			if got := tt.com.SpeedNormalGraphic(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.SpeedNormalGraphic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_Ver(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.Ver(); got != tt.want {
				t.Errorf("dmsoft.Ver() = %v, want %v", got, tt.want)
			}
		})
	}
}
