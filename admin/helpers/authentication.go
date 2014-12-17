package helpers

import(
    "github.com/dgrijalva/jwt-go"
    "fmt"
)

func GetToken(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, keyFn)
    if err == nil {
        token, err = validate(token)
    }
    return token, err
}

func validate(token *jwt.Token) (*jwt.Token, error) {
    var err error
    if !token.Valid {
        err = fmt.Errorf("invalid token")
    }
    return token, err
}

func keyFn(token *jwt.Token) (interface{}, error) {
    return []byte("WOW,MuchShibe,ToDogge"), nil
}