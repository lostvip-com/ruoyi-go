package types

import "fmt"

const (
	defaultMsgProtocol = "tcp"
)

// HostInfo is the URL information of the host as the following scheme:
// <Protocol>://<Host>:<Port>
type HostInfo struct {
	// Host is the hostname or IP address of the messaging broker, if applicable.
	Host string
	// Port defines the port on which to access the message queue.
	Port int
	// Protocol indicates the protocol to use when accessing the message queue.
	Protocol string
}

// GetHostURL returns the complete URL for the host-info configuration
func (info *HostInfo) GetHostURL() string {

	protocol := info.Protocol
	if info.Protocol == "" {
		protocol = defaultMsgProtocol
	}
	return fmt.Sprintf("%s://%s:%d", protocol, info.Host, info.Port)
}

// IsHostInfoEmpty returns whether the host-info has been initialized or not
func (info *HostInfo) IsHostInfoEmpty() bool {
	return info == nil || info.Host == "" || info.Port == 0
}
