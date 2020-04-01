// 图色

package dmsoft

import "testing"

func Test_dmsoft_AppendPicAddr(t *testing.T) {
	type args struct {
		picInfo string
		addr    int
		size    int
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
			if got := tt.com.AppendPicAddr(tt.args.picInfo, tt.args.addr, tt.args.size); got != tt.want {
				t.Errorf("dmsoft.AppendPicAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_Capture(t *testing.T) {
	type args struct {
		x1   int
		y1   int
		x2   int
		y2   int
		file string
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
			if got := tt.com.Capture(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.file); got != tt.want {
				t.Errorf("dmsoft.Capture() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_CaptureGif(t *testing.T) {
	type args struct {
		x1    int
		y1    int
		x2    int
		y2    int
		file  string
		delay int
		time  int
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
			if got := tt.com.CaptureGif(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.file, tt.args.delay, tt.args.time); got != tt.want {
				t.Errorf("dmsoft.CaptureGif() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_CaptureJpg(t *testing.T) {
	type args struct {
		x1      int
		y1      int
		x2      int
		y2      int
		file    string
		quality int
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
			if got := tt.com.CaptureJpg(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.file, tt.args.quality); got != tt.want {
				t.Errorf("dmsoft.CaptureJpg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_CapturePng(t *testing.T) {
	type args struct {
		x1   int
		y1   int
		x2   int
		y2   int
		file string
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
			if got := tt.com.CapturePng(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.file); got != tt.want {
				t.Errorf("dmsoft.CapturePng() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_CapturePre(t *testing.T) {
	type args struct {
		file string
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
			if got := tt.com.CapturePre(tt.args.file); got != tt.want {
				t.Errorf("dmsoft.CapturePre() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_CmpColor(t *testing.T) {
	type args struct {
		x     int
		y     int
		color string
		sim   float32
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
			if got := tt.com.CmpColor(tt.args.x, tt.args.y, tt.args.color, tt.args.sim); got != tt.want {
				t.Errorf("dmsoft.CmpColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableDisplayDebug(t *testing.T) {
	type args struct {
		enableDebug int
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
			if got := tt.com.EnableDisplayDebug(tt.args.enableDebug); got != tt.want {
				t.Errorf("dmsoft.EnableDisplayDebug() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableFindPicMultithread(t *testing.T) {
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
			if got := tt.com.EnableFindPicMultithread(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.EnableFindPicMultithread() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableGetColorByCapture(t *testing.T) {
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
			if got := tt.com.EnableGetColorByCapture(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.EnableGetColorByCapture() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindColor(t *testing.T) {
	type args struct {
		x1    int
		y1    int
		x2    int
		y2    int
		color string
		sim   float32
		dir   int
		intX  *int
		intY  *int
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
			if got := tt.com.FindColor(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.sim, tt.args.dir, tt.args.intX, tt.args.intY); got != tt.want {
				t.Errorf("dmsoft.FindColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindColorBlock(t *testing.T) {
	type args struct {
		x1     int
		y1     int
		x2     int
		y2     int
		color  string
		sim    float32
		count  int
		width  int
		height int
		intX   *int
		intY   *int
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
			if got := tt.com.FindColorBlock(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.sim, tt.args.count, tt.args.width, tt.args.height, tt.args.intX, tt.args.intY); got != tt.want {
				t.Errorf("dmsoft.FindColorBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindColorBlockEx(t *testing.T) {
	type args struct {
		x1     int
		y1     int
		x2     int
		y2     int
		color  string
		sim    float32
		count  int
		width  int
		height int
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
			if got := tt.com.FindColorBlockEx(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.sim, tt.args.count, tt.args.width, tt.args.height); got != tt.want {
				t.Errorf("dmsoft.FindColorBlockEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindColorE(t *testing.T) {
	type args struct {
		x1    int
		y1    int
		x2    int
		y2    int
		color string
		sim   float32
		dir   int
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
			if got := tt.com.FindColorE(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.sim, tt.args.dir); got != tt.want {
				t.Errorf("dmsoft.FindColorE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindColorEx(t *testing.T) {
	type args struct {
		x1    int
		y1    int
		x2    int
		y2    int
		color string
		sim   float32
		dir   int
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
			if got := tt.com.FindColorEx(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.sim, tt.args.dir); got != tt.want {
				t.Errorf("dmsoft.FindColorEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindMulColor(t *testing.T) {
	type args struct {
		x1    int
		y1    int
		x2    int
		y2    int
		color string
		sim   float32
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
			if got := tt.com.FindMulColor(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.sim); got != tt.want {
				t.Errorf("dmsoft.FindMulColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindMultiColor(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		firstColor  string
		offsetColor string
		sim         float32
		dir         int
		intX        *int
		intY        *int
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
			if got := tt.com.FindMultiColor(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.firstColor, tt.args.offsetColor, tt.args.sim, tt.args.dir, tt.args.intX, tt.args.intY); got != tt.want {
				t.Errorf("dmsoft.FindMultiColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindMultiColorE(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		firstColor  string
		offsetColor string
		sim         float32
		dir         int
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
			if got := tt.com.FindMultiColorE(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.firstColor, tt.args.offsetColor, tt.args.sim, tt.args.dir); got != tt.want {
				t.Errorf("dmsoft.FindMultiColorE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindMultiColorEx(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		firstColor  string
		offsetColor string
		sim         float32
		dir         int
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
			if got := tt.com.FindMultiColorEx(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.firstColor, tt.args.offsetColor, tt.args.sim, tt.args.dir); got != tt.want {
				t.Errorf("dmsoft.FindMultiColorEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindPic(t *testing.T) {
	type args struct {
		x1         int
		y1         int
		x2         int
		y2         int
		picName    string
		deltaColor string
		sim        float32
		dir        int
		intX       *int
		intY       *int
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
			if got := tt.com.FindPic(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.picName, tt.args.deltaColor, tt.args.sim, tt.args.dir, tt.args.intX, tt.args.intY); got != tt.want {
				t.Errorf("dmsoft.FindPic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindPicE(t *testing.T) {
	type args struct {
		x1         int
		y1         int
		x2         int
		y2         int
		picName    string
		deltaColor string
		sim        float32
		dir        int
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
			if got := tt.com.FindPicE(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.picName, tt.args.deltaColor, tt.args.sim, tt.args.dir); got != tt.want {
				t.Errorf("dmsoft.FindPicE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindPicEx(t *testing.T) {
	type args struct {
		x1         int
		y1         int
		x2         int
		y2         int
		picName    string
		deltaColor string
		sim        float32
		dir        int
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
			if got := tt.com.FindPicEx(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.picName, tt.args.deltaColor, tt.args.sim, tt.args.dir); got != tt.want {
				t.Errorf("dmsoft.FindPicEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindPicExS(t *testing.T) {
	type args struct {
		x1         int
		y1         int
		x2         int
		y2         int
		picName    string
		deltaColor string
		sim        float32
		dir        int
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
			if got := tt.com.FindPicExS(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.picName, tt.args.deltaColor, tt.args.sim, tt.args.dir); got != tt.want {
				t.Errorf("dmsoft.FindPicExS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindPicS(t *testing.T) {
	type args struct {
		x1         int
		y1         int
		x2         int
		y2         int
		picName    string
		deltaColor string
		sim        float32
		dir        int
		intX       *int
		intY       *int
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
			if got := tt.com.FindPicS(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.picName, tt.args.deltaColor, tt.args.sim, tt.args.dir, tt.args.intX, tt.args.intY); got != tt.want {
				t.Errorf("dmsoft.FindPicS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetAveHSV(t *testing.T) {
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
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetAveHSV(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("dmsoft.GetAveHSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetAveRGB(t *testing.T) {
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
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetAveRGB(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("dmsoft.GetAveRGB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetColor(t *testing.T) {
	type args struct {
		x int
		y int
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
			if got := tt.com.GetColor(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("dmsoft.GetColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetColorHSV(t *testing.T) {
	type args struct {
		x int
		y int
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
			if got := tt.com.GetColorHSV(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("dmsoft.GetColorHSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetPicSize(t *testing.T) {
	type args struct {
		picName string
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
			if got := tt.com.GetPicSize(tt.args.picName); got != tt.want {
				t.Errorf("dmsoft.GetPicSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_IsDisplayDead(t *testing.T) {
	type args struct {
		x1   int
		y1   int
		x2   int
		y2   int
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
			if got := tt.com.IsDisplayDead(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.time); got != tt.want {
				t.Errorf("dmsoft.IsDisplayDead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_MatchPicName(t *testing.T) {
	type args struct {
		picName string
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
			if got := tt.com.MatchPicName(tt.args.picName); got != tt.want {
				t.Errorf("dmsoft.MatchPicName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetExcludeRegion(t *testing.T) {
	type args struct {
		mode int
		info string
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
			if got := tt.com.SetExcludeRegion(tt.args.mode, tt.args.info); got != tt.want {
				t.Errorf("dmsoft.SetExcludeRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetPicPwd(t *testing.T) {
	type args struct {
		pwd string
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
			if got := tt.com.SetPicPwd(tt.args.pwd); got != tt.want {
				t.Errorf("dmsoft.SetPicPwd() = %v, want %v", got, tt.want)
			}
		})
	}
}
