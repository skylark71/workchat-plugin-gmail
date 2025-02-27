// This file is automatically generated. Do not modify it manually.

package main

import (
	"strings"

	"gitlab.com/w1572/backend/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "workchat-plugin-gmail",
  "name": "Workchat Gmail Bot",
  "description": "Gmail Integration for Workchat",
  "homepage_url": "https://gitlab.com/w1572/workchat-plugin-gmail/blob/master/README.md",
  "support_url": "https://gitlab.com/w1572/workchat-plugin-gmail/issues",
  "release_notes_url": "https://gitlab.com/w1572/workchat-plugin-gmail/blob/master/CHANGELOG.md",
  "version": "0.1.1",
  "min_server_version": "5.19.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "settings_schema": {
    "header": "The Gmail plugin for Workchat",
    "footer": "Made with Love and Support from Workchat by Abdul Sattar Mapara",
    "settings": [
      {
        "key": "GmailOAuthClientID",
        "display_name": "Client ID",
        "type": "text",
        "help_text": "The client ID for the OAuth app registered with Google Cloud",
        "placeholder": "Please copy client ID over from Google API console for Gmail API",
        "default": null
      },
      {
        "key": "GmailOAuthSecret",
        "display_name": "Client Secret",
        "type": "text",
        "help_text": "The client secret for the OAuth app registered with Google Cloud.",
        "placeholder": "Please copy secret over from Google (gmail) OAuth application",
        "default": null
      },
      {
        "key": "TopicName",
        "display_name": "Topic Name",
        "type": "text",
        "help_text": "Topic Name is used to subscribe user for notifications from Gmail.",
        "placeholder": "Create a topic in Google Cloud pubsub",
        "default": null
      },
      {
        "key": "EncryptionKey",
        "display_name": "Plugin Encryption Key",
        "type": "generated",
        "help_text": "The AES encryption key internally used in plugin to encrypt stored access tokens.",
        "placeholder": "Generate the key and store before connecting the account",
        "default": null
      }
    ]
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}
