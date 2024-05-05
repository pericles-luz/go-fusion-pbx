package fusionbox

import "encoding/json"

type Credential struct {
	Username string
	Password string
	BaseLink string
}

func NewCredential() *Credential {
	return &Credential{}
}

func (c *Credential) FillFromJSONFile(raw []byte) error {
	err := json.Unmarshal(raw, c)
	if err != nil {
		return err
	}
	return nil
}
