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
		return "InProduction"
	case Created:
		return "Created"
	case Payed:
		return "Payed"
	case Delivered:
		return "Delivered"
	}

	return "InvalidState"
}

func ProductStateFromString(strState string) (ProductState, error) {
	var state ProductState
	switch strState {
	case "InProduction":
		state = InProduction
	case "Created":
		state = Created
	case "Payed":
		state = Payed
	case "Delivered":
		state = Delivered
	default:
		return state, fmt.Errorf("passed state is invalid")
	}

	return state, nil
}
