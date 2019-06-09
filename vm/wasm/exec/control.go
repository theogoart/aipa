// Copyright 2018~2022 The aipa Authors
// This file is part of the aipa Chain library.
// Created by  Team of aipa.

// This program is free software: you can distribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with aipa.  If not, see <http://www.gnu.org/licenses/>.

// Copyright 2018 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * file description:  control instructions
 * @Author: 
 * @Date:   2018-12-07
 * @Last Modified by:
 * @Last Modified time:
 */

package exec

import "errors"

// ErrUnreachable is the error value used while trapping the VM when
// an unreachable operator is reached during execution.
var ErrUnreachable = errors.New("exec: reached unreachable")

func (vm *VM) unreachable() {
	//panic(ErrUnreachable)
}

func (vm *VM) nop() {}
