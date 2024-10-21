/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   conversion.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/17 22:51:55 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/17 22:56:51 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package conversion

import (
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for _ , value := range strings {
		floatPrice, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, floatPrice)
	}
	return floats, nil
}