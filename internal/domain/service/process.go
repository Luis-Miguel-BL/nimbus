package service

import "github.com/Luis-Miguel-BL/nimbus/internal/adapter"

type ProcessService struct {
	targetAdapter adapter.TargetAdapter[any]
}
