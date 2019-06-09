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
 * file description:  producer actor
 * @Author: 
 * @Date:   2018-12-06
 * @Last Modified by:
 * @Last Modified time:
 */

package context

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/aipadad/aipa/action/message"
)

type ProtocolInstance interface {
	ProtocolInterface
	Start()
	SetChainActor(tpid *actor.PID)
	SetTrxPreHandleActor(tpid *actor.PID)
	SetConsensusActor(tpid *actor.PID)

	SendNewTrx(notify *message.NotifyTrx)
	SendNewBlock(notify *message.NotifyBlock)
}

type ProtocolInterface interface {
	GetBlockSyncState() bool
	GetPeerInfo()(uint64,[]*PeersInfo)
}
