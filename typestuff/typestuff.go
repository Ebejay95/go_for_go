/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   typestuff.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/15 20:39:08 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/15 20:45:32 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

type customString string

func (text customString) log() {
	fmt.Println(text)
}

func main() {
	var name customString = "jonathaneberle"
	name.log()
}