package components

import (
	"testing"

	"github.com/logos/ecs/components"
	"github.com/logos/ecs/components/attributes"
	"github.com/logos/internal/utils"
)

func TestNewElementAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value [5]attributes.DiagramsAttributes
		want  components.ElementAttributes
	}{
		{
			name: "default values",
			value: [5]attributes.DiagramsAttributes{
				{Sky: 1},
				{Earth: 2},
				{Thunder: 3},
				{Wind: 4},
				{Water: 5},
			},
			want: components.ElementAttributes{
				Value: [5]attributes.DiagramsAttributes{
					{Sky: 1},
					{Earth: 2},
					{Thunder: 3},
					{Wind: 4},
					{Water: 5},
				},
			},
		},
		{
			name: "all zeros",
			value: [5]attributes.DiagramsAttributes{
				{Sky: 0},
				{Earth: 0},
				{Thunder: 0},
				{Wind: 0},
				{Water: 0},
			},
			want: components.ElementAttributes{
				Value: [5]attributes.DiagramsAttributes{
					{Sky: 0},
					{Earth: 0},
					{Thunder: 0},
					{Wind: 0},
					{Water: 0},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := components.NewElementAttributes(tt.value)
			if got != tt.want {
				t.Errorf("NewElementAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElementAttributes_EncodeByDiagrams(t *testing.T) {
	tests := []struct {
		name       string
		value      [5]attributes.DiagramsAttributes
		wantEncode uint64
	}{
		{
			name: "default values",
			value: [5]attributes.DiagramsAttributes{
				{Sky: 1},
				{Earth: 2},
				{Thunder: 3},
				{Wind: 4},
				{Water: 5},
			},
			wantEncode: 68853957121,
		},
		{
			name: "all zeros",
			value: [5]attributes.DiagramsAttributes{
				{Sky: 0},
				{Earth: 0},
				{Thunder: 0},
				{Wind: 0},
				{Water: 0},
			},
			wantEncode: 0x00000000000000000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := components.NewElementAttributes(tt.value)
			e.EncodeByDiagrams()
			t.Logf("want:\t%v", utils.FormatBinary(tt.wantEncode, 40))
			t.Logf("got:\t%v", utils.FormatBinary(e.EncodeValue, 40))
			if e.EncodeValue != tt.wantEncode {
				t.Errorf("EncodeByDiagrams() = %v, want %v", e.EncodeValue, tt.wantEncode)
			}
		})
	}
}
