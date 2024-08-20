package types

// TopicChannel is the data structure for subscriber
type TopicChannel struct {
	// Topic for subscriber to filter on if any
	Topic string
	// ConsumerGroup
	Group string
	// Messages is the returned message channel for the subscriber
	Messages chan MessageEnvelope
}

// MessageBusConfig defines the messaging information need to connect to the message bus
// in a publish-subscribe pattern
type MessageBusConfig struct {
	// PublishHost contains the connection information for a publishing on 0mq
	PublishHost HostInfo
	// SubscribeHost contains the connection information for a subscribing on 0mq
	SubscribeHost HostInfo
	// Type indicates the message queue platform being used. eg. "zero" for 0mq
	Type string
	// Optional contains all other properties of message bus that is specific to
	// certain concrete implementation like MQTT's QoS, for example
	Optional map[string]string
}
