package identity_manager

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUUIDV4(t *testing.T) {
	uuid := NewUUIDV4()
	require.True(t, IsValidUUID(uuid))
}

func TestIsValidUUID(t *testing.T) {
	valid := "a3bb189e-8bf9-3888-9912-ace4e6543002"
	invalid := "not-a-uuid"
	require.True(t, IsValidUUID(valid))
	require.False(t, IsValidUUID(invalid))
}

func TestIsNotValidUUID(t *testing.T) {
	valid := "a3bb189e-8bf9-3888-9912-ace4e6543002"
	invalid := "not-a-uuid"
	require.False(t, IsNotValidUUID(valid))
	require.True(t, IsNotValidUUID(invalid))
}
