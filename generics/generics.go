/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   generics.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/16 12:40:16 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/16 12:49:52 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"


func main() {
	result := add(1, 2)
	fmt.Println(result)
	smartResult := smartAdd("Hello", "World")
	fmt.Println(smartResult)
}

func smartAdd[T int | float64 | string](a, b T) T {
	return a + b
}

func add(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}
	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}
	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	if aIsString && bIsString {
		return aString + bString
	}
	return nil
}