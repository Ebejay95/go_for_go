/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   todo.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/16 11:50:37 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/16 12:14:28 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid input")
	}
	return Todo{
		Text: content,
	}, nil
}

func (note *Todo) Display () {
	fmt.Println(note.Text)
}

func (note *Todo) Save () error {
	fileName := "todo.json"
	jsonContent, err := json.Marshal(note)
	if err != nil {
		return err
	}
	os.WriteFile(fileName, jsonContent, 0644)
	return nil
}