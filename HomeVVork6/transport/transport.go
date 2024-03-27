package transport

type PublicTransport interface {
	GetType() string
	BoardPassengers()
	DisembarkPassengers()
}

type Bus struct{}

func (b Bus) GetType() string {
	return "Автобус"
}

func (b Bus) BoardPassengers() {
	println("Пасажири сідають у автобус")
}

func (b Bus) DisembarkPassengers() {
	println("Пасажири виходять з автобуса")
}

type Train struct{}

func (t Train) GetType() string {
	return "Потяг"
}

func (t Train) BoardPassengers() {
	println("Пасажири сідають у потяг")
}

func (t Train) DisembarkPassengers() {
	println("Пасажири виходять з потягу")
}

type Plane struct{}

func (p Plane) GetType() string {
	return "Літак"
}

func (p Plane) BoardPassengers() {
	println("Пасажири сідають у літак")
}

func (p Plane) DisembarkPassengers() {
	println("Пасажири виходять з літака")
}

