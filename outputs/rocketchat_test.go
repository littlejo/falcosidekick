package outputs

import (
	"encoding/json"
	"testing"
	"text/template"

	"github.com/falcosecurity/falcosidekick/types"

	"github.com/stretchr/testify/require"
)

func TestNewRocketchatPayload(t *testing.T) {
	expectedOutput := slackPayload{
		Text:     "Rule: Test rule Priority: Debug",
		Username: "Falcosidekick",
		IconURL:  DefaultIconURL,
		Attachments: []slackAttachment{
			{
				Fallback: "This is a test from falcosidekick",
				Color:    PaleCyan,
				Text:     "This is a test from falcosidekick",
				Footer:   "",
				Fields: []slackAttachmentField{
					{
						Title: "proc.name",
						Value: "falcosidekick",
						Short: true,
					},
					{
						Title: "rule",
						Value: "Test rule",
						Short: true,
					},
					{
						Title: "priority",
						Value: "Debug",
						Short: true,
					},
					{
						Title: "time",
						Value: "2001-01-01 01:10:00 +0000 UTC",
						Short: false,
					},
				},
			},
		},
	}

	var f types.FalcoPayload
	require.Nil(t, json.Unmarshal([]byte(falcoTestInput), &f))
	config := &types.Configuration{
		Rocketchat: types.RocketchatOutputConfig{
			Username: "Falcosidekick",
			Icon:     DefaultIconURL,
		},
	}

	var err error
	config.Rocketchat.MessageFormatTemplate, err = template.New("").Parse("Rule: {{ .Rule }} Priority: {{ .Priority }}")
	require.Nil(t, err)

	output := newRocketchatPayload(f, config)
	require.Equal(t, output, expectedOutput)
}
