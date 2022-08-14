package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFooter_SetNumTransactions(t *testing.T) {
	f := NewFooter()
	f.SetNumTransactions(5)
	assert.Equal(t, "5 Transactions", f.Label.Text)
}
