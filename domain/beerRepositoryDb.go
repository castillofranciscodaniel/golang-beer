package domain

import (
	"database/sql"
	"github.com/castillofranciscodaniel/golang-beers/config"
	err2 "github.com/castillofranciscodaniel/golang-beers/err"
	"github.com/castillofranciscodaniel/golang-beers/utils"
	"github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type BeerRepositoryDb struct {
	log       zerolog.Logger
	dbManager config.DbManager
}

func NewBeersRepositoryDb(dbManager config.DbManager) BeerRepositoryDb {
	return BeerRepositoryDb{
		dbManager: dbManager,
		log:       log.With().Str(utils.Struct, "BeerRepositoryDb").Logger(),
	}
}

// Get -
func (b BeerRepositoryDb) Get() ([]BeerSql, error) {
	b.log.Info().Str(utils.Method, utils.GetFunc).Msgf(utils.InitStr)

	query := `SELECT id, name, brewery, country, price, currency FROM beer`
	rows, err := b.dbManager.DB().Query(query)

	if rows.Err() != nil && err != nil {
		b.log.Error().Err(err).Err(rows.Err()).Str(utils.Method, utils.PostFunc).Send()
		return nil, err
	}

	var beersSql []BeerSql

	for rows.Next() {
		var beer BeerSql

		if err := rows.Scan(
			&beer.Id,
			&beer.Name,
			&beer.Brewery,
			&beer.Country,
			&beer.Price,
			&beer.Currency,
		); err != nil {
			b.log.Error().Err(err).Send()
			return nil, err
		}

		beersSql = append(beersSql, beer)
	}

	b.log.Info().Str(utils.Method, utils.GetFunc).Send()
	return beersSql, nil
}

// Post -
func (b BeerRepositoryDb) Post(beer Beer) error {

	b.log.Info().Str(utils.Method, utils.PostFunc).Msgf(utils.InitStr)

	_, err := b.dbManager.DB().Exec(`insert into beer (id, name, brewery, country, price, currency) 
		values ($1, $2, $3, $4, $5, $6);`,
		sql.Named("ID", beer.GetId()),
		sql.Named("NAME", beer.GetName()),
		sql.Named("BREWERY", beer.GetBrewery()),
		sql.Named("COUNTRY", beer.GetCountry()),
		sql.Named("PRICE", beer.GetPrice()),
		sql.Named("CURRENCY", beer.GetCurrency()),
	)

	if err != nil {
		if pqErr, isOk := err.(*pq.Error); isOk && pqErr.Code.Name() == utils.UniqueViolationSql {
			return err2.DuplicatedIdError
		}
		b.log.Error().Err(err).Str(utils.Method, utils.PostFunc).Send()
		return err
	}

	//id, _ := rows.LastInsertId() --> not support
	b.log.Info().Str(utils.Method, utils.PostFunc).Send()
	return nil
}

func (b BeerRepositoryDb) ById(id int64) (*BeerSql, error) {
	b.log.Info().Str(utils.Method, utils.ByIdFunc).Msgf(utils.InitStr)

	var beer BeerSql

	query := `SELECT id, name, brewery, country, price, currency FROM beer WHERE id = $1`
	err := b.dbManager.DB().QueryRow(query, id).Scan(
		&beer.Id,
		&beer.Name,
		&beer.Brewery,
		&beer.Country,
		&beer.Price,
		&beer.Currency,
	)

	switch {
	case err == sql.ErrNoRows:
		b.log.Error().Err(err).Str(utils.Method, utils.ByIdFunc).Send()
		return nil, err2.NotFoundError
	case err != nil:
		b.log.Error().Err(err).Str(utils.Method, utils.ByIdFunc).Send()
		return nil, err
	default:
		b.log.Info().Str(utils.Method, utils.ByIdFunc).Send()
		return &beer, nil
	}
}
