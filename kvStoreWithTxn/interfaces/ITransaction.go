package interfaces

type ITransaction interface {
	GET(transactionId int, key string) (string, error)
	PUT(transactionId int, key string, value string) error
	DELETE(transactionId int, key string) error
	BEGIN() (int, error)
	COMMIT(transactionId int) error
	ROLLBACK(transactionId int) error
}
