/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   lists.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/16 13:23:47 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/16 13:58:36 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"


func main() {
	var productNames [8]string
	prices := [4]float64{4.3,23.2,253,232}
	fmt.Println(prices)
	fmt.Println(prices[2])
	fmt.Println(prices[1:3])
	fmt.Println(prices[:3])
	fmt.Println(prices[2:])
	highPrices := prices[1:3]
	fmt.Println(len(highPrices))
	fmt.Println(cap(highPrices))
	fmt.Println(productNames)
	
	dynamicPrices := []float64{4.3, 435}
	fmt.Println(dynamicPrices[0])
	fmt.Println(dynamicPrices[1])
	updatedPrices := append(dynamicPrices, 5.99)
	fmt.Println(updatedPrices)
	updatedMPrices := append(dynamicPrices, 5.25, 83.53)
	fmt.Println(updatedMPrices)
	testPrices := []float64{4.3, 435}
	updatedMPricesU := append(updatedMPrices, testPrices...)
	fmt.Println(updatedMPricesU)
}