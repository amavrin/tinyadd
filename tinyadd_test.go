package tinyadd

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		{
			args: []int{1, 2, 3},
			want: 6,
		},
		{
			args: []int{},
			want: 0,
		},
	}

	for _, test := range tests {
		res := Add(test.args...)
		assert.Equal(t, test.want, res)
	}
}
