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

package dtos

import (
	"encoding/json"
	"github.com/winc-link/hummingbird/internal/models"
)

type StatsResp struct {
	MemoryStats MemoryStats `json:"memory_stats"`
	CpuStats    CpuStats    `json:"cpu_stats"`
	PrecpuStats CpuStats    `json:"precpu_stats"`
}

type CpuStats struct {
	CpuUsage       CpuUsage `json:"cpu_usage"`
	SystemCpuUsage int64    `json:"system_cpu_usage"`
	OnlineCpus     int64    `json:"online_cpus"`
}

type CpuUsage struct {
	PercpuUsage []int64 `json:"percpu_usage"`
	TotalUsage  int64   `json:"total_usage"`
}

type MemoryStats struct {
	Usage    int64             `json:"usage"`
	Limit    int64             `json:"limit"`
	MaxUsage int64             `json:"max_usage"`
	Stats    MemoryStatsDetail `json:"stats"`
}

type MemoryStatsDetail struct {
	Cache int64 `json:"cache"`
	Rss   int64 `json:"rss"`
}

func (s StatsResp) UsedMemory() int64 {
	return s.MemoryStats.Usage - s.MemoryStats.Stats.Cache
}

// MemoryUsage %
func (s StatsResp) MemoryUsage() float64 {
	if s.MemoryStats.Limit <= 0 {
		return 0
	}
	return float64(s.UsedMemory()/s.MemoryStats.Limit) * 100.0
}

func (s StatsResp) CpuDelta() float64 {
	return float64(s.CpuStats.CpuUsage.TotalUsage - s.PrecpuStats.CpuUsage.TotalUsage)
}

func (s StatsResp) SystemCpuDelta() float64 {
	return float64(s.CpuStats.SystemCpuUsage - s.PrecpuStats.SystemCpuUsage)
}

// CpuUsage %
func (s StatsResp) CpuUsage() float64 {
	scd := s.SystemCpuDelta()
	if scd <= 0 {
		return 0
	}
	return (s.CpuDelta() / scd) * float64(s.CpuStats.OnlineCpus) * 100.0
}

type SystemMetrics struct {
	Timestamp      int64                    `json:"timestamp"`        // 时间戳
	CpuUsedPercent float64                  `json:"cpu_used_percent"` // cpu 使用率百分比
	CpuAvg         float64                  `json:"cpu_avg"`          // cpu 负载，1分钟
	Memory         SystemMemory             `json:"memory"`           // 内存
	Disk           SystemDisk               `json:"disk"`             // 磁盘使用率
	Network        map[string]SystemNetwork `json:"network"`          // 网卡en/eth的IO
	Openfiles      int                      `json:"openfiles"`        // 文件数，linux 才有
}

type SystemMemory struct {
	Total       uint64  `json:"total"`        // 大小
	Used        uint64  `json:"used"`         // 使用大小 bytes
	UsedPercent float64 `json:"used_percent"` // 百分比
}

type SystemDisk struct {
	Path        string  `json:"path"`         // 获取 / 目录信息
	Total       uint64  `json:"total"`        // 大小  bytes
	Used        uint64  `json:"used"`         // 使用值
	UsedPercent float64 `json:"used_percent"` // 使用百分比
}

type SystemNetwork struct {
	Name         string `json:"name"`
	BytesSent    uint64 `json:"bytes_sent"`     // 总发送字节
	BytesRecv    uint64 `json:"bytes_recv"`     // 总接收字节
	BytesSentPre uint64 `json:"bytes_sent_pre"` // 单位时间内发送的字节，1分钟
	BytesRecvPre uint64 `json:"bytes_recv_pre"` // 单位时间内接收的字节，1分钟
	Last         int64  `json:"-"`              // 采集时记录，不做输出
}

func FromModelsSystemMetricsToDTO(m models.SystemMetrics) (SystemMetrics, error) {
	var s SystemMetrics
	if err := json.Unmarshal([]byte(m.Data), &s); err != nil {
		return SystemMetrics{}, err
	}
	return s, nil
}

func (s SystemNetwork) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s SystemMetrics) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

// Response

type SystemMetricsResponse struct {
	Total   int                  `json:"total"`
	Metrics []SystemStatResponse `json:"metrics"`
}

type SystemStatResponse struct {
	Timestamp         int64   `json:"timestamp"`           // 时间戳
	CpuUsedPercent    float64 `json:"cpu_used_percent"`    // cpu 使用率百分比
	MemoryTotal       uint64  `json:"memory_total"`        // 内存使用
	MemoryUsed        uint64  `json:"memory_used"`         // 内存使用
	MemoryUsedPercent float64 `json:"memory_used_percent"` // 内存使用率
	DiskTotal         uint64  `json:"disk_total"`
	DiskUsed          uint64  `json:"disk_used"`
	DiskUsedPercent   float64 `json:"disk_used_percent"` // 磁盘使用率
	NetSentBytes      uint64  `json:"net_sent_bytes"`    // 网卡发送字节
	NetRecvBytes      uint64  `json:"net_recv_bytes"`    // 网卡接收字节
	Openfiles         int     `json:"openfiles"`         // 文件句柄数，linux 才有
}
