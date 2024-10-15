/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   investment_calculator.go                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/14 10:34:47 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/15 10:42:41 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64
	
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	
	fmt.Print("Years: ")
	fmt.Scan(&years)
	
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}