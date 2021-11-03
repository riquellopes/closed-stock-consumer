package repository

import "fmt"

const (
	MONGO = "mongo"
)

func getRepository(repositoryType string) (StockInterface, error) {
	if repositoryType == MONGO {
		database := mongoDataBase()
		return &StockMongo{
			Collection: database.Collection("stock"),
		}, nil
	}

	return nil, fmt.Errorf("Repository notfound.")
}
