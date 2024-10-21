/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   concurrency.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/21 19:38:05 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/21 19:41:49 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"time"
)

// func greet(phrase string, doneChan chan bool) {
// 	fmt.Println("Hello!", phrase)
// 	doneChan <- true
// }

// func slowGreet(phrase string, doneChan chan bool) {
// 	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
// 	fmt.Println("Hello!", phrase)
// 	doneChan <- true
// }

// func main() {
// 	dones := make([]chan bool, 4)
// 	dones[0] = make(chan bool)
// 	go greet("Nice to meet you!", dones[0])
// 	dones[1] = make(chan bool)
// 	go greet("How are you?",  dones[1])
// 	dones[2] = make(chan bool)
// 	go slowGreet("How ... are ... you ...?",  dones[2])
// 	dones[3] = make(chan bool)
// 	go greet("I hope you're liking the course!", dones[3])

// 	for _, done := range dones {
// 		<-done
// 	}
// }


func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan) // only makes sence when we do know which operation takes the most time
}

func main() {
	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)
	for doneChan := range done {
		fmt.Println(doneChan)
	}
}
