package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

const TokenExpireDuration = time.Hour * 24

var mySecret = []byte("夏天夏天悄悄过去")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type MyClaims1 struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userID,
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(), // 过期时间
			Issuer:    "bluebell",                                                                        // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	//if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
	//	return claims, nil
	//}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

// GenToken1 生产access token 和 refresh token
func GenToken1(userID int64) (aToken, rToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims1{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
			Issuer:    "bluebell",
		},
	}
	// 加密并获取完整的编码后的字符串Token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)

	// refresh token 不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
		Issuer:    "bluebell",
	}).SignedString(mySecret)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

func KeyFunc(token *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

// ParseToken1 解析aToken
func ParseToken1(tokenString string) (claim *MyClaims1, err error) {
	// 解析token
	var token *jwt.Token
	claim = new(MyClaims1)
	token, err = jwt.ParseWithClaims(tokenString, claim, KeyFunc)
	if err != nil {
		return
	}
	if !token.Valid { // 校验token
		err = errors.New("invalid Token")
	}
	return
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token无效直接返回
	if _, err = jwt.Parse(rToken, KeyFunc); err != nil {
		return
	}
	// 从旧的access token中解析claims数据
	var claims MyClaims1
	_, err = jwt.ParseWithClaims(aToken, &claims, KeyFunc)
	v, _ := err.(jwt.ValidationError)

	// 当access token是过期的错误，并且refresh token 没有过期时就会创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken1(claims.UserID)
	}
	return
}
