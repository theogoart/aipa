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
 * file description: code database test
 * @Author: 
 * @Date:   2018-12-04
 * @Last Modified by:
 * @Last Modified time:
 */
package role

import (
	//	"encoding/json"
	log "github.com/cihub/seelog"
	"testing"

	"github.com/aipadad/aipa/db"
)

func TestDelegate_writedb(t *testing.T) {
	ins := db.NewDbService("./file1", "./file1/db.db", "")
	err := CreateDelegateRole(ins)
	if err != nil {
		log.Error(err)
	}
	value := &Delegate{
		AccountName:           "lmq",
		LastSlot:              3,
		ReportKey:             "0xaaaaaaaaaaaaaaaaaa",
		TotalMissed:           0,
		LastConfirmedBlockNum: 2}
	err = SetDelegateRole(ins, value.AccountName, value)
	if err != nil {
		log.Error("SetDelegateRole", err)
	}

	value, err = GetDelegateRoleByAccountName(ins, value.AccountName)
	if err != nil {
		log.Error("GetDelegateRoleByAccountName", err)
	}
	log.Info(value)

	value, err = GetDelegateRoleBySignKey(ins, value.ReportKey)
	if err != nil {
		log.Error("GetDelegateRoleByAccountName", err)
	}
	log.Info(value)
}

func TestDelegate_WritedbTheSameKey(t *testing.T) {
	ins := db.NewDbService("./file2", "./file2/db2.db", "")
	err := CreateDelegateRole(ins)
	if err != nil {
		log.Error(err)
	}
	value1 := &Delegate{
		AccountName:           "lmq1",
		LastSlot:              4,
		ReportKey:             "0xaaaaaaaaaaaaaaaaaa",
		TotalMissed:           0,
		LastConfirmedBlockNum: 3}
	value2 := &Delegate{
		AccountName:           "lmq1",
		LastSlot:              3,
		ReportKey:             "0xbbbbbb",
		TotalMissed:           0,
		LastConfirmedBlockNum: 2}
	err = SetDelegateRole(ins, value1.AccountName, value1)
	if err != nil {
		log.Error("SetDelegateRole", err)
	}

	err = SetDelegateRole(ins, value2.AccountName, value2)
	if err != nil {
		log.Error("SetDelegateRole", err)
	}

	value, err1 := GetDelegateRoleByAccountName(ins, value1.AccountName)
	if err1 != nil {
		log.Error("GetDelegateRoleByAccountName", err)
	}
	log.Info(value)

	value, err = GetDelegateRoleBySignKey(ins, value2.ReportKey)
	if err != nil {
		log.Error("GetDelegateRoleByAccountName", err)
	}
	log.Info(value)
}
