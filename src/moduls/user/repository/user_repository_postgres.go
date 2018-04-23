package repository

import (
	"database/sql"
	"fmt"

	"github.com/jemmycalak/calak_chatdate_postgre/src/moduls/user/model"
)

type userReposityPostgres struct {
	db *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) *userReposityPostgres {
	return &userReposityPostgres{db}
}

func (r *userReposityPostgres) Save(user *model.User) error {
	query := `INSERT INTO "t_user"("firtsname", "lastname", "email", "password", "imageprofile", "createat", "updateat") VALUES ($1, $2, $3, $4, $5, $6, $7)`

	statement, err := r.db.Prepare(query)

	if err != nil {
		fmt.Println(" error Prepare query")
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Lastname, user.Lastname, user.Email, user.Password, user.Imageprofile, user.CreateAt, user.UpdateAt)

	if err != nil {
		fmt.Println(" error Execute query")
		return err
	}
	return nil
}

func (r *userReposityPostgres) Update(id string, user *model.User) error {
	query := `UPDATE "t_user" SET "firtsname"=$1, "lastname"=$2, "email"=$3, "password"=$4, "imageprofile"=$5, "updateat"=$6 WHERE "iduser"=$7 `

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Firstname, user.Lastname, user.Email, user.Password, user.Imageprofile, user.UpdateAt, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *userReposityPostgres) Delete(id string) error {
	query := `DELETE FROM "t_user" WHERE "iduser"= $1`

	statement, err := r.db.Prepare(query)

	if err != nil {
		fmt.Println("error prepare delete")
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(id)

	if err != nil {
		fmt.Println("error exc query")
		return err
	}

	return nil
}

func (r *userReposityPostgres) FindById(id string) (*model.User, error) {
	query := `SELECT * FROM "t_user" WHERE "iduser" = $1`

	var muser model.User

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&muser.Id, &muser.Firstname, &muser.Lastname, &muser.Email, &muser.Password, &muser.Imageprofile, &muser.CreateAt, &muser.UpdateAt)

	if err != nil {
		return nil, err
	}

	return &muser, nil
}

func (r *userReposityPostgres) FindAll() (model.Users, error) {
	query := `SELECT * FROM "t_user"`

	var musers model.Users

	rows, err := r.db.Query(query)

	if err != nil {
		fmt.Println("error query")
		// return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var nuser model.User

		err := rows.Scan(&nuser.Id, &nuser.Firstname, &nuser.Lastname, &nuser.Email, &nuser.Password, &nuser.Imageprofile, &nuser.CreateAt, &nuser.UpdateAt)

		if err != nil {
			fmt.Println("error loop data")
			return nil, err
		}
		musers = append(musers, nuser)
	}

	return musers, nil
}
