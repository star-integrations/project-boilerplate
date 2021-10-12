package log2slack

import (
	"context"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/star-integrations/project-boilerplate/back/pkg/config"
)

// Do - log to slack
func Do(_ context.Context, cfg *config.Config, attachments []slack.Attachment) error {
	payload := slack.Payload{
		Username:    cfg.SlackUserName,
		IconEmoji:   cfg.SlackIconEmoji,
		Attachments: attachments,
	}

	if errs := slack.Send(cfg.SlackWebhookURL, "", payload); len(errs) > 0 {
		return errs[0]
	}

	return nil
}
