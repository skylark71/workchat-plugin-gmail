package main

import (
	"io/ioutil"
	"path/filepath"
	"sync"

	"github.com/pkg/errors"
	"gitlab.com/w1572/backend/model"
	"gitlab.com/w1572/backend/plugin"
)

// Plugin implements the interface expected by the Workchat server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.WorkchatPlugin

	// UserID of gmail bot
	gmailBotID string

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

// OnActivate is invoked when the plugin is activated. If an error is returned, the plugin will be terminated.
// The plugin will not receive hooks until after OnActivate returns without error.
// https://developers.workchat.com/extend/plugins/server/reference/#Hooks.OnActivate
func (p *Plugin) OnActivate() error {
	// Retrieves the active configuration under lock.
	config := p.getConfiguration()

	err := config.IsValid()
	if err != nil {
		return err
	}

	// Register the command commandGmail
	if err = p.API.RegisterCommand(&model.Command{
		Trigger:          commandGmail,
		AutoComplete:     true,
		AutoCompleteHint: "[command]",
		AutoCompleteDesc: "Available Commands: connect, disconnect, subscribe, unsubscribe, import, subscriptions, help",
	}); err != nil {
		errorMessage := "failed to register command " + commandGmail
		p.API.LogError(errorMessage, "err", err.Error())
		return errors.Wrapf(err, "failed to register %s command", commandGmail)
	}

	p.API.LogInfo(commandGmail + " command registered")

	gmailBot := &model.Bot{
		Username:    "gmail",
		DisplayName: "Gmail Bot",
		Description: "Created by Workchat Gmail Plugin.",
	}

	// Ensure the bot. If not present create Gmail bot
	gmailBotID, err := p.Helpers.EnsureBot(gmailBot)
	if err != nil {
		p.API.LogError("Failed to ensure gmail bot ", "err", err.Error())
		return errors.Wrap(err, "Failed to ensure gmail bot")
	}

	p.API.LogInfo("Gmail Bot ensured")

	// Store created ID in Plugin struct
	p.gmailBotID = gmailBotID

	// Get the plugin file path
	bundlePath, err := p.API.GetBundlePath()
	if err != nil {
		return errors.Wrap(err, "Could not get bundle path")
	}

	botProfileImageName := "profile-image.png"

	// Retrieve Bot profile image from assets file folder
	botProfileImage, err := ioutil.ReadFile(filepath.Join(bundlePath, "assets", botProfileImageName))
	if err != nil {
		return errors.Wrap(err, "Could not get the profile image")
	}

	// Set the profile image to bot via API
	errInSetProfileImage := p.API.SetProfileImage(gmailBotID, botProfileImage)
	if errInSetProfileImage != nil {
		return errors.Wrap(err, "Could not set the profile image")
	}

	return nil
}
