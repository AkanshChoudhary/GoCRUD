package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int64  `json:"age"`
}

func getUsers(c *gin.Context) {
	jsonFile, err := os.Open("db.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result Users
	json.Unmarshal([]byte(byteValue), &result)
	c.JSON(200, result)
}

func createUser(c *gin.Context) {
	var newUser = User{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(newUser)
	jsonFile, err := os.Open("db.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Users
	json.Unmarshal([]byte(byteValue), &result)

	result.Users = append(result.Users, newUser)

	file, _ := json.MarshalIndent(result, "", " ")

	_ = ioutil.WriteFile("db.json", file, 0644)
	fmt.Print(result)
	c.JSON(200, result)
}

func updateUser(c *gin.Context) {
	var newUser = User{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(newUser)
	jsonFile, err := os.Open("db.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Users
	var newRes Users
	json.Unmarshal([]byte(byteValue), &result)

	for i := 0; i < len(result.Users); i++ {
		if result.Users[i].Name == c.Param("name") {
			newRes.Users = append(result.Users[:i], result.Users[i+1:]...)
		}
	}

	newRes.Users = append(newRes.Users, newUser)

	file, _ := json.MarshalIndent(newRes, "", " ")

	_ = ioutil.WriteFile("db.json", file, 0644)
	fmt.Print(newRes)
	c.JSON(200, newRes)
}

func deleteUser(c *gin.Context) {
	jsonFile, err := os.Open("db.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Users
	json.Unmarshal([]byte(byteValue), &result)

	for i := 0; i < len(result.Users); i++ {
		if result.Users[i].Name == c.Param("name") {
			result.Users = append(result.Users[:i], result.Users[i+1:]...)
		}
	}
	file, _ := json.MarshalIndent(result, "", " ")
	_ = ioutil.WriteFile("db.json", file, 0644)
	fmt.Print(result)
	c.JSON(200, result)
}
func main() {
	//CREATE
	srv := gin.Default()
	srv.POST("/create", createUser)

	//READ
	srv.GET("/getUsers", getUsers)

	//UPDATE
	srv.POST("/update/:name", updateUser)

	//DELETE
	srv.DELETE("/delete/:name", deleteUser)
	srv.Run(":9080")
}
