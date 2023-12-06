package adapter

import "github.com/Luis-Miguel-BL/nimbus/internal/domain/vo"

type TargetAdapter[T TargetDetails] interface {
	Send(vo.FileItem, T) error
}

type TargetDetails any
