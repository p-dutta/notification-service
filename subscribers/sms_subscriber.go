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
	"time"

	"cloud.google.com/go/pubsub"
	"go.uber.org/zap"
)

func ProcessSMS(sub *pubsub.Subscription, ctx context.Context) {
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		select {
		case <-ctx.Done():
			// Context is cancelled or timed out
			zapLog := logger.GetForFile("send-sms")
			zapLog.Info("cancelled or timed-out",
				zap.String("cancel", "Context cancelled or timed out, stopping message processing"),
			)
			return
		default:
			// Process message
			log.Printf("SMSSubscriber: Received message: %s\n", string(msg.Data))
			var pubSubMsg utils.PubSubMessage
			if err := json.Unmarshal(msg.Data, &pubSubMsg); err != nil {
				log.Printf("Error unmarshalling JSON: %v", err)
				msg.Ack()
				return
			}

			smsTemplate := `Submit OTP :OTP_HERE to enter Toffee’s exciting world. This OTP will expire within 5 minutes. X7n5ycNnKFa`
			smsTemplate = strings.Replace(smsTemplate, "OTP_HERE", pubSubMsg.OTP, 1)
			currentTime := time.Now()
			formattedTime := currentTime.Format("2006-01-02 15:04:05")

			err := notificationSender.SendSms(os.Getenv("SMS_API"), os.Getenv("SMS_USER"), os.Getenv("SMS_PASSWORD"), "+"+pubSubMsg.Target, smsTemplate, os.Getenv("SMS_SID"), formattedTime, "1")

			if err != nil {
				zapLog := logger.GetForFile("send-sms")
				zapLog.Error("Could not send sms",
					zap.Error(errors.New("OTP could not be sent to: "+pubSubMsg.Target)),
				)
			} else {
				msg.Ack()
			}
		}

	})

	if err != nil {
		zapLog := logger.GetForFile("send-sms")
		zapLog.Error("SMSSubscriber: Error receiving messages: %v\n",
			zap.Error(err),
		)
		return
	}

}
