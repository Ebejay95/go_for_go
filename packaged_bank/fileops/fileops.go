/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   fileops.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/15 14:57:46 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/15 15:23:27 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(filename string, defaultVal float64) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return defaultVal, errors.New("Failed to find file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return defaultVal, errors.New("Failed convert float from file")
	}
	return value, nil
}

func WriteFloatToFile(filename string, value float64) {
	valueString := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueString), 0644);
}