package main

import "fmt"

func main() {

	var selectOperation int = 0
	productFamilyMap := make(map[string][]string)
	productFamilyMap["SR"] = []string{"7750", "7250"}
	productFamilyMap["SAR"] = []string{"7705"}
	productFamilyMap["OMNI"] = []string{"6900"}

	fmt.Println("The Family before Modifying", productFamilyMap)

	fmt.Println("Enter Operation: ")
	fmt.Scanln(&selectOperation)

	switch selectOperation {
	case 1:
		addProductFamily(productFamilyMap)
	case 2:
		deleteProductFamily(productFamilyMap)
	case 3:
		addNodeToFamily(productFamilyMap)
	case 4:
		deleteNodeFromFamily(productFamilyMap)
	default:
		fmt.Println("Invalid Option")
	}

}

func addProductFamily(productFamilyMap map[string][]string) {
	var newFamily string
	fmt.Println("Enter new Family")
	fmt.Scanln(&newFamily)

	_, ok := productFamilyMap[newFamily]
	if !ok {
		productFamilyMap[newFamily] = make([]string, 0, 20)
		fmt.Println("The Family Map After adding Family", productFamilyMap)

	}
}

func deleteProductFamily(productFamilyMap map[string][]string) {
	var depricatedfamily string
	fmt.Println("Enter Family to be deleted")
	fmt.Scanln(&depricatedfamily)

	_, ok := productFamilyMap[depricatedfamily]
	if ok {
		delete(productFamilyMap, depricatedfamily)
		fmt.Println("The Family Map After deleting Family", productFamilyMap)
	}
}

func addNodeToFamily(productFamilyMap map[string][]string) {
	var productFamily string
	var nodeType string

	fmt.Println("Enter productFamily")
	fmt.Scanln(&productFamily)

	fmt.Println("Enter Node type")
	fmt.Scanln(&nodeType)

	_, ok := productFamilyMap[productFamily]
	if !ok {
		productFamilyMap[productFamily] = make([]string, 0, 20)
		productFamilyMap[productFamily] = append(productFamilyMap[productFamily], nodeType)
		fmt.Println("The Family Map After adding Family", productFamilyMap)
	} else {

		fmt.Println("Family is already present")
		fmt.Println("The Family Map is", productFamilyMap)
	}

}

func deleteNodeFromFamily(productFamilyMap map[string][]string) {
	var deleteFamily string
	fmt.Println("Enter the Family from which the node to be deleted")
	fmt.Scanln(&deleteFamily)

	var deleteNodeFromFamily string
	fmt.Println("Enter the node to be deleted")
	fmt.Scanln(&deleteNodeFromFamily)

	var deleteNodeIndex int
	var nodeTypeSlice []string

	_, ok := productFamilyMap[deleteFamily]
	if !ok {
		fmt.Println("No such Family")
		fmt.Println("The Family Map is", productFamilyMap)
	} else {
		nodeTypeSlice = productFamilyMap[deleteFamily]

		for index, nodeType := range nodeTypeSlice {
			if nodeType == deleteNodeFromFamily {
				deleteNodeIndex = index
			}
		}
	}
	productFamilyMap[deleteFamily] = append(nodeTypeSlice[:deleteNodeIndex], nodeTypeSlice[deleteNodeIndex+1:]...)
	fmt.Println("The Family Map is", productFamilyMap)
}
