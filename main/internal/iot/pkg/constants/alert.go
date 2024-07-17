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

package constants

type Trigger string

const (
	DeviceDataTrigger   Trigger = "设备数据触发"
	DeviceEventTrigger  Trigger = "设备事件触发"
	DeviceStatusTrigger Trigger = "设备状态触发"
)

type RuleStatus string

const (
	RuleStart RuleStatus = "running"
	RuleStop  RuleStatus = "stopped"
)

type AlertType string

const DeviceAlertType AlertType = "设备告警"

type AlertLevel string

const (
	Urgent        AlertLevel = "紧急"
	Important     AlertLevel = "重要"
	LessImportant AlertLevel = "次要"
	Remind        AlertLevel = "提示"
)

type AlertListStatus string

const (
	Ignore    AlertListStatus = "忽略"
	Treated   AlertListStatus = "已处理"
	Untreated AlertListStatus = "未处理"
)

type WorkerCondition string

const (
	WorkerConditionAnyone WorkerCondition = "anyone"
	WorkerConditionAll    WorkerCondition = "all"
)

type AlertWay string

const (
	SMS      AlertWay = "sms"
	PHONE    AlertWay = "语音告警"
	QYweixin AlertWay = "企业微信机器人"
	DingDing AlertWay = "钉钉机器人"
	FeiShu   AlertWay = "飞书机器人"
	WEBAPI   AlertWay = "API接口"
)

func GetAlertWays() []string {
	return []string{string(SMS), string(PHONE), string(QYweixin), string(DingDing), string(FeiShu), string(WEBAPI)}
}

const (
	Original = "original"
	Avg      = "avg"
	Max      = "max"
	Min      = "min"
	Sum      = "sum"
)

var (
	ValueTypes = []string{Original, Avg, Max, Min, Sum}
)

var (
	DecideConditions = []string{">", ">=", "<", "<=", "=", "!="}
)
