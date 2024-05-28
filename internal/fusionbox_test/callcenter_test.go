package fusionbox_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-fusion-pbx/internal/fusionbox"
	"github.com/pericles-luz/go-fusion-pbx/internal/gear"
	"github.com/stretchr/testify/require"
)

func TestCallcenterMustFindTierUUID(t *testing.T) {
	callcenter := fusionbox.NewCallcenter()
	callcenter.SetExtension("200115")
	callcenter.SetAgent("1638")
	tierUUID := callcenter.TierUUID(callcenterWithBranchPageContent())
	require.Len(t, tierUUID, 36)
	tierUUID = callcenter.TierUUID(callcenterWithoutBranchPageContent())
	require.Empty(t, tierUUID)
}

func TestCallcenterMustFindAllFormFields(t *testing.T) {
	callcenter := fusionbox.NewCallcenter()
	callcenter.SetExtension("200115")
	callcenter.SetAgent("1638")
	formFields := gear.ExtractFromDataFromHTML(callcenterWithBranchPageContent())
	require.Len(t, formFields, 35)
}

func TestCallcenterMustFindAgentID(t *testing.T) {
	callcenter := fusionbox.NewCallcenter()
	callcenter.SetExtension("200115")
	callcenter.SetAgent("1638")
	agentID := callcenter.GetAgentID(callcenterWithBranchPageContent())
	require.True(t, utils.ValidateUUID(agentID))
	agentID = callcenter.GetAgentID(callcenterWithoutBranchPageContent())
	require.True(t, utils.ValidateUUID(agentID))
}

func TestCallcenterMustSetAsAlreadyInCallcenter(t *testing.T) {
	callcenter := fusionbox.NewCallcenter()
	callcenter.SetExtension("200115")
	callcenter.SetAgent("1638")
	require.True(t, callcenter.IsAlreadyInCallcenter(callcenterWithBranchPageContent()))
}

func TestCallcenterMustSetAsNotAlreadyInCallcenter(t *testing.T) {
	callcenter := fusionbox.NewCallcenter()
	callcenter.SetExtension("200115")
	callcenter.SetAgent("1670")
	require.False(t, callcenter.IsAlreadyInCallcenter(callcenterWithBranchPageContent()))
}
