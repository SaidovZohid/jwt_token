package utils

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := "1234"
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckHashedPassword(password, hashedPassword)
	require.NoError(t, err)
}