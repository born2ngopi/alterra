package response

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
)

type Meta struct {
	Success bool                `json:"success" default:"true"`
	Message string              `json:"message" default:"true"`
	Info    *dto.PaginationInfo `json:"info"`
}
