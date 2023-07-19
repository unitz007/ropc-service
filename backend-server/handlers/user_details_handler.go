package handlers

import (
	"net/http"
	"ropc-service/model/dto"
	"ropc-service/utils"
)

func UserDetailsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		accessToken := r.Header.Get("Authorization")
		claims, err := utils.ValidateToken(accessToken)
		if err != nil {
			panic(err)
		}

		userDetails := dto.UserDetails{
			Username: claims["username"].(string),
			Email:    claims["email"].(string),
			ClientId: claims["client_id"].(string),
		}

		_ = PrintResponse(http.StatusOK, w, dto.NewResponse("User details fetched successfully", userDetails))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
