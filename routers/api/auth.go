package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameschun7/go-gin-sample/pkg/app"
	"github.com/jameschun7/go-gin-sample/pkg/e"
	"github.com/jameschun7/go-gin-sample/service/auth_service"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/idtoken"
)

type socialAuth struct {
	LoginToken string `json:"login_token" binding:"required"`
	LoginType  string `json:"login_type" binding:"required"`
}

type TokenInfo struct {
	Sub           string
	Email         string
	AtHash        string
	Aud           string
	EmailVerified bool
	Name          string
	GivenName     string
	FamilyName    string
	Picture       string
	Local         string
	Iss           string
	Azp           string
	Iat           int
	Exp           int
}

// @Summary Post Auth
// @Produce  json
// @Param auth body socialAuth true "socialAuth Body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		authJson socialAuth
	)

	httpCode, errCode := app.BindAndValid(c, &authJson)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	log.Printf("token: %s", authJson.LoginToken)

	tokenInfo, errCode := verifyGoogleIDToken(c.Request.Context(), authJson.LoginToken)
	if errCode == e.ERROR_AUTH_CHECK_TOKEN_FAIL {
		appG.Response(http.StatusOK, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if errCode == e.ERROR_AUTH_INVALID_TOKEN {
		appG.Response(http.StatusOK, e.ERROR_AUTH_INVALID_TOKEN, nil)
		return
	}

	authService := auth_service.Auth{
		UserID:      tokenInfo.Sub,
		Username:    tokenInfo.Name,
		LoginType:   "google",
		Country:     "Korea",
		Email:       tokenInfo.Email,
		UserPicture: tokenInfo.Picture,
	}
	exists, err := authService.ExistByUserID()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_DB, nil)
		return
	}
	if exists {
		//get
		log.Println("exists true")
	} else {
		//add
		err = authService.Add()
		if err != nil {
			appG.Response(http.StatusOK, e.ERROR_DB, nil)
			return
		}
	}

	// log.Printf("Name: %s", tokenInfo.Name)
	// log.Printf("Email: %s", tokenInfo.Email)
	// log.Printf("Exp: %d", tokenInfo.Exp)

	//response
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func verifyGoogleIDToken(ctx context.Context, token string) (*TokenInfo, int) {
	aud := "526488632616-h18cgl28r8o4jvubm8nv7jbatl6pcdi5.apps.googleusercontent.com"
	validTok, err := idtoken.Validate(ctx, token, aud)
	if err != nil {
		return nil, e.ERROR_AUTH_CHECK_TOKEN_FAIL
	}
	if validTok.Audience != aud {
		return nil, e.ERROR_AUTH_INVALID_TOKEN
	}
	//tokenInfo := new(TokenInfo)
	var tokenInfo TokenInfo
	err = mapstructure.Decode(validTok.Claims, &tokenInfo)
	if err != nil {
		log.Println("Decode err")
	}

	return &tokenInfo, e.SUCCESS
}
