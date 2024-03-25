package route

import "awesomeProject/transport" 

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
		switch t.(type) {
		case transport.Bus:
			println("Тип: Автобус")
		case transport.Train:
			println("Тип: Потяг")
		case transport.Plane:
			println("Тип: Літак")
		default:
			println("Тип: Невідомий")
		}
	}
}

