package gear_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-fusion-pbx/internal"
	"github.com/pericles-luz/go-fusion-pbx/internal/fusionbox"
	"github.com/pericles-luz/go-fusion-pbx/internal/gear"
	"github.com/stretchr/testify/require"
)

func TestGearMustAddAgentToCallcenterFusionBox(t *testing.T) {
	// t.Skip("Use only if necessary")
	configContent, err := internal.ReadConfigFile(utils.GetBaseDirectory("config") + "/fusionbox.json")
	require.NoError(t, err)
	credential := fusionbox.NewCredential()
	require.NoError(t, credential.FillFromJSONFile(configContent))
	g := gear.NewGear()
	callcenter := fusionbox.NewCallcenter()
	callcenter.SetExtension("200115")
	callcenter.SetAgent("1601")
	require.NoError(t, g.Get(credential.BaseLink))
	require.NoError(t, g.Post(credential.BaseLink+"/login.php", map[string]string{
		"username": credential.Username,
		"password": credential.Password,
	}))
	g.ShowContent()
	g.OnHTML("tr[href]", callcenter.VisitCallcenter)
	g.Collector.OnResponse(callcenter.AddAgent)
	require.NoError(t, g.Get(credential.BaseLink+"/app/call_centers/call_center_queues.php"))
}

func TestGearMustRemoveAgentToCallcenterFusionBox(t *testing.T) {
	// t.Skip("Use only if necessary")
	configContent, err := internal.ReadConfigFile(utils.GetBaseDirectory("config") + "/fusionbox.json")
	require.NoError(t, err)
	credential := fusionbox.NewCredential()
	require.NoError(t, credential.FillFromJSONFile(configContent))
	g := gear.NewGear()
	callcenter := fusionbox.NewCallcenter()
	callcenter.SetExtension("200115")
	callcenter.SetAgent("1601")
	require.NoError(t, g.Get(credential.BaseLink))
	require.NoError(t, g.Post(credential.BaseLink+"/login.php", map[string]string{
		"username": credential.Username,
		"password": credential.Password,
	}))
	g.ShowContent()
	g.OnHTML("tr[href]", callcenter.VisitCallcenter)
	g.Collector.OnResponse(callcenter.RemoveAgent)
	require.NoError(t, g.Get(credential.BaseLink+"/app/call_centers/call_center_queues.php"))
}
