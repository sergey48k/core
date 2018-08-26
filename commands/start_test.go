package commands

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStartCmdFlags(t *testing.T) {
	c := newStartCmd(nil)
	require.Equal(t, "text", string(c.lfv))
	require.Equal(t, "info", string(c.llv))

	require.NoError(t, c.cmd.Flags().Set("log-format", "json"))
	require.Equal(t, "json", string(c.lfv))

	require.NoError(t, c.cmd.Flags().Set("log-level", "debug"))
	require.Equal(t, "debug", string(c.llv))
}
