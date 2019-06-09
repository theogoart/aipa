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
 * file description: account role test
 * @Author: 
 * @Date:   2018-12-13
 * @Last Modified by:
 * @Last Modified time:
 */
package chain

import (
	log "github.com/cihub/seelog"
	"sort"
	"testing"

	"github.com/aipadad/aipa/config"
)

func TestBlockChain_ConfirmedSort(t *testing.T) {
	delegateNum := config.BLOCKS_PER_ROUND
	lastConfirmedNums := make(ConfirmedNum, delegateNum)
	var i uint32
	for i = 0; i < delegateNum; i++ {
		lastConfirmedNums[i] = delegateNum - i
	}

	log.Info(lastConfirmedNums)

	consensusIndex := (100 - int(config.CONSENSUS_BLOCKS_PERCENT)) * len(lastConfirmedNums) / 100
	sort.Sort(lastConfirmedNums)
	log.Info(lastConfirmedNums)
	log.Info(lastConfirmedNums[consensusIndex])
}
