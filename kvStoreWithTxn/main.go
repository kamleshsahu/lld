package main

import (
	"fmt"
	"lld/kvStoreWithTxn/service"
)

func main() {

	kvStore := service.GetKVStore()

	kvStore.PUT("nilesh", "sahu")

	val, _ := kvStore.GET("nilesh")
	fmt.Println(val)

	kvStore.PUT("nilesh", "sahu2")

	val, _ = kvStore.GET("nilesh")
	fmt.Println(val)

	ssIsoTxn := service.NewTxn(kvStore)

	txnId, _ := ssIsoTxn.BEGIN()
	txnId2, _ := ssIsoTxn.BEGIN()

	ssIsoTxn.PUT(txnId, "kamlesh", "sahu")
	ssIsoTxn.PUT(txnId, "nilesh", "tikesh3")
	ssIsoTxn.PUT(txnId2, "nilesh", "tikesh2")
	ssIsoTxn.PUT(txnId, "kamlesh", "tikesh")
	err := ssIsoTxn.ROLLBACK(txnId2)
	fmt.Println(err)
	err = ssIsoTxn.COMMIT(txnId)
	fmt.Println(err)

	val, _ = kvStore.GET("nilesh")
	fmt.Println("after txn nilesh:", val)
	val, _ = kvStore.GET("kamlesh")
	fmt.Println("after txn kamlesh:", val)
}
