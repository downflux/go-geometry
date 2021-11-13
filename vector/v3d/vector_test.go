package vector

import (
	"testing"
)

func TestCross(t *testing.T) {
	testConfigs := []struct {
		name string
		v    V
		u    V
		want V
	}{
		{
			name: "TrivialBasis",
			v:    *New(1, 0, 0),
			u:    *New(0, 1, 0),
			want: *New(0, 0, 1),
		},
		{
			name: "TrivialBasis/Anticommutative",
			v:    *New(0, 1, 0),
			u:    *New(1, 0, 0),
			want: *New(0, 0, -1),
		},
	}
	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := Cross(c.v, c.u); !Within(got, c.want) {
				t.Errorf("Cross() = %v, want = %v", got, c.want)
			}
		})
	}
}
