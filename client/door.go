package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"

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

	cardNumber := pickCardNumber()
	db.Update(strconv.Itoa(cardNumber), "door")

	return c.JSON(http.StatusOK, deck.Cards[cardNumber])

}

func pickCardNumber() int {
	cardNumber := rand.Intn(2)
	val, err := db.Read(strconv.Itoa(cardNumber))
	if err != nil {
		fmt.Println(err)
	}
	if val == "door" {
		return pickCardNumber()
	}
	return cardNumber
}
