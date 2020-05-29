package main

import (
	"fmt"
	"os"

	"tesla"
)

func main() {
	client, err := tesla.NewClient(
		&tesla.Auth{
			ClientID:     "81527cff06843c8634fdc09e8ac0abefb46ac849f38fe1e431c2ef2106796384",
			ClientSecret: "c7257eb71a564034f9419ee651c7d0e5f7aa6bfbd18bafb5c5c033b093bb2fa3",
			Email:        os.Getenv("TESLA_SENTRY_EMAIL"),
			Password:     os.Getenv("TESLA_SENTRY_PASSWORD"),
		})
	if err != nil {
		fmt.Println("HIER")

	} else {
		vehicles, err := client.Vehicles()
		fmt.Println("Car is currently ", vehicles[0].State)
		if vehicles[0].State == "Asleep" {
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
							vehicle.EnableSentry()
						}
					}
				} else {
					//Car is probably sleeping
					fmt.Print(err)
				}
			}
		}
	}

	// vehicles, err := client.Vehicles()
	// if err != nil {

	// 	panic(err)
	// }

	// vehicle := vehicles[0]
	// status, err := vehicle.MobileEnabled()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(status)
	// fmt.Println(vehicle.HonkHorn())

	// Autopark
	// Use with care, as this will move your car
	// vehicle.AutoparkForward()
	// vehicle.AutoparkReverse()
	// Use with care, as this will move your car

	// // Stream vehicle events
	// eventChan, errChan, err := vehicle.Stream()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	for {
	// 		select {
	// 		case event := <-eventChan:
	// 			eventJSON, _ := json.Marshal(event)
	// 			fmt.Println(string(eventJSON))
	// 		case err = <-errChan:
	// 			fmt.Println(err)
	// 			if err.Error() == "HTTP stream closed" {
	// 				fmt.Println("Reconnecting!")
	// 				// eventChan, errChan, err := vehicle.Stream()
	// 				if err != nil {
	// 					fmt.Println(err)
	// 					return
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}
