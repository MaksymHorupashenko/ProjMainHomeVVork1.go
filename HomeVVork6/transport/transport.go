package transport

type PublicTransport interface {
	BoardPassengers()    
	DisembarkPassengers()
}

type Bus struct{}

func (b Bus) BoardPassengers() {
	println("Пасажири сідають у автобус")
}

func (b Bus) DisembarkPassengers() {
	println("Пасажири виходять з автобуса")
}

type Train struct{}

func (t Train) BoardPassengers() {
	println("Пасажири сідають у потяг")
}

func (t Train) DisembarkPassengers() {
	println("Пасажири виходять з потягу")
}

type Plane struct{}

func (p Plane) BoardPassengers() {
	println("Пасажири сідають у літак")
}

func (p Plane) DisembarkPassengers() {
	println("Пасажири виходять з літаку")
}

