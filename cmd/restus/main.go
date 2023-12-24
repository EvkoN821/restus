package main

import (
	"context"
	"flag"
	"github.com/IlyaZayats/restus/internal/db"
	"github.com/IlyaZayats/restus/internal/handlers"
	"github.com/IlyaZayats/restus/internal/repository"
	"github.com/IlyaZayats/restus/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var (
		dbUrl    string
		listen   string
		logLevel string
	)
	//postgres://dynus:dynus@localhost:5555/dynus
	//postgres://restus:restus@postgres_restus:5432/restus
	//postgres://restus:restus@postgres:5555/restus
	flag.StringVar(&dbUrl, "db", "postgres://restus:restus@postgres_restus:5432/restus", "database connection url")
	flag.StringVar(&listen, "listen", ":8080", "server listen interface")
	flag.StringVar(&logLevel, "log-level", "debug", "log level: panic, fatal, error, warning, info, debug, trace")

	flag.Parse()

	ctx := context.Background()

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Panicf("unable to get log level: %v", err)
	}
	logrus.SetLevel(level)

	dbc, err := db.NewPostgresPool(dbUrl)
	if err != nil {
		logrus.Panicf("unable get postgres pool: %v", err)
	}

	restaurantRepo, err := repository.NewPostgresRestaurantRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build restaurant repo: %v", err)
	}

	courseRepo, err := repository.NewPostgresCourseRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build course repo: %v", err)
	}

	foodRepo, err := repository.NewPostgresFoodRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build food repo: %v", err)
	}

	userRepo, err := repository.NewPostgresUserRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build user repo: %v", err)
	}

	restaurantService, err := services.NewRestaurantService(restaurantRepo)
	if err != nil {
		logrus.Panicf("unable build restaurant service: %v", err)
	}

	courseService, err := services.NewCourseService(courseRepo)
	if err != nil {
		logrus.Panicf("unable build course service: %v", err)
	}

	foodService, err := services.NewFoodService(foodRepo)
	if err != nil {
		logrus.Panicf("unable build food service: %v", err)
	}

	userService, err := services.NewUserService(userRepo)
	if err != nil {
		logrus.Panicf("unable build user service: %v", err)
	}

	g := gin.New()

	_, err = handlers.NewRestaurantHandlers(g, restaurantService)
	if err != nil {
		logrus.Panicf("unable build slug handlers: %v", err)
	}

	_, err = handlers.NewCourseHandlers(g, courseService)
	if err != nil {
		logrus.Panicf("unable build course handlers: %v", err)
	}

	_, err = handlers.NewFoodHandlers(g, foodService)
	if err != nil {
		logrus.Panicf("unable build food handlers: %v", err)
	}

	_, err = handlers.NewUserHandlers(g, userService)
	if err != nil {
		logrus.Panicf("unable build slug handlers: %v", err)
	}

	doneC := make(chan error)

	go func() { doneC <- g.Run(listen) }()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGABRT, syscall.SIGHUP, syscall.SIGTERM)

	childCtx, cancel := context.WithCancel(ctx)
	go func() {
		sig := <-signalChan
		logrus.Debugf("exiting with signal: %v", sig)
		cancel()
	}()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				doneC <- ctx.Err()
			}
		}
	}(childCtx)

	<-doneC

}
