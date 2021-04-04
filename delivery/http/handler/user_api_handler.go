package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/solkn/soccer/api/auth"
	"github.com/solkn/soccer/api/entity"
	user "github.com/solkn/soccer/api/user"
	"github.com/solkn/soccer/api/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserApiHandler struct {
	userService user.UsersService
}

func NewUserApiHandler(userServices user.UsersService) *UserApiHandler {
	return &UserApiHandler{userService: userServices}
}
func (uph *UserApiHandler) Authenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
func (uph *UserApiHandler) Authorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userX entity.User
		uid, err := auth.ExtractTokenID(r)
		userX.Id = uid
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		user, err := uph.userService.User(uid)
		fmt.Println(user.Role)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if strings.ToUpper(user.Role.Name) != strings.ToUpper("ADMIN") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
		next(w, r)
	}

}
func (uph *UserApiHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	users, errs := uph.userService.User(uint32(id))

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(users, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uph *UserApiHandler) GetUserRoles(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var user1 entity.User
	err := json.Unmarshal(body, &user1)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return

	}
	userRoles, errs := uph.userService.UserRoles(&user1)
	if len(errs) > 0 {

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(userRoles, "", "\t\t")

	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uph *UserApiHandler) IsEmailExists(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]

	ok := uph.userService.EmailExists(email)
	output, err := json.MarshalIndent(ok, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uph *UserApiHandler) IsPhoneExists(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	phone := params["phone"]

	ok := uph.userService.PhoneExists(phone)
	output, err := json.MarshalIndent(ok, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uph *UserApiHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, errs := uph.userService.Users()
	if errs != nil {
		fmt.Println("err1")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(users, "", "\t\t")

	if err != nil {
		fmt.Println("err2")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(output)
	if err != nil {
		fmt.Println("err3")
		fmt.Println(err.Error())
	}
	return

}
func (uph *UserApiHandler) GetUserByUsernameAndPassword(w http.ResponseWriter, r *http.Request) {

	body := utils.BodyParser(r)
	var user1 entity.User
	err := json.Unmarshal(body, &user1)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	user2, errs := uph.userService.UserByUserName(user1)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	output, err := json.MarshalIndent(*user2, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return

}
func (uph *UserApiHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var user1 entity.User
	err := json.Unmarshal(body, &user1)
	if err != nil {
		print(err.Error())
		utils.ToJson(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user1.Password), 12)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user1.Password = string(hashedPassword)
	user2, errs := uph.userService.StoreUser(&user1)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}
	output, err := json.MarshalIndent(user2, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
func (uph *UserApiHandler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--logging---")
	body := utils.BodyParser(r)
	var user1 entity.User

	var token string
	var expiry int64

	err := json.Unmarshal(body, &user1)
	if err != nil {
		fmt.Println("error 263")
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	fmt.Println(user1)
	userFromDatabase, errs := uph.userService.UserByUserName(user1)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	err = bcrypt.CompareHashAndPassword([]byte(user1.Password), []byte(userFromDatabase.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	token, err = auth.CreateToken(userFromDatabase.Id)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	expiry = time.Now().Add(time.Minute * 30).Unix()

	output, err := json.MarshalIndent(userFromDatabase, "", "\t\t")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	expiryString := strconv.Itoa(int(expiry))
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("token", token)
	w.Header().Add("expiry_date", expiryString)
	_, _ = w.Write(output)
	return
}

func (uph *UserApiHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var token string
	var expiry int64
	body := utils.BodyParser(r)
	var user1 entity.User
	err := json.Unmarshal(body, &user1)
	if err != nil {
		print(err.Error())
		utils.ToJson(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user1.Password), 12)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user1.Password = string(hashedPassword)
	user2, errs := uph.userService.StoreUser(&user1)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}

	output, err := json.MarshalIndent(user2, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	token, err = auth.CreateToken(user2.Id)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	expiry = time.Now().Add(time.Minute * 30).Unix()
	expiryString := strconv.Itoa(int(expiry))
	w.Header().Add("token", token)
	w.Header().Add("token", expiryString)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
func (uph *UserApiHandler) PutUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user1, errs := uph.userService.User(uint32(id))

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	l := r.ContentLength

	body := make([]byte, l)

	_, _ = r.Body.Read(body)

	_ = json.Unmarshal(body, &user1)
	user1, errs = uph.userService.UpdateUser(user1)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(user1, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
func (uph *UserApiHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := uph.userService.DeleteUser(uint32(id))

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
