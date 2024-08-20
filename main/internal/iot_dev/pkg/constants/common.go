/*******************************************************************************
 * Copyright 2023 Winc link Inc.
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
package constants

const (
	ApiVersion = "v2"
	ApiBase    = "/api/v2"
)

type InstanceType string

const (
	CloudInstance  InstanceType = "cloudInstanceService"
	DriverInstance InstanceType = "driverInstanceService"
)

type ResourceType string

const (
	DriverResource  ResourceType = "driver"
	DeviceResource  ResourceType = "device"
	ServiceResource ResourceType = "service"
	OtherResource   ResourceType = "other"
)

// Constants related to defined url path names and parameters in the v2 service APIs
const (
	All         = "all"
	Id          = "id"
	Created     = "created"
	Modified    = "modified"
	Pushed      = "pushed"
	Count       = "count"
	Device      = "device"
	DeviceId    = "deviceId"
	DeviceName  = "deviceName"
	Check       = "check"
	Product     = "product"
	ProductId   = "productId"
	Service     = "service"
	Command     = "command"
	ProductName = "productName"
	ServiceName = "serviceName"
	//ResourceName = "resourceName"
	ResourceId   = "resourceId"
	Start        = "start"
	End          = "end"
	Age          = "age"
	Scrub        = "scrub"
	Type         = "type"
	Name         = "name"
	Label        = "label"
	Manufacturer = "manufacturer"
	Model        = "model"
	ValueType    = "valueType"
	Offset       = "offset"         //query string to specify the number of items to skip before starting to collect the result set.
	Limit        = "limit"          //query string to specify the numbers of items to return
	Labels       = "labels"         //query string to specify associated user-defined labels for querying a given object. More than one label may be specified via a comma-delimited list
	PushEvent    = "ds-pushevent"   //query string to specify if an event should be pushed to the EdgeX system
	ReturnEvent  = "ds-returnevent" //query string to specify if an event should be returned from device service
	Search       = "search"
	MarkCode     = "markCode" //标示符
	Status       = "status"
	Exist        = "exist"
	FuncPointId  = "funcPointId"
)

const (
	BootTimeoutDefault        = BootTimeoutSecondsDefault * 1000
	BootTimeoutSecondsDefault = 30
	BootRetrySecondsDefault   = 1
	ConfigFileName            = "configuration.toml"
	ConfigStemCore            = "hummingbird/core/"
	ConfigMajorVersion        = "1.0/"
	LogDurationKey            = "duration"
)

const (
	CorrelationHeader = "X-Correlation-ID" // Sets the key of the Correlation ID HTTP header
)

const (
	CoreServiceKey = "hummingbird-core"
)

type MetricsType string

// 性能采集监控类型
const (
	HourMetricsType    = "hour"
	HalfDayMetricsType = "halfday"
	DayMetricsType     = "day"
)

func (m MetricsType) String() string {
	return string(m)
}
