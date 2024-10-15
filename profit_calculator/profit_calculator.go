/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   profit_calculator.go                               :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/15 10:44:11 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/15 11:06:06 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main () {
	const taxRate float64 = 0.19
	var amount float64
	var result float64

	fmt.Print("Amount: ")
	fmt.Scan(&amount)
	
	result = amount * (1 - taxRate)
	//fmt.Println("Real revenue: ", result)
	fmt.Printf("Real revenue: %.2f\n", result)
}