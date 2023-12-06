package target

import (
	"github.com/Luis-Miguel-BL/nimbus/internal/domain/model"
	"github.com/Luis-Miguel-BL/nimbus/internal/domain/vo"
)

const ProcessTargetTypeSQS model.TargetType = "sqs-queue"

type SQSTargetAdapter struct {
	sqsClient any
}

type SQSTargetDetails struct {
	QueueUrl string
}

func (t *SQSTargetAdapter) Send(fileItem vo.FileItem, targetDetails SQSTargetDetails) (err error) {
	return

}
