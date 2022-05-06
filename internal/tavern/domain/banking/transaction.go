package transaction

type Transaction struct{}

// Сейчас фабрика взяла на себя всю ответственность по валидации входных данных, созданию нового ID и заданию всех начальных значений
func NewTransaction(name string) (Transaction, error) {
	return Transaction{}, nil
}
