package recommendfood

type mockDb struct {
	findFoodFun func() []food
}

func (m mockDb) findFoods() []food {
	return m.findFoodFun()
}

const (
	egg     int = 1
	chicken     = 2
	beef        = 3
	salmon      = 4
	salad       = 5
	cheese      = 6
	apple       = 7
	noodle      = 80
	bread       = 90
)

var mockedFoods = []food{
	{"first", []int{egg, chicken, salmon}},
	{"second", []int{egg, chicken, salad}},
	{"third", []int{salad, cheese, apple}},
}

func createMockDb() mockDb {
	return mockDb{
		findFoodFun: func() []food {
			return mockedFoods
		},
	}
}

type ingredientsMock struct {
	getIdFun func(string) (int, bool)
}

func (i ingredientsMock) getId(name string) (int, bool) {
	if i.getIdFun == nil {
		return 0, false
	}
	return i.getIdFun(name)
}
