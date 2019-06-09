// Copyright 2018~2022 The aipa Authors
// This file is part of the aipa Chain library.
// Created by  Team of aipa.

//This program is free software: you can distribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
// along with aipa.  If not, see <http://www.gnu.org/licenses/>.

/*
 * file description:  provide a interface such as time to seconds and epoch time etc.
 * @Author: 
 * @Date:   2018-12-01
 * @Last Modified by:
 * @Last Modified time:
 */

package common

//Filter filter out and return the rest
func Filter(sources []string, filters []string) []string {
	var tmpList []string
	var has bool
	for _, src := range sources {
		has = false
		for _, filter := range filters {
			if src == filter {
				has = true
				break
			}
		}
		if has == false {
			tmpList = append(tmpList, src)
		}
	}
	return tmpList
}
