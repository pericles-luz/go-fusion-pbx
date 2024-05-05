package dialer_test

import (
	"testing"

	"github.com/pericles-luz/go-fusion-pbx/pkg/dialer"
	"github.com/stretchr/testify/require"
)

func TestFusionboxLoadCredential(t *testing.T) {
	t.Skip("Use only if necessary")
	credential := dialer.LoadCredential("fusionbox.json")
	require.NotNil(t, credential)
}

func TestFusionboxAddAgent(t *testing.T) {
	t.Skip("Use only if necessary")
	credential := dialer.LoadCredential("fusionbox.json")
	require.NotNil(t, credential)
	require.NoError(t, dialer.AddAgent(credential, "200115", "1601"))
}

func TestFusionboxRemoveAgent(t *testing.T) {
	t.Skip("Use only if necessary")
	credential := dialer.LoadCredential("fusionbox.json")
	require.NotNil(t, credential)
	require.NoError(t, dialer.RemoveAgent(credential, "200115", "1601"))
}
