/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   structs_practise.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/15 20:47:32 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/16 12:32:14 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs_practise/note"
	"example.com/structs_practise/todo"
)

type saver interface {
	Save() error
}
type displayer interface {
	Display()
}

type outputable interface {
	saver
	displayer
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the Data failed")
		return err
	}
	fmt.Println("Saving the Data succeeded")
	return nil
}

func main () {
	printer("Test")
	printer(1)
	printer(3.1)
	title, content := getNoteData()
	text := getTodoData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userTodo, err := todo.New(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	output(&userNote)
	output(&userTodo)
}

func printer(value interface{}) {
	switch value.(type) {
		case int:
			fmt.Println("Integer:", value)
		case float64:
			fmt.Println("float64:", value)
	}
	typedVal, ok := value.(int)
	typedVal += 1
	if (ok) {
		fmt.Println("Integer:", typedVal)
	}
}

func output(data outputable) {
	data.Display()
	err := saveData(data)
	if err != nil {
		return
	}
}

func getTodoData() (string) {
	text := getUserInput("Todo Text:")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) (string) {
	fmt.Print(prompt)
	
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")
	return value
}