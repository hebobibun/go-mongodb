package main

import (
	"context"
	"fmt"
	"go-mongodb/config"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/create-user", createUser)
	e.GET("/users", getUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Age       string             `bson:"age" json:"age"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

func createUser(c echo.Context) error {
	clt := config.MgConnect()
	defer clt.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var table = config.MgCollection("users", clt)

	user1 := User{
		Name:      "Pevita P",
		Age:       "22",
		CreatedAt: time.Now(),
	}

	_, err := table.InsertOne(ctx, &user1)
	if err != nil {
		fmt.Println("Failed to insert user")
	}

	fmt.Println(user1)

	return c.String(http.StatusOK, "Inserted a new user successsfully")
}

func getUser(c echo.Context) error {
	age := c.QueryParam("age")

	client := config.MgConnect()
	defer client.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var users = make([]User, 0)

	coll := config.MgCollection("users", client)

	res, err := coll.Find(
		ctx,
		bson.M{
			"age": age,
		},
	)
	if err != nil {
		return err
	}

	defer res.Close(ctx)

	fmt.Println(res)
	for res.Next(ctx) {
		fmt.Println(1)
		var user User
		err = res.Decode(&user)
		if err != nil {
			return err
		}
		fmt.Println("user :", user)

		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}
