package app

import (
	"github.com/castillofranciscodaniel/golang-beers/pkg/domain"
	"github.com/castillofranciscodaniel/golang-beers/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/json-iterator/go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

type BeerHandler struct {
	log          zerolog.Logger
	beersService domain.BeerService
}

func NewBeersHandler(beersService domain.BeerService) BeerHandler {
	return BeerHandler{
		beersService: beersService,
		log:          log.With().Str(utils.Struct, "BeerHandler").Logger(),
	}
}

// Get -
func (b *BeerHandler) Get(w http.ResponseWriter, r *http.Request) {
	b.log.Info().Str(utils.Thread, middleware.GetReqID(r.Context())).Str(utils.Method, utils.GetFunc).Msg(utils.InitStr)

	beersResponse, err := b.beersService.Get(r.Context())
	if err != nil {
		b.log.Error().Err(err).Str(utils.Thread, middleware.GetReqID(r.Context())).Str(utils.Method, utils.GetFunc).Send()
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := jsoniter.NewEncoder(w).Encode(beersResponse); err != nil {
		b.log.Error().Err(err).Str(utils.Thread, middleware.GetReqID(r.Context())).Str(utils.Method, utils.GetFunc).Send()
		w.WriteHeader(http.StatusConflict)
		return
	}

	b.log.Info().Str(utils.Thread, middleware.GetReqID(r.Context())).Str(utils.Method, utils.GetFunc).Msg(utils.EndStr)
	w.WriteHeader(http.StatusOK)
}

// Post -
func (b *BeerHandler) Post(w http.ResponseWriter, r *http.Request) {
	b.log.Info().Str(utils.Thread, middleware.GetReqID(r.Context())).Str(utils.Method, utils.PostFunc).Msg(utils.InitStr)

	var beerDto BeerDto

	err := jsoniter.NewDecoder(r.Body).Decode(&beerDto)
	if err != nil {
		WriteErrorResponse(r.Context(), w, err)
		return
	}

	err = b.beersService.Post(r.Context(), beerDto.MapToDomain())
	if err != nil {
		b.log.Error().Str(utils.Thread, middleware.GetReqID(r.Context())).Str(utils.Method, utils.PostFunc).Send()

		WriteErrorResponse(r.Context(), w, err)
		return
	}
	b.log.Info().Str(utils.Thread, middleware.GetReqID(r.Context())).Str(utils.Method, utils.PostFunc).Msg(utils.EndStr)
	WriteResponse(r.Context(), w, http.StatusCreated, nil)
}

func (b *BeerHandler) ById(w http.ResponseWriter, r *http.Request) {
	b.log.Info().Str(utils.Thread, middleware.GetReqID(r.Context())).Str(utils.Method, utils.ByIdFunc).Msg(utils.InitStr)

	idParam := chi.URLParam(r, "beerId")

	id, err := strconv.ParseInt(idParam, 0, 64)
	if err != nil {
		WriteInvalidRequest(w)
		return
	}

	beer, err := b.beersService.ById(r.Context(), id)
	if err != nil {
		WriteErrorResponse(r.Context(), w, err)
		return
	}

	b.log.Info().Str(utils.Thread, middleware.GetReqID(r.Context())).Str(utils.Method, utils.ByIdFunc).Msg(utils.EndStr)
	WriteResponse(r.Context(), w, http.StatusOK, beer)
}
