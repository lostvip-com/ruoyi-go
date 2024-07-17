//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/json"
	//"gitlab.com/tedge/edgex/internal/models"
)

// swagger:response ServicesStats
type ServicesStats []ServiceStats

func (s ServicesStats) Len() int {
	return len(s)
}

func (s ServicesStats) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

func (s ServicesStats) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ServiceStats struct {
	Id          string `json:"id" binding:"required"` //
	Name        string `json:"name" binding:"required"`
	Healthy     bool   `json:"healthy"` // 健康状态
	Created     string `json:"created"`
	LogPath     string `json:"log_path"` // 日志地址, 宿主主机日志路径
	Started     string `json:"started"`  // 服务最近启动时间
	ServiceType string `json:"service_type" binding:"required"`
}

type Logging struct {
	Log string `json:"log"`
}

type MetricsQuery struct {
	Service     string `form:"service" binding:"required"`                           // 服务ID
	MetricsType string `form:"metrics_type" binding:"oneof=minute hour halfday day"` // 监控类型，范围: minute hour halfday day
	MetricsRangeQuery
}

type SystemMetricsQuery struct {
	Iface       string `form:"iface"`
	MetricsType string `form:"metrics_type" binding:"oneof=hour halfday day"`
}

type MetricsRangeQuery struct {
	Start int64 `form:"start" binding:"gt=0"` // 开始时间戳
	End   int64 `form:"end" binding:"gt=0"`   // 结束时间戳
}

// swagger:response MetricsResult
type MetricsResult struct {
	Total   int       `json:"total"`
	Metrics []Metrics `json:"metrics"` // 性能点列表
}

type Metrics struct {
	Timestamp      int64   `json:"timestamp"`      // 时间戳
	CpuUsedPercent float64 `json:"cpuUsedPercent"` // cpu 使用率百分比
	MemoryUsed     int64   `json:"memoryUsed"`     // 内存使用大小，单位:字节
}

func (m Metrics) ToJSON() string {
	marshal, _ := json.Marshal(m)
	return string(marshal)
}

type LogParam struct {
	Line int `form:"line"`
}

type TerminalParams struct {
	Cmd            string   `json:"cmd" binding:"required"`
	Args           []string `json:"args" binding:"required"`
	TimeoutSeconds int      `json:"timeout_seconds" binding:"min=1,max=60"`
}

type AgentRequest struct {
	Cmd            string   `json:"cmd"`
	Args           []string `json:"args"`
	TimeoutSeconds int      `json:"timeout_seconds"`
}

type AgentResponse struct {
	Operation    string `json:"operation"`
	Service      string `json:"service"`
	Executor     string `json:"executor"`
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage"`
}
