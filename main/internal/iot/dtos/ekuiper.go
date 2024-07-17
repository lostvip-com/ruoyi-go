/*******************************************************************************
 * Copyright 2017.
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

package dtos

type GetRuleInfoResponse struct {
	Triggered bool                     `json:"triggered"`
	Id        string                   `json:"id"`
	Name      string                   `json:"name"`
	Sql       string                   `json:"sql"`
	Actions   []map[string]interface{} `json:"actions"`
	Options   map[string]interface{}   `json:"options"`
}

type GetRuleStatusResponse struct {
	Status string `json:"status"`
}

type CreateRule struct {
	Triggered bool      `json:"triggered"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Sql       string    `json:"sql"`
	Actions   []Actions `json:"actions"`
}

type Actions struct {
	Rest     map[string]interface{} `json:"rest,omitempty"`
	MQTT     map[string]interface{} `json:"mqtt,omitempty"`
	Kafka    map[string]interface{} `json:"kafka,omitempty"`
	Zmq      map[string]interface{} `json:"zmq,omitempty"`
	Redis    map[string]interface{} `json:"redis,omitempty"`
	Influx   map[string]interface{} `json:"influx,omitempty"`
	Tdengine map[string]interface{} `json:"tdengine,omitempty"`
}

type Rest struct {
	Method       string `json:"method,omitempty"`
	Url          string `json:"url,omitempty"`
	BodyType     string `json:"bodyType,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	RunAsync     bool   `json:"runAsync,omitempty"`
	OmitIfEmpty  bool   `json:"omitIfEmpty,omitempty"`
	SendSingle   bool   `json:"sendSingle,omitempty"`
	BufferLength int    `json:"bufferLength,omitempty"`
	EnableCache  bool   `json:"enableCache,omitempty"`
	Format       string `json:"format,omitempty"`
}

type MQTT struct {
	Server             string `json:"server,omitempty"`
	Topic              string `json:"topic,omitempty"`
	ClientId           string `json:"clientId,omitempty"`
	ProtocolVersion    string `json:"protocolVersion,omitempty"`
	Qos                int    `json:"qos,omitempty"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	CertificationPath  string `json:"certificationPath,omitempty"`
	PrivateKeyPath     string `json:"privateKeyPath,omitempty"`
	RootCaPath         string `json:"rootCaPath,omitempty"`
	InsecureSkipVerify string `json:"insecureSkipVerify,omitempty"`
	Retained           bool   `json:"retained,omitempty"`
	//compression        string `json:"compression"`
}

func (b CreateRule) BuildCreateRuleParam(actions []Actions, ruleId, sql string) CreateRule {
	b.Triggered = false
	b.Id = ruleId
	b.Name = ruleId
	b.Sql = sql
	b.Actions = actions
	return b
}

func GetRuleAlertEkuiperActions(actionUrl string) []Actions {
	var a []Actions
	rest := make(map[string]interface{})
	rest["method"] = "POST"
	//bug-fix
	rest["url"] = "http://hummingbird-core:58081" + "/api/v1/ekuiper/alert"
	rest["bodyType"] = "json"
	rest["timeout"] = 5000
	rest["runAsync"] = false
	rest["omitIfEmpty"] = false
	rest["sendSingle"] = true
	rest["enableCache"] = false
	rest["format"] = "json"
	a = append(a, Actions{
		Rest: rest,
	})
	return a
}

func GetRuleSceneEkuiperActions(actionUrl string) []Actions {
	var a []Actions
	rest := make(map[string]interface{})
	rest["method"] = "POST"
	rest["url"] = "http://hummingbird-core:58081" + "/api/v1/ekuiper/scene"
	rest["bodyType"] = "json"
	rest["timeout"] = 5000
	rest["runAsync"] = false
	rest["omitIfEmpty"] = false
	rest["sendSingle"] = true
	rest["enableCache"] = false
	rest["format"] = "json"
	a = append(a, Actions{
		Rest: rest,
	})
	return a
}
