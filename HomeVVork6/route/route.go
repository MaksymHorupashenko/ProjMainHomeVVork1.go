package route

import (
	"fmt"

	"awesomeProject/transport"
)

type PublicTransport interface {
	GetType() string
	BoardPassengers()
	DisembarkPassengers()
}

type Route struct {
	Transports []PublicTransport
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) AddTransport(t PublicTransport) {
	r.Transports = append(r.Transports, t)
}

func (r *Route) ShowTransports() {
	for i, t := range r.Transports {
		fmt.Printf("Транспортний засіб %d\n", i+1)
		fmt.Printf("Тип: %s\n", t.GetType())
	}
}
