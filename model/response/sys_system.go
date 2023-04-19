package response

import (
	"cms/common/config"
)

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
