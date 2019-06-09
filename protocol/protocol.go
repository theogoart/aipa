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

package protocol

import (
	"net"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/aipadad/aipa/action/message"
	"github.com/aipadad/aipa/chain"
	"github.com/aipadad/aipa/config"
	"github.com/aipadad/aipa/context"
	"github.com/aipadad/aipa/p2p"
	"github.com/aipadad/aipa/role"
	"github.com/aipadad/aipa/protocol/block"
	"github.com/aipadad/aipa/protocol/common"
	"github.com/aipadad/aipa/protocol/consensus"
	"github.com/aipadad/aipa/protocol/discover"
	"github.com/aipadad/aipa/protocol/transaction"
	log "github.com/cihub/seelog"
)

type protocol struct {
	d *discover.Discover
	t *transaction.Transaction
	b *block.Block
	c *consensus.Consensus
}

func MakeProtocol(config *config.P2PConfig, chain chain.BlockChainInterface,roleIntf role.RoleInterface ) context.ProtocolInstance {
	runner := p2p.MakeP2PServer(config,roleIntf)

	p := &protocol{
		d: discover.MakeDiscover(config),
		t: transaction.MakeTransaction(roleIntf),
		b: block.MakeBlock(chain, true), //we should get node type here
		c: consensus.MakeConsensus(),
	}

	sendup := func(index uint16, packet *p2p.Packet) {
		if packet.H.ProtocolType == common.P2P_PACKET {
			p.d.Dispatch(index, packet)
		} else if packet.H.ProtocolType == common.TRX_PACKET {
			p.t.Dispatch(index, packet)
		} else if packet.H.ProtocolType == common.BLOCK_PACKET {
			p.b.Dispatch(index, packet)
		} else if packet.H.ProtocolType == common.CONSENSUS_PACKET {
			p.c.Dispatch(index, packet)
		} else {
			log.Errorf("PROTOCOL wrong packet type")
		}
	}

	p.d.SetSendupCallback(sendup)

	conn := func(conn net.Conn) {
		p.d.NewConnCb(conn, sendup)
	}

	runner.SetCallback(conn)

	return p
}

func (p *protocol) Start() {
	p2p.Runner.Start()
	p.d.Start()
	p.c.Start()
	p.b.Start()
	p.t.Start()
}
func (p *protocol) GetPeerInfo()( uint64, []*context.PeersInfo) {
	var peersInfo []*context.PeersInfo
	var peerCount uint64 = 0
	pr := p2p.Runner.GetPeerP2PInfo()
	lengthPeers := len(pr)

	for _, info := range p.b.S.Peers {
		temp := info
		for j := 0; j < lengthPeers; j++ {
			if temp.Index == pr[j].Index {
				peer := &context.PeersInfo{
					LastLib:   temp.LastLib,
					LastBlock: temp.LastBlock,
					Addr:      pr[j].Info.Addr,
					Port:      pr[j].Info.Port,
					ChainId:   pr[j].Info.ChainId,
					Version:   pr[j].Info.Version,
				}
				peersInfo = append(peersInfo, peer)
				peerCount++
				break
			}

		}
	}
	return peerCount, peersInfo
}
func (p *protocol) GetBlockSyncState() bool {

	if config.BtoConfig.Delegate.Solo == true {
		return true
	}
	return p.b.GetSyncState()
}

func (p *protocol) SetChainActor(tpid *actor.PID) {
	p.b.SetActor(tpid)
}

func (p *protocol) SetTrxPreHandleActor(tpid *actor.PID) {
	p.t.SetActor(tpid)
}

func (p *protocol) SetConsensusActor(tpid *actor.PID) {
	p.c.SetActor(tpid)
}

func (p *protocol) SendNewTrx(notify *message.NotifyTrx) {
	p.t.SendNewTrx(notify)
}

func (p *protocol) SendNewBlock(notify *message.NotifyBlock) {
	p.b.UpdateHeadNumber()
	p.b.UpdateNumber()
	p.b.SendNewBlock(notify)
}

func (p *protocol) SendPrevote(notify *message.SendPrevote) {
	p.c.SendPrevote(notify)
}
func (p *protocol) SendPrecommit(notify *message.SendPrecommit) {
	p.b.UpdateHeadNumber()
	p.c.SendPrecommit(notify)
}
func (p *protocol) SendCommit(notify *message.SendCommit) {
	p.b.UpdateNumber()
	p.c.SendCommit(notify)
}