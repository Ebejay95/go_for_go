/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   funcdeepdive.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/16 15:15:37 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/16 17:20:35 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func mapnumbers(numbers *[]int, cb func(int) int) []int {
	tNumbers := []int{}

	for _, val := range *numbers {
		tNumbers = append(tNumbers, cb(val))
	}
	return tNumbers
}

func main () {
	numbers := []int{1, 2, 3, 4}
	doubled := mapnumbers(&numbers, double)
	fmt.Println(doubled)
	tripled := mapnumbers(&numbers, triple)
	fmt.Println(tripled)
}