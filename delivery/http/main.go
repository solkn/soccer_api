package main

import (
	"github.com/gorilla/mux"
	"github.com/solkn/soccer/api/delivery/http/handler"
	"github.com/solkn/soccer/api/entity"
	ur "github.com/solkn/soccer/api/user/repository"
	us "github.com/solkn/soccer/api/user/services"

	cr "github.com/solkn/soccer/api/club/repository"
	cs "github.com/solkn/soccer/api/club/services"
	fr "github.com/solkn/soccer/api/fixture/repository"
	fs "github.com/solkn/soccer/api/fixture/services"

	rr "github.com/solkn/soccer/api/result/repository"
	rs "github.com/solkn/soccer/api/result/services"

	sr "github.com/solkn/soccer/api/scorer/repository"
	ss "github.com/solkn/soccer/api/scorer/services"

	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"

	_ "github.com/lib/pq"
)

func createTables(dbConn *gorm.DB) []error {
	dbConn.DropTableIfExists(&entity.Role{}, &entity.User{}, &entity.Club{}, &entity.Fixture{}, &entity.Result{},
		&entity.Scorer{}).GetErrors()
	errs := dbConn.CreateTable(&entity.Role{}, &entity.User{}, &entity.Club{}, &entity.Fixture{}, &entity.Result{},
		&entity.Scorer{}).GetErrors()
	dbConn.Debug().Model(&entity.User{}).AddForeignKey("role_id", "roles(Id)", "cascade", "cascade")
	dbConn.Debug().Model(&entity.Result{}).AddForeignKey("fixture_id", "fixtures(Id)", "cascade", "cascade")
	dbConn.Debug().Model(&entity.Scorer{}).AddForeignKey("result_id", "results(Id)", "cascade", "cascade")
	dbConn.Debug().Model(&entity.Scorer{}).AddForeignKey("club_id", "clubs(Id)", "cascade", "cascade")

	if len(errs) > 0 {
		return errs
	}
	return nil
}
func main() {

	dbconn, err := gorm.Open("postgres",
		"postgres://postgres:solomon@localhost/soccer?sslmode=disable")
	if dbconn != nil {
		defer dbconn.Close()
	}
	if err != nil {
		panic(err)
	}
	//createTables(dbconn)

	router := mux.NewRouter()

	roleGormRepo := ur.NewRoleGormRepo(dbconn)
	roleService := us.NewRoleService(roleGormRepo)
	roleHandler := handler.NewRoleApiHandler(roleService)

	usersRepo := ur.NewUserGormRepo(dbconn)
	usersService := us.NewUserService(usersRepo)
	usersHandler := handler.NewUserApiHandler(usersService)

	router.HandleFunc("/v1/role", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.GetRoles))).Methods("GET")
	router.HandleFunc("/v1/roles/{name}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.GetRoleByName))).Methods("GET")
	router.HandleFunc("/v1/role/{id}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.GetRoleByID))).Methods("GET")
	router.HandleFunc("/v1/role", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.PostRole))).Methods("POST")
	router.HandleFunc("/v1/role/{id}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.PutRole))).Methods("PUT")
	router.HandleFunc("/v1/role/{id}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.DeleteRole))).Methods("DELETE")

	router.HandleFunc("/v1/admin/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUser))).Methods("GET")
	router.HandleFunc("/v1/admin/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUsers))).Methods("GET")
	router.HandleFunc("/v1/admin/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PutUser))).Methods("PUT")
	router.HandleFunc("/v1/admin/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PostUser))).Methods("POST")
	router.HandleFunc("/v1/admin/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.DeleteUser))).Methods("DELETE")
	router.HandleFunc("/v1/admin/email/{email}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.IsEmailExists))).Methods("GET")
	router.HandleFunc("/v1/admin/phone/{phone}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.IsPhoneExists))).Methods("GET")
	router.HandleFunc("/v1/admin/check", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUserByUsernameAndPassword))).Methods("POST")

	router.HandleFunc("/v1/user/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PostUser))).Methods("POST")
	router.HandleFunc("/v1/user/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUsers))).Methods("GET")
	router.HandleFunc("/v1/user/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUser))).Methods("GET")
	router.HandleFunc("/v1/user/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PutUser))).Methods("PUT")
	router.HandleFunc("/v1/user/email/{email}", usersHandler.IsEmailExists).Methods("GET")
	router.HandleFunc("/v1/user/phone/{phone}", usersHandler.IsPhoneExists).Methods("GET")
	router.HandleFunc("/v1/user/check", usersHandler.GetUserByUsernameAndPassword).Methods("POST")
	router.HandleFunc("/v1/user/login", usersHandler.Login).Methods("POST")
	router.HandleFunc("/v1/user/signup", usersHandler.SignUp).Methods("POST")

	clubsRepo := cr.NewClubGormRepo(dbconn)
	clubsService := cs.NewClubService(clubsRepo)
	clubsHandler := handler.NewClubApiHandler(clubsService)

	router.HandleFunc("/v1/club", clubsHandler.GetClubs).Methods("GET")
	router.HandleFunc("/v1/club/{id}", clubsHandler.GetClub).Methods("GET")
	router.HandleFunc("/v1/club", usersHandler.Authenticated(usersHandler.Authorized(clubsHandler.PostClub))).Methods("POST")
	router.HandleFunc("/v1/club/{id}", usersHandler.Authenticated(usersHandler.Authorized(clubsHandler.PutClub))).Methods("PUT")
	router.HandleFunc("/v1/club/{id}", usersHandler.Authenticated(usersHandler.Authorized(clubsHandler.DeleteClub))).Methods("DELETE")

	fixturesRepo := fr.NewFixtureGormRepo(dbconn)
	fixturesService := fs.NewClubService(fixturesRepo)
	fixturesHandler := handler.NewFixtureApiHandler(fixturesService)

	router.HandleFunc("/v1/fixture", fixturesHandler.GetFixtures).Methods("GET")
	router.HandleFunc("/v1/fixture/{id}", fixturesHandler.GetFixture).Methods("GET")
	router.HandleFunc("/v1/fixture", usersHandler.Authenticated(usersHandler.Authorized(fixturesHandler.PostFixture))).Methods("POST")
	router.HandleFunc("/v1/fixture/{id}", usersHandler.Authenticated(usersHandler.Authorized(fixturesHandler.PutFixture))).Methods("PUT")
	router.HandleFunc("/v1/fixture/{id}", usersHandler.Authenticated(usersHandler.Authorized(fixturesHandler.DeleteFixture))).Methods("DELETE")

	resultsRepo := rr.NewResultGormRepo(dbconn)
	resultsService := rs.NewClubService(resultsRepo)
	resultsHandler := handler.NewResultApiHandler(resultsService)

	router.HandleFunc("/v1/result", resultsHandler.GetResults).Methods("GET")
	router.HandleFunc("/v1/result/{id}", resultsHandler.GetResult).Methods("GET")
	router.HandleFunc("/v1/result", usersHandler.Authenticated(usersHandler.Authorized(resultsHandler.PostResult))).Methods("POST")
	router.HandleFunc("/v1/result/{id}", usersHandler.Authenticated(usersHandler.Authorized(resultsHandler.PutResult))).Methods("PUT")
	router.HandleFunc("/v1/result/{id}", usersHandler.Authenticated(usersHandler.Authorized(resultsHandler.DeleteResult))).Methods("DELETE")

	scorersRepo := sr.NewClubGormRepo(dbconn)
	scorersService := ss.NewScorerService(scorersRepo)
	scorersHandler := handler.NewScorerApiHandler(scorersService)

	router.HandleFunc("/v1/scorer", scorersHandler.GetScorers).Methods("GET")
	router.HandleFunc("/v1/scorer/{id}", scorersHandler.GetScorer).Methods("GET")
	router.HandleFunc("/v1/scorer", usersHandler.Authenticated(usersHandler.Authorized(scorersHandler.PostScorer))).Methods("POST")
	router.HandleFunc("/v1/scorer/{id}", usersHandler.Authenticated(usersHandler.Authorized(scorersHandler.PutScorer))).Methods("PUT")
	router.HandleFunc("/v1/scorer/{id}", usersHandler.Authenticated(usersHandler.Authorized(scorersHandler.DeleteScorer))).Methods("DELETE")
	err = http.ListenAndServe("192.168.56.1:8080", router)

	if err != nil {
		panic(err)
	}

}
