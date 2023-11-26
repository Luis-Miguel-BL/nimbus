package domain

type AggregateID string
type AggregateType string

type AggregateRoot struct {
	aggregateID   AggregateID
	aggregateType AggregateType
}

func NewAggregateRoot(aggregateType AggregateType, aggregateID AggregateID) *AggregateRoot {
	return &AggregateRoot{
		aggregateID:   aggregateID,
		aggregateType: aggregateType,
	}
}
func (a *AggregateRoot) AggregateID() AggregateID {
	return a.aggregateID
}

func (a *AggregateRoot) AggregateType() AggregateType {
	return a.aggregateType
}
