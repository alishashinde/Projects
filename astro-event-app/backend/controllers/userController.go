package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "time"
    "os"
    "astro-event-app/backend/config"
    "astro-event-app/backend/models"
    "github.com/gorilla/mux"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
    UserID uint
    Email  string
    jwt.RegisteredClaims
}

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    hashedPassword, err := hashPassword(user.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    user.Password = hashedPassword
    user.DateJoined = time.Now()
    if err := config.DB.Create(&user).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
    var creds struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var user models.User
    if err := config.DB.Where("email = ?", creds.Email).First(&user).Error; err != nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }

    if !checkPasswordHash(creds.Password, user.Password) {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: user.ID,
        Email:  user.Email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        http.Error(w, "Could not generate token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var user models.User
    if err := config.DB.First(&user, id).Error; err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var user models.User
    if err := config.DB.First(&user, id).Error; err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    var updatedData models.User
    if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user.Username = updatedData.Username
    user.Location = updatedData.Location
    config.DB.Save(&user)

    json.NewEncoder(w).Encode(user)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
    var payload struct {
        Email       string `json:"email"`
        NewPassword string `json:"new_password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var user models.User
    if err := config.DB.Where("email = ?", payload.Email).First(&user).Error; err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    hashedPassword, _ := hashPassword(payload.NewPassword)
    user.Password = hashedPassword
    config.DB.Save(&user)

    json.NewEncoder(w).Encode(map[string]string{"message": "Password reset successful"})
}
