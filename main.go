package main

import (
	"flag"
	"fmt"
	"os"
	"tesla"
	"time"
)

func teslaFunc() {
	e := os.Getenv("TESLA_SENTRY_EMAIL")
	p := os.Getenv("TESLA_SENTRY_PASSWORD")
	if e == "" || p == "" {
		fmt.Println("E-mailaddress or password not set in environment vars.")

	}
	client, err := tesla.NewClient(
		&tesla.Auth{
			ClientID:     "81527cff06843c8634fdc09e8ac0abefb46ac849f38fe1e431c2ef2106796384",
			ClientSecret: "c7257eb71a564034f9419ee651c7d0e5f7aa6bfbd18bafb5c5c033b093bb2fa3",
			Email:        e,
			Password:     p,
		})
	if err != nil {
		fmt.Println(err)
	} else {
		vehicles, err := client.Vehicles()
		fmt.Println("Car is currently", vehicles[0].State)
		if vehicles[0].State == "sleep" {
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
									fmt.Println("Sentry mode is turned off but car is charing. Enabling!")
									vehicle.EnableSentry()
								} else {
									fmt.Println("Car is charging and Sentrymode is already on!")
								}
							} else {
								fmt.Println(err)
							}
						} else {
							fmt.Println("Car is online but disconnected from charger. Not doing anything!")
							result := vehicle.EnableSentry()
							fmt.Println(result)
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
		tick := time.Tick(2 * time.Minute)
		for range tick {
			teslaFunc()
		}
	} else {
		fmt.Println("Starting single run")
		teslaFunc()
	}
}
