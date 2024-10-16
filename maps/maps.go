/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   maps.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/16 14:45:38 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/16 14:52:36 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func main () {
	websites := []string{"https://google.com", "c"}
	fmt.Println(websites)

	mapOfWebsites := map[string]string{
		"Google": "https://google.com",
		"Amazon Web Services": "https://google.com",
	}
	fmt.Println(mapOfWebsites)
	mapOfWebsites["linkedin"] = "https://linkedin.com"
	fmt.Println(mapOfWebsites)
	delete(mapOfWebsites, "Google")
	fmt.Println(mapOfWebsites)
}