package dtos

import (
	//"gitlab.com/tedge/edgex/internal/models"
	//"gitlab.com/tedge/edgex/proto/devicelibrary"
	"github.com/winc-link/hummingbird/internal/models"
)

type DeviceLibrarySupportVersionSimple struct {
	Version    string `json:"version"`
	IsDefault  bool   `json:"is_default"`
	ConfigFile string `json:"config_file"`
}

func DeviceLibrarySupportVersionSimpleFromModel(versions models.SupportVersions) []DeviceLibrarySupportVersionSimple {
	ret := make([]DeviceLibrarySupportVersionSimple, len(versions))
	for i, v := range versions {
		ret[i] = DeviceLibrarySupportVersionSimple{
			Version:   v.Version,
			IsDefault: v.IsDefault,
			//ConfigFile: v.ConfigFile,
		}
	}
	return ret
}
