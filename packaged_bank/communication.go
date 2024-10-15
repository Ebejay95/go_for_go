/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   communication.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/15 14:56:48 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/15 14:56:50 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func presentOptions() {
	fmt.Println(C + "Welcome to Go Bank!" + D)
	fmt.Println(C + "What do you want to do?" + D)
	fmt.Println(C + "1. Check balance" + D)
	fmt.Println(C + "2. Deposit money" + D)
	fmt.Println(C + "3. Withdraw money" + D)
	fmt.Println(C + "4. Exit\n" + D)
}