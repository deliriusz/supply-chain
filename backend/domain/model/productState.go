package model

import "fmt"

type ProductState int64

const (
	InProduction ProductState = iota
	Created
	Payed
	Delivered
)

func (s ProductState) String() string {
	switch s {
	case InProduction:
		return "IN_PRODUCTION"
	case Created:
		return "CREATED"
	case Payed:
		return "PAYED"
	case Delivered:
		return "DELIVERED"
	}

	return "INVALID_STATE"
}

func ProductStateFromString(strState string) (ProductState, error) {
	var state ProductState
	switch strState {
	case "IN_PRODUCTION":
		state = InProduction
	case "CREATED":
		state = Created
	case "PAYED":
		state = Payed
	case "DELIVERED":
		state = Delivered
	default:
		return state, fmt.Errorf("passed state is invalid")
	}

	return state, nil
}
