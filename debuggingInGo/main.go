package main

import "fmt"

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	itemOpeningStock := 500
	itemSold := 100
	itemReturned := 50
	itemMissing := 5

	var itemClosingStockCalc int

	itemActualClosingStock := 445

	fmt.Println("Available Inventory (check): ", itemActualClosingStock)
	fmt.Println("-----------------------------")

	itemClosingStockCalc = finalCalculation(itemOpeningStock, itemSold, itemReturned, itemMissing)
	fmt.Println("Closing Stock Calculation: ", itemClosingStockCalc)

	if itemClosingStockCalc != itemActualClosingStock {
		fmt.Println("Warning: Calculation Error!")
	} else {
		fmt.Println("Great Work: All Good!")
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
