package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	fiber "github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Ticker represents the structure of the API response
type Ticker struct {
	Symbol     string `json:"symbol"`
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
	OpenPrice  string `json:"openPrice"`
	LowPrice   string `json:"lowPrice"`
	HighPrice  string `json:"highPrice"`
	LastPrice  string `json:"lastPrice"`
	Volume     string `json:"volume"`
	BidPrice   string `json:"bidPrice"`
	AskPrice   string `json:"askPrice"`
	At         int64  `json:"at"`
}

func fetchWazirXData() ([]Ticker, error) {
	var tickers []Ticker

	resp, err := http.Get("https://api.wazirx.com/sapi/v1/tickers/24hr")
	if err != nil {
		return tickers, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tickers, err
	}

	err = json.Unmarshal(body, &tickers)
	if err != nil {
		return tickers, err
	}

	return tickers, nil
}

func setupRoutes(app *fiber.App) {

	//db connection
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	rawUsername :=    // Replace with your actual username
	rawPassword :=  // Replace with your actual password

	encodedUsername := url.QueryEscape(rawUsername)
	encodedPassword := url.QueryEscape(rawPassword)
	//connectionString := "mongodb+srv://" + encodedUsername + ":" + encodedPassword + "@yourclusterurl/test?retryWrites=true&w=majority"
	opts := options.Client().ApplyURI("mongodb+srv://" + encodedUsername + ":" + encodedPassword + "@cluster0.juskofl.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// Route for WazirX API Data
	app.Get("/api/wazirx", func(c *fiber.Ctx) error {
		data, err := fetchWazirXData()
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(data)
	})

	// Root Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Crypto Dashboard API!")
	})
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
