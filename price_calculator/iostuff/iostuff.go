/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   io_stuff.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/18 00:05:14 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/18 00:05:19 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package iostuff

type IOStuff interface {
	WriteResult(data any) error
	ReadLines() ([]string, error)
}