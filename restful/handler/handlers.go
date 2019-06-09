package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"errors"

	"github.com/gorilla/mux"
	"github.com/aipadad/aipa/action/env"
	"github.com/aipadad/aipa/action/message"
	"github.com/AsynkronIT/protoactor-go/actor"


	"time"

	aipaErr "github.com/aipadad/aipa/common/errors"
	"github.com/aipadad/aipa/api"
	"github.com/aipadad/aipa/common"
	"github.com/aipadad/aipa/common/types"
	"github.com/aipadad/aipa/role"
	service "github.com/aipadad/aipa/action/actor/api"
	log "github.com/cihub/seelog"
	"github.com/aipadad/aipa/contract/abi"
	"github.com/aipadad/aipa/config"
	"regexp"
	"runtime"
)

//ApiService is actor service
type ApiService struct {
	env *env.ActorEnv
}

var roleIntf role.RoleInterface
//SetChainActorPid set chain actor pid
func SetRoleIntf(tpid role.RoleInterface) {
	roleIntf = tpid
}

var chainActorPid *actor.PID

//SetChainActorPid set chain actor pid
func SetChainActorPid(tpid *actor.PID) {
	chainActorPid = tpid
}

var trxactorPid *actor.PID

//SetTrxActorPid set trx actor pid
func SetTrxActorPid(tpid *actor.PID) {
	trxactorPid = tpid
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := ResponseStructs{
		ResponseStruct{Msg: "Write presentation"},
		ResponseStruct{Msg: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintf(w, "Todo show: %s\n", todoId)
}

//Node
func GetGenerateBlockTime(w http.ResponseWriter, r *http.Request) {
/*	//func GetGenerateBlockTime(cmd map[string]interface{}) map[string]interface{} {
	resp := ResponsePack(error.SUCCESS)
	resp["Result"] = "aq"
	//fmt.Fprint(w, "Welcome!\n",resp	)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}*/
	//resp["Result"] = config.DEFAULT_GEN_BLOCK_TIME
	//return resp
}

//GetInfo query chain info
func GetInfo(w http.ResponseWriter, r *http.Request) {
	msgReq := &message.QueryChainInfoReq{}
	res, err := chainActorPid.RequestFuture(msgReq, 500*time.Millisecond).Result()

	var resp ResponseStruct
	if err != nil {
		resp.Errcode = uint32(aipaErr.ErrApiQueryChainInfoError)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrApiQueryChainInfoError)
		encoderRestResponse(w, resp)
		return
		}

	if resp := checkNil(res, 1); resp.Errcode != 0 {
		encoderRestResponse(w, resp)
		return
	}

	response := res.(*message.QueryChainInfoResp)
	if response.Error != nil {
		resp.Errcode = uint32(aipaErr.ErrApiQueryChainInfoError)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrApiQueryChainInfoError)
		encoderRestResponse(w, resp)
		return
	}

	result := &api.GetInfoResponse_Result{}
	result.HeadBlockNum = response.HeadBlockNum
	result.LastConsensusBlockNum = response.LastConsensusBlockNum
	result.HeadBlockHash = response.HeadBlockHash.ToHexString()
	result.HeadBlockTime = response.HeadBlockTime
	result.HeadBlockDelegate = response.HeadBlockDelegate
	result.CursorLabel = response.HeadBlockHash.Label()
	result.ChainId=common.BytesToHex(config.GetChainID())

	resp.Errcode = uint32(aipaErr.ErrNoError)
	resp.Msg = aipaErr.GetCodeString(aipaErr.ErrNoError)
	resp.Result = result
	encoderRestResponse(w, resp)
}
func checkNil(req interface{}, flag int8) (ResponseStruct) {
	var resp ResponseStruct

	if req == nil {
		if flag == 0 {
			resp.Errcode = uint32(aipaErr.RestErrReqNil)
			resp.Msg = aipaErr.GetCodeString(aipaErr.RestErrReqNil)
		} else {
			resp.Errcode = uint32(aipaErr.RestErrResultNil)
			resp.Msg = aipaErr.GetCodeString(aipaErr.RestErrResultNil)
		}

		funcName, _, _, _ := runtime.Caller(1)
		log.Errorf("%s errcode: %d checkNil error:%s", runtime.FuncForPC(funcName).Name(), resp.Errcode, resp.Msg)
		//encoderRestResponse(w, resp)
		return resp
	}
	return resp
}

