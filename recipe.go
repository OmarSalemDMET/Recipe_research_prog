package main

// Define the struct for recipe
type recipe struct {
	Title              string
	Ingredients        string
	Instructions       string
	ImageName          string
	CleanedIngredients string
}

func fillRecipe(records [][]string) []recipe {
	var recipes []recipe
	for _, row := range records {
		recipe := recipe{
			Title:              row[0],
			Ingredients:        row[1],
			Instructions:       row[2],
			ImageName:          row[3],
			CleanedIngredients: row[4],
		}
		recipes = append(recipes, recipe)
	}
	return recipes
}

func chaeckRecipe(recipes recipe, ingredient string) bool {
	return SearchString(recipes.Ingredients, ingredient)

}

func SearchString(str1 string, str2 string) bool {
	len1 := len(str1)
	len2 := len(str2)

	if len1 > len2 {
		for i := 0; i < len1; i++ {
			if !(i+len2 > len1) {
				temp := str1[i : i+len2]
				if str2 == temp {
					return true
				}
			} else {
				return false
			}
		}
	} else {

		for i := 0; i < len2; i++ {
			if !(i+len1 > len2) {
				temp := str2[i : i+len1]
				if str1 == temp {
					return true
				}
			} else {
				return false
			}
		}
	}

	return false
}

func getRecipe(recipes *[]recipe, ingredient string) []recipe {

	var result []recipe

	for _, r := range *recipes {
		if chaeckRecipe(r, ingredient) {

			result = append(result, r)
		}

	}
	return result
}
