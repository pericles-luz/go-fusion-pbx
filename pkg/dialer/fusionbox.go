package dialer

import (
	"github.com/pericles-luz/go-fusion-pbx/internal/fusionbox"
	"github.com/pericles-luz/go-fusion-pbx/internal/gear"
)

func AddAgent(credential *fusionbox.Credential, extension, agent string) error {
	callcenter := fusionbox.NewCallcenter()
	callcenter.SetExtension(extension)
	callcenter.SetAgent(agent)
	g := gear.NewGear()
	g.Get(credential.BaseLink)
	Logon(credential, g)
	g.OnHTML("tr[href]", callcenter.VisitCallcenter)
	g.Collector.OnResponse(callcenter.AddAgent)
	g.Get(credential.BaseLink + "/app/call_centers/call_center_queues.php")
	return nil
}

func RemoveAgent(credential *fusionbox.Credential, extension, agent string) error {
	callcenter := fusionbox.NewCallcenter()
	callcenter.SetExtension(extension)
	callcenter.SetAgent(agent)
	g := gear.NewGear()
	g.Get(credential.BaseLink)
	Logon(credential, g)
	g.OnHTML("tr[href]", callcenter.VisitCallcenter)
	g.Collector.OnResponse(callcenter.RemoveAgent)
	g.Get(credential.BaseLink + "/app/call_centers/call_center_queues.php")
	return nil
}

func Logon(credential *fusionbox.Credential, g *gear.Gear) {
	g.Post(credential.BaseLink+"/login.php", map[string]string{
		"username": credential.Username,
		"password": credential.Password,
	})
}
