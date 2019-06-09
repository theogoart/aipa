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
 * file description: database interface
 * @Author: 
 * @Date:   2018-12-04
 * @Last Modified by:
 * @Last Modified time:
 */
package codedb

import (
	"fmt"
	"testing"
)

func TestMultindexDB_AddObject(t *testing.T) {
	fmt.Println("abc")
	ins, _ := NewMultindexDB("./file1")
	ins.CallAddObject("abc")
	ins.CallAddObject("elf")
}

func TestMultindexDB_Commit(t *testing.T) {
	fmt.Println("abc")
	ins, _ := NewMultindexDB("./file1")
	ins.CallAddObject("abc")
	ins.CallAddObject("elf")
	ins.CallSetRevision(1)
	ins.CallSetObject("abc", "111", "222")
	ins.CallSetObject("elf", "112", "elf222")
	rl := ins.CallGetRevision()
	ins.CallCommit(2)
	fmt.Println(rl)
}

func TestMultindexDB_Undo(t *testing.T) {
	fmt.Println("abcdo")
	ins, _ := NewMultindexDB("./file1")
	ins.CallAddObject("abcundo")
	ins.CallAddObject("elfundo")
	ins.CallSetRevision(1)
	ins.CallSetObject("abcundo", "111", "222")
	ins.CallSetObject("elfundo", "112", "elf222")
	rl := ins.CallGetRevision()
	ins.CallRollback()
	fmt.Println(rl)
	val, err := ins.CallGetObject("abcundo", "111")
	fmt.Println("value", val, err)
	val, err = ins.CallGetObject("elfundo", "112")
	fmt.Println("value", val, err)
}
