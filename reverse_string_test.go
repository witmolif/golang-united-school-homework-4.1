package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{"t1", args{"runes"}, "senur"},
		{"t1", args{"The quick bròwn 狐 jumped over the lazy 犬"}, "犬 yzal eht revo depmuj 狐 nẁorb kciuq ehT"},
		{"t1", args{`A
B`},
			`B
A`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := ReverseString(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("ReverseString() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
