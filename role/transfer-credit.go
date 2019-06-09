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
 * file description:  balance role
 * @Author: 
 * @Date:   2018-01-20
 * @Last Modified by:
 * @Last Modified time:
 */

package role

import (
	"encoding/json"
	"math/big"

	"github.com/aipadad/aipa/db"
	log "github.com/cihub/seelog"
)

//TransferCreditObjectName is credit name
const TransferCreditObjectName string = "credit"

//TransferCredit is the struct for transferring credit
type TransferCredit struct {
	Name    string `json:"name"`
	Spender string `json:"spender"`
	Limit   *big.Int `json:"limit"`
}

//CreateTransferCreditRole is create transfer credits
func CreateTransferCreditRole(ldb *db.DBService) error {
	ldb.AddObject(TransferCreditObjectName)
	return nil
}

func generateKey(name string, spender string) string {
	return name + string("-") + spender
}

//SafeSub is safe sub
func (credit *TransferCredit) SafeSub(amount *big.Int) error {
	result := big.NewInt(0)
	result, err := safeSub(result, credit.Limit, amount)
	if err != nil {
		log.Error("ROLE safeSub failed ", err)
		return err
	}
	credit.Limit.Set(result)
	return nil

}

//SetTransferCreditRole is setting transfer credit role
func SetTransferCreditRole(ldb *db.DBService, name string, value *TransferCredit) error {
	key := generateKey(name, value.Spender)
	jsonvalue, err := json.Marshal(value)
	if err != nil {
		log.Error("ROLE Marshal failed ", err)
		return err
	}
	return ldb.SetObject(TransferCreditObjectName, key, string(jsonvalue))
}

//GetTransferCreditRole is geting transfer credit role
func GetTransferCreditRole(ldb *db.DBService, name string, spender string) (*TransferCredit, error) {
	key := generateKey(name, spender)
	value, err := ldb.GetObject(TransferCreditObjectName, key)
	if err != nil {
		return nil, err
	}

	res := &TransferCredit{}
	err = json.Unmarshal([]byte(value), res)
	if err != nil {
		log.Error("ROLE Unmarshal failed ", err)
		return nil, err
	}

	return res, nil
}

//DeleteTransferCreditRole is deleting transfer credit role
func DeleteTransferCreditRole(ldb *db.DBService, name string, spender string) error {
	key := generateKey(name, spender)
	_, err := ldb.DeleteObject(TransferCreditObjectName, key)
	return err
}
