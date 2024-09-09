package router

import (
	"log/slog"
	"net/http"

	"FITNESS-TRACKING-APP/internal/http/handlers"
	"FITNESS-TRACKING-APP/internal/http/middleware"
	"FITNESS-TRACKING-APP/storage"
)

func NewMux(logger *slog.Logger, storage storage.Queries) http.Handler {
	mux := http.NewServeMux()
	
	handler :=handlers.NewHandler(logger, storage)

	// USERS
	mux.HandleFunc("POST /api/users/register", handler.UserRegister)
	mux.HandleFunc("POST /api/users/login", handler.UserLogin)

	//WORKOUTS
	mux.Handle("POST /api/workouts",  middleware.ConfirmTokenMiddleware(http.HandlerFunc(handler.CreateWorkouts)))
	mux.Handle("GET /api/workouts/{id}", middleware.ConfirmTokenMiddleware(http.HandlerFunc(handler.GetWorkoutsByUserID)))
	mux.Handle("GET /api/workouts", middleware.ConfirmTokenMiddleware(http.HandlerFunc(handler.GetWorkoutsByID)))
	mux.Handle("PUT /api/workouts/{id}", middleware.ConfirmTokenMiddleware(http.HandlerFunc(handler.UpdateWorkoutsByUserID)))
	mux.Handle("DELETE /api/workouts", middleware.ConfirmTokenMiddleware(http.HandlerFunc(handler.DeleteWorkoutsByID)))

	return mux
}



