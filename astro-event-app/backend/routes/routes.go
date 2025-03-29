package routes

import (
    "net/http"
    "astro-event-app/backend/controllers"
    "astro-event-app/backend/middleware"
    "github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Home route
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to the Astronomical Event API"))
    }).Methods("GET")

    // User routes
    router.HandleFunc("/api/users/register", controllers.RegisterUser).Methods("POST")
    router.HandleFunc("/api/users/login", controllers.LoginUser).Methods("POST")
    router.HandleFunc("/api/users/{id}", controllers.GetUser).Methods("GET")
    router.HandleFunc("/api/users/{id}/update", controllers.UpdateUserProfile).Methods("PUT")
    router.HandleFunc("/api/users/reset-password", controllers.ResetPassword).Methods("POST")

    // Protected Event routes
    eventRouter := router.PathPrefix("/api/events").Subrouter()
    eventRouter.HandleFunc("", controllers.CreateEvent).Methods("POST")
    eventRouter.Use(middleware.JWTAuth)

    router.HandleFunc("/api/events", controllers.GetAllEvents).Methods("GET")
    router.HandleFunc("/api/events/{id}", controllers.GetEventByID).Methods("GET")

    // Protected Observation routes
    obsRouter := router.PathPrefix("/api/observations").Subrouter()
    obsRouter.HandleFunc("", controllers.CreateObservation).Methods("POST")
    obsRouter.Use(middleware.JWTAuth)
    router.HandleFunc("/api/observations", controllers.GetAllObservations).Methods("GET")

    // Protected Comment routes
    comRouter := router.PathPrefix("/api/comments").Subrouter()
    comRouter.HandleFunc("", controllers.CreateComment).Methods("POST")
    comRouter.Use(middleware.JWTAuth)
    router.HandleFunc("/api/comments/observation/{id}", controllers.GetCommentsByObservation).Methods("GET")

    // Protected Notification routes
    notRouter := router.PathPrefix("/api/notifications").Subrouter()
    notRouter.HandleFunc("", controllers.CreateNotification).Methods("POST")
    notRouter.Use(middleware.JWTAuth)
    router.HandleFunc("/api/notifications/user/{id}", controllers.GetUserNotifications).Methods("GET")

    return router
}
