package app

import (
	"context"
	err2 "github.com/castillofranciscodaniel/golang-beers/infrastructure/err"
	"github.com/castillofranciscodaniel/golang-beers/utils"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/json-iterator/go"
	"github.com/rs/zerolog/log"
	"net/http"
)

func WriteErrorResponse(ctx context.Context, w http.ResponseWriter, err error) {
	var httpError *err2.HttpError

	if castErr, isOK := err.(*err2.HttpError); isOK {
		httpError = castErr
	} else {
		httpError = err2.BadRequestError
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpError.Code)

	if err := jsoniter.NewEncoder(w).Encode(httpError); err != nil {
		log.Error().Err(err).Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.ByIdFunc).Send()
		w.WriteHeader(err2.BadRequestError.Code)
		return
	}
}

func WriteInvalidRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

func WriteResponse(ctx context.Context, w http.ResponseWriter, statusExpected int, i interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusExpected)

	if i != nil {
		if err := jsoniter.NewEncoder(w).Encode(i); err != nil {
			log.Error().Err(err).Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.ByIdFunc).Send()
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
