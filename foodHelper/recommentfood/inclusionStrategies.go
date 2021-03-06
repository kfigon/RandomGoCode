package recommentfood

type inclusionStrategy interface {
	shouldBeIncluded(usersIngredients *set, requiredIngredients *set) bool
	calcFitness(usersIngredients *set, requiredIngredients *set) int
}

type fitnessInclusionStrategy struct {
	persentThreshold int
}

func (f fitnessInclusionStrategy) shouldBeIncluded(usersIngredients *set, requiredIngredients *set) bool {
	fit := f.calcCommonAndFitness(usersIngredients, requiredIngredients)
	return fit >= f.persentThreshold
}

func (f fitnessInclusionStrategy) calcFitness(usersIngredients *set, requiredIngredients *set) int {
	return f.calcCommonAndFitness(usersIngredients, requiredIngredients)
}

func (f fitnessInclusionStrategy) calcCommonAndFitness(users *set, required *set) int {
	commonIngredients := users.intersection(required)
	return calcFitnessFun(commonIngredients, required)
}

func calcFitnessFun(commonIngredients *set, required *set) int {
	licznik := float64((commonIngredients.size()))
	mianownik := float64((required.size()))
	res := licznik / mianownik
	return int(res * 100)
}
