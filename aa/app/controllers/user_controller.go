package controllers

import (
	"log"
	"strconv"

	"github.com/revel/revel"
)

type UserController struct {
	*revel.Controller
}

func (c UserController) GetAllUser() revel.Result {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM users"

	rows, err := db.Query(query)
	if err != nil {
		var response ErrorResponse
		response.Status = 500
		response.Message = "Internal server error"
		log.Println(err)
	}

	var user User
	var users []User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address)
		if err != nil {
			var response ErrorResponse
			response.Status = 500
			response.Message = "Internal server error"
			log.Fatal(err.Error())
		} else {
			users = append(users, user)
		}
	}

	var response UsersResponse
	response.Status = 200
	response.Message = "Request success"
	response.Data = users
	return c.RenderJSON(response)
}

func (c UserController) InsertNewUser() revel.Result {
	db := connect()
	defer db.Close()

	name := c.Params.Form.Get("name")
	age, _ := strconv.Atoi(c.Params.Form.Get("age"))
	address := c.Params.Form.Get("address")

	_, errQuery := db.Exec("INSERT INTO users(name, age, address) VALUES (?, ?, ?)",
		name,
		age,
		address,
	)

	if errQuery == nil {
		var response SuccessResponse
		response.Status = 200
		response.Message = "Success"
		return c.RenderJSON(response)
	} else {
		var response ErrorResponse
		response.Status = 400
		response.Message = "Bad Request"
		return c.RenderJSON(response)
	}
}

func (c UserController) UpdateUser() revel.Result {
	db := connect()
	defer db.Close()

	id, _ := strconv.Atoi(c.Params.Form.Get("id"))
	name := c.Params.Form.Get("name")
	age, _ := strconv.Atoi(c.Params.Form.Get("age"))
	address := c.Params.Form.Get("address")

	_, errQuery := db.Exec("UPDATE users SET name = ?, age = ?, address = ? WHERE id =?",
		name,
		age,
		address,
		id,
	)

	if errQuery == nil {
		var response SuccessResponse
		response.Status = 200
		response.Message = "Success"
		return c.RenderJSON(response)
	} else {
		var response ErrorResponse
		response.Status = 400
		response.Message = "Bad Request"
		return c.RenderJSON(response)
	}
}

func (c UserController) DeleteUser() revel.Result {
	db := connect()
	defer db.Close()

	id, _ := strconv.Atoi(c.Params.Form.Get("id"))

	_, errQuery := db.Exec("DELETE FROM users WHERE id = ?",
		id,
	)

	if errQuery == nil {
		var response SuccessResponse
		response.Status = 200
		response.Message = "Success"
		return c.RenderJSON(response)
	} else {
		var response ErrorResponse
		response.Status = 400
		response.Message = "Bad Request"
		return c.RenderJSON(response)
	}
}
