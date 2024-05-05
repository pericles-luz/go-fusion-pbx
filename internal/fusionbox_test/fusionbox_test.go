package fusionbox_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-fusion-pbx/internal"
	"github.com/pericles-luz/go-fusion-pbx/internal/fusionbox"
	"github.com/stretchr/testify/require"
)

func TestFusionboxAddAgent(t *testing.T) {
	t.Skip("Use only if necessary")
	configContent, err := internal.ReadConfigFile(utils.GetBaseDirectory("config") + "/fusionbox.json")
	require.NoError(t, err)
	credential := fusionbox.NewCredential()
	require.NoError(t, credential.FillFromJSONFile(configContent))
	require.NoError(t, fusionbox.AddAgent(credential, "200115", "1601"))
}

func TestFusionboxRemoveAgent(t *testing.T) {
	t.Skip("Use only if necessary")
	configContent, err := internal.ReadConfigFile(utils.GetBaseDirectory("config") + "/fusionbox.json")
	require.NoError(t, err)
	credential := fusionbox.NewCredential()
	require.NoError(t, credential.FillFromJSONFile(configContent))
	require.NoError(t, fusionbox.RemoveAgent(credential, "200115", "1601"))
}
