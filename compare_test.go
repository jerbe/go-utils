package goutils

import "testing"

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/5 19:45
  @describe :
*/

func TestEQ(t *testing.T) {
	type args struct {
		left  interface{}
		right interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "A EQ A",
			args: args{
				left:  "A",
				right: "A",
			},
			want: true,
		},
		{
			name: "1 EQ 1",
			args: args{
				left:  "1",
				right: "1",
			},
			want: true,
		},
		{
			name: "a EQ 1",
			args: args{
				left:  "a",
				right: "1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EQ(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("EQ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualAll(t *testing.T) {
	type args struct {
		obj     interface{}
		target  interface{}
		targets []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal all,a,a",
			args: args{
				obj:     "a",
				target:  "a",
				targets: nil,
			},
			want: true,
		},
		{
			name: "equal all,a,nil",
			args: args{
				obj:     "a",
				target:  nil,
				targets: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualAll(tt.args.obj, tt.args.target, tt.args.targets...); got != tt.want {
				t.Errorf("EqualAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGT(t *testing.T) {
	type args struct {
		left  interface{}
		right interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1 GT 2",
			args: args{
				left:  1,
				right: 2,
			},
			want: false,
		},
		{
			name: "2 GT 1",
			args: args{
				left:  2,
				right: 1,
			},
			want: true,
		},
		{
			name: "string 2 GT 1",
			args: args{
				left:  "2",
				right: "1",
			},
			want: true,
		},
		{
			name: "string 1 GT 2",
			args: args{
				left:  "1",
				right: "2",
			},
			want: false,
		},
		{
			name: "2 GT '1'",
			args: args{
				left:  2,
				right: "1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GT(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("GT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGTE(t *testing.T) {
	type args struct {
		left  interface{}
		right interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2 GTE 1",
			args: args{
				left:  2,
				right: 1,
			},
			want: true,
		},
		{
			name: "1 GTE 1",
			args: args{
				left:  1,
				right: 1,
			},
			want: true,
		},
		{
			name: "'1' GTE 1",
			args: args{
				left:  "1",
				right: 1,
			},
			want: false,
		},
		{
			name: "'1' GTE '1'",
			args: args{
				left:  "1",
				right: "1",
			},
			want: true,
		},
		{
			name: "'2' GTE '1'",
			args: args{
				left:  "2",
				right: "1",
			},
			want: true,
		},
		{
			name: "2 GTE '1'",
			args: args{
				left:  2,
				right: "1",
			},
			want: false,
		},
		{
			name: "2 GTE nil",
			args: args{
				left:  2,
				right: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GTE(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("GTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIn(t *testing.T) {
	type args struct {
		obj     interface{}
		target  interface{}
		targets []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a",
			args: args{
				obj:    "a",
				target: "a",
				targets: []interface{}{
					"1", "2", 2, 1, "0", 0, "-1",
				},
			},
			want: true,
		},
		{
			name: "b",
			args: args{
				obj:    "b",
				target: "a",
				targets: []interface{}{
					1, 2, 3, 4, 5, "5", "6", nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := In(tt.args.obj, tt.args.target, tt.args.targets...); got != tt.want {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNil(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is nil",
			args: args{v: func() interface{} { return nil }()},
			want: true,
		},
		{
			name: "not nil",
			args: args{v: func() interface{} { return "1" }()},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNil(tt.args.v); got != tt.want {
				t.Errorf("IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLT(t *testing.T) {
	type args struct {
		left  interface{}
		right interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LT(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("LT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLTE(t *testing.T) {
	type args struct {
		left  interface{}
		right interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LTE(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("LTE() = %v, want %v", got, tt.want)
			}
		})
	}
}