//GetBlock query block
func GetBlock(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var msgReq *api.GetBlockRequest
	err := json.NewDecoder(r.Body).Decode(&msgReq)
	if err != nil {
		log.Errorf("request error: %s", err)
		panic(err)
		return
	}

	msgReq2 := &message.QueryBlockReq{common.HexToHash(msgReq.BlockHash),
		msgReq.BlockNum}

	res, err := chainActorPid.RequestFuture(msgReq2, 500*time.Millisecond).Result()
	var resp ResponseStruct
	if err != nil {
		resp.Errcode = uint32(aipaErr.ErrApiBlockNotFound)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrApiBlockNotFound)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	response := res.(*message.QueryBlockResp)
	if response.Block == nil {
		resp.Errcode = uint32(aipaErr.ErrApiBlockNotFound)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrApiBlockNotFound)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	result := &api.GetBlockResponse_Result{}
	hash := response.Block.Hash()
	result.PrevBlockHash = response.Block.GetPrevBlockHash().ToHexString()
	result.BlockNum = response.Block.GetNumber()
	result.BlockHash = hash.ToHexString()
	result.CursorBlockLabel = hash.Label()
	result.BlockTime = response.Block.GetTimestamp()
	result.TrxMerkleRoot = response.Block.ComputeMerkleRoot().ToHexString()
	result.Delegate = string(response.Block.GetDelegate())
	result.DelegateSign = response.Block.GetDelegateSign().ToHexString()
	resp.Result = result

	resp.Errcode = 0
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

func SendTransaction(w http.ResponseWriter, r *http.Request) {
	var trx *api.Transaction
	err := json.NewDecoder(r.Body).Decode(&trx)
	if err != nil {
		log.Errorf("request error: %s", err)
		panic(err)
		return
	}
	var resp ResponseStruct
	if trx != nil {
		//verity Sender
		match, err := regexp.MatchString("^[a-z1-9][a-z1-9.-]{2,20}$", trx.Sender)
		if err != nil {
			if err := json.NewEncoder(w).Encode(err); err != nil {
				//panic(err)
				log.Error(err)
			}
			return
		}
		if !match {
			resp.Errcode = uint32(aipaErr.ErrTrxAccountError)
			resp.Msg = aipaErr.GetCodeString((aipaErr.ErrCode)(resp.Errcode))
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				panic(err)
			}
			return
		}
		//verity Contract
		match, err = regexp.MatchString("^[a-z1-9][a-z1-9.-]{2,20}$", trx.Contract)
		if err != nil {
			if err := json.NewEncoder(w).Encode(err); err != nil {
				//panic(err)
				log.Error(err)
			}
			return
		}
		if !match {
			resp.Errcode = uint32(aipaErr.ErrTrxAccountError)
			resp.Msg = aipaErr.GetCodeString((aipaErr.ErrCode)(resp.Errcode))
			if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}
		//verity Method
		match, err = regexp.MatchString("^[a-z1-9][a-z1-9.-]{2,20}$", trx.Method)
		if err != nil {
			if err := json.NewEncoder(w).Encode(err); err != nil {
				//panic(err)
				log.Error(err)
			}
			return
		}
		if !match {
			resp.Errcode = uint32(aipaErr.ErrTrxAccountError)
			resp.Msg = aipaErr.GetCodeString((aipaErr.ErrCode)(resp.Errcode))
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				panic(err)
			}
			return
		}
	} else {
		//rsp.retCode = ??
	}

	intTrx, err := service.ConvertApiTrxToIntTrx(trx)
	if err != nil {
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	reqMsg := &message.PushTrxReq{
		Trx: intTrx,
	}

	handlerErr, err := trxactorPid.RequestFuture(reqMsg, 500*time.Millisecond).Result() // await result

	if nil != err {
		resp.Errcode = uint32(aipaErr.ErrActorHandleError)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrActorHandleError)

		log.Errorf("trx: %x actor process failed", intTrx.Hash(), )

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	result := &api.SendTransactionResponse_Result{}
	if aipaErr.ErrNoError == handlerErr {
		result.TrxHash = intTrx.Hash().ToHexString()
		result.Trx = service.ConvertIntTrxToApiTrx(intTrx)
		resp.Result = result
		resp.Msg = "trx receive succ"
		resp.Errcode = 0
	} else {
		result.TrxHash = intTrx.Hash().ToHexString()
		result.Trx = service.ConvertIntTrxToApiTrx(intTrx)
		resp.Result = result
		//resp.Msg = handlerErr.(string)GetCodeString
		//resp.Msg = "to be add detail error description"
		var tempErr aipaErr.ErrCode
		tempErr = handlerErr.(aipaErr.ErrCode)

		resp.Errcode = (uint32)(tempErr)
		resp.Msg = aipaErr.GetCodeString((aipaErr.ErrCode)(resp.Errcode))
	}

	log.Infof("trx: %v %s", result.TrxHash, resp.Msg)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

type reqStruct struct {
	TrxHash string `json:"trx_hash,omitemty"`
}

type Transaction struct {
	Version     uint32 `json:"version"`
	CursorNum   uint64 `json:"cursor_num"`
	CursorLabel uint32 `json:"cursor_label"`
	Lifetime    uint64 `json:"lifetime"`
	Sender      string `json:"sender"`
	Contract    string `json:"contract"`
	Method      string `json:"method"`
	Param       interface{} `json:"param"`
	SigAlg      uint32 `json:"sig_alg"`
	Signature   string `json:"signature"`
}

func getContractAbi(r role.RoleInterface, contract string) (*abi.ABI, error) {
	account, err := r.GetAccount(contract)
	if err != nil {
		return nil, errors.New("Get account fail")
	}

	Abi, err := abi.ParseAbi(account.ContractAbi)
	if err != nil {
		return nil, err
	}

	return Abi, nil
}

func ParseTransactionParam(r role.RoleInterface, Param []byte, Contract string, Method string) (interface{}, error) {
	var Abi *abi.ABI = nil
	if Contract != "aipa" {
		var err error
		Abi, err = getContractAbi(r, Contract)
		if  err != nil {
			return nil, errors.New("External Abi is empty!")
		}
	} else {
		Abi = abi.GetAbi()
	}

	if Abi == nil {
		return nil, errors.New("Abi is empty!")
	}

	decodedParam := abi.UnmarshalAbiEx(Contract, Abi, Method, Param)
	if decodedParam == nil || len(decodedParam) <= 0 {
		return nil, errors.New("ParseTransactionParam: FAILED")
	}
	return decodedParam, nil
}

func convertIntTrxToApiTrxInter(trx *types.Transaction,r role.RoleInterface) interface{} {
	parmConvered, err := ParseTransactionParam(r, trx.Param, trx.Contract, trx.Method)
	if err != nil {
		log.Errorf("role.ParseParam: %s", err)
		panic(err)
	}

	apiTrx := &Transaction{
		Version:     trx.Version,
		CursorNum:   trx.CursorNum,
		CursorLabel: trx.CursorLabel,
		Lifetime:    trx.Lifetime,
		Sender:      trx.Sender,
		Contract:    trx.Contract,
		Method:      trx.Method,
		Param:       parmConvered,
		SigAlg:      trx.SigAlg,
		Signature:   common.BytesToHex(trx.Signature),
	}

	return apiTrx
}

//GetTransaction get transaction by Trx hash
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	var req *reqStruct
	var resp comtool.ResponseStruct

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Errorf("REST:json Decoder failed: %v", err)
		resp.Errcode = uint32(aipaErr.RestErrJsonNewEncoder)
		resp.Msg = aipaErr.GetCodeString(aipaErr.RestErrJsonNewEncoder)
		resp.Result = err

		encoderRestResponse(w, resp)
		return
	}

	if resp := checkNil(req, 0); resp.Errcode != 0 {
		encoderRestResponse(w, resp)
		return
	}

	msgReq := &message.QueryTrxReq{
		TrxHash: common.HexToHash(req.TrxHash),
}
	res, err := chainActorPid.RequestFuture(msgReq, 500*time.Millisecond).Result()
	if err != nil {
		log.Errorf("REST:chainActor process failed: %v", err)
		resp.Errcode = uint32(aipaErr.ErrActorHandleError)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrActorHandleError)
		encoderRestResponse(w, resp)
		return
		}

	if resp := checkNil(res, 1); resp.Errcode != 0 {
		encoderRestResponse(w, resp)
		return
	}

	response := res.(*message.QueryTrxResp)
	if response.Trx == nil {
		resp.Errcode = uint32(aipaErr.ErrApiTrxNotFound)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrApiTrxNotFound)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	//resp.Result = service.ConvertIntTrxToApiTrx(response.Trx)
	resp.Result = convertIntTrxToApiTrxInter(response.Trx, roleIntf)

	resp.Errcode = uint32(aipaErr.ErrNoError)
	resp.Msg = aipaErr.GetCodeString(aipaErr.ErrNoError)
	encoderRestResponse(w, resp)
}

