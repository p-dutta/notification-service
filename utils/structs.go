package utils

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type Subscriber struct {
	ctx context.Context
	//done chan struct{}
}

type PubSubMessage struct {
	Target  string `json:"target"`
	OTP     string `json:"otp"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type AdminSmsPubSubMessage struct {
	Type            string `json:"type"`
	Target          string `json:"target"`
	OTP             string `json:"otp"`
	Message         string `json:"message"`
	VerificationUrl string `json:"verification_url"`
	UserName        string `json:"user_name"`
}

//attributes := make(map[string]string)

/*func NewSubscriber(ctx context.Context, done chan struct{}) *Subscriber {
	return &Subscriber{
		ctx:  ctx,
		done: done,
	}
}*/

func NewSubscriber(ctx context.Context) *Subscriber {
	return &Subscriber{
		ctx: ctx,
	}
}

func (subs *Subscriber) CreateSubscription(subscriptionID string, client *pubsub.Client) *pubsub.Subscription {
	sub := client.Subscription(subscriptionID)
	return sub
}
