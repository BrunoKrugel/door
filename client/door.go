package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"

	"github.com/BrunoKrugel/door/db"
	"github.com/BrunoKrugel/door/model"

	"github.com/labstack/echo"
)

func PickDoor(c echo.Context) error {
	jsonFile, err := os.Open("resources/door.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var deck model.Deck
	json.Unmarshal(byteValue, &deck)
	defer jsonFile.Close()

	cardNumber := rand.Intn(2)
	db.Write(cardNumber)
	aux := db.Read(3)
	if aux != nil {
		return c.String(http.StatusInternalServerError, "Error reading from database")
	}

	return c.JSON(http.StatusOK, deck.Cards[cardNumber])

}
