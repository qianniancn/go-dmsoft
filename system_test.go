// 系统

package dmsoft

import "testing"

func Test_dmsoft_Beep(t *testing.T) {
	type args struct {
		f        int
		duration int
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
			if got := tt.com.Beep(tt.args.f, tt.args.duration); got != tt.want {
				t.Errorf("dmsoft.Beep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_CheckFontSmooth(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.CheckFontSmooth(); got != tt.want {
				t.Errorf("dmsoft.CheckFontSmooth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_CheckUAC(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.CheckUAC(); got != tt.want {
				t.Errorf("dmsoft.CheckUAC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_Delay(t *testing.T) {
	type args struct {
		mis int
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
			if got := tt.com.Delay(tt.args.mis); got != tt.want {
				t.Errorf("dmsoft.Delay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_Delays(t *testing.T) {
	type args struct {
		misMin int
		misMax int
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
			if got := tt.com.Delays(tt.args.misMin, tt.args.misMax); got != tt.want {
				t.Errorf("dmsoft.Delays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_DisableCloseDisplayAndSleep(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.DisableCloseDisplayAndSleep(); got != tt.want {
				t.Errorf("dmsoft.DisableCloseDisplayAndSleep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_DisableFontSmooth(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.DisableFontSmooth(); got != tt.want {
				t.Errorf("dmsoft.DisableFontSmooth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_DisablePowerSave(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.DisablePowerSave(); got != tt.want {
				t.Errorf("dmsoft.DisablePowerSave() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_DisableScreenSave(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.DisableScreenSave(); got != tt.want {
				t.Errorf("dmsoft.DisableScreenSave() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableFontSmooth(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.EnableFontSmooth(); got != tt.want {
				t.Errorf("dmsoft.EnableFontSmooth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_ExitOs(t *testing.T) {
	type args struct {
		types int
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
			if got := tt.com.ExitOs(tt.args.types); got != tt.want {
				t.Errorf("dmsoft.ExitOs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetClipboard(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetClipboard(); got != tt.want {
				t.Errorf("dmsoft.GetClipboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetCPUType(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetCPUType(); got != tt.want {
				t.Errorf("dmsoft.GetCPUType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetCPUUsage(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetCPUUsage(); got != tt.want {
				t.Errorf("dmsoft.GetCPUUsage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetDir(t *testing.T) {
	type args struct {
		types int
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
			if got := tt.com.GetDir(tt.args.types); got != tt.want {
				t.Errorf("dmsoft.GetDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetDiskModel(t *testing.T) {
	type args struct {
		index int
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
			if got := tt.com.GetDiskModel(tt.args.index); got != tt.want {
				t.Errorf("dmsoft.GetDiskModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetDiskReversion(t *testing.T) {
	type args struct {
		index int
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
			if got := tt.com.GetDiskReversion(tt.args.index); got != tt.want {
				t.Errorf("dmsoft.GetDiskReversion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetDiskSerial(t *testing.T) {
	type args struct {
		index int
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
			if got := tt.com.GetDiskSerial(tt.args.index); got != tt.want {
				t.Errorf("dmsoft.GetDiskSerial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetDisplayInfo(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetDisplayInfo(); got != tt.want {
				t.Errorf("dmsoft.GetDisplayInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetDPI(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetDPI(); got != tt.want {
				t.Errorf("dmsoft.GetDPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetLocale(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetLocale(); got != tt.want {
				t.Errorf("dmsoft.GetLocale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetMachineCode(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetMachineCode(); got != tt.want {
				t.Errorf("dmsoft.GetMachineCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetMachineCodeNoMac(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetMachineCodeNoMac(); got != tt.want {
				t.Errorf("dmsoft.GetMachineCodeNoMac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetMemoryUsage(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetMemoryUsage(); got != tt.want {
				t.Errorf("dmsoft.GetMemoryUsage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetNetTime(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetNetTime(); got != tt.want {
				t.Errorf("dmsoft.GetNetTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetNetTimeByIP(t *testing.T) {
	type args struct {
		ip string
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
			if got := tt.com.GetNetTimeByIP(tt.args.ip); got != tt.want {
				t.Errorf("dmsoft.GetNetTimeByIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetOsBuildNumber(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetOsBuildNumber(); got != tt.want {
				t.Errorf("dmsoft.GetOsBuildNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetOsType(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetOsType(); got != tt.want {
				t.Errorf("dmsoft.GetOsType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetScreenDepth(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetScreenDepth(); got != tt.want {
				t.Errorf("dmsoft.GetScreenDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetScreenHeight(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetScreenHeight(); got != tt.want {
				t.Errorf("dmsoft.GetScreenHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetScreenWidth(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetScreenWidth(); got != tt.want {
				t.Errorf("dmsoft.GetScreenWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetTime(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetTime(); got != tt.want {
				t.Errorf("dmsoft.GetTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_Is64Bit(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.Is64Bit(); got != tt.want {
				t.Errorf("dmsoft.Is64Bit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_IsSurrpotVt(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.IsSurrpotVt(); got != tt.want {
				t.Errorf("dmsoft.IsSurrpotVt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_Play(t *testing.T) {
	type args struct {
		mediaFile string
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
			if got := tt.com.Play(tt.args.mediaFile); got != tt.want {
				t.Errorf("dmsoft.Play() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_RunApp(t *testing.T) {
	type args struct {
		appPath string
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
			if got := tt.com.RunApp(tt.args.appPath, tt.args.mode); got != tt.want {
				t.Errorf("dmsoft.RunApp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetClipboard(t *testing.T) {
	type args struct {
		value string
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
			if got := tt.com.SetClipboard(tt.args.value); got != tt.want {
				t.Errorf("dmsoft.SetClipboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetDisplayAcceler(t *testing.T) {
	type args struct {
		level int
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
			if got := tt.com.SetDisplayAcceler(tt.args.level); got != tt.want {
				t.Errorf("dmsoft.SetDisplayAcceler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetLocale(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.SetLocale(); got != tt.want {
				t.Errorf("dmsoft.SetLocale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetScreen(t *testing.T) {
	type args struct {
		width  int
		height int
		depth  int
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
			if got := tt.com.SetScreen(tt.args.width, tt.args.height, tt.args.depth); got != tt.want {
				t.Errorf("dmsoft.SetScreen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetUAC(t *testing.T) {
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
			if got := tt.com.SetUAC(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.SetUAC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_ShowTaskBarIcon(t *testing.T) {
	type args struct {
		hwnd   int
		isShow int
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
			if got := tt.com.ShowTaskBarIcon(tt.args.hwnd, tt.args.isShow); got != tt.want {
				t.Errorf("dmsoft.ShowTaskBarIcon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_Stop(t *testing.T) {
	type args struct {
		id int
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
			if got := tt.com.Stop(tt.args.id); got != tt.want {
				t.Errorf("dmsoft.Stop() = %v, want %v", got, tt.want)
			}
		})
	}
}
