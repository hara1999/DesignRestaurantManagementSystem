package main

import "errors"

type MenuType string

const (
	MainCourse MenuType = "MainCourse"
	Starter    MenuType = "Starter"
)

type Menu struct {
	Type MenuType
	Dish []*Dish
}

func NewMenu(Type MenuType) *Menu {
	return &Menu{
		Type: Type,
		Dish: make([]*Dish, 0),
	}
}

func (m *Menu) AddDish(dish *Dish) {
	m.Dish = append(m.Dish, dish)
}

func (m *Menu) GetDishByName(dishName string) (*Dish, error) {
	for _, dish := range m.Dish {
		if dish.DishName == dishName {
			return dish, nil
		}
	}

	return nil, errors.New("no dish found")
}