type TransactionStatus struct {
	Status string `json:"status"`
}

func GetTransactionStatus(w http.ResponseWriter, r *http.Request) {
	var req *reqStruct
	var resp comtool.ResponseStruct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Errorf("REST:json Decoder failed: %v", err)
		resp.Errcode = uint32(aipaErr.RestErrJsonNewEncoder)
		resp.Msg = aipaErr.GetCodeString(aipaErr.RestErrJsonNewEncoder)
		resp.Result = err

		encoderRestResponse(w, resp)
		return
	}

	if resp := checkNil(req, 0); resp.Errcode != 0 {
		encoderRestResponse(w, resp)
		return
	}

	tx := actorenv.Chain.GetCommittedTransaction(common.HexToHash(req.TrxHash))
	if tx != nil {
		resp.Errcode = uint32(aipaErr.ErrNoError)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrNoError)
		resp.Result = &TransactionStatus{Status: "committed"}
		encoderRestResponse(w, resp)
		return
	}

	tx = actorenv.Chain.GetTransaction(common.HexToHash(req.TrxHash))
	if tx != nil {
		resp.Errcode = uint32(aipaErr.RestErrTxPacked)
		resp.Msg = aipaErr.GetCodeString(aipaErr.RestErrTxPacked)
		resp.Result = &TransactionStatus{Status: "packed"}
		encoderRestResponse(w, resp)
		return
	}

	trxApply := transaction.NewTrxApplyService()
	errCode := trxApply.GetTrxErrorCode(common.HexToHash(req.TrxHash))
	if aipaErr.ErrNoError != errCode {

		resp.Errcode = uint32(aipaErr.ErrNoError)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrNoError)
		resp.Result = &TransactionStatus{Status: aipaErr.GetCodeString(errCode)}
		encoderRestResponse(w, resp)
		return
	}

	isInPool := trxApply.IsTrxInPendingPool(common.HexToHash(req.TrxHash))
	if true == isInPool {
		resp.Errcode = uint32(aipaErr.RestErrTxPending)
		resp.Msg = aipaErr.GetCodeString(aipaErr.RestErrTxPending)
		resp.Result = &TransactionStatus{Status: "pending"}
		encoderRestResponse(w, resp)
		return
	} else {
		resp.Errcode = uint32(aipaErr.RestErrTxNotFound)
		resp.Msg = aipaErr.GetCodeString(aipaErr.RestErrTxNotFound)
		resp.Result = &TransactionStatus{Status: "not found"}
		encoderRestResponse(w, resp)
		return
	}
}

