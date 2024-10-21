/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   cmd_manager.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/17 23:55:15 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/18 00:01:45 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package cmdmanager

import "fmt"

type CMDManager struct {
	
}

func (fm CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter prices. Confirm every price with ENTER (EXIT exits)")
	var prices []string
	for {
		var price string
		fmt.Print("price: ")
		fmt.Scan(&price)
		if price == "EXIT" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (fm CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
