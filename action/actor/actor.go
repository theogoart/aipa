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
 * file description:  actor entry
 * @Author:
 * @Date:   2018-12-06
 * @Last Modified by:
 * @Last Modified time:
 */

package actionactor

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	apiactor "github.com/aipadad/aipa/action/actor/api"
	chainactor "github.com/aipadad/aipa/action/actor/chain"
	netactor "github.com/aipadad/aipa/action/actor/net"
	produceractor "github.com/aipadad/aipa/action/actor/producer"
	"github.com/aipadad/aipa/action/actor/transaction/trxpoolactor"
	"github.com/aipadad/aipa/action/actor/transaction/trxhandleactor"
	"github.com/aipadad/aipa/action/actor/transaction/trxprehandleactor"
	log "github.com/cihub/seelog"

	"github.com/aipadad/aipa/action/env"
	restactor"github.com/aipadad/aipa/restful/handler"
)

var apiActorPid *actor.PID
var netActorPid *actor.PID
var trxActorPid *actor.PID
var chainActorPid *actor.PID

//MultiActor actor group
type MultiActor struct {
	apiActorPid      *actor.PID
	netActorPid      *actor.PID
	trxActorPid      *actor.PID
	trxPoolActorPid      *actor.PID
	trxPreHandleActor *actor.PID
	chainActorPid    *actor.PID
	producerActorPid *actor.PID
}

func (m *MultiActor) GetTrxActor() *actor.PID {
	return m.trxActorPid
}

func (m *MultiActor) GetTrxPreHandleActor() *actor.PID {
	return m.trxPreHandleActor
}


//GetNetActor get net actor PID
func (m *MultiActor) GetNetActor() *actor.PID {
	return m.netActorPid
}

//InitActors init all actor
func InitActors(env *env.ActorEnv) *MultiActor {

	mActor := &MultiActor{
		apiactor.NewApiActor(),
		netactor.NewNetActor(env),
		trxactor.NewTrxActor(),
		trxpoolactor.NewTrxPoolActor(),
		trxprehandleactor.NewTrxPreHandleActor(),
		chainactor.NewChainActor(env),
		produceractor.NewProducerActor(env),
	}
	registerActorMsgTbl(mActor)
	return mActor
}

func registerActorMsgTbl(m *MultiActor) {

	log.Info("RegisterActorMsgTbl")
	apiactor.SetTrxPreHandleActorPid(m.trxPreHandleActor) // api --> trx
	restactor.SetTrxPreHandleActorPid(m.trxPreHandleActor) // api --> trx
	apiactor.SetChainActorPid(m.chainActorPid) // api --> chain
	restactor.SetChainActorPid(m.chainActorPid)	// restapi --> chain
	trxactor.SetApiActorPid(m.apiActorPid)          // trx --> api
	produceractor.SetChainActorPid(m.chainActorPid) // producer --> chain
	produceractor.SetTrxPoolActorPid(m.trxPoolActorPid)     // producer --> trx
	produceractor.SetNetActorPid(m.netActorPid)     // producer --> chain
	chainactor.SetTrxPoolActorPid(m.trxPoolActorPid)        // chain --> trx
	chainactor.SetNetActorPid(m.netActorPid)        // chain --> net

	netactor.SetTrxPreHandleActorPid(m.trxPreHandleActor)     //p2p --> trx
	netactor.SetChainActorPid(m.chainActorPid) //p2p --> chain
}

//GetTrxActorPID get trx actor pid
func (m *MultiActor) GetTrxActorPID() *actor.PID {
	return m.trxActorPid
}

//ActorsStop stop all actor
func (m *MultiActor) ActorsStop() {
	m.chainActorPid.Stop()
	m.producerActorPid.Stop()
	m.apiActorPid.Stop()
	m.netActorPid.Stop()
	m.trxActorPid.Stop()
	m.trxPoolActorPid.Stop()

}
