// 其他

package dmsoft

import "testing"

func Test_dmsoft_DmGuard(t *testing.T) {
	type args struct {
		enable int
		lType  int
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
			if got := tt.com.DmGuard(tt.args.enable, tt.args.lType); got != tt.want {
				t.Errorf("dmsoft.DmGuard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_DmGuardParams(t *testing.T) {
	type args struct {
		cmd    string
		subcmd string
		param  string
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
			if got := tt.com.DmGuardParams(tt.args.cmd, tt.args.subcmd, tt.args.param); got != tt.want {
				t.Errorf("dmsoft.DmGuardParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_UnLoadDriver(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.UnLoadDriver(); got != tt.want {
				t.Errorf("dmsoft.UnLoadDriver() = %v, want %v", got, tt.want)
			}
		})
	}
}
