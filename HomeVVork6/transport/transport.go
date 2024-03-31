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
	fmt.Println("Пасажири сідають у автобус")
}

func (b Bus) DisembarkPassengers() {
	fmt.Println("Пасажири виходять з автобуса")
}

type Train struct{}

func (t Train) GetType() string {
	return "Потяг"
}

func (t Train) BoardPassengers() {
	fmt.Println("Пасажири сідають у потяг")
}

func (t Train) DisembarkPassengers() {
	fmt.Println("Пасажири виходять з потягу")
}

type Plane struct{}

func (p Plane) GetType() string {
	return "Літак"
}

func (p Plane) BoardPassengers() {
	fmt.Println("Пасажири сідають у літак")
}

func (p Plane) DisembarkPassengers() {
	fmt.Println("Пасажири виходять з літака")
}

