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

import (
	"time"
)

//Microsecond micro second type
type Microsecond struct {
	Count uint64
}

func maximum() Microsecond {
	return Microsecond{Count: 0x7fffffffffffffff}
}

//ToMicroseconds current time to micro second
func ToMicroseconds(current time.Time) uint64 {
	cur := current.Unix()
	return uint64(cur * 1000)
}

//ToSeconds micro second to second
func ToSeconds(m Microsecond) uint64 {
	return m.Count / 1000000
}

//ToMilliseconds micro second to milli second
func ToMilliseconds(m Microsecond) uint64 {
	return m.Count / 1000
}

//SecondsToMicro second to micro second
func SecondsToMicro(s uint64) Microsecond {
	return Microsecond{s * 1000000}
}

//MilliSecToMicro milli second to micro second
func MilliSecToMicro(m uint64) Microsecond {
	return Microsecond{m * 1000}
}
