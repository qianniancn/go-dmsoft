// 窗口

package dmsoft

import "testing"

func Test_dmsoft_ClientToScreen(t *testing.T) {
	type args struct {
		hwnd int
		x    *int
		y    *int
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
			if got := tt.com.ClientToScreen(tt.args.hwnd, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("dmsoft.ClientToScreen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnumProcess(t *testing.T) {
	type args struct {
		name string
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
			if got := tt.com.EnumProcess(tt.args.name); got != tt.want {
				t.Errorf("dmsoft.EnumProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnumWindow(t *testing.T) {
	type args struct {
		parent    int
		title     string
		className string
		filter    int
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
			if got := tt.com.EnumWindow(tt.args.parent, tt.args.title, tt.args.className, tt.args.filter); got != tt.want {
				t.Errorf("dmsoft.EnumWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnumWindowByProcess(t *testing.T) {
	type args struct {
		processName string
		title       string
		className   string
		filter      int
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
			if got := tt.com.EnumWindowByProcess(tt.args.processName, tt.args.title, tt.args.className, tt.args.filter); got != tt.want {
				t.Errorf("dmsoft.EnumWindowByProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnumWindowByProcessId(t *testing.T) {
	type args struct {
		pid       int
		title     string
		className string
		filter    int
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
			if got := tt.com.EnumWindowByProcessId(tt.args.pid, tt.args.title, tt.args.className, tt.args.filter); got != tt.want {
				t.Errorf("dmsoft.EnumWindowByProcessId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnumWindowSuper(t *testing.T) {
	type args struct {
		spec1 string
		flag1 int
		type1 int
		spec2 string
		flag2 int
		type2 int
		sort  int
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
			if got := tt.com.EnumWindowSuper(tt.args.spec1, tt.args.flag1, tt.args.type1, tt.args.spec2, tt.args.flag2, tt.args.type2, tt.args.sort); got != tt.want {
				t.Errorf("dmsoft.EnumWindowSuper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindWindow(t *testing.T) {
	type args struct {
		class string
		title string
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
			if got := tt.com.FindWindow(tt.args.class, tt.args.title); got != tt.want {
				t.Errorf("dmsoft.FindWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindWindowByProcess(t *testing.T) {
	type args struct {
		processName string
		class       string
		title       string
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
			if got := tt.com.FindWindowByProcess(tt.args.processName, tt.args.class, tt.args.title); got != tt.want {
				t.Errorf("dmsoft.FindWindowByProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindWindowByProcessId(t *testing.T) {
	type args struct {
		processId int
		class     string
		title     string
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
			if got := tt.com.FindWindowByProcessId(tt.args.processId, tt.args.class, tt.args.title); got != tt.want {
				t.Errorf("dmsoft.FindWindowByProcessId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindWindowEx(t *testing.T) {
	type args struct {
		parent int
		class  string
		title  string
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
			if got := tt.com.FindWindowEx(tt.args.parent, tt.args.class, tt.args.title); got != tt.want {
				t.Errorf("dmsoft.FindWindowEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindWindowSuper(t *testing.T) {
	type args struct {
		spec1 string
		flag1 int
		type1 int
		spec2 string
		flag2 int
		type2 int
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
			if got := tt.com.FindWindowSuper(tt.args.spec1, tt.args.flag1, tt.args.type1, tt.args.spec2, tt.args.flag2, tt.args.type2); got != tt.want {
				t.Errorf("dmsoft.FindWindowSuper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetClientRect(t *testing.T) {
	type args struct {
		hwnd int
		x1   *int
		y1   *int
		x2   *int
		y2   *int
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
			if got := tt.com.GetClientRect(tt.args.hwnd, tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("dmsoft.GetClientRect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetClientSize(t *testing.T) {
	type args struct {
		hwnd   int
		width  *int
		height *int
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
			if got := tt.com.GetClientSize(tt.args.hwnd, tt.args.width, tt.args.height); got != tt.want {
				t.Errorf("dmsoft.GetClientSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetForegroundFocus(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetForegroundFocus(); got != tt.want {
				t.Errorf("dmsoft.GetForegroundFocus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetForegroundWindow(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetForegroundWindow(); got != tt.want {
				t.Errorf("dmsoft.GetForegroundWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetMousePointWindow(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetMousePointWindow(); got != tt.want {
				t.Errorf("dmsoft.GetMousePointWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetPointWindow(t *testing.T) {
	type args struct {
		x int
		y int
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
			if got := tt.com.GetPointWindow(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("dmsoft.GetPointWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetProcessInfo(t *testing.T) {
	type args struct {
		pid int
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
			if got := tt.com.GetProcessInfo(tt.args.pid); got != tt.want {
				t.Errorf("dmsoft.GetProcessInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetSpecialWindow(t *testing.T) {
	type args struct {
		flag int
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
			if got := tt.com.GetSpecialWindow(tt.args.flag); got != tt.want {
				t.Errorf("dmsoft.GetSpecialWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetWindow(t *testing.T) {
	type args struct {
		hwnd int
		flag int
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
			if got := tt.com.GetWindow(tt.args.hwnd, tt.args.flag); got != tt.want {
				t.Errorf("dmsoft.GetWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetWindowClass(t *testing.T) {
	type args struct {
		hwnd int
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
			if got := tt.com.GetWindowClass(tt.args.hwnd); got != tt.want {
				t.Errorf("dmsoft.GetWindowClass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetWindowProcessId(t *testing.T) {
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
			if got := tt.com.GetWindowProcessId(tt.args.hwnd); got != tt.want {
				t.Errorf("dmsoft.GetWindowProcessId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetWindowProcessPath(t *testing.T) {
	type args struct {
		hwnd int
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
			if got := tt.com.GetWindowProcessPath(tt.args.hwnd); got != tt.want {
				t.Errorf("dmsoft.GetWindowProcessPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetWindowRect(t *testing.T) {
	type args struct {
		hwnd int
		x1   *int
		y1   *int
		x2   *int
		y2   *int
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
			if got := tt.com.GetWindowRect(tt.args.hwnd, tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("dmsoft.GetWindowRect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetWindowState(t *testing.T) {
	type args struct {
		hwnd int
		flag int
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
			if got := tt.com.GetWindowState(tt.args.hwnd, tt.args.flag); got != tt.want {
				t.Errorf("dmsoft.GetWindowState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetWindowTitle(t *testing.T) {
	type args struct {
		hwnd int
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
			if got := tt.com.GetWindowTitle(tt.args.hwnd); got != tt.want {
				t.Errorf("dmsoft.GetWindowTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_MoveWindow(t *testing.T) {
	type args struct {
		hwnd int
		x    int
		y    int
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
			if got := tt.com.MoveWindow(tt.args.hwnd, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("dmsoft.MoveWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_ScreenToClient(t *testing.T) {
	type args struct {
		hwnd int
		x    *int
		y    *int
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
			if got := tt.com.ScreenToClient(tt.args.hwnd, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("dmsoft.ScreenToClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SendPaste(t *testing.T) {
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
			if got := tt.com.SendPaste(tt.args.hwnd); got != tt.want {
				t.Errorf("dmsoft.SendPaste() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SendString(t *testing.T) {
	type args struct {
		hwnd int
		str  string
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
			if got := tt.com.SendString(tt.args.hwnd, tt.args.str); got != tt.want {
				t.Errorf("dmsoft.SendString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SendString2(t *testing.T) {
	type args struct {
		hwnd int
		str  string
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
			if got := tt.com.SendString2(tt.args.hwnd, tt.args.str); got != tt.want {
				t.Errorf("dmsoft.SendString2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SendStringIme(t *testing.T) {
	type args struct {
		str string
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
			if got := tt.com.SendStringIme(tt.args.str); got != tt.want {
				t.Errorf("dmsoft.SendStringIme() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SendStringIme2(t *testing.T) {
	type args struct {
		hwnd int
		str  string
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
			if got := tt.com.SendStringIme2(tt.args.hwnd, tt.args.str, tt.args.mode); got != tt.want {
				t.Errorf("dmsoft.SendStringIme2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetClientSize(t *testing.T) {
	type args struct {
		hwnd   int
		width  int
		height int
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
			if got := tt.com.SetClientSize(tt.args.hwnd, tt.args.width, tt.args.height); got != tt.want {
				t.Errorf("dmsoft.SetClientSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetWindowSize(t *testing.T) {
	type args struct {
		hwnd   int
		width  int
		height int
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
			if got := tt.com.SetWindowSize(tt.args.hwnd, tt.args.width, tt.args.height); got != tt.want {
				t.Errorf("dmsoft.SetWindowSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetWindowState(t *testing.T) {
	type args struct {
		hwnd int
		flag int
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
			if got := tt.com.SetWindowState(tt.args.hwnd, tt.args.flag); got != tt.want {
				t.Errorf("dmsoft.SetWindowState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetWindowText(t *testing.T) {
	type args struct {
		hwnd  int
		title string
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
			if got := tt.com.SetWindowText(tt.args.hwnd, tt.args.title); got != tt.want {
				t.Errorf("dmsoft.SetWindowText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetWindowTransparent(t *testing.T) {
	type args struct {
		hwnd  int
		trans int
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
			if got := tt.com.SetWindowTransparent(tt.args.hwnd, tt.args.trans); got != tt.want {
				t.Errorf("dmsoft.SetWindowTransparent() = %v, want %v", got, tt.want)
			}
		})
	}
}
