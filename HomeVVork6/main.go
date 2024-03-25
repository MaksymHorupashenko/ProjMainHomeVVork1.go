package main

import (
	"awesomeProject/route"
	"awesomeProject/transport"
	"time"
)

type Passenger struct{}

func (p Passenger) Travel(r *route.Route) {
	println("Початок подорожі:")
	r.ShowTransports()

	for _, t := range r.Transports {
		println("-------------------------")
		switch t.(type) {
		case transport.Bus:
			println("Подорож на автобусі:")
		case transport.Train:
			println("Подорож на потязі:")
		case transport.Plane:
			println("Подорож на літаку:")
		default:
			println("Подорож на невідомому транспорті:")
		}

		t.BoardPassengers()
		time.Sleep(2 * time.Second)
		t.DisembarkPassengers()
		time.Sleep(2 * time.Second)
	}

	println("Кінець подорожі")
}

func main() {
	r := route.NewRoute()

	b := transport.Bus{}
	t := transport.Train{}
	p := transport.Plane{}
	r.AddTransport(b)
	r.AddTransport(t)
	r.AddTransport(p)

	passenger := Passenger{}

	passenger.Travel(r)
}

