package main

import (
	"fmt"
	"lld/rideHailing/service"
)

func main() {

	dm := service.NewDriverManager()

	d1, _ := dm.AddDriver(service.Driver{CurrentLocation: service.Location{Lat: 1, Long: 2}, DriverStatus: service.AVAILABLE})

	d2, _ := dm.AddDriver(service.Driver{CurrentLocation: service.Location{Lat: 1, Long: 2}, DriverStatus: service.OFFLINE})

	d3, _ := dm.AddDriver(service.Driver{CurrentLocation: service.Location{Lat: 1, Long: 2}, DriverStatus: service.OFFLINE})

	rm := service.NewRideManager(dm)
	rm.Subscribe(dm)

	ride := rm.CreateRide(1, service.Location{1, 2}, service.Location{4, 5})

	fmt.Println(d1, d2, d3)

	status := rm.AcceptRide(d1.ID, ride.Id)

	fmt.Println("accept status ", status)

	rm.CompleteRide(ride.Id)

	status = rm.AcceptRide(d2.ID, ride.Id)
	fmt.Println("accept status ", status)

}
