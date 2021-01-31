package main

type inclusionStrategy interface {
	shouldBeIncluded(usersIngredients *set, requiredIngredients *set) bool
	calcFitness(usersIngredients *set, requiredIngredients *set) int
}

type fitnessInclusionStrategy struct {
	persentThreshold int
}

func (f fitnessInclusionStrategy) shouldBeIncluded(usersIngredients *set, requiredIngredients *set) bool {
	commonIngredients := usersIngredients.intersection(requiredIngredients)
	fit := calcFitnessFun(commonIngredients, requiredIngredients)
	return fit >= f.persentThreshold
}

func (f fitnessInclusionStrategy) calcFitness(usersIngredients *set, requiredIngredients *set) int {
	commonIngredients := usersIngredients.intersection(requiredIngredients)
	return calcFitnessFun(commonIngredients, requiredIngredients)
}

func calcFitnessFun(commonIngredients *set, required *set) int {
	return commonIngredients.size()/required.size() * 100
}