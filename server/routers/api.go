package routers

import (
	"Gopress/server/database"
	"Gopress/server/jsonWebToken"
	"Gopress/server/models"
	"fmt"
	"regexp"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
)

// RegisterStruct need documantion
type RegisterStruct struct {
	Name     string
	Surname  string
	Email    string
	Password string
}

// LoginStruct need documantion
type LoginStruct struct {
	Email    string
	Password string
}

// RecaptchaResponse need documantion
type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []int     `json:"error-codes"`
}

// ContactStruct need documantion
type ContactStruct struct {
	Name    string
	Email   string
	Message string
	Captcha string
}

// SimpleResponse need documantion
type SimpleResponse struct {
	Success bool
	Message string
}

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// APIHandler function handling every single API request coming from client.
func APIHandler(api iris.Party) {
	api.Post("/register", func(ctx iris.Context) {
		req := &RegisterStruct{}
		err := ctx.ReadJSON(req)
		if err != nil {
			res := &SimpleResponse{false, "Someting happend while reading your informations. Please try again..."}
			ctx.JSON(res)
			return
		}

		if req.Name == "" || req.Surname == "" || req.Email == "" || req.Password == "" {
			res := &SimpleResponse{false, "Missing fields. Please try again..."}
			ctx.JSON(res)
			return
		}

		if !emailRegexp.MatchString(req.Email) {
			res := &SimpleResponse{false, "Unsupported email address."}
			ctx.JSON(res)
			return
		}

		db, err := database.GetConnection()

		if err != nil {
			res := &SimpleResponse{false, "We have a problem with our database. Please try again later..."}
			ctx.JSON(res)
			return
		}

		// Is user already registered ?
		var isUserExist string
		db.QueryRow("SELECT u.email FROM users u WHERE u.email = ?", req.Email).Scan(&isUserExist)

		// If yes...
		if isUserExist != "" {
			res := &SimpleResponse{false, "User already registered."}
			ctx.JSON(res)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

		if err != nil {
			res := &SimpleResponse{false, "Invalid password. Please use a different password."}
			ctx.JSON(res)
			return
		}

		_, err = db.Query("INSERT INTO users SET name=?, surname=?, email=?, password=?, email_verification=?, date=NOW(), admin=?", req.Name, req.Surname, req.Email, hash, 0, 0)
		defer db.Close()
		if err != nil {
			res := &SimpleResponse{false, "The error occurred during registration. Please try again..."}
			ctx.JSON(res)
			return
		}

		res := &SimpleResponse{true, "Registration Successful."}
		ctx.JSON(res)
	})

	api.Post("/login", func(ctx iris.Context) {

		req := &LoginStruct{}
		err := ctx.ReadJSON(req)
		if err != nil {
			res := &SimpleResponse{false, "Someting happend while reading your informations. Please try again..."}
			ctx.JSON(res)
			return
		}
		if req.Email == "" || req.Password == "" {
			res := &SimpleResponse{false, "Missing fields. Please try again..."}
			ctx.JSON(res)
			return
		}

		if !emailRegexp.MatchString(req.Email) {
			res := &SimpleResponse{false, "Unsupported email address."}
			ctx.JSON(res)
			return
		}

		db, err := database.GetConnection()

		if err != nil {
			res := &SimpleResponse{false, "We have a problem with our database. Please try again later..."}
			ctx.JSON(res)
			return
		}

		rows, err := db.Query("SELECT u.name, u.surname, u.email, u.password, u.email_verification, u.date, u.admin FROM users u WHERE u.email=?", req.Email)
		defer db.Close()
		if err != nil {
			res := &SimpleResponse{false, "The error occurred during registration. Please try again..."}
			ctx.JSON(res)
			return
		}

		// UserInfo need documantion
		type UserInfo struct {
			jwt.StandardClaims
			Name              string `json:"name"`
			Surname           string `json:"surname"`
			Email             string `json:"email"`
			Password          string `json:"password"`
			EmailVerification int    `json:"email_verification"`
			Date              string `json:"date"`
			Admin             int    `json:"admin"`
		}
		n := UserInfo{}

		users := []UserInfo{}
		for rows.Next() {
			// Armazena os valores em variaveis
			var name, surname, email, password, date string
			var email_verification, admin int

			// Faz o Scan do SELECT
			err = rows.Scan(&name, &surname, &email, &password, &email_verification, &date, &admin)
			if err != nil {
				panic(err.Error())
			}

			// Envia os resultados para a struct
			n.Name = name
			n.Surname = surname
			n.Email = email
			n.Password = password
			n.EmailVerification = email_verification
			n.Date = date
			n.Admin = admin

			// Junta a Struct com Array
			users = append(users, n)
		}

		if len(users) == 0 {
			res := &SimpleResponse{false, "User not found. Please try again later..."}
			ctx.JSON(res)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(req.Password)); err != nil {
			// TODO: Properly handle error
			res := &SimpleResponse{false, "Wrong password or email address. Please try again later..."}
			ctx.JSON(res)
			return
		}

		token, err := jsonWebToken.Make(&users[0])
		if err != nil {
			res := &SimpleResponse{false, "Something happend while creating session for you. Please try again later..."}
			ctx.JSON(res)
			return
		}

		fmt.Println("%v\n", token)
	})

	api.Post("/contact", func(ctx iris.Context) {
		req := &ContactStruct{}
		err := ctx.ReadJSON(req)
		if err != nil {
			res := &SimpleResponse{false, "Bad request. Please try again later..."}
			ctx.JSON(res)
			return
		}

		if req.Captcha == "" || req.Email == "" || req.Message == "" || req.Name == "" {
			res := &SimpleResponse{false, "Missing fields. Please try again later..."}
			ctx.JSON(res)
			return
		}

		captcha := Models.CaptchaConfirm(req.Captcha)

		if captcha == false {
			res := &SimpleResponse{false, "Please prove that you are not a robot."}
			ctx.JSON(res)
			return
		}

		db, err := database.GetConnection()

		if err != nil {
			res := &SimpleResponse{false, "Database connection error. Please contact us with email."}
			ctx.JSON(res)
			return
		}

		_, err = db.Query("INSERT INTO contacts SET email=?, message=?, name=?", req.Email, req.Message, req.Name)
		defer db.Close()
		if err != nil {
			res := &SimpleResponse{false, "Error occurred while retrieving your request."}
			ctx.JSON(res)
			return
		}
		res := &SimpleResponse{true, "Your request was received. I will communicate with you for the shortest time."}
		ctx.JSON(res)
	})
	api.Get("/parse", func(ctx iris.Context) {
		token, err := jsonWebToken.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoib3RpYWkxMCIsImFnZSI6MzB9.Rs-LpJmqeg8dvj7ft4K1FS7y73kd2BcN4NmsEap31yU")
		if err != nil {
			fmt.Println(err)
		}
		claim, ok := token.(jwt.MapClaims)
		if !ok {
			fmt.Println(err)
		}
		fmt.Println(claim)
		ctx.View("index.html")
	})
}

func ParseJwt(tk string) {
	token, err := jsonWebToken.ParseToken(tk)
	if err != nil {
		fmt.Println(err)
	}
	claim, ok := token.(jwt.MapClaims)
	if !ok {
		fmt.Println(err)
	}
	fmt.Println(claim)
}

// Password...
/** [NOTES:]
userPassword1 := "some user-provided password"

hash, err := bcrypt.GenerateFromPassword([]byte(userPassword1), bcrypt.DefaultCost)
if err != nil {
	// TODO: Properly handle error
	log.Fatal(err)
}
fmt.Println("Hash to store:", string(hash))

userPassword2 := "some user-provided passssssword"
hashFromDatabase := hash

if err := bcrypt.CompareHashAndPassword(hashFromDatabase, []byte(userPassword2)); err != nil {
	// TODO: Properly handle error
	log.Fatal(err)
}

fmt.Println("Password was correct!")
**/
