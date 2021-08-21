package main

import (
	"testing"
)

/*func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}*/

func TestSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"正数",
			args{int(2)},
			4,
		},
		{"零数",
			args{int(0)},
			0,
		},
		{"负数",
			args{int(-2)},
			4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Square(tt.args.num); got != tt.want {
				t.Errorf("Square() = %v, want %v", got, tt.want)
			}
		})
	}
}