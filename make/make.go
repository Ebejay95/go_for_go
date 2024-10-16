/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   make.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/16 14:56:58 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/16 15:12:08 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

type floatMap map[string]float64

func main() {
	userNames := make([]string, 2, 5)

	userNames = append(userNames, "Jonathan")
	userNames = append(userNames, "Anna")
	userNames = append(userNames, "Bernd")
	userNames = append(userNames, "Sarah")
	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 5.0
	courseRatings["c"] = 4.0
	fmt.Println(courseRatings)

	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}
	for key, value := range courseRatings {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
