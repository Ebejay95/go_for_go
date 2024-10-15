/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   user.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/15 18:22:28 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/15 18:33:52 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthDate: "------",
			createdAt: time.Now(),
		},
	}
}

func NewUser(userFirstName, userLastName, userBirthDate string) (*User, error) {
	
	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		return nil, errors.New("firstname, lastname and birtdate are required")
	}
	
	return &User{
		firstName: userFirstName,
		lastName: userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}, nil
}

func (u *User) Print() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearName() {
	u.firstName = ""
	u.lastName = ""
}