package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Message string `json:"message"`
}

type RegisterRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Email    string `json:"email"`
}

type ProfileResponse struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Age      int    `json:"age"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    var req LoginRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Validate login credentials (this is just a placeholder, replace with your logic)
    if req.Username == "user@gmail.com" && req.Password == "password" {
        res := LoginResponse{Message: "Login successful"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(res)
    } else {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    }
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Register the user (this is just a placeholder, replace with your logic)
    res := LoginResponse{Message: "Registration successful"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
    // Get user profile (this is just a placeholder, replace with your logic)
    res := ProfileResponse{
        Username: "user",
        Email:    "user@example.com",
        Age:      30,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}

func main() {
    fs := http.FileServer(http.Dir("./web/build"))
    http.Handle("/", fs)

    http.HandleFunc("/api/login", loginHandler)
    http.HandleFunc("/api/register", registerHandler)
    http.HandleFunc("/api/profile", profileHandler)

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}