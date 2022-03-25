package main

import (
	"context"
	"encoding/json"
	"fmt"
	"grpc-protobuf/common/config"
	model "grpc-protobuf/common/model"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "z01",
		Name:     "Zakir",
		Password: "!ZX404s",
		Gender:   model.UserGender_FEMALE,
	}

	user2 := model.User{
		Id:       "n002",
		Name:     "Nabila Rozan",
		Password: "PasswordTralala",
		Gender:   model.UserGender(model.UserGender_value["FEMALE"]),
	}

	garage1 := model.Garage{
		Id:   "G01",
		Name: "Gundam",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}

	garage2 := model.Garage{
		Id:   "f001",
		Name: "Frostwing",
		Coordinate: &model.GarageCoordinate{
			Latitude:  32.123123123,
			Longitude: 11.1231313123,
		},
	}
	garage3 := model.Garage{
		Id:   "u001",
		Name: "Undercity",
		Coordinate: &model.GarageCoordinate{
			Latitude:  22.123123123,
			Longitude: 123.1231313123,
		},
	}

	user := serviceUser()
	fmt.Println("\n", "===========> user test")

	// register user1
	user.Register(context.Background(), &user1)

	// register user2
	user.Register(context.Background(), &user2)

	// show all registered users
	res1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))

	garage := serviceGarage()

	fmt.Println("\n", "===========> garage test A")

	// add garage1 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})

	// add garage2 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage2,
	})

	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage3,
	})

}
