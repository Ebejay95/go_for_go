/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   note.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/15 23:02:52 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/16 12:14:05 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input")
	}
	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (note *Note) Display () {
	fmt.Println(note.Title, "has the following content:")
	fmt.Println(note.Content)
}

func (note *Note) Save () error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName = fileName + ".json"
	jsonContent, err := json.Marshal(note)
	if err != nil {
		return err
	}
	os.WriteFile(fileName, jsonContent, 0644)
	return nil
}