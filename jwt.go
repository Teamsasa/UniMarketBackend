package main

import (
	"fmt"
	"net/http"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

// cookieからトークンを取得
func getTokenFromCookie(r *http.Request, cookieName string) (string, error) {
	// クッキーからトークンを取得
	c, err := r.Cookie(cookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			return "", fmt.Errorf("%vが見つかりません: %v", cookieName, err)
		}
		return "", fmt.Errorf("不正なリクエスト：クッキーの取得エラー: %v", err)
	}

	return c.Value, nil
}

// トークンを取得し検証
func parseToken(tknStr string) (*jwt.Token, error) {
	// JWKセットを指定されたURLから取得
	jwks, err := keyfunc.NewDefault([]string{jwksURL})
	if err != nil {
		return nil, fmt.Errorf("指定されたURLのリソースからJWKセットを作成できませんでした: %v", err)
	}

	// トークンをパース
	token, err := jwt.Parse(tknStr, jwks.Keyfunc)
	if err != nil {
		return nil, fmt.Errorf("JWTのパースに失敗しました: %v", err)
	}

	// トークンが有効かどうかをチェック
	if !token.Valid {
		return nil, fmt.Errorf("認証失敗：無効なトークン")
	}

	return token, nil
}

// トークンを検証する
func validateAccessToken(r *http.Request) error {
	accessToken, err := getTokenFromCookie(r, "accessToken")
	if err != nil {
		return err
	}
	_, err = parseToken(accessToken)
	if err != nil {
		return err
	}
	return nil
}

/** トークンから指定されたユーザーデータを取得
 * 取得したい値を引数に設定、"tokenはidToken"を使用
 * id -> "sub"
 * username -> "cognito:username"
 * email -> "email"
 */
func getValueFromToken(token *jwt.Token, claimKey string) (string, error) {
	// クレームから指定されたキーの値を取得
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if value, ok := claims[claimKey].(string); ok {
			return value, nil
		}
		return "", fmt.Errorf("クレームからキーの値を取得できませんでした: claims: %v, key: %s", claims, claimKey)
	}
	return "", fmt.Errorf("トークンが無効です: token: %v", token)
}
