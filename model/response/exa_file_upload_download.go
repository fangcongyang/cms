package response

import "cms/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
