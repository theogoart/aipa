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
 * file description:  test time point
 * @Author: 
 * @Date:   2018-12-01
 * @Last Modified by:
 * @Last Modified time:
 */
package common

import (
	log "github.com/cihub/seelog"
	"math/big"
	"testing"
	"time"
)

func Test_Now(t *testing.T) {

	ab := Now()
	log.Info(ab)

}

func Test_MeasureStart(t *testing.T) {
	now := time.Now().Unix()
	log.Info(now)
	m := MeasureStart()
	log.Info(m)
}

func Test_NanoSeconds(t *testing.T) {
	m := MeasureStart()
	an := big.NewInt(0)
	mb := an.SetUint64(m)
	log.Infof("ddd", mb)
	bp := bigpow(10, 9)
	log.Info(bp)
	value := mb.Div(mb, bp)

	log.Info(value)
	b := ToNanoseconds(time.Now())
	log.Info(b)
}
