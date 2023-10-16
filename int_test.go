package goutils

import "testing"

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/16 10:36
  @describe :
*/

func TestIntBitSet(t *testing.T) {
	type args struct {
		val      int
		position int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntBitSet(tt.args.val, tt.args.position); got != tt.want {
				t.Errorf("IntBitSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
