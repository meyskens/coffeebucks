package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/OGKevin/go-bunq/bunq"
)

var fromIBAN, toIBAN, amount, token string
var c *bunq.Client

func main() {
	apiKey := os.Getenv("BUNQ_KEY")
	fromIBAN = os.Getenv("FROM_IBAN")
	toIBAN = os.Getenv("TO_IBAN")
	amount = os.Getenv("AMOUNT")
	token = os.Getenv("TOKEN")

	if amount == "" {
		amount = "2.00" // yeah i know quite cheap...
	}
	if apiKey == "" || fromIBAN == "" || toIBAN == "" || token == "" {
		log.Println("You need a user API key set as BUNQ_KEY, IBANs (without spaces) in FROM_IBAN and TO_IBAN and a TOKEN for auth of the ESP32")
		os.Exit(1)
	}
	key, err := bunq.CreateNewKeyPair()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c = bunq.NewClient(ctx, bunq.BaseURLProduction, key, apiKey, "Coffeebucks")
	err = c.Init()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/coffee-payment", coffeePayment)

	s := &http.Server{
		Addr: ":80",
	}
	go func() {
		log.Println("Serving coffee!")
		log.Println(s.ListenAndServe())
		os.Exit(1)
	}()

	// waiting for a signal to exit
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}

func coffeePayment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid form"))
		return
	}

	t := r.Form.Get("token")
	if t != token {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid token"))
		return
	}

	fromID, err := GetAccountIDForIBAN(c, fromIBAN)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	coffeePtr, err := GetAccountPoinerForIBAN(c, toIBAN)
	if err != nil || coffeePtr == nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = c.PaymentService.CreatePaymentBatch(
		fromID,
		bunq.PaymentBatchCreate{
			Payments: []bunq.PaymentCreate{
				{
					Amount: bunq.Amount{
						Currency: "EUR",
						Value:    "2.00",
					},
					CounterpartyAlias: *coffeePtr,
					Description:       "Coffeebucks coffee",
					AllowBunqto:       true,
				},
			},
		},
	)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Enjoy!"))
}
