// 后台

package dmsoft

import "testing"

func Test_dmsoft_BindWindow(t *testing.T) {
	type args struct {
		hwnd    int
		display string
		mouse   string
		keypad  string
		mode    int
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
			if got := tt.com.BindWindow(tt.args.hwnd, tt.args.display, tt.args.mouse, tt.args.keypad, tt.args.mode); got != tt.want {
				t.Errorf("dmsoft.BindWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_BindWindowEx(t *testing.T) {
	type args struct {
		hwnd    int
		display string
		mouse   string
		keypad  string
		public  string
		mode    int
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
			if got := tt.com.BindWindowEx(tt.args.hwnd, tt.args.display, tt.args.mouse, tt.args.keypad, tt.args.public, tt.args.mode); got != tt.want {
				t.Errorf("dmsoft.BindWindowEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_DownCPU(t *testing.T) {
	type args struct {
		rate int
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
			if got := tt.com.DownCPU(tt.args.rate); got != tt.want {
				t.Errorf("dmsoft.DownCPU() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableBind(t *testing.T) {
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
			if got := tt.com.EnableBind(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.EnableBind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableFakeActive(t *testing.T) {
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
			if got := tt.com.EnableFakeActive(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.EnableFakeActive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableIme(t *testing.T) {
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
			if got := tt.com.EnableIme(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.EnableIme() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableKeypadMsg(t *testing.T) {
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
			if got := tt.com.EnableKeypadMsg(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.EnableKeypadMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableKeypadPatch(t *testing.T) {
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
			if got := tt.com.EnableKeypadPatch(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.EnableKeypadPatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableKeypadSync(t *testing.T) {
	type args struct {
		enable  int
		timeOut int
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
			if got := tt.com.EnableKeypadSync(tt.args.enable, tt.args.timeOut); got != tt.want {
				t.Errorf("dmsoft.EnableKeypadSync() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableMouseMsg(t *testing.T) {
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
			if got := tt.com.EnableMouseMsg(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.EnableMouseMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableMouseSync(t *testing.T) {
	type args struct {
		enable  int
		timeOut int
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
			if got := tt.com.EnableMouseSync(tt.args.enable, tt.args.timeOut); got != tt.want {
				t.Errorf("dmsoft.EnableMouseSync() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableRealKeypad(t *testing.T) {
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
			if got := tt.com.EnableRealKeypad(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.EnableRealKeypad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableRealMouse(t *testing.T) {
	type args struct {
		enable     int
		mousedelay int
		mousestep  int
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
			if got := tt.com.EnableRealMouse(tt.args.enable, tt.args.mousedelay, tt.args.mousestep); got != tt.want {
				t.Errorf("dmsoft.EnableRealMouse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableSpeedDx(t *testing.T) {
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
			if got := tt.com.EnableSpeedDx(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.EnableSpeedDx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_ForceUnBindWindow(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.ForceUnBindWindow(); got != tt.want {
				t.Errorf("dmsoft.ForceUnBindWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetBindWindow(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetBindWindow(); got != tt.want {
				t.Errorf("dmsoft.GetBindWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetFps(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetFps(); got != tt.want {
				t.Errorf("dmsoft.GetFps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_HackSpeed(t *testing.T) {
	type args struct {
		rate int
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
			if got := tt.com.HackSpeed(tt.args.rate); got != tt.want {
				t.Errorf("dmsoft.HackSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_IsBind(t *testing.T) {
	type args struct {
		hwnd int
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
			if got := tt.com.IsBind(tt.args.hwnd); got != tt.want {
				t.Errorf("dmsoft.IsBind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_LockDisplay(t *testing.T) {
	type args struct {
		lock int
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
			if got := tt.com.LockDisplay(tt.args.lock); got != tt.want {
				t.Errorf("dmsoft.LockDisplay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_LockInput(t *testing.T) {
	type args struct {
		lock int
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
			if got := tt.com.LockInput(tt.args.lock); got != tt.want {
				t.Errorf("dmsoft.LockInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_LockMouseRect(t *testing.T) {
	type args struct {
		x1 int
		y1 int
		x2 int
		y2 int
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
			if got := tt.com.LockMouseRect(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("dmsoft.LockMouseRect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetAero(t *testing.T) {
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
			if got := tt.com.SetAero(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.SetAero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetDisplayDelay(t *testing.T) {
	type args struct {
		time int
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
			if got := tt.com.SetDisplayDelay(tt.args.time); got != tt.want {
				t.Errorf("dmsoft.SetDisplayDelay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetDisplayRefreshDelay(t *testing.T) {
	type args struct {
		time int
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
			if got := tt.com.SetDisplayRefreshDelay(tt.args.time); got != tt.want {
				t.Errorf("dmsoft.SetDisplayRefreshDelay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SwitchBindWindow(t *testing.T) {
	type args struct {
		hwnd int
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
			if got := tt.com.SwitchBindWindow(tt.args.hwnd); got != tt.want {
				t.Errorf("dmsoft.SwitchBindWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_UnBindWindow(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.UnBindWindow(); got != tt.want {
				t.Errorf("dmsoft.UnBindWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
