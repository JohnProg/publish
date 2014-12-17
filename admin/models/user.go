package models

import (
  "code.google.com/p/go.crypto/bcrypt"
  "fmt"
  "encoding/json"
  "github.com/dgrijalva/jwt-go"
  "time"
  //"publish/admin/helpers"
  "log"
  // "io"
  // "io/ioutil"
  // "os"
  //"labix.org/v2/mgo/bson"
)

type User struct {
  Id int `json:"id"`
  Username string `json:"username,omitempty"`
  FirstName string `json:"first_name,omitempty"`
  LastName string `json:"last_name,omitempty"`
  Password []byte `json:"-"`
}


//SetPassword takes a plaintext password and hashes it with bcrypt and sets the
//password field to the hash.
func (u *User) SetPassword(password string) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err) //this is a panic because bcrypt errors on invalid costs
	}
	u.Password = hpass
	fmt.Println("hpass: " + string(hpass))
}
//Login validates and returns a user object if they exist in the database.
// func Login(password string) (u *User, err error) {
// 	// err = ctx.C("users").Find(bson.M{"username": username}).One(&u)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
// 	if err != nil {
// 		u = nil
// 	}
// 	return
// }

func (u *User) Login(password string) (tokenString string, err error) {
	err = nil
	fmt.Println([]byte(password))

	// err = ctx.C("users").Find(bson.M{"username": username}).One(&u)
	// if err != nil {
	// 	return
	// }
	
	// err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	// if err != nil {

	// }

	// return

	// Hashing the password with the cost of 10

    // Comparing the password with the hash
    err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
    if err != nil {
    	//fmt.Printf("[%s] != \n[%s]\n", u.Password, []byte(password))
    	//tokenString = ""
      log.Println("login failed, try again.")
    } else {
    	t, err := u.CreateToken()
    	//helpers.PanicIf(err)
      log.Println(err)
    	tokenString = t
    }
    //fmt.Println(err)
    return
}

func (u *User) CreateToken() (tokenString string, err error){
	var secretKey string = "WOW,MuchShibe,ToDogge"
	// get the key
	// keyData, err := loadData(key)
	// if err != nil {
	// 	return fmt.Errorf("Couldn't read key: %v", err)
	// }
    // Create the token
    token := jwt.New(jwt.GetSigningMethod("HS256"))
    // Set some claims
    user, err := json.Marshal(u)
    fmt.Println("create token ---")
    // fmt.Println(user)
    fmt.Println(string(user))

    token.Claims["user"] = string(user)
    token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
    // Sign and get the complete encoded token as a string
    t, err := token.SignedString([]byte(secretKey))
    //helpers.PanicIf(err)
    log.Println(err)
    tokenString = t
    return
}
