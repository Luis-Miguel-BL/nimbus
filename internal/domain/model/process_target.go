package model

import (
	"github.com/Luis-Miguel-BL/nimbus/internal/adapter"
)

type ProcessTargetID string
type ProcessTarget struct {
	processTargetID ProcessTargetID
	targetType      TargetType
	targetDetails   adapter.TargetDetails
}

type TargetType string
