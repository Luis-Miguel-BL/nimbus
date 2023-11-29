package model

type ProcessTargetID string
type ProcessTarget struct {
	processTargetID ProcessTargetID
	targetType      ProcessTargetType
	sqsDetails      TargetSQSDetails
	httpDetails     TargetHttpDetails
}

type ProcessTargetType string

const (
	ProcessTargetTypeSQS  ProcessTargetType = "sqs-queue"
	ProcessTargetTypeHttp ProcessTargetType = "http-endpoint"
)

type TargetSQSDetails struct {
	queueUrl string
}

type TargetHttpDetails struct {
	endpoint string
	method   string
}
