package auth

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/utils/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var JwtWrapper = &JwtInfo{
	AccessTokenKey:  config.AccessTokenKey,
	RefreshTokenKey: config.RefreshTokenKey,
	Issuer:          config.Issuer,
}

var (
	AtExpiresTime = time.Now().Add(15 * time.Minute)
	RtExpiresTime = time.Now().Add(24 * 7 * time.Hour)
)

type JwtInfo struct {
	AccessTokenKey  string
	RefreshTokenKey string
	Issuer          string
}

type Claims struct {
	UserID uint64
	jwt.StandardClaims
}

func (j *JwtInfo) GenerateAccessToken(userID uint64) (*entity.AccessToken, error) {
	atClaims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			Issuer:    j.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	log.Println(atClaims.ExpiresAt)

	tokenSgined, err := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims).SignedString([]byte(j.AccessTokenKey))
	if err != nil {
		log.Println("error when trying to genereate access token with claims,", err)
		return nil, err
	}

	at := &entity.AccessToken{
		AccessToken: tokenSgined,
		ExpiresAt:   atClaims.ExpiresAt,
	}
	return at, nil
}

func (j *JwtInfo) GenerateRefreshToken(userID uint64) (*entity.RefreshToken, error) {
	rtClaims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * 7 * time.Hour).Unix(),
			Issuer:    j.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	tokenSigned, err := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims).SignedString([]byte(j.RefreshTokenKey))
	if err != nil {
		log.Println("error when trying to generate refresh token with claims, ", err)
		return nil, err
	}

	rt := &entity.RefreshToken{
		Uuid:         uuid.New().String(),
		RefreshToken: tokenSigned,
		UserID:       userID,
		ExpiresAt:    rtClaims.ExpiresAt,
	}

	return rt, nil
}

func (j *JwtInfo) GenerateTokenPair(userID uint64) (map[string]interface{}, error) {
	at, err := j.GenerateAccessToken(userID)
	if err != nil {
		return nil, err
	}

	rt, err := j.GenerateRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"access_token":  at,
		"refresh_token": rt,
	}, nil
}

func (j *JwtInfo) ValidateToken(token string) (*Claims, error) {
	parseToken, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.AccessTokenKey), nil
		},
	)

	if err != nil {
		log.Println("error when parsing token with claims, ", err)
		return nil, err
	}

	claims, ok := parseToken.Claims.(*Claims)
	if !ok {
		log.Println("token is not valid")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		log.Println("access token expired, ", err)
		err = errors.New("access token expired")
		return nil, err
	}

	if claims.IssuedAt > time.Now().Unix() {
		log.Println("token is issued at after now, ", err)
		err = errors.New("token is issued at after now")
		return nil, err
	}

	if claims.Issuer != j.Issuer {
		log.Println("token issuer is wrong, ", err)
		err = errors.New("token issuer is wrong")
		return nil, err
	}

	return claims, nil
}

func ExtractToken(bearerToken string) string {
	clientToken := ""
	//Authorization Bearer xxx
	strToken := strings.Split(bearerToken, "Bearer ")
	if len(strToken) == 2 {
		clientToken = strings.TrimSpace(strToken[1])
	}

	return clientToken
}
