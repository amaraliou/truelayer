# Truelayer Go client

## Usage Example
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"

	"github.com/amaraliou/truelayer"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

const redirectURI = "http://localhost:3000/callback"

var (
	auth  truelayer.Authenticator
	ch    = make(chan *truelayer.Client)
	state = uuid.NewV4().String()
)

func main() {

	err := godotenv.Load(os.ExpandEnv(".env"))
	if err != nil {
		fmt.Printf("Error getting env %v\n", err)
	}

	auth = truelayer.NewAuthenticator(redirectURI, os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), "accounts", "balance")

    // Callback
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go http.ListenAndServe(":3000", nil)

	authURL, err := url.PathUnescape(auth.AuthURL(state))
	if err != nil {
		log.Fatal(err)
	}

	openbrowser(authURL)

	// wait for auth to complete
	client := <-ch

	accounts, err := client.GetAccounts()
	if err != nil {
		log.Fatal(err)
	}

	for _, account := range accounts {
		fmt.Println(account.ID)

		accountBalance, err := client.GetAccountBalance(account.ID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(accountBalance.Available)
		fmt.Println(accountBalance.Overdraft)

		fmt.Printf("%s\t%0.2f\t%0.2f",
			account.ID,
			accountBalance.Available,
			accountBalance.Overdraft,
		)
	}
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	client := auth.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
```