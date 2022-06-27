package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nestrip/request-data-service/db"
	"github.com/nestrip/request-data-service/pkg"
	"github.com/nestrip/request-data-service/services"
	"github.com/procyon-projects/chrono"
	"runtime"
	"time"
)

func main() {
	_ = godotenv.Load()
	fmt.Println("Starting request-data-service, to manage data requests!")

	db.Connect()
	pkg.ConnectToMinio()

	taskScheduler := chrono.NewDefaultTaskScheduler()
	// ignoring the error, since there are no real cases that an error happens
	_, _ = taskScheduler.ScheduleAtFixedRate(services.DataRequestService, 1*time.Minute) //TODO: revert this to prod

	runtime.Goexit()

}
