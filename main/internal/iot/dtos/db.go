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

package dtos

type Configuration struct {
	DbType       string
	Host         string
	Port         string
	Timeout      int
	DatabaseName string
	Username     string
	Password     string
	BatchSize    int
	Dsn          string
	// 添加sqlite数据库存储地址
	DataSource string
	// 添加tqlite集群地址
	Cluster []string
}
