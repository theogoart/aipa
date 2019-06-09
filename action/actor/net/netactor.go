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
 * @Date:   2018-12-06
 * @Last Modified by:
 * @Last Modified time:
 */

package netactor

import (
	log "github.com/cihub/seelog"
	//"encoding/json"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/aipadad/aipa/action/env"
	"github.com/aipadad/aipa/action/message"
	"github.com/aipadad/aipa/config"
	"github.com/aipadad/aipa/context"
	netprotocol "github.com/aipadad/aipa/protocol"
)

type NetActor struct {
	actorEnv *env.ActorEnv
	protocol context.ProtocolInstance
}

var netactor *NetActor

func NewNetActor(env *env.ActorEnv) *actor.PID {

	netactor = &NetActor{
		actorEnv: env,
	}

	props := actor.FromProducer(func() actor.Actor { return netactor })

	pid, err := actor.SpawnNamed(props, "NetActor")
	if err == nil {

		netactor.protocol = netprotocol.MakeProtocol(&config.Param, env.Chain)
		netactor.protocol.Start()

		env.Protocol = netactor.protocol
		return pid
	} else {
		panic(log.Errorf("NetActor SpawnNamed error: ", err))
	}

	return nil
}

//main loop
func (n *NetActor) handleSystemMsg(context actor.Context) {
	switch msg := context.Message().(type) {

	case *actor.Started:
		log.Infof("NetActor received started msg", msg)

	case *actor.Stopping:
		log.Info("NetActor received stopping msg")

	case *actor.Restart:
		log.Info("NetActor received restart msg")

	case *actor.Restarting:
		log.Info("NetActor received restarting msg")

	case *actor.Stop:
		log.Info("NetActor received Stop msg")

	case *actor.Stopped:
		log.Info("NetActor received Stopped msg")

	case *message.NotifyTrx:
		n.protocol.SendNewTrx(msg)

	case *message.NotifyBlock:
		n.protocol.SendNewBlock(msg)

	default:
		log.Error("netactor receive unknown message")
	}

}

func (n *NetActor) Receive(context actor.Context) {
	n.handleSystemMsg(context)
}

func (n *NetActor) setChainActor(tpid *actor.PID) {
	n.protocol.SetChainActor(tpid)
}

func (n *NetActor) setTrxPreHandleActor(tpid *actor.PID) {
	n.protocol.SetTrxPreHandleActor(tpid)
}

func (n *NetActor) setConsensusActor(tpid *actor.PID) {
	n.protocol.SetConsensusActor(tpid)
}

func SetChainActorPid(tpid *actor.PID) {
	netactor.setChainActor(tpid)
}

func SetTrxPreHandleActorPid(tpid *actor.PID) {
	netactor.setTrxPreHandleActor(tpid)
}

func SetConsensusActorPid(tpid *actor.PID) {
	netactor.setConsensusActor(tpid)
}
