// 鼠键

package dmsoft

import "testing"

func Test_dmsoft_EnableMouseAccuracy(t *testing.T) {
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
			if got := tt.com.EnableMouseAccuracy(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.EnableMouseAccuracy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetCursorPos(t *testing.T) {
	type args struct {
		x *int
		y *int
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
			if got := tt.com.GetCursorPos(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("dmsoft.GetCursorPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetCursorShape(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetCursorShape(); got != tt.want {
				t.Errorf("dmsoft.GetCursorShape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetCursorShapeEx(t *testing.T) {
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
			if got := tt.com.GetCursorShapeEx(tt.args.types); got != tt.want {
				t.Errorf("dmsoft.GetCursorShapeEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetCursorSpot(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetCursorSpot(); got != tt.want {
				t.Errorf("dmsoft.GetCursorSpot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetKeyState(t *testing.T) {
	type args struct {
		vkCode int
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
			if got := tt.com.GetKeyState(tt.args.vkCode); got != tt.want {
				t.Errorf("dmsoft.GetKeyState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetMouseSpeed(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetMouseSpeed(); got != tt.want {
				t.Errorf("dmsoft.GetMouseSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_KeyDown(t *testing.T) {
	type args struct {
		vkCode int
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
			if got := tt.com.KeyDown(tt.args.vkCode); got != tt.want {
				t.Errorf("dmsoft.KeyDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_KeyDownChar(t *testing.T) {
	type args struct {
		keyStr string
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
			if got := tt.com.KeyDownChar(tt.args.keyStr); got != tt.want {
				t.Errorf("dmsoft.KeyDownChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_KeyPress(t *testing.T) {
	type args struct {
		vkCode int
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
			if got := tt.com.KeyPress(tt.args.vkCode); got != tt.want {
				t.Errorf("dmsoft.KeyPress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_KeyPressChar(t *testing.T) {
	type args struct {
		keyStr string
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
			if got := tt.com.KeyPressChar(tt.args.keyStr); got != tt.want {
				t.Errorf("dmsoft.KeyPressChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_KeyPressStr(t *testing.T) {
	type args struct {
		keyStr string
		delay  int
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
			if got := tt.com.KeyPressStr(tt.args.keyStr, tt.args.delay); got != tt.want {
				t.Errorf("dmsoft.KeyPressStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_KeyUp(t *testing.T) {
	type args struct {
		vkCode int
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
			if got := tt.com.KeyUp(tt.args.vkCode); got != tt.want {
				t.Errorf("dmsoft.KeyUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_KeyUpChar(t *testing.T) {
	type args struct {
		keyStr string
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
			if got := tt.com.KeyUpChar(tt.args.keyStr); got != tt.want {
				t.Errorf("dmsoft.KeyUpChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_LeftClick(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.LeftClick(); got != tt.want {
				t.Errorf("dmsoft.LeftClick() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_LeftDoubleClick(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.LeftDoubleClick(); got != tt.want {
				t.Errorf("dmsoft.LeftDoubleClick() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_LeftDown(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.LeftDown(); got != tt.want {
				t.Errorf("dmsoft.LeftDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_LeftUp(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.LeftUp(); got != tt.want {
				t.Errorf("dmsoft.LeftUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_MiddleClick(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.MiddleClick(); got != tt.want {
				t.Errorf("dmsoft.MiddleClick() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_MiddleDown(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.MiddleDown(); got != tt.want {
				t.Errorf("dmsoft.MiddleDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_MiddleUp(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.MiddleUp(); got != tt.want {
				t.Errorf("dmsoft.MiddleUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_MoveR(t *testing.T) {
	type args struct {
		rx int
		ry int
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
			if got := tt.com.MoveR(tt.args.rx, tt.args.ry); got != tt.want {
				t.Errorf("dmsoft.MoveR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_MoveTo(t *testing.T) {
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
			if got := tt.com.MoveTo(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("dmsoft.MoveTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_MoveToEx(t *testing.T) {
	type args struct {
		x int
		y int
		w int
		h int
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
			if got := tt.com.MoveToEx(tt.args.x, tt.args.y, tt.args.w, tt.args.h); got != tt.want {
				t.Errorf("dmsoft.MoveToEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_RightClick(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.RightClick(); got != tt.want {
				t.Errorf("dmsoft.RightClick() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_RightDown(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.RightDown(); got != tt.want {
				t.Errorf("dmsoft.RightDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_RightUp(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.RightUp(); got != tt.want {
				t.Errorf("dmsoft.RightUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetKeypadDelay(t *testing.T) {
	type args struct {
		types string
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
			if got := tt.com.SetKeypadDelay(tt.args.types, tt.args.delay); got != tt.want {
				t.Errorf("dmsoft.SetKeypadDelay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetMouseDelay(t *testing.T) {
	type args struct {
		types string
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
			if got := tt.com.SetMouseDelay(tt.args.types, tt.args.delay); got != tt.want {
				t.Errorf("dmsoft.SetMouseDelay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetMouseSpeed(t *testing.T) {
	type args struct {
		speed int
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
			if got := tt.com.SetMouseSpeed(tt.args.speed); got != tt.want {
				t.Errorf("dmsoft.SetMouseSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetSimMode(t *testing.T) {
	type args struct {
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
			if got := tt.com.SetSimMode(tt.args.mode); got != tt.want {
				t.Errorf("dmsoft.SetSimMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_WaitKey(t *testing.T) {
	type args struct {
		vkCode  int
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
			if got := tt.com.WaitKey(tt.args.vkCode, tt.args.timeOut); got != tt.want {
				t.Errorf("dmsoft.WaitKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_WheelDown(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.WheelDown(); got != tt.want {
				t.Errorf("dmsoft.WheelDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_WheelUp(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.WheelUp(); got != tt.want {
				t.Errorf("dmsoft.WheelUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
