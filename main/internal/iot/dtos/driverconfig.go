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
	"main/internal/iot/pkg/constants"
	"main/internal/iot/pkg/utils"
	"sync"
)

var dcMutex sync.Mutex

type DriverConfigManage struct {
	HostRootDir string // 宿主主机dir
	NetWorkName string
	DockerManageConfig
}

type DockerManageConfig struct {
	ContainerConfigPath string // 容器内部驱动运行启动的配置文件路径
	DockerApiVersion    string // docker 版本号
	DockerRunMode       string // 默认运行模式 默认是host
	DockerSelfName      string // edge启动的容器名
	Privileged          bool   // 容器特权
}

func (m *DriverConfigManage) SetHostRootDir(dir string) {
	dcMutex.Lock()
	defer dcMutex.Unlock()
	m.HostRootDir = dir
}

func (m *DriverConfigManage) SetNetworkName(networkName string) {
	dcMutex.Lock()
	defer dcMutex.Unlock()
	m.NetWorkName = networkName
}

func (m *DriverConfigManage) GetLibraryDir() string {
	return utils.GetPwdDir() + "/" + constants.DriverLibraryDir + "/"
}

func (m *DriverConfigManage) GetBinPath(serverName string) string {
	return utils.GetPwdDir() + "/" + constants.DriverBinDir + "/" + serverName
}

func (m *DriverConfigManage) GetRunConfigPath(serviceName string) string {
	return constants.DockerHummingbirdRootDir + "/" + constants.DriverRunConfigDir + "/" + serviceName + constants.ConfigSuffix
}

func (m *DriverConfigManage) GetHostRunConfigPath(serviceName string) string {
	return m.HostRootDir + "/" + constants.DriverRunConfigDir + "/" + serviceName + constants.ConfigSuffix
}

func (m *DriverConfigManage) GetMntDir(serviceName string) string {
	return constants.DockerHummingbirdRootDir + "/" + constants.DriverMntDir + "/" + serviceName
}

func (m *DriverConfigManage) GetHostMntDir(serviceName string) string {
	return m.HostRootDir + "/" + constants.DriverMntDir + "/" + serviceName
}

func (m *DriverConfigManage) GetLogFilePath(serviceName string) string {
	return utils.GetPwdDir() + "/" + constants.DriverMntDir + "/" + serviceName + "/" + constants.DriverDefaultLogPath
}

func (m *DriverConfigManage) GetHostLogFilePath(serviceName string) string {
	return constants.DockerHummingbirdRootDir + "/" + constants.DriverMntDir + "/" + serviceName + "/" + constants.DriverDefaultLogPath
}
