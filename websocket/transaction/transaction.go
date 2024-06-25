package transaction

import (
	"fmt"
	"r3_client/tools"
	"r3_client/types"
	"sync"
)

var (
	transactionNrMap    = make(map[uint64]types.RequestTransaction)
	transactionNrMap_mx = sync.Mutex{}
)

func Get(transNr uint64) (types.RequestTransaction, bool) {
	transactionNrMap_mx.Lock()
	defer transactionNrMap_mx.Unlock()
	trans, exists := transactionNrMap[transNr]
	return trans, exists
}

func Deregister(transNr uint64) {
	transactionNrMap_mx.Lock()
	defer transactionNrMap_mx.Unlock()

	delete(transactionNrMap, transNr)
}
func Register(requests []types.Request) (types.RequestTransaction, error) {
	trans := types.RequestTransaction{
		Requests: requests,
	}

	transactionNrMap_mx.Lock()
	defer transactionNrMap_mx.Unlock()

	var number uint64
	var maxAttempts = 10000
	for maxAttempts > 0 {
		number = uint64(tools.RandNumber(100000, 499999))

		if _, exists := transactionNrMap[number]; !exists {
			trans.TransactionNr = number
			transactionNrMap[number] = trans
			return trans, nil
		}
		maxAttempts--
	}
	return trans, fmt.Errorf("failed to register transaction number")
}
