package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"

	"github.com/faroukelkholy/bank/internal/storage"
)

//Options contains db connection options
type Options struct {
	Debug    string
	DBHost   string
	DBPort   int
	DBName   string
	DBUser   string
	DBPass   string
	DBSchema string
}

// Driver holds db connection
// repos implementation
type Driver struct {
	storage.AccountRepository
	storage.CustomerRepository
	storage.TransactionRepository
}

//New return repos instance
func New(ctx context.Context, opt *Options) (repo storage.Repository, err error) {
	DB, err := connect(ctx, opt)

	repo = &Driver{
		AccountRepository:     NewAccountRepo(DB),
		CustomerRepository:    NewCustomerRepo(DB),
		TransactionRepository: NewTransactionRepo(DB),
	}
	return
}

//connect hold the db connection code using the provided options
func connect(ctx context.Context, opt *Options) (DB *pg.DB, err error) {
	DB = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%v", opt.DBHost, opt.DBPort),
		User:     opt.DBUser,
		Password: opt.DBPass,
		Database: opt.DBName,
		OnConnect: func(ctx context.Context, cn *pg.Conn) error {
			// on db connection define the schema that hold the tables
			if _, errExec := cn.Exec("set search_path=?", opt.DBSchema); errExec != nil {
				return errExec
			}
			return nil
		},
	})

	if opt.Debug == "dev" {
		DB.AddQueryHook(queryLogger{})
	}

	err = DB.Ping(ctx)
	return
}

//queryLogger implements the QueryHook interface that enables a hook to run before and after a query is executed
type queryLogger struct{}

func (queryLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	fmt.Printf("db query: <%v>\n ", q.Query)
	return c, nil
}

func (queryLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	return nil
}
