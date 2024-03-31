package main

import (
	"fmt"

	"awesomeProject/route"
	"awesomeProject/transport"
)

type Passenger struct{}

func (p Passenger) Travel(r *route.Route) {
	fmt.Println("Початок подорожі:")
	r.ShowTransports()

	for _, t := range r.Transports {
		fmt.Println("-------------------------")
		fmt.Printf("Подорож на %s:\n", t.GetType())
		t.BoardPassengers()
		time.Sleep(2 * time.Second)
		t.DisembarkPassengers()
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Кінець подорожі")
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


