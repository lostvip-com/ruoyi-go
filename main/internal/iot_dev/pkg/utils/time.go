//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"time"
)

func FormatDockerTime(value string) (*time.Time, error) {
	if value == "" {
		return nil, nil
	}

	parse, err := parseTime(value)
	if err != nil {
		return nil, err
	}
	t := formatZoneTime(parse)
	return &t, nil
}

func parseTime(value string) (time.Time, error) {
	// capture of 2021-07-14T07:49:24.050553098Z
	return time.Parse("2006-01-02T15:04:05", value[:19])
}

func formatZoneTime(created time.Time) time.Time {
	// zone UTC to CST
	return created.Add(8 * time.Hour)
}

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 获取当天0时的时间戳
func Get0ClockTimeStamp(d time.Time) int64 {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local).UnixNano() / int64(time.Millisecond)
}

// 生成当前时间戳
func GetNowTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}
