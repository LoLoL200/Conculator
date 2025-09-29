package main

import (
	"fmt"
	"math"
)

func main() {
	// Basic meaning
	var startSum float64 = 5000.00
	var monthAccumulation float64 = 1500.00
	var bankYearsProcent float64 = 15
	var accumulatorYears int = 2

	// Action
	accumulatorMonth := accumulatorYears * 12         // transtlation in month
	bankMonthProcent := bankYearsProcent / (100 * 12) // transtlation in month bank procent

	// Print basic meaning
	const Title = "\n=== КАЛЬКУЛЯТОР НАКОПИЧЕНЬ ===\n"
	fmt.Printf(Title)
	fmt.Printf(" \n" + "Початкові дані:\n")
	fmt.Printf("● Початкова сума: %.2f грн\n", startSum)
	fmt.Printf("● Щомісячні накопичення: %.2f грн\n", monthAccumulation)
	fmt.Printf("● Річна відсоткова ставка банку: %.2f %s  \n", bankYearsProcent, "%")
	fmt.Printf("● Термін накопичень: %d роки (%d місяців) \n", accumulatorYears, accumulatorMonth)

	fmt.Printf(" \n" + "Результати:\n")
	// Total contributions
	totalContributions := startSum + (monthAccumulation * float64(accumulatorMonth))
	fmt.Printf("● Загальна сума внесків: %.2f грн \n", totalContributions)

	// Final summa
	compoundInitial := startSum * math.Pow(1+bankMonthProcent, float64(accumulatorMonth))                                        // Initial amount with compound interest
	contributionEffect := monthAccumulation * ((math.Pow(1+bankMonthProcent, float64(accumulatorMonth)) - 1) / bankMonthProcent) // Future value of monthly payments:
	finalAmount := compoundInitial + contributionEffect                                                                          // Final summa
	fmt.Printf("● Фінальна сума: %.2f грн\n", finalAmount)

	// Accrued interest
	insert := finalAmount - totalContributions
	fmt.Printf("● Нараховані відсотки %.2f грн"+" \n", insert)

	//Convert in dollar
	fmt.Printf("\nУ доларах США:\n")
	var dollar float64 = 38.5
	convertTotalContributionsDollar := totalContributions / dollar //Convert in dollar 'Total contributions'
	fmt.Printf("● Загальна сума внесків: %.2f $\n", convertTotalContributionsDollar)
	convertFinalAmountDollar := finalAmount / dollar //Convert in dollar 'Final summa'
	fmt.Printf("● Фінальна сума: %.2f $\n", convertFinalAmountDollar)

	//Convert in euro
	fmt.Printf("\nУ доларах Євро:\n")
	var euro float64 = 42.1
	convertTotalContributionsEuro := totalContributions / euro //Convert in euro 'Total contributions'
	fmt.Printf("● Загальна сума внесків: %.2f €\n", convertTotalContributionsEuro)
	convertFinalAmountEuro := finalAmount / euro //Convert in euro 'Final summa'
	fmt.Printf("● Фінальна сума: %.2f €\n", convertFinalAmountEuro)
}
