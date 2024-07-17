/*******************************************************************************
 * Copyright 2017 Dell Inc.
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

package models

type AlertType int

type AlertLevel int

type AlertContent struct {
	ServiceName string     `json:"name"`                        // 服务名
	Type        AlertType  `json:"type" binding:"oneof=1 2"`    // 告警类型
	Level       AlertLevel `json:"level" binding:"oneof=1 2 3"` // 告警级别
	T           int64      `json:"time"`                        // 告警时间
	Content     string     `json:"content"`                     // 告警内容
}

func (table *AlertContent) TableName() string {
	return "alert_content"
}

func (table *AlertContent) Get() interface{} {
	return *table
}
