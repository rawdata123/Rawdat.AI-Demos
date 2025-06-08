package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/coreos/go-oidc/v3/oidc"
    "golang.org/x/oauth2"
)

var (
    clientID     = os.Getenv("OIDC_CLIENT_ID")
    clientSecret = os.Getenv("OIDC_CLIENT_SECRET")
    redirectURL  = "http://localhost:8080/callback"
    providerURL  = "https://accounts.google.com" // Change to your provider
    state        = "random-state-string" // Use a secure random string in production
)

func main() {
    ctx := context.Background()

    provider, err := oidc.NewProvider(ctx, providerURL)
    if err != nil {
        log.Fatalf("Failed to get provider: %v", err)
    }

    oauth2Config := oauth2.Config{
        ClientID:     clientID,
        ClientSecret: clientSecret,
        RedirectURL:  redirectURL,
        Endpoint:     provider.Endpoint(),
        Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
    }

    verifier := provider.Verifier(&oidc.Config{ClientID: clientID})

    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
    })

    http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Query().Get("state") != state {
            http.Error(w, "state did not match", http.StatusBadRequest)
            return
        }
        code := r.URL.Query().Get("code")
        oauth2Token, err := oauth2Config.Exchange(ctx, code)
        if err != nil {
            http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
            return
        }
        rawIDToken, ok := oauth2Token.Extra("id_token").(string)
        if !ok {
            http.Error(w, "No id_token field in oauth2 token", http.StatusInternalServerError)
            return
        }
        idToken, err := verifier.Verify(ctx, rawIDToken)
        if err != nil {
            http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
            return
        }
        var claims struct {
            Email string `json:"email"`
        }
        if err := idToken.Claims(&claims); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Welcome, %s!", claims.Email)
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, `<a href="/login">Log in with OpenID Connect</a>`)
    })
    log.Println("Listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

