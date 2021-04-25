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
	DB *pg.DB
}

//New return repos instance
func New(opt *Options) (storage.Repository, error) {
	d := new(Driver)
	if err := d.connect(opt); err != nil {
		return nil, err
	}
	return d, nil
}

//connect hold the db connection code using the provided options
func (d *Driver) connect(opt *Options) (err error) {
	d.DB = pg.Connect(&pg.Options{
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
	defer fmt.Println("db connected success")

	if opt.Debug == "dev" {
		d.DB.AddQueryHook(queryLogger{})
	}

	if err = d.DB.Ping(context.Background()); err != nil {
		return err
	}
	return nil
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
