package api

import (
	"context"
	"fmt"
	"io/ioutil"
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

type AccessToken struct {
	Token  string
	Expiry int64
}

// @Summary Post Auth
// @Produce  json
// @Param auth body socialAuth true "socialAuth Body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [post]
func GetAuth(c *gin.Context) {

	// //test log
	// body := c.Request.Body
	// value, err := ioutil.ReadAll(body)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// var data map[string]interface{}
	// json.Unmarshal([]byte(value), &data)
	// doc, _ := json.Marshal(data)
	// log.Printf("%s", doc)

	var (
		appG      = app.Gin{C: c}
		authJson  socialAuth
		tokenInfo *TokenInfo
	)

	errCode := app.BindAndValid(c, &authJson)
	if errCode != e.SUCCESS {
		appG.Response(errCode, nil)
		return
	}

	log.Printf("token: %s", authJson.LoginToken)
	switch authJson.LoginType {
	case "google":
		tokenInfo, errCode = verifyGoogleIDToken(c.Request.Context(), authJson.LoginToken)
		if errCode == e.ERROR_AUTH_CHECK_TOKEN_FAIL {
			appG.Response(e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			return
		}
		if errCode == e.ERROR_AUTH_INVALID_TOKEN {
			appG.Response(e.ERROR_AUTH_INVALID_TOKEN, nil)
			return
		}
	case "facebook":
		verifyFacebookAccessToken(authJson.LoginToken)
		return
	default:
		appG.Response(e.ERROR_INVALID_LOGIN_TYPE, nil)
		return
	}

	authService := auth_service.Auth{
		UserID:      tokenInfo.Sub,
		Username:    tokenInfo.Name,
		LoginType:   authJson.LoginType,
		Country:     "Korea",
		Email:       tokenInfo.Email,
		UserPicture: tokenInfo.Picture,
	}
	exists, err := authService.ExistByUserID()
	if err != nil {
		appG.Response(e.ERROR_DB, nil)
		return
	}
	if exists {
		//get
		log.Println("exists true")
	} else {
		//add
		err = authService.Add()
		if err != nil {
			appG.Response(e.ERROR_DB, nil)
			return
		}
	}

	// log.Printf("Name: %s", tokenInfo.Name)
	// log.Printf("Email: %s", tokenInfo.Email)
	// log.Printf("Exp: %d", tokenInfo.Exp)

	//response
	appG.Response(e.SUCCESS, map[string]string{
		"userID":   tokenInfo.Sub,
		"username": tokenInfo.Name,
	})
}

func verifyGoogleIDToken(ctx context.Context, token string) (*TokenInfo, int) {
	aud := "cliend-id"
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

func verifyFacebookAccessToken(token string) {
	clientId := "app-id"
	clientSecret := "app-secret-code"

	appAccessToken := clientId + "|" + clientSecret

	resp, err := http.Get("https://graph.facebook.com/debug_token?input_token=" +
		token + "&access_token=" + appAccessToken)

	if err != nil {
		fmt.Printf("Get: %s\n", err)

		return
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll: %s\n", err)
		return
	}

	log.Printf("parseResponseBody: %s\n", string(response))
}
