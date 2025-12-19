package api

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStartAPI(t *testing.T) {
	tests := []struct {
		name string // description of this test case
	}{
		{
			name: "Try start api",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StartAPI()
			if got != nil {
				t.Error("the engine expected to be nil")
			}
			StartEngine()
			got = StartAPI()
			if got == nil {
				t.Error("can't start engine")
			} else {
				other := StartAPI()
				if got != other {
					t.Error("can't get singleton instance of api")
				}
			}
		})
	}
}

func Test_startEngine(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want *gin.Engine
	}{
		{
			name: "Test Call twice GoEngine",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StartEngine()
			if got == nil {
				t.Errorf("can't start engine")
			} else {
				other := StartEngine()
				if got != other {
					t.Errorf("can't load the singleton statement engine")
				}
			}
		})
	}
}
