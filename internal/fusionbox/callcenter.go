package fusionbox

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-fusion-pbx/internal/gear"
)

type Callcenter struct {
	extension   string
	extensionID string
	agent       string
	visited     bool
	acted       bool
}

func NewCallcenter() *Callcenter {
	return &Callcenter{}
}

func (c *Callcenter) SetExtension(extension string) {
	c.extension = extension
}

func (c *Callcenter) SetAgent(agent string) {
	c.agent = agent
}

func (c *Callcenter) Extension() string {
	return c.extension
}

func (c *Callcenter) Agent() string {
	return c.agent
}

func (c *Callcenter) VisitCallcenter(e *colly.HTMLElement) {
	if c.visited {
		return
	}
	println(e.Text)
	if !strings.Contains(e.Text, c.extension) {
		return
	}
	// extensionID is the last 36 characters of the href
	c.extensionID = e.Attr("href")[len(e.Attr("href"))-36:]
	c.visited = true
	utils.ManageError(e.Request.Visit(e.Attr("href")))
}

func (c *Callcenter) AddAgent(r *colly.Response) {
	if !c.visited {
		return
	}
	if c.acted {
		return
	}
	fields := gear.ExtractFromDataFromHTML(string(r.Body))
	tierUUID := c.TierUUID(string(r.Body))
	if utils.ValidateUUID(tierUUID) {
		return
	}
	agentID := c.GetAgentID(string(r.Body))
	nextTierKey := c.NextTierKey(fields)
	if nextTierKey == "" {
		return
	}
	println("Adding agent to callcenter")
	fields[nextTierKey] = agentID
	utils.ManageError(r.Request.Post(fmt.Sprintf("call_center_queue_edit.php?id=%s", c.extensionID), fields))
	c.acted = true
	utils.ManageError(r.Request.Visit(fmt.Sprintf("cmd.php?cmd=reload&id=%s", c.extensionID)))
}

func (c *Callcenter) NextTierKey(fields map[string]string) string {
	for k, v := range fields {
		if strings.HasSuffix(k, "[call_center_agent_uuid]") && v == "" {
			return k
		}
	}
	return ""
}

func (c *Callcenter) RemoveAgent(r *colly.Response) {
	if !c.visited {
		return
	}
	if c.acted {
		return
	}
	tierUUID := c.TierUUID(string(r.Body))
	if tierUUID == "" {
		return
	}
	println(tierUUID)
	println(fmt.Sprintf("call_center_queue_edit.php?id=%s&call_center_tier_uuid=%s&a=delete", c.extensionID, tierUUID))
	utils.ManageError(r.Request.Visit(fmt.Sprintf("call_center_queue_edit.php?id=%s&call_center_tier_uuid=%s&a=delete", c.extensionID, tierUUID)))
	c.acted = true
	utils.ManageError(r.Request.Visit(fmt.Sprintf("cmd.php?cmd=reload&id=%s", c.extensionID)))
}

func (c *Callcenter) TierUUID(source string) string {
	agentPosition := strings.Index(source, c.agent)
	if agentPosition == -1 {
		return ""
	}
	tierUUIDPosition := strings.Index(source[:agentPosition], "call_center_tier_uuid")
	if tierUUIDPosition == -1 {
		return ""
	}
	tierUUIDStart := strings.Index(source[tierUUIDPosition:], `"`) + 1
	if tierUUIDStart == -1 {
		return ""
	}
	if (agentPosition - tierUUIDPosition) > 280 {
		return ""
	}
	tierUUIDEnd := strings.Index(source[tierUUIDPosition+tierUUIDStart:], `"`)
	println("tierUUIDPosition", tierUUIDPosition)
	println("tierUUIDStart", tierUUIDStart)
	println("diference", agentPosition-tierUUIDPosition)
	println("tierUUIDEnd", tierUUIDEnd)
	return source[tierUUIDPosition+tierUUIDStart : tierUUIDPosition+tierUUIDStart+tierUUIDEnd]
}

// gets agentid from html option tag
func (c *Callcenter) GetAgentID(source string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(source))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return ""
	}
	agentID := ""
	doc.Find("option").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), c.agent) && agentID == "" {
			agentID = s.AttrOr("value", "")
			return
		}
	})
	return agentID
}
