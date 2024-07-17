package dtos

import (
	"context"
)

type ConnectHandler func(ctx context.Context)

type CallbackHandler func(context.Context, CallbackMessage)

type CallbackMessage struct {
	Error error
}

type NewMQTTClient struct {
	Broker    string   `json:"broker"`
	ClientId  string   `json:"client_id"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	PubTopic  string   `json:"pub_topic"`
	SubTopics []string `json:"sub_topics"`
}
