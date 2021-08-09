package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/bogosj/tesla"
)

func teslaFunc() {
	ctx := context.Background()
	tokenPath := os.Getenv("TESLA_SENTRY_TOKENFILE")
	if tokenPath == "" {
		fmt.Println("TESLA_SENTRY_TOKENFILE environment variable not set or empty.")
	}

	client, err := tesla.NewClient(ctx, tesla.WithTokenFile(tokenPath))
	if err != nil {
		fmt.Println(err)
	} else {
		vehicles, err := client.Vehicles()
		fmt.Println("Car is currently", vehicles[0].State)
		if vehicles[0].State == "asleep" {
			fmt.Println("Not waking up car!")
		}
		if vehicles[0].State == "online" {
			if err != nil {
				panic(err)
			} else {
				vehicle := vehicles[0]

				if err == nil {
					chargestate, err := vehicle.ChargeState()
					if err == nil {
						if chargestate.ChargingState == "Complete" || chargestate.ChargingState == "Charging" {
							vehicleState, err := vehicle.VehicleState()
							if err == nil {
								if vehicleState.SentryMode == false {
									fmt.Println("Sentry mode is turned off but car is charging. Enabling!")
									vehicle.EnableSentry()
								} else {
									fmt.Println("Car is charging and Sentrymode is already on!")
								}
							} else {
								fmt.Println(err)
							}
						} else {
							fmt.Println("Car is online but disconnected from charger. Not doing anything!")
						}
					} else {
						fmt.Println(err)
					}
				} else {
					//Car is probably sleeping
					fmt.Print(err)
				}
			}
		}
	}
}

func main() {
	daemon := flag.Bool("daemon", false, "Enables daemon mode if true")
	flag.Parse()
	if *daemon {
		fmt.Println("Starting in daemon mode")
		teslaFunc()
		tick := time.Tick(16 * time.Minute)
		for range tick {
			teslaFunc()
		}
	} else {
		fmt.Println("Starting single run")
		teslaFunc()
	}
}
