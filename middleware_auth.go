package main

import (
	"fmt"
	"net/http"

	"github.com/manikant8/GoServer-rssagg/internal/auth"
	"github.com/manikant8/GoServer-rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithERROR(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		// getting user by Api key
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithERROR(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}

}
