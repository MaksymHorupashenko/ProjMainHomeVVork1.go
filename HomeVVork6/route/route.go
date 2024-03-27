package route

import "awesomeProject/transport"

type PublicTransport interface {
	GetType() string
	BoardPassengers()
	DisembarkPassengers()
}

type Route struct {
	Transports []transport.PublicTransport
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) AddTransport(t transport.PublicTransport) {
	r.Transports = append(r.Transports, t)
}

func (r *Route) ShowTransports() {
	for i, t := range r.Transports {
		println("Транспортний засіб", i+1)
		println("Тип:", t.GetType())
	}
}
