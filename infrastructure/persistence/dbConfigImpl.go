package persistence

import (
	"context"
	"database/sql"
	"github.com/castillofranciscodaniel/golang-beers/utils"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/url"
	"time"
)

type DbManagerPostgres struct {
	log zerolog.Logger
	db  *sql.DB
}

func NewDbManagerImpl() DbManagerPostgres {
	igDbConfigImpl := DbManagerPostgres{
		log: log.With().Str(utils.Struct, "DbManagerPostgres").Logger(),
	}

	igDbConfigImpl.config()
	return igDbConfigImpl
}

func (i DbManagerPostgres) DB() *sql.DB {
	defer i.log.Info().Interface("statusBd: ", i.db.Stats()).Msg(utils.Data)
	return i.db
}

func (i *DbManagerPostgres) config() {
	i.log.Info().Str(utils.Struct, "DbManagerPostgres").Msg(utils.InitStr)

	query := url.Values{}
	//query.Add("app name", "ms-payment-neg")
	query.Add("database", "test")
	query.Add("sslmode", "disable")

	u := &url.URL{
		Scheme: "postgresql",
		User:   url.UserPassword("postgres", "postgres"),
		Host:   "database:5432",
		// Path:  instance, // if connecting to an instance instead of a app
		RawQuery: query.Encode(),
	}

	var err error

	// Create connection pool
	db, err := sql.Open("postgres", u.String())

	if err != nil {
		i.log.Error().Str(utils.Struct, "DbManagerPostgres").Err(err).Send()
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		i.log.Error().Str(utils.Struct, "DbManagerPostgres").Err(err).Send()
	}
	i.log.Info().Str(utils.Struct, "DbManagerPostgres").Msgf("Connected %v", utils.EndStr)

	maxOpenConnes := 25

	maxIdleConns := 25

	maxLifeTime := 300

	db.SetMaxOpenConns(maxOpenConnes)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(time.Duration(maxLifeTime) * time.Second)

	i.db = db
}
