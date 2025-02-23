package service

import (
	"fmt"
	"time"
)

type User struct {
	ID   int
	Name string
}

type Rider struct {
	User
}

type DriverStatus int

const (
	AVAILABLE = iota
	BUSY
	OFFLINE
)

type IObserver interface {
	Notify(rideId int, startLocation Location, endLocation Location)
}

type Driver struct {
	User
	Address         string
	Vehicle         Vehicle
	DriverStatus    DriverStatus
	CurrentLocation Location
}

func (d *driverManager) Notify(rideId int, startLocation Location, endLocation Location) {
	drivers := d.GetDrivers(startLocation, endLocation)
	for i := 0; i < len(drivers); i++ {
		fmt.Println(fmt.Sprintf("Driver %d notified for rideId: %d", drivers[i].ID, rideId))
	}
}

type VehicleType int

const (
	BIKE = iota
	AUTO
	ALTO
	SEDAN
	XUV
)

type Vehicle struct {
	Number string
	Type   VehicleType
}

type Location struct {
	Lat  float64
	Long float64
}

type RideStatus int

const (
	REQUESTED = iota
	ACCEPTED
	CANCELLED
	STARTED
	COMPLETED
)

type Ride struct {
	Id         int
	DriverId   *int
	RiderId    int
	Start      Location
	End        Location
	StartTime  time.Time
	RideStatus RideStatus
}

type IObservable interface {
	Subscribe(observer IObserver)
	Unsubscribe(observer IObserver)
	Fire(rideId int, start Location, end Location)
}

type Observable struct {
	observers []IObserver
}

func (o *Observable) Subscribe(observer IObserver) {
	o.observers = append(o.observers, observer)
}

func (o *Observable) Unsubscribe(observer IObserver) {
	for i, obs := range o.observers {
		if obs == observer {
			o.observers = append(o.observers[:i], o.observers[i+1:]...)
			break
		}
	}
}

func (o *Observable) Fire(rideId int, startLocation Location, endLocation Location) {
	for _, observer := range o.observers {
		observer.Notify(rideId, startLocation, endLocation)
	}
}

type IRideManager interface {
	IObservable
	CreateRide(riderId int, start Location, end Location) Ride
	AcceptRide(driverId int, rideId int) bool
	StartRide(rideId int) bool
	CompleteRide(rideId int) bool
	CancelRide(rideId int) bool
	GetRide(rideId int) Ride
	GetRidesForRider(riderId int) []Ride
	GetRidesForDriver(driverId int) []Ride
}

type IDriverManager interface {
	IObserver
	AddDriver(driver Driver) (*Driver, error)
	GetDriverById(id int) *Driver
	GetDrivers(start Location, end Location) []*Driver
	ChangeDriverStatus(id int, status DriverStatus)
}

type driverManager struct {
	nextId  int
	drivers map[int]*Driver
}

func (d *driverManager) AddDriver(driver Driver) (*Driver, error) {
	d.nextId++
	driver.ID = d.nextId
	d.drivers[d.nextId] = &driver

	return &driver, nil
}

func (d *driverManager) GetDrivers(start Location, end Location) []*Driver {
	var availableDrivers []*Driver
	for i, driver := range d.drivers {
		if driver.DriverStatus == AVAILABLE {
			availableDrivers = append(availableDrivers, d.drivers[i])
		}
	}
	return availableDrivers
}

func (d *driverManager) ChangeDriverStatus(id int, status DriverStatus) {
	d.drivers[id].DriverStatus = status
}

func (d driverManager) GetDriverById(id int) *Driver {
	return d.drivers[id]
}

func NewDriverManager() IDriverManager {
	return &driverManager{
		drivers: make(map[int]*Driver),
	}
}

type rideManager struct {
	Observable
	nextId int
	rides  map[int]*Ride
	dm     IDriverManager
}

func (r *rideManager) CreateRide(riderId int, start Location, end Location) Ride {
	r.nextId++
	ride := Ride{
		DriverId: nil,
		RiderId:  riderId,
		Id:       r.nextId,
	}
	r.rides[r.nextId] = &ride

	r.Fire(r.nextId, start, end)
	return ride
}

func (r *rideManager) AcceptRide(driverId int, rideId int) bool {
	ride := r.rides[rideId]
	if ride.RideStatus != REQUESTED {
		return false
	}
	ride.RideStatus = ACCEPTED
	ride.DriverId = &driverId
	r.dm.ChangeDriverStatus(driverId, BUSY)
	return true
}

func (r *rideManager) StartRide(rideId int) bool {
	ride := r.rides[rideId]
	if ride.RideStatus != ACCEPTED {
		return false
	}
	ride.RideStatus = STARTED
	return true
}

func (r *rideManager) CompleteRide(rideId int) bool {
	ride := r.rides[rideId]
	if ride.RideStatus != STARTED {
		return false
	}
	ride.RideStatus = COMPLETED
	return true
}

func (r *rideManager) CancelRide(rideId int) bool {
	ride := r.rides[rideId]
	if ride.RideStatus == COMPLETED {
		return false
	}
	ride.RideStatus = CANCELLED

	return true
}

func (r *rideManager) GetRide(rideId int) Ride {
	ride := r.rides[rideId]
	return *ride
}

func (r *rideManager) GetRidesForRider(riderId int) []Ride {
	var rides []Ride
	for i, ride := range r.rides {
		if ride.RiderId != riderId {
			continue
		}
		rides = append(rides, *r.rides[i])
	}
	return rides
}

func (r *rideManager) GetRidesForDriver(driverId int) []Ride {
	var rides []Ride
	for i, ride := range r.rides {
		if *ride.DriverId != driverId {
			continue
		}
		rides = append(rides, *r.rides[i])
	}
	return rides
}

func NewRideManager(manager IDriverManager) IRideManager {
	return &rideManager{
		rides: make(map[int]*Ride),
		dm:    manager,
	}
}
