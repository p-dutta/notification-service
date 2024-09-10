package subscribers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"notificationSubscriber/logger"
	"notificationSubscriber/notificationSender"
	templates "notificationSubscriber/templates"
	"notificationSubscriber/utils"
	"os"
	"strings"

	"cloud.google.com/go/pubsub"
	"go.uber.org/zap"
)

func ProcessEmail(sub *pubsub.Subscription, ctx context.Context) {
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {

		select {
		case <-ctx.Done():
			// Context is cancelled or timed out
			zapLog := logger.GetForFile("send-mail")
			zapLog.Info("cancelled or timed-out",
				zap.String("cancel", "Context cancelled or timed out, stopping message processing"),
			)
			return
		default:
			log.Printf("EmailSubscriber: Received message: %s\n", string(msg.Data))
			var pubSubMsg utils.PubSubMessage
			if err := json.Unmarshal(msg.Data, &pubSubMsg); err != nil {
				log.Printf("Error unmarshalling JSON: %v", err)
				msg.Ack() // Acknowledge the message to prevent redelivery
				return
			}

			htmlTemplate := strings.Replace(templates.SignInHtmlTemplate, "OTP_HERE", pubSubMsg.OTP, 1)
			err := notificationSender.SendEmail(os.Getenv("API_KEY"), "no-reply@toffeelive.com", "No reply", pubSubMsg.Target, "", "Toffee OTP", htmlTemplate)
			if err != nil {
				zapLog := logger.GetForFile("send-mail")
				zapLog.Error("Could not send email",
					zap.Error(errors.New("OTP could not be sent to: {{.pubSubMsg.Target}}")),
				)
			} else {
				msg.Ack()
			}
		}
	})
	if err != nil {
		zapLog := logger.GetForFile("send-mail")
		zapLog.Error("Error receiving email messages",
			zap.Error(err),
		)
	}
}
