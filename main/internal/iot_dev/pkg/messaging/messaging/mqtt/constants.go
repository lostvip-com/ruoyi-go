/********************************************************************************
 *  Copyright 2020 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

// Package mqttclient provides functionality useful for interacting with the MQTT client implementation.
package mqtt

const (
	// Constants for configuration properties provided via the MessageBusConfig's Optional field.

	// Client identifier configurations
	Username = "Username"
	Password = "Password"
	ClientId = "ClientId"

	// Connection configuration names
	Qos            = "Qos"
	KeepAlive      = "KeepAlive"
	Retained       = "Retained"
	AutoReconnect  = "AutoReconnect"
	ConnectTimeout = "ConnectTimeout"

	// TLS configuration names
	SkipCertVerify = "SkipCertVerify"
	CertFile       = "CertFile"
	KeyFile        = "KeyFile"
	KeyPEMBlock    = "KeyPEMBlock"
	CertPEMBlock   = "CertPEMBlock"
)
