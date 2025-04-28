type Transaction struct {
	commands [][]string
	state TransactionState
}

type TransactionState int

const (
	TX_OPEN TransactionState = iota
	TX_COMMITED
	TX_DISCARDED
)

func(s *Server) Multi(conn net.Conn) {
 // Start a transaction
}

func(s *Server) Exec(conn net.Conn) {
// Execute all commands the transaction
}