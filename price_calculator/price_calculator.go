/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   price_calculator.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/17 20:47:56 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/18 00:14:22 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	cmdmanager "example.com/price_calculator/cmd_manager"
	"example.com/price_calculator/prices"
)

func main () {
	taxRates := []float64{0, 0.07, 0.01, 0.15}

	for _ , taxRate := range taxRates {
		cmd := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		//job := prices.NewTaxIncludedPriceJob(cmd.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100)), taxRate)
		err := job.Process()
		if (err != nil) {
			return 
		}
	}
}