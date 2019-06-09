
package apiactor

import (
	log "github.com/cihub/seelog"
	"path/filepath"
	"testing"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/aipadad/aipa/action/actor/transaction/trxhandleactor"
	"github.com/aipadad/aipa/action/message"
	"github.com/aipadad/aipa/common/types"
	"github.com/aipadad/aipa/db"
	"github.com/aipadad/aipa/transaction"
	//"github.com/aipadad/aipa/config"
)

var trxActorPid *actor.PID

func TestPushTrxTest(t *testing.T) {

	// init testing
	dbInst := db.NewDbService("./datadir/", filepath.Join("./datadir/", "blockchain"))
	if dbInst == nil {
		log.Info("Create DB service fail")
		//os.Exit(1)
	}
	trxActorPid = trxactor.NewTrxActor()

	//InitTrxActorAgent()
	var trxPool = transaction.InitTrxPool(dbInst)
	trxactor.SetTrxPool(trxPool)

	log.Info("Test PushTrxTest will called")
	trxTest := &types.Transaction{
		Cursor:      11,
		CursorLabel: 22,
	}

	reqMsg := &message.PushTrxReq{
		Trx:       trxTest,
		TrxSender: message.TrxSenderTypeFront,
	}

	// push trx
	result, err := trxActorPid.RequestFuture(reqMsg, 500*time.Millisecond).Result()

	if nil == err {
		log.Info("push trx req exec result:")
		log.Infof("rusult is =======", result)
		log.Infof("error  is =======", err)
	} else {
		t.Error("push trx failed, trx:", trxTest)
	}

	getTrxsReq := &message.GetAllPendingTrxReq{}

	// get all trx
	getTrxsResult, getTrxsErr := trxActorPid.RequestFuture(getTrxsReq, 500*time.Millisecond).Result()

	if nil == err {
		log.Info("push trx req exec result:")
		log.Infof("rusult is =======", getTrxsResult)
		log.Infof("error  is =======", getTrxsErr)
	} else {
		t.Error("get all trx req exec error")
	}

	var removeTrxs []*types.Transaction

	removeTrxs = append(removeTrxs, trxTest)

	removeTrxsReq := &message.RemovePendingTrxsReq{
		Trxs: removeTrxs,
	}

	// remove trx
	trxActorPid.Tell(removeTrxsReq)

	// get all trxs after remove ,should be empty
	getTrxsAfterRemoveResult, getTrxsAfterRemoveErr := trxActorPid.RequestFuture(getTrxsReq, 500*time.Millisecond).Result()

	if nil == err {
		log.Info("get all trx req after remove exec result:")
		log.Infof("rusult is =======", getTrxsAfterRemoveResult)
		log.Infof("error  is =======", getTrxsAfterRemoveErr)
	} else {
		t.Error("get all trx req after remove exec error")
	}
}
