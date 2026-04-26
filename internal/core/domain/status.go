package domain

const (
	StatusCreated            = "created"
	StatusCanceled           = "canceled"
	StatusCooking            = "cooking"
	StatusWaitingForCustomer = "waiting for customer"
	StatusCompleted          = "completed"
)

var AllowedTransitions = map[string][]string{
	StatusCreated:            {StatusCanceled, StatusCooking},
	StatusCooking:            {StatusWaitingForCustomer, StatusCanceled},
	StatusWaitingForCustomer: {StatusCompleted},
	StatusCompleted:          {},
	StatusCanceled:           {},
}
