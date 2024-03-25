package main

import "fmt"

type Animal struct {
	Name string
}

type Cage struct {
	Animal *Animal 
	Empty  bool    
}

func (c *Cage) AddAnimal(animal *Animal) {
	c.Animal = animal
	c.Empty = false 
}

func (c *Cage) InfoCage() {
	if c.Animal != nil {
		fmt.Printf("звіра: %s\n", c.Animal.Name)
	} else {
		fmt.Println("порожня")
	}
}

type Zookeeper struct {
	Name  string
	Cages []*Cage 
}

func (z *Zookeeper) CollectAnimal(animal *Animal, cageIndex int) {
	if cageIndex < len(z.Cages) {
		fmt.Printf("Клітка звіра %s ", animal.Name)
		z.Cages[cageIndex].InfoCage() 
		z.Cages[cageIndex].AddAnimal(animal)
		fmt.Printf("%s закриває %s у клітку\n", z.Name, animal.Name)
		z.Cages[cageIndex].InfoCage() 
	} else {
		fmt.Println("Немає такої клітки")
	}
}

func main() {
	fmt.Println("Звірі втекли з кліток!")

	cage1 := &Cage{Empty: true}
	cage2 := &Cage{Empty: true}
	cage3 := &Cage{Empty: true}
	cage4 := &Cage{Empty: true}
	cage5 := &Cage{Empty: true}

	zookeeper := &Zookeeper{
		Name:  "Максим",
		Cages: []*Cage{cage1, cage2, cage3, cage4, cage5},
	}

	animal1 := &Animal{"Лев"}
	animal2 := &Animal{"Тигр"}
	animal3 := &Animal{"Зебра"}
	animal4 := &Animal{"Жираф"}
	animal5 := &Animal{"Мавпа"}

	zookeeper.CollectAnimal(animal1, 0)
	zookeeper.CollectAnimal(animal2, 1)
	zookeeper.CollectAnimal(animal3, 2)
	zookeeper.CollectAnimal(animal4, 3)
	zookeeper.CollectAnimal(animal5, 4)
}
