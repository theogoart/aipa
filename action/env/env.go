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
 * @Date:   2018-12-20
 * @Last Modified by:
 * @Last Modified time:
 */

package env

import (
	"github.com/aipadad/aipa/chain"
	"github.com/aipadad/aipa/context"
	"github.com/aipadad/aipa/contract"
	"github.com/aipadad/aipa/db"
	"github.com/aipadad/aipa/db/platform/codedb"
	"github.com/aipadad/aipa/role"
)

//ActorEnv actor external interface
type ActorEnv struct {
	RoleIntf   role.RoleInterface
	Chain      chain.BlockChainInterface
	NcIntf     contract.NativeContractInterface
	Protocol   context.ProtocolInterface
	Db               *db.DBService
	PendingTxSession *codedb.UndoSession
}
