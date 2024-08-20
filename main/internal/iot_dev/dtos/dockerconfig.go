/*******************************************************************************
 * Copyright 2017 Dell Inc.
 * Copyright (c) 2019 Intel Corporation
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
	"main/internal/iot_dev/pkg/constants"
)

type DockerConfig struct {
	Id       string
	Address  string
	Account  string
	Password string
	SaltKey  string
}

func WincLinkDockerConfig() DockerConfig {
	return DockerConfig{
		Address:  constants.Address,
		Password: constants.Password,
		Account:  constants.Account,
		SaltKey:  constants.SaltKey,
	}
}

type DockerConfigAddRequest struct {
	Id       string `json:"id"`
	Address  string `json:"address" binding:"required"` // 仓库地址 true
	Account  string `json:"account,omitempty"`          // 账户
	Password string `json:"password,omitempty"`         // 密码
}

type DockerConfigSearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	Address                  string `schema:"address"`
	Account                  string `schema:"account"`
}

type DockerConfigResponse struct {
	Id      string `json:"id"`
	Address string `json:"address"`
	Account string `json:"account"`
}

type DockerConfigUpdateRequest struct {
	Id       string  `json:"id" binding:"required"`
	Address  *string `json:"address"`
	Account  *string `json:"account"`
	Password *string `json:"password"`
}
