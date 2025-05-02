package menu

import (
	"DesignRestaurantManagementSystem/models/dish"
	"errors"
)

type MenuType string

const (
	MainCourse MenuType = "MainCourse"
	Starter    MenuType = "Starter"
)

type Menu struct {
	Type MenuType
	Dish []*dish.Dish
}

func NewMenu(Type MenuType) *Menu {
	return &Menu{
		Type: Type,
		Dish: make([]*dish.Dish, 0),
	}
}

func (m *Menu) AddDish(dish *dish.Dish) {
	m.Dish = append(m.Dish, dish)
}

func (m *Menu) GetDishByName(dishName string) (*dish.Dish, error) {
	for _, dish := range m.Dish {
		if dish.DishName == dishName {
			return dish, nil
		}
	}

	return nil, errors.New("no dish found")
}

func (m *Menu) RemoveDish(dish *dish.Dish) {
	for i, d := range m.Dish {
		if d == dish {
			m.Dish = append(m.Dish[:i], m.Dish[i+1:]...)
		}
	}
}