//GetAccountBrief query account public key
func GetAccountBrief(w http.ResponseWriter, r *http.Request) {
	var msgReq api.GetAccountBriefRequest
	var resp comtool.ResponseStruct
	err := json.NewDecoder(r.Body).Decode(&msgReq)
	if err != nil {
		log.Errorf("REST:json Decoder failed: %v", err)
		resp.Errcode = uint32(aipaErr.RestErrJsonNewEncoder)
		resp.Msg = aipaErr.GetCodeString(aipaErr.RestErrJsonNewEncoder)
		resp.Result = err

		encoderRestResponse(w, resp)
		return
	}

	if resp := checkNil(msgReq, 0); resp.Errcode != 0 {
		encoderRestResponse(w, resp)
		return
	}
	name := msgReq.AccountName

	result := &api.GetAccountBriefResponse_Result{}

	accountType, _ := common.AnalyzeName(name)
	if common.NameTypeUnknown == accountType {
		resp.Errcode = uint32(aipaErr.ErrApiAccountNameIllegal)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrApiAccountNameIllegal)
		encoderRestResponse(w, resp)
		return
	} else if common.NameTypeAccount == accountType {
		account, err := roleIntf.GetAccount(name)
		if err != nil {
			resp.Errcode = uint32(aipaErr.ErrApiAccountNotFound)
			resp.Msg = aipaErr.GetCodeString(aipaErr.ErrApiAccountNotFound)

			encoderRestResponse(w, resp)
			return
		}
		if resp := checkNil(account, 1); resp.Errcode != 0 {
			encoderRestResponse(w, resp)
			return
		}

		balance, err := roleIntf.GetBalance(name)
		if err != nil {
			log.Errorf("DB:GetBalance failed: %v", err)

			resp.Errcode = uint32(aipaErr.ErrApiAccountNotFound)
			resp.Msg = aipaErr.GetCodeString(aipaErr.ErrApiAccountNotFound)
			encoderRestResponse(w, resp)
			return
		}

		result.AccountName = name
		result.Pubkey = common.BytesToHex(account.PublicKey)
		result.Balance = balance.Balance.String()
	}

	resp.Errcode = uint32(aipaErr.ErrNoError)
	resp.Msg = aipaErr.GetCodeString(aipaErr.ErrNoError)
	resp.Result = result
	encoderRestResponse(w, resp)
}

