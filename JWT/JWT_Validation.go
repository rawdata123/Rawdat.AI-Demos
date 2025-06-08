package main

import (
    "fmt"
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("fantasticjwt") // Replace with your secure key

func main() {
    http.HandleFunc("/protected", jwtMiddleware(protectedHandler))
    fmt.Println("Server running on :8080")
    http.ListenAndServe(":8080", nil)
}

func jwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
            return
        }

        parts := strings.Fields(authHeader)
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
            return
        }
        tokenString := parts[1]

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Token is valid, proceed
        next(w, r)
    }
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
    // Redirect to a new URL upon successful authentication
    http.Redirect(w, r, "https://third.run.place", http.StatusFound)
}

