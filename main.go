package main

import (
	"crypto/hmac"
	"crypto/rand" // most random your computer can generate
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func (u *UserClaims) Valid() error {
	// check if token expired
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expried")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("invalid session ID")
	}

	return nil
}

// `curl -u username:password -v domain.com`
func main() {
	pass := "123456789"

	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in")
	}

	log.Println("Logged in!")

}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while generating bcrypt bash from password %w", err)
	}

	return bs, nil
}

// comparing a hash with the original pass code is a potential source of vulnerabilities.
func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}
	return nil
}

func signMessage(msg []byte) ([]byte, error) {
	hash := hmac.New(sha512.New, keys[currentKid].key)
	_, err := hash.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("error in signMessage while hashing message: %w", err)
	}

	signature := hash.Sum(nil)

	return signature, nil
}

// checks if a signature matches
func checkSignature(msg, sig []byte) (bool, error) {
	newSignature, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("error in checkSig while getting signature of message: %w", err)
	}

	same := hmac.Equal(newSignature, sig)
	return same, nil
}

func createToken(c *UserClaims) (string, error) {
	// create base token
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := t.SignedString(keys[currentKid])
	if err != nil {
		return "", fmt.Errorf("error in createToken when signing token: %w", err)
	}
	return signedToken, nil
}

// cron job
func generateNewKey() error {
	newKey := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, newKey)
	if err != nil {
		return fmt.Errorf("error in generateNewKey while generating key: %w", err)
	}

	uid := uuid.New()
	if err != nil {
		return fmt.Errorf("error in generateNewKey while generating key: %w", err)
	}

	// add the new key to the map of valid keys
	keys[uid.String()] = key{
		key:     newKey,
		created: time.Now(),
	}

	currentKid = uid.String()

	return nil
}

type key struct {
	key []byte

	// if the key is older than a week, make user log in again.
	created time.Time
}

var currentKid = "" // current non-expired key id

// map of keys you used so far. (even expired keys)
// could be cleared out periodically.
var keys = map[string]key{}

func parseToken(signedToken string) (*UserClaims, error) {
	// ParseWithClaims passes the token into anon the function for you to verify
	// the `t` in the anonymous function is the unverified token.
	t, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		// jwt probably verifies the signing methods match, but can do it anyways.
		if t.Method.Alg() != jwt.SigningMethodES512.Alg() {
			return nil, fmt.Errorf("invalid signing algorithm")
		}

		// kid = key id
		kid, ok := t.Header["kid"].(string) // assert its a string
		if !ok {
			return nil, fmt.Errorf("invlaid key ID")
		}

		// when you parse a token, it will look up the key id from within the token,
		// to figure which key was used to sign the message.
		k, ok := keys[kid]
		if !ok {
			return nil, fmt.Errorf("invalid key ID")
		}

		return k, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error in parseToken while parsing token: %w", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("error in parseToken, token is not valid")
	}

	// assert that t.Claims if of type pointer to UserClaims
	return t.Claims.(*UserClaims), nil
}
