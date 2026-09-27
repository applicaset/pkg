package initials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	t.Parallel()

	cases := []struct{ name, fallback, want string }{
		{"Ada Lovelace", "ada", "AL"},
		{"Ada King Lovelace", "ada", "AL"},
		{"ada", "x", "A"},
		{"  ", "grace", "G"},
		{"", "", ""},
		{"élodie durand", "", "ÉD"},
		{"ناصر میرزائی", "", "نم"},
	}
	for _, c := range cases {
		assert.Equal(t, c.want, Of(c.name, c.fallback), c.name)
	}
}
