//
// Copyright (c) 2019 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package types

import (
	"context"
)

const (
	checksum      = "payload-checksum"
	correlationId = "X-Correlation-ID"
	contentType   = "Content-Type"
)

type PayloadType int

const (
	Invalid = iota
	Event
	DeviceStatus
	Command
	WithoutDPReport
)

type AckCommit func(id ...string) error

// MessageEnvelope is the data structure for messages. It wraps the generic message payload with attributes.
type MessageEnvelope struct {
	// message envelop id
	ID string
	// Checksum used to communicate to core-data that an Event message has been received via the message bus
	Checksum string
	// CorrelationID is an object id to identify the envelop
	CorrelationID string
	// Payload is byte representation of the data being transferred.
	Payload []byte
	// ContentType is the marshaled type of payload, i.e. application/json, application/xml, application/cbor, etc
	ContentType string
	//PayloadType is Payload's Object Type,.
	PayloadType PayloadType
	//
	AckCommit AckCommit `json:"-"`
}

// NewMessageEnvelope creates a new MessageEnvelope for the specified payload with attributes from the specified context
func NewMessageEnvelope(payloadType PayloadType, payload []byte, ctx context.Context) MessageEnvelope {
	envelope := MessageEnvelope{
		// TODO: Remove Checksum for V2.0 release.
		//       Also consider just passing correlationId & contentType as parameters instead of Context
		Checksum:      fromContext(ctx, checksum),
		CorrelationID: fromContext(ctx, correlationId),
		ContentType:   fromContext(ctx, contentType),
		Payload:       payload,
		PayloadType:   payloadType,
	}

	return envelope
}

func fromContext(ctx context.Context, key string) string {
	hdr, ok := ctx.Value(key).(string)
	if !ok {
		hdr = ""
	}
	return hdr
}
