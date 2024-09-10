package subscribers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"notificationSubscriber/logger"
	"notificationSubscriber/notificationSender"
	"notificationSubscriber/utils"
	"os"
	"strings"

	"notificationSubscriber/common"
	templates "notificationSubscriber/templates"

	"cloud.google.com/go/pubsub"
	"go.uber.org/zap"
)

func AdminProcessEmail(sub *pubsub.Subscription, ctx context.Context) {

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

			var pubSubMsg utils.PubSubMessage
			if err := json.Unmarshal(msg.Data, &pubSubMsg); err != nil {
				log.Printf("Error unmarshalling JSON: %v", err)
				msg.Ack() // Acknowledge the message to prevent redelivery
				return
			}
			var htmlTemplate string

			switch pubSubMsg.Type {
			case common.SmsAdmin:
				htmlTemplate = templates.SmsAdminHtmlTemplate
			case common.EmailVerification:
				htmlTemplate = templates.EmailVerificationHtmlTemplate
			case common.EmailAdmin:
				htmlTemplate = templates.EmailAdminHtmlTemplate
			case common.SignIn:
				htmlTemplate = strings.Replace(templates.AdminSignInHtmlTemplate, "OTP_HERE", pubSubMsg.OTP, 1)
			case common.ForgetPassword:
				htmlTemplate = templates.ForgetPasswordHtmlTemplate
			default:
				log.Printf("AdminEmailSubscriber: Received message type not found: %s\n", string(msg.Data))
			}

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
