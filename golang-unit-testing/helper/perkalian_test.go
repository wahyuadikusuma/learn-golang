package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTableNilaiKuadrat(t *testing.T) {
	tests := []struct{
		name string
		request int
		expected int
	}{
		{
			name: "10",
			request: 10,
			expected: 100,
		},
		{
			name: "50",
			request: 50,
			expected: 2500,
		},
		{
			name: "7",
			request: 7,
			expected: 49,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := NilaiKuadrat(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}