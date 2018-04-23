package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/jemmycalak/calak_chatdate_postgre/config"
	"github.com/jemmycalak/calak_chatdate_postgre/src/moduls/user/model"
	"github.com/jemmycalak/calak_chatdate_postgre/src/moduls/user/repository"
)

var stts, msg string

func main() {

	router := mux.NewRouter()
	fmt.Println("App is runing now !")

	router.HandleFunc("/api/v1/addUser", addUser).Methods("POST")
	router.HandleFunc("/api/v1/showUsers", showUsers).Methods("GET")
	router.HandleFunc("/api/v1/updateUser", updateUser).Methods("POST")
	router.HandleFunc("/api/v1/deleteUser", deleteUser).Methods("POST")
	router.HandleFunc("/api/v1/showUser", findById).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func showUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("showUser was called")

	db, err := config.GetPostgresDB()
	if err != nil {
		responseError(w, http.StatusBadRequest, "Database not connected")
	}

	repositoryUser := repository.NewUserRepositoryPostgres(db)

	dtUsers, err := repositoryUser.FindAll()
	if err != nil {
		responseError(w, http.StatusBadRequest, "No data found")
	}

	responseJSON(w, http.StatusOK, dtUsers)

}

func findById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("findById was called")

	muser := model.NewUser()
	_ = json.NewDecoder(r.Body).Decode(&muser)

	db, err := config.GetPostgresDB()
	if err != nil {
		responseError(w, http.StatusBadRequest, "Database not connected")
	}

	repositoryUser := repository.NewUserRepositoryPostgres(db)

	dtUser, errs := repositoryUser.FindById(muser.Id)
	if errs != nil {
		responseError(w, http.StatusBadRequest, "Data not found")
	} else {
		responseJSON(w, http.StatusOK, dtUser)
	}

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("updateuser was called")

	muser := model.NewUser()
	_ = json.NewDecoder(r.Body).Decode(&muser)

	db, err := config.GetPostgresDB()
	if err != nil {
		responseError(w, http.StatusBadRequest, "Database not connected")
	}

	repositoryUser := repository.NewUserRepositoryPostgres(db)
	errs := repositoryUser.Update(muser.Id, muser)

	if errs != nil {
		responseError(w, http.StatusBadRequest, "Faild update user")
	} else {
		msg = "data updated"
		stts = "true"
		result := map[string]string{"msg": msg, "status": stts}
		responseJSON(w, http.StatusOK, result)
	}

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("delete user was called")

	muser := model.NewUser()
	_ = json.NewDecoder(r.Body).Decode(&muser)

	db, err := config.GetPostgresDB()
	if err != nil {
		responseError(w, http.StatusBadRequest, "Database not connected")
	}

	fmt.Println("iduser:", muser.Id)

	repositoryUser := repository.NewUserRepositoryPostgres(db)
	errs := repositoryUser.Delete(muser.Id)
	if errs != nil {
		responseError(w, http.StatusBadRequest, "faild delete user")
	} else {
		msg = "user deleted"
		stts = "true"
		result := map[string]string{"msg": msg, "status": stts}
		responseJSON(w, http.StatusOK, result)
	}
}

func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("adduser was called")

	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	////////////////this method for test apps was connected to db or not////////////////
	// err = db.Ping()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	///////////////////////////////////////////////////////////////////////////////////

	muser := model.NewUser()

	_ = json.NewDecoder(r.Body).Decode(&muser)
	repositoryUser := repository.NewUserRepositoryPostgres(db)

	errs := repositoryUser.Save(muser)

	if errs != nil {
		responseError(w, http.StatusBadRequest, "Faild save data")
	} else {
		fmt.Println("data saved")
		msg = "data saved"
		stts = "true"
		result := map[string]string{"status": stts, "msg": msg}
		responseJSON(w, http.StatusOK, result)
	}

}

func responseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func responseError(w http.ResponseWriter, code int, msg string) {
	responseJSON(w, code, map[string]string{"msg": msg, "status": "false"})
}
