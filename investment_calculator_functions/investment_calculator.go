/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   investment_calculator.go                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/14 10:34:47 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/15 12:40:27 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64
	
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	
	fmt.Print("Years: ")
	fmt.Scan(&years)
	
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calc_results(investmentAmount, expectedReturnRate, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func calc_results(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {

	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	frv = fv / math.Pow(1 + inflationRate/100, years)
	return fv, frv
}