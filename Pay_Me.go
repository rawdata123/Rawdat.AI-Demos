package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

var (
	stripeSecretKey = "MY_STRIPE_API_KEY"
)

func main() {
	if stripeSecretKey == "" {
		log.Fatal("STRIPE_SECRET_KEY must be set in the environment")
	}

	stripe.Key = stripeSecretKey

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create-checkout-session", checkoutHandler)
	http.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "✅ Payment successful! Thanks for supporting my services!")
	})
	http.HandleFunc("/cancel", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "❌ Payment canceled. Maybe next time!")
	})

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	const tpl = `
<!DOCTYPE html>
<html>
<head>
  <title>Support My Work</title>
  <style>
    body {
      background: #f7fafc;
      font-family: Arial, sans-serif;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      height: 100vh;
      margin: 0;
    }
    h1 {
      color: #2d3748;
    }
    p {
      font-size: 1.2em;
      color: #4a5568;
    }
    button {
      padding: 12px 24px;
      font-size: 1em;
      background-color: #667eea;
      color: white;
      border: none;
      border-radius: 6px;
      cursor: pointer;
      margin-top: 20px;
    }
    button:hover {
      background-color: #5a67d8;
    }
    form {
      margin-top: 30px;
    }
  </style>
</head>
<body>
  <h1>💸 Pay Me for My Services 😄</h1>
  <p>I appreciate your support. Click below to make a payment of $20.</p>
  <form action="/create-checkout-session" method="POST">
    <button type="submit">Pay $20</button>
  </form>
</body>
</html>
`
	t := template.Must(template.New("page").Parse(tpl))
	t.Execute(w, nil)
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("My Awesome Service"),
					},
					UnitAmount: stripe.Int64(2000), // $20.00
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String("http://localhost:8080/success"),
		CancelURL:  stripe.String("https://third.run.place/cancel"),
	}

	s, err := session.New(params)
	if err != nil {
		http.Error(w, "Stripe error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}

