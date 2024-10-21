/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   file_manager.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/17 23:00:44 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/17 23:45:36 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package file_manager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println("Could not open file", fm.InputFilePath)
		fmt.Println(err)
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Could not scan file", fm.InputFilePath)
		fmt.Println(err)
		file.Close()
		return nil, err
	}
	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		fmt.Println("Could not create json file", fm.OutputFilePath)
		fmt.Println(err)
		return err
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Could not convert data to json")
		fmt.Println(err)
		file.Close()
		return err
	}
	file.Close()
	return nil
}

func New(in string, out string) FileManager {
	return FileManager{
		InputFilePath: in,
		OutputFilePath: out,
	}
}