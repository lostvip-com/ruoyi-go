/*******************************************************************************
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

type HomePageRequest struct {
}

type HomePageResponse struct {
	PageInfo struct {
		Product struct {
			Total uint32 `json:"total"`
			Self  uint32 `json:"self"`
			Other uint32 `json:"other"`
		} `json:"product"`
		Device struct {
			Total uint32 `json:"total"`
			Self  uint32 `json:"self"`
			Other uint32 `json:"other"`
		} `json:"device"`
		CloudInstance struct {
			Count     uint32 `json:"count"`
			RunCount  uint32 `json:"run_count"`
			StopCount uint32 `json:"stop_count"`
			//InstanceName string `json:"instanceName"`
			//Status       string `json:"status"`
		} `json:"cloudInstance"`
		Alert struct {
			Total uint32 `json:"total"`
		} `json:"alert"`
	} `json:"pageInfo"`
	QuickNavigation []QuickNavigation         `json:"quickNavigation"`
	Docs            Docs                      `json:"docs"`
	AlertPlate      []AlertPlateQueryResponse `json:"alertPlate"`
	MsgGather       []MsgGather               `json:"msg_gather"`
}

type PlatformInfo struct {
	VersionName string `json:"version_name"`
	DbName      string `json:"db_name"`
}

type NetWorkInfo struct {
	NewWork []NewWork `json:"newWork"`
	Dns     string    `json:"dns"`
}
type NewWork struct {
	NcId    string `json:"ncId"`
	LocalIp string `json:"localIp,omitempty"`
	GwIp    string `json:"gwIp,omitempty"`
	SmIp    string `json:"smIp,omitempty"`
	Netlink bool   `json:"netlink"`
}

type QuickNavigation struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	//JumpLink string `json:"jumpLink,omitempty"`
}

type Doc struct {
	Name     string `json:"name"`
	JumpLink string `json:"jumpLink,omitempty"`
}

type Docs struct {
	More string `json:"more"`
	Doc  []Doc  `json:"info"`
}

type Alert struct {
	Count      int    `json:"count"`
	AlertLevel string `json:"alert_level"`
}

type MsgGather struct {
	Count int    `json:"count"`
	Date  string `json:"date"`
}