//GetAccount query account info
func GetAccount(w http.ResponseWriter, r *http.Request) {
	var msgReq api.GetAccountRequest
	err := json.NewDecoder(r.Body).Decode(&msgReq)
	if err != nil {
		log.Errorf("request error: %s", err)
		panic(err)
		return
	}
	name := msgReq.AccountName

	account, err := roleIntf.GetAccount(name)
	var resp ResponseStruct
	if err != nil {
		resp.Errcode = uint32(aipaErr.ErrApiAccountNotFound)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrApiAccountNotFound)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	balance, err := roleIntf.GetBalance(name)
	if err != nil {
		resp.Errcode = uint32(aipaErr.ErrApiAccountNotFound)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrApiAccountNotFound)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	stakedBalance, err := roleIntf.GetStakedBalance(name)
	if err != nil {
		resp.Errcode = uint32(aipaErr.ErrApiAccountNotFound)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrApiAccountNotFound)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	result := &api.GetAccountResponse_Result{}
	result.AccountName = name
	result.Pubkey = common.BytesToHex(account.PublicKey)
	result.Balance = balance.Balance.String()
	result.StakedBalance = stakedBalance.StakedBalance.String()
	resp.Result=result
	resp.Errcode = 0

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

func GetKeyValue(w http.ResponseWriter, r *http.Request) {
	var req *api.GetKeyValueRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Errorf("request error: %s", err)
		panic(err)
		return
	}

	contract := req.Contract
	object := req.Object
	key := req.Key
	value, err := roleIntf.GetBinValue(contract, object, key)
	var resp ResponseStruct
	if err != nil {
		resp.Errcode = uint32(aipaErr.ErrApiObjectNotFound)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrApiObjectNotFound)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	result := &api.GetKeyValueResponse_Result{}
	result.Contract = contract
	result.Object = object
	result.Key = key
	result.Value = common.BytesToHex(value)
	resp.Result = result
	resp.Errcode = 0

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

func GetContractAbi(w http.ResponseWriter, r *http.Request) {
	var req *api.GetAbiRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Errorf("request error: %s", err)
		panic(err)
		return
	}
	//contract := req.Contract
	account, err := roleIntf.GetAccount(req.Contract)
	var resp ResponseStruct
	if err != nil {
		resp.Errcode = uint32(aipaErr.ErrApiAccountNotFound)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrApiAccountNotFound)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	if len(account.ContractAbi) > 0 {
		resp.Result = string(account.ContractAbi)
	} else {
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

func GetContractCode(w http.ResponseWriter, r *http.Request) {
	var req *api.GetAbiRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Errorf("request error: %s", err)
		panic(err)
		return
	}
	//contract := req.Contract
	account, err := roleIntf.GetAccount(req.Contract)
	var resp ResponseStruct
	if err != nil {
		resp.Errcode = uint32(aipaErr.ErrApiAccountNotFound)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrApiAccountNotFound)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	if len(account.ContractCode) > 0 {
		resp.Result = common.BytesToHex(account.ContractCode)
	} else {
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

func GetTransferCredit(w http.ResponseWriter, r *http.Request) {
	var req *api.GetTransferCreditRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Errorf("request error: %s", err)
		panic(err)
		return
	}
	name := req.Name
	spender := req.Spender
	credit, err := roleIntf.GetTransferCredit(name, spender)
	var resp ResponseStruct
	if err != nil {
		resp.Errcode = uint32(aipaErr.ErrTransferCreditNotFound)
		resp.Msg = aipaErr.GetCodeString(aipaErr.ErrTransferCreditNotFound)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	result := &api.GetTransferCreditResponse_Result{}
	result.Name = credit.Name
	result.Spender = credit.Spender
	result.Limit = credit.Limit.String()
	resp.Result = result

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}



func encoderRestResponse(w http.ResponseWriter, resp ResponseStruct) (http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Credentials", "true")
	//w.Header().Set("Access-Control-Expose-Headers", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	//w.Header().Set("Access-Control-Allow-Methods", "POST")

	//w.WriteHeader(404)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		resp.Errcode = uint32(aipaErr.RestErrJsonNewEncoder)
		resp.Msg = aipaErr.GetCodeString(aipaErr.RestErrJsonNewEncoder)

		funcName, _, _, _ := runtime.Caller(1)
		log.Errorf("%s errcode: %d json.NewEncoder(w).Encode(resp) error:%s", runtime.FuncForPC(funcName).Name(), resp.Errcode, err)
	}

	return w
}

func encoderRestRequest(r *http.Request, req interface{}) (interface{}, error) {
	var resp ResponseStruct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Errorf("request error: %s", err)
		resp.Errcode = uint32(aipaErr.RestErrJsonNewEncoder)
		resp.Msg = aipaErr.GetCodeString(aipaErr.RestErrJsonNewEncoder)
		resp.Result = err

		funcName, _, _, _ := runtime.Caller(1)
		log.Errorf("%s errcode: %d json.NewEncoder(w).Encode(resp) error:%s", runtime.FuncForPC(funcName).Name(), resp.Errcode, err)

		return resp, err
	}

	if req == nil {
		resp.Errcode = uint32(aipaErr.RestErrReqNil)
		resp.Msg = aipaErr.GetCodeString(aipaErr.RestErrReqNil)

		funcName, _, _, _ := runtime.Caller(1)
		log.Errorf("%s errcode: %d json.NewEncoder(w).Encode(resp) error:%s", runtime.FuncForPC(funcName).Name(), resp.Errcode, err)
		return resp, errors.New("request is nil")
	}

	return req, nil
}

func GetTrxHashForSign(sender, contract, method string, param []byte, h *api.GetInfoResponse) ([]byte, *types.Transaction, error) {
	var blockHeader *api.GetInfoResponse_Result
	if h == nil {
		var err error
		blockHeader, err = GetBlockHeader()
		if err != nil {
			return nil, nil, err
		}
	} else {
		blockHeader = h.Result
	}

	trx := &types.BasicTransaction{
		Version:     blockHeader.HeadBlockVersion,
		CursorNum:   blockHeader.HeadBlockNum,
		CursorLabel: blockHeader.CursorLabel,
		Lifetime:    blockHeader.HeadBlockTime + 100,
		Sender:      sender,
		Contract:    contract,
		Method:      method,
		Param:       param,
		SigAlg:      config.SIGN_ALG,
	}
	msg, err := bpl.Marshal(trx)
	if nil != err {
		log.Errorf("REST:bpl Marshal failed: %v", err)
		return nil, nil, err
	}

	//Add chainID Flag
	chainID, _ := hex.DecodeString(blockHeader.ChainId)
	msg = bytes.Join([][]byte{msg, chainID}, []byte{})

	intTrx := &types.Transaction{
		Version:     trx.Version,
		CursorNum:   trx.CursorNum,
		CursorLabel: trx.CursorLabel,
		Lifetime:    trx.Lifetime,
		Sender:      trx.Sender,
		Contract:    trx.Contract,
		Method:      trx.Method,
		Param:       trx.Param,
		SigAlg:      config.SIGN_ALG,
	}
	return comtool.Sha256(msg), intTrx, err
}