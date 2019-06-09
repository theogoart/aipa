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
 * file description:  net actor
 * @Author:
 * @Date:   2018-12-07
 * @Last Modified by:
 * @Last Modified time:
 */

package produceractor

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/aipadad/aipa/action/message"
	"github.com/aipadad/aipa/common/types"
)

var netActorPid *actor.PID

// SetNetActorPid is to set net actor pid
func SetNetActorPid(tpid *actor.PID) {
	netActorPid = tpid
}

// BroadCastBlock is to broadcast blocks
func BroadCastBlock(block *types.Block) {

	broadCastBlock := &message.NotifyBlock{block}
	netActorPid.Tell(broadCastBlock)
	return
}
