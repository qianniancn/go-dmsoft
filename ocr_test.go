// 文字识别

package dmsoft

import "testing"

func Test_dmsoft_AddDict(t *testing.T) {
	type args struct {
		index    int
		dictInfo string
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
			if got := tt.com.AddDict(tt.args.index, tt.args.dictInfo); got != tt.want {
				t.Errorf("dmsoft.AddDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_ClearDict(t *testing.T) {
	type args struct {
		index int
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
			if got := tt.com.ClearDict(tt.args.index); got != tt.want {
				t.Errorf("dmsoft.ClearDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_EnableShareDict(t *testing.T) {
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
			if got := tt.com.EnableShareDict(tt.args.enable); got != tt.want {
				t.Errorf("dmsoft.EnableShareDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FetchWord(t *testing.T) {
	type args struct {
		x1    int
		y1    int
		x2    int
		y2    int
		color string
		word  string
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
			if got := tt.com.FetchWord(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.word); got != tt.want {
				t.Errorf("dmsoft.FetchWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindStr(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		str         string
		colorFormat string
		sim         float32
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
			if got := tt.com.FindStr(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.str, tt.args.colorFormat, tt.args.sim, tt.args.intX, tt.args.intY); got != tt.want {
				t.Errorf("dmsoft.FindStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindStrE(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		str         string
		colorFormat string
		sim         float32
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
			if got := tt.com.FindStrE(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.str, tt.args.colorFormat, tt.args.sim); got != tt.want {
				t.Errorf("dmsoft.FindStrE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindStrEx(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		str         string
		colorFormat string
		sim         float32
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
			if got := tt.com.FindStrEx(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.str, tt.args.colorFormat, tt.args.sim); got != tt.want {
				t.Errorf("dmsoft.FindStrEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindStrExS(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		str         string
		colorFormat string
		sim         float32
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
			if got := tt.com.FindStrExS(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.str, tt.args.colorFormat, tt.args.sim); got != tt.want {
				t.Errorf("dmsoft.FindStrExS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindStrFast(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		str         string
		colorFormat string
		sim         float32
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
			if got := tt.com.FindStrFast(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.str, tt.args.colorFormat, tt.args.sim, tt.args.intX, tt.args.intY); got != tt.want {
				t.Errorf("dmsoft.FindStrFast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindStrFastE(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		str         string
		colorFormat string
		sim         float32
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
			if got := tt.com.FindStrFastE(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.str, tt.args.colorFormat, tt.args.sim); got != tt.want {
				t.Errorf("dmsoft.FindStrFastE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindStrFastEx(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		str         string
		colorFormat string
		sim         float32
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
			if got := tt.com.FindStrFastEx(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.str, tt.args.colorFormat, tt.args.sim); got != tt.want {
				t.Errorf("dmsoft.FindStrFastEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindStrFastExS(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		str         string
		colorFormat string
		sim         float32
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
			if got := tt.com.FindStrFastExS(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.str, tt.args.colorFormat, tt.args.sim); got != tt.want {
				t.Errorf("dmsoft.FindStrFastExS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindStrFastS(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		str         string
		colorFormat string
		sim         float32
		intX        *int
		intY        *int
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
			if got := tt.com.FindStrFastS(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.str, tt.args.colorFormat, tt.args.sim, tt.args.intX, tt.args.intY); got != tt.want {
				t.Errorf("dmsoft.FindStrFastS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindStrS(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		str         string
		colorFormat string
		sim         float32
		intX        *int
		intY        *int
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
			if got := tt.com.FindStrS(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.str, tt.args.colorFormat, tt.args.sim, tt.args.intX, tt.args.intY); got != tt.want {
				t.Errorf("dmsoft.FindStrS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindStrWithFont(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		str         string
		colorFormat string
		sim         float32
		fontName    string
		fontSize    int
		flag        int
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
			if got := tt.com.FindStrWithFont(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.str, tt.args.colorFormat, tt.args.sim, tt.args.fontName, tt.args.fontSize, tt.args.flag, tt.args.intX, tt.args.intY); got != tt.want {
				t.Errorf("dmsoft.FindStrWithFont() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindStrWithFontE(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		str         string
		colorFormat string
		sim         float32
		fontName    string
		fontSize    int
		flag        int
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
			if got := tt.com.FindStrWithFontE(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.str, tt.args.colorFormat, tt.args.sim, tt.args.fontName, tt.args.fontSize, tt.args.flag); got != tt.want {
				t.Errorf("dmsoft.FindStrWithFontE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_FindStrWithFontEx(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		str         string
		colorFormat string
		sim         float32
		fontName    string
		fontSize    int
		flag        int
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
			if got := tt.com.FindStrWithFontEx(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.str, tt.args.colorFormat, tt.args.sim, tt.args.fontName, tt.args.fontSize, tt.args.flag); got != tt.want {
				t.Errorf("dmsoft.FindStrWithFontEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetDict(t *testing.T) {
	type args struct {
		index     int
		fontIndex int
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
			if got := tt.com.GetDict(tt.args.index, tt.args.fontIndex); got != tt.want {
				t.Errorf("dmsoft.GetDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetDictCount(t *testing.T) {
	type args struct {
		index int
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
			if got := tt.com.GetDictCount(tt.args.index); got != tt.want {
				t.Errorf("dmsoft.GetDictCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetDictInfo(t *testing.T) {
	type args struct {
		str      string
		fontName string
		fontSize int
		flag     int
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
			if got := tt.com.GetDictInfo(tt.args.str, tt.args.fontName, tt.args.fontSize, tt.args.flag); got != tt.want {
				t.Errorf("dmsoft.GetDictInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetNowDict(t *testing.T) {
	tests := []struct {
		name string
		com  *dmsoft
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetNowDict(); got != tt.want {
				t.Errorf("dmsoft.GetNowDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetResultCount(t *testing.T) {
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
			if got := tt.com.GetResultCount(tt.args.str); got != tt.want {
				t.Errorf("dmsoft.GetResultCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetResultPos(t *testing.T) {
	type args struct {
		str   string
		index int
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
			if got := tt.com.GetResultPos(tt.args.str, tt.args.index, tt.args.intX, tt.args.intY); got != tt.want {
				t.Errorf("dmsoft.GetResultPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetWordResultCount(t *testing.T) {
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
			if got := tt.com.GetWordResultCount(tt.args.str); got != tt.want {
				t.Errorf("dmsoft.GetWordResultCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetWordResultPos(t *testing.T) {
	type args struct {
		str   string
		index int
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
			if got := tt.com.GetWordResultPos(tt.args.str, tt.args.index, tt.args.intX, tt.args.intY); got != tt.want {
				t.Errorf("dmsoft.GetWordResultPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetWordResultStr(t *testing.T) {
	type args struct {
		str   string
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
			if got := tt.com.GetWordResultStr(tt.args.str, tt.args.index); got != tt.want {
				t.Errorf("dmsoft.GetWordResultStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetWords(t *testing.T) {
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
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.com.GetWords(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.sim); got != tt.want {
				t.Errorf("dmsoft.GetWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_GetWordsNoDict(t *testing.T) {
	type args struct {
		x1    int
		y1    int
		x2    int
		y2    int
		color string
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
			if got := tt.com.GetWordsNoDict(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color); got != tt.want {
				t.Errorf("dmsoft.GetWordsNoDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_Ocr(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		colorFormat string
		sim         float32
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
			if got := tt.com.Ocr(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.colorFormat, tt.args.sim); got != tt.want {
				t.Errorf("dmsoft.Ocr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_OcrEx(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		colorFormat string
		sim         float32
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
			if got := tt.com.OcrEx(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.colorFormat, tt.args.sim); got != tt.want {
				t.Errorf("dmsoft.OcrEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_OcrExOne(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		colorFormat string
		sim         float32
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
			if got := tt.com.OcrExOne(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.colorFormat, tt.args.sim); got != tt.want {
				t.Errorf("dmsoft.OcrExOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_OcrInFile(t *testing.T) {
	type args struct {
		x1          int
		y1          int
		x2          int
		y2          int
		picName     string
		colorFormat string
		sim         float32
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
			if got := tt.com.OcrInFile(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.picName, tt.args.colorFormat, tt.args.sim); got != tt.want {
				t.Errorf("dmsoft.OcrInFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SaveDict(t *testing.T) {
	type args struct {
		index int
		file  string
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
			if got := tt.com.SaveDict(tt.args.index, tt.args.file); got != tt.want {
				t.Errorf("dmsoft.SaveDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetColGapNoDict(t *testing.T) {
	type args struct {
		colGap int
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
			if got := tt.com.SetColGapNoDict(tt.args.colGap); got != tt.want {
				t.Errorf("dmsoft.SetColGapNoDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetDict(t *testing.T) {
	type args struct {
		index int
		file  string
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
			if got := tt.com.SetDict(tt.args.index, tt.args.file); got != tt.want {
				t.Errorf("dmsoft.SetDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetDictMem(t *testing.T) {
	type args struct {
		index int
		addr  int
		size  int
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
			if got := tt.com.SetDictMem(tt.args.index, tt.args.addr, tt.args.size); got != tt.want {
				t.Errorf("dmsoft.SetDictMem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetDictPwd(t *testing.T) {
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
			if got := tt.com.SetDictPwd(tt.args.pwd); got != tt.want {
				t.Errorf("dmsoft.SetDictPwd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetExactOcr(t *testing.T) {
	type args struct {
		exactOcr int
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
			if got := tt.com.SetExactOcr(tt.args.exactOcr); got != tt.want {
				t.Errorf("dmsoft.SetExactOcr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetMinColGap(t *testing.T) {
	type args struct {
		minColGap int
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
			if got := tt.com.SetMinColGap(tt.args.minColGap); got != tt.want {
				t.Errorf("dmsoft.SetMinColGap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetMinRowGap(t *testing.T) {
	type args struct {
		minRowGap int
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
			if got := tt.com.SetMinRowGap(tt.args.minRowGap); got != tt.want {
				t.Errorf("dmsoft.SetMinRowGap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetRowGapNoDict(t *testing.T) {
	type args struct {
		RowGap int
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
			if got := tt.com.SetRowGapNoDict(tt.args.RowGap); got != tt.want {
				t.Errorf("dmsoft.SetRowGapNoDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetWordGap(t *testing.T) {
	type args struct {
		wordGap int
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
			if got := tt.com.SetWordGap(tt.args.wordGap); got != tt.want {
				t.Errorf("dmsoft.SetWordGap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetWordGapNoDict(t *testing.T) {
	type args struct {
		wordGap int
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
			if got := tt.com.SetWordGapNoDict(tt.args.wordGap); got != tt.want {
				t.Errorf("dmsoft.SetWordGapNoDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetWordLineHeight(t *testing.T) {
	type args struct {
		lineHeight int
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
			if got := tt.com.SetWordLineHeight(tt.args.lineHeight); got != tt.want {
				t.Errorf("dmsoft.SetWordLineHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_SetWordLineHeightNoDict(t *testing.T) {
	type args struct {
		lineHeight int
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
			if got := tt.com.SetWordLineHeightNoDict(tt.args.lineHeight); got != tt.want {
				t.Errorf("dmsoft.SetWordLineHeightNoDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dmsoft_UseDict(t *testing.T) {
	type args struct {
		index int
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
			if got := tt.com.UseDict(tt.args.index); got != tt.want {
				t.Errorf("dmsoft.UseDict() = %v, want %v", got, tt.want)
			}
		})
	}
}
