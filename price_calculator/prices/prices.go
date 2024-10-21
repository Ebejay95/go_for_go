/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   prices.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/17 21:19:13 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/18 00:13:54 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package prices

import (
	"fmt"

	"example.com/price_calculator/conversion"
	"example.com/price_calculator/iostuff"
)


type TaxIncludedPriceJob struct {
	Io iostuff.IOStuff `json:"-"`
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_incl_prices"`
}

func NewTaxIncludedPriceJob(io iostuff.IOStuff, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob {
		Io: io,
		TaxRate: taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxInclPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxInclPrice)
	}
	job.TaxIncludedPrices = result
	job.Io.WriteResult(job)
	return nil
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.Io.ReadLines()
	if err != nil {
		fmt.Println(err)
		return err
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("Could not convert to floats")
		fmt.Println(err)
		return err
	}
	job.InputPrices = prices;
	return nil
}

