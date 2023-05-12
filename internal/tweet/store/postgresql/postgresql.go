package postgresql

import (
	"context"
	"log"
	"time"

	"github.com/alwisisva/twitter-app/internal/tweet/entity"
	"github.com/alwisisva/twitter-app/pkg/helper"
	"github.com/jmoiron/sqlx"
)

// storeClient implements service.StoreClient.
type storeClient struct {
	Db *sqlx.DB
}

// New constructs a new storeClient
func New(connectionString string) (*storeClient, error) {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(helper.EnvInt("POSTGRES_MIN_IDLE_CONNECTION"))

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(helper.EnvInt("POSTGRES_MAX_OPEN_CONNECTION"))

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	maxLifeTime := time.Duration(helper.EnvInt("POSTGRES_MAX_LIFETIME")) * time.Second
	db.SetConnMaxLifetime(maxLifeTime)

	log.Printf("⇨ Psql status is connected to %s:%s database %s \n", helper.EnvString("POSTGRES_HOST"), helper.EnvString("POSTGRES_PORT"), helper.EnvString("POSTGRES_DATABASE"))
	return &storeClient{
		Db: db,
	}, nil
}

// TODO: implements service.StoreClient with storeClient.
func (s *storeClient) Post(ctx context.Context, params entity.Posts) (res int64, err error) {
	_, err = s.Db.NamedExecContext(ctx, "insert into posts (id,text,created_by,user_id) values(:id,:text,:created_by,:user_id)", params)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return params.ID, nil
}

func (s *storeClient) GetList(ctx context.Context, filter entity.FilterPosts) (res []entity.Posts, count int64, err error) {
	countQuery := "SELECT count(*) from posts"
	err = s.Db.GetContext(ctx, &count, countQuery)
	if err != nil {
		log.Println(err)
		return res, count, err
	}

	err = s.Db.SelectContext(ctx, &res, "SELECT id,text,created_by,user_id,created_at,updated_at from posts order by created_at desc limit $1 offset $2", filter.Limit, filter.Page)
	if err != nil {
		log.Println(err)
		return res, count, err
	}
	return
}

func (s *storeClient) GetByID(ctx context.Context, id string) (res entity.Posts, err error) {
	err = s.Db.GetContext(ctx, &res, "SELECT id,text,created_by,user_id,created_at,updated_at from posts where id = $1 ", id)
	if err != nil {
		log.Println(err)
		return res, err
	}
	return
}
