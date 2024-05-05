package dialer

import (
	"github.com/pericles-luz/go-fusion-pbx/internal"
	"github.com/pericles-luz/go-fusion-pbx/internal/fusionbox"
)

func LoadCredential(path string) *fusionbox.Credential {
	configContent, err := internal.ReadConfigFile(path)
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
