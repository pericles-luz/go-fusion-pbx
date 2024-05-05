package dialer

import (
	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-fusion-pbx/internal"
	"github.com/pericles-luz/go-fusion-pbx/internal/fusionbox"
)

func LoadCredential(filename string) *fusionbox.Credential {
	configContent, err := internal.ReadConfigFile(utils.GetBaseDirectory("config") + "/" + filename)
	if err != nil {
		return nil
	}
	credential := fusionbox.NewCredential()
	err = credential.FillFromJSONFile(configContent)
	if err != nil {
		return nil
	}
	return credential
}
