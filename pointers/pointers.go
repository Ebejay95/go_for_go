/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   pointers.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/15 15:42:45 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/15 15:47:08 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func main() {
	age := 18
	fmt.Println("Age:", age);
	fmt.Println("Adult since:", getAdultY(&age));
}

func getAdultY(age *int) int {
	return *age - 18
}