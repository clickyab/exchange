package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"services/config"
	"services/safe"

	"github.com/Sirupsen/logrus"
	"gopkg.in/fzerorubigd/onion.v2"
)

var (
	userName   = config.RegisterString("services.slack.username", "")
	postIcon   = config.RegisterString("services.slack.icon", ":shit:")
	channel    = config.RegisterString("services.slack.channel", "")
	webHookURL = config.RegisterString("services.slack.webhook_url", "")
	active     = config.RegisterBoolean("services.slack.active", false)
)

type reporter struct {
}

func (reporter) Initialize(*onion.Onion) []onion.Layer {
	return nil
}

// Loaded is called after config loading, so the active is ready here
func (r *reporter) Loaded() {
	if *active {
		safe.Register(r)
	}
}

// payload the slack payload
type payload struct {
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	Username    string       `json:"username"`
	IconURL     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Parse       string       `json:"parse"`
	Attachments []attachment `json:"attachments"`
}

// attachment the attachment
type attachment struct {
	Color   string `json:"color"`
	Text    string `json:"text"`
	PreText string `json:"pretext,omitempty"`
	Title   string `json:"title,omitempty"`
}

// SlackDoMessage Try to send message to configured slack channel
func (reporter) Recover(err error, stack []byte, extra ...interface{}) {
	payload := &payload{}
	payload.Channel = *channel

	payload.Text = err.Error()
	payload.Username = *userName
	payload.Parse = "full" // WTF?
	icon := *postIcon
	if icon != "" {
		if icon[0] == ':' {
			payload.IconEmoji = icon
		} else {
			payload.IconURL = icon
		}
	}

	at := []attachment{}
	for i := range extra {
		if t, ok := extra[i].(*http.Request); ok {
			if b, err := httputil.DumpRequest(t, true); err != nil {
				at = append(at, attachment{
					Title: "Request dump",
					Text:  string(b),
				})
				continue
			}
		}
		at = append(at, attachment{
			Title: fmt.Sprintf("%T", extra[i]),
			Text:  fmt.Sprintf("%+v", extra[i]),
		})
	}

	payload.Attachments = at

	encoded, err := json.Marshal(payload)
	if err != nil {
		logrus.WithField("payload", payload).Warn(err)
		return
	}

	resp, err := http.PostForm(*webHookURL, url.Values{"payload": {string(encoded)}})
	if err != nil {
		logrus.WithField("payload", payload).Warn(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		logrus.WithField("response", resp).Warn("sending payload to slack failed")
		return
	}
}

func init() {
	config.Register(&reporter{})
}