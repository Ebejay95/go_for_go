/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   bank.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/15 13:01:23 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/15 14:11:46 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

const (
	D = "\033[0m"
	R = "\033[31m"
	G = "\033[32m"
	Y = "\033[33m"
	B = "\033[34m"
	P = "\033[35m"
	C = "\033[36m"
	W = "\033[37m"
)

func main() {

	accountBalance := 1000.0
	
	for {
		fmt.Println(C + "Welcome to Go Bank!" + D)
		fmt.Println(C + "What do you want to do?" + D)
		fmt.Println(C + "1. Check balance" + D)
		fmt.Println(C + "2. Deposit money" + D)
		fmt.Println(C + "3. Withdraw money" + D)
		fmt.Println(C + "4. Exit\n" + D)
	
		var choice int
		var changeAmount float64
		fmt.Print(B + "Your Choice: " + D)
		fmt.Scan(&choice)
	
		if choice == 1 {
			fmt.Println(G + "Your balance is", accountBalance, "" + D)
		} else if choice == 2 {
			fmt.Print(B + "Your Deposit: " + D)
			fmt.Scan(&changeAmount)
			if changeAmount <= 0 {
				fmt.Println(R + "Deposit must be greater that 0." + D)
				continue
			}
			accountBalance += changeAmount
			fmt.Println(G + "Your new balance is", accountBalance, "" + D)
		} else if choice == 3 {
			fmt.Print(B + "Your Withdraw: " + D)
			fmt.Scan(&changeAmount)
			if changeAmount <= 0 {
				fmt.Println(R + "Withdrawal must be greater that 0." + D)
				continue
			}
			if changeAmount > accountBalance {
				fmt.Println(R + "Cant withdraw more than you have." + D)
				continue
			}
			accountBalance -= changeAmount
			fmt.Println(G + "Your new balance is", accountBalance, "" + D)
		} else {
			fmt.Println(C + "Goodbye from Go Bank!" + D)
			break
		}	
	}
}