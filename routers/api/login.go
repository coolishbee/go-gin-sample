package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jameschun7/go-gin-sample/pkg/app"
	"github.com/jameschun7/go-gin-sample/pkg/e"
	"github.com/jameschun7/go-gin-sample/pkg/logging"
	"github.com/jameschun7/go-gin-sample/service/auth_service"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/idtoken"
)

type LoginInfo struct {
	LoginType  string `json:"login_type" binding:"required"`
	LoginToken string `json:"login_token" binding:"required"`
	DeviceID   string `json:"device_id" binding:"required"`
}

type GoogleTokenInfo struct {
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

// @Summary Post Login
// @Produce  json
// @Param auth body LoginInfo true "LoginInfo Body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/login [post]
func Login(c *gin.Context) {

	var (
		appG      = app.Gin{C: c}
		session   = sessions.Default(c)
		loginInfo LoginInfo
		tokenInfo *GoogleTokenInfo
	)

	errCode := app.BindAndValid(c, &loginInfo)
	if errCode != e.SUCCESS {
		appG.Response(errCode, nil)
		return
	}

	log.Printf("token: %s", loginInfo.LoginToken)
	logging.Info("api", "test")
	switch loginInfo.LoginType {
	case "google":
		tokenInfo, errCode = verifyGoogleIDToken(c.Request.Context(), loginInfo.LoginToken)
		if errCode == e.ERROR_AUTH_CHECK_TOKEN_FAIL {
			appG.Response(e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			return
		}
		if errCode == e.ERROR_AUTH_INVALID_TOKEN {
			appG.Response(e.ERROR_AUTH_INVALID_TOKEN, nil)
			return
		}
	case "facebook":
		verifyFacebookAccessToken(loginInfo.LoginToken)
		return
	case "apple":
		verifyAppleIDToken()
	default:
		appG.Response(e.ERROR_INVALID_LOGIN_TYPE, nil)
		return
	}

	authService := auth_service.Auth{
		UserID:      tokenInfo.Sub,
		Username:    tokenInfo.Name,
		LoginType:   loginInfo.LoginType,
		Country:     "Korea",
		Email:       tokenInfo.Email,
		UserPicture: tokenInfo.Picture,
	}
	account, err := authService.ExistByUserID()
	if err != nil {
		appG.Response(e.ERROR_DB, nil)
		return
	}
	//redis session create
	session.Set("userID", tokenInfo.Sub)
	session.Set("loginType", loginInfo.LoginType)
	session.Set("deviceID", loginInfo.DeviceID)
	session.Set("socialToken", loginInfo.LoginToken)
	session.Save()

	//response
	if account != nil {
		//account load
		appG.Response(e.SUCCESS, map[string]string{
			"userID":      account.UserID,
			"userName":    account.Username,
			"loginType":   account.LoginType,
			"userPicture": account.UserPicture,
			"teamID":      account.TeamID,
			"sessionKey":  session.ID(),
		})
		return
	} else {
		//account create
		err = authService.Add()
		if err != nil {
			appG.Response(e.ERROR_DB, nil)
			return
		} else {
			appG.Response(e.SUCCESS, map[string]string{
				"userID":     tokenInfo.Sub,
				"username":   tokenInfo.Name,
				"sessionKey": session.ID(),
			})
		}
	}
}

func verifyGoogleIDToken(ctx context.Context, token string) (*GoogleTokenInfo, int) {
	aud := "526488632616-h18cgl28r8o4jvubm8nv7jbatl6pcdi5.apps.googleusercontent.com"
	validTok, err := idtoken.Validate(ctx, token, aud)
	if err != nil {
		return nil, e.ERROR_AUTH_CHECK_TOKEN_FAIL
	}
	if validTok.Audience != aud {
		return nil, e.ERROR_AUTH_INVALID_TOKEN
	}
	//tokenInfo := new(TokenInfo)
	var tokenInfo GoogleTokenInfo
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

func verifyAppleIDToken() {

	resp, err := http.Get("https://appleid.apple.com/auth/keys")
	if err != nil {
		fmt.Printf("Get: %s\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll: %s\n", err)
		return
	}
	log.Printf("parseBody: %s\n", string(body))
}
