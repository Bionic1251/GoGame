package init

import ()

/**
* Global settings
	Time tics  		<const maxAge = 100>
	Agent count
		Id
		Location?

* Instanse settings

	Coordinate
		X float64
		Y float64

	Location struct{
		Coordinate
		Radius float64

	Person
		Coordinate
		FirstName string
		LastName string
		Age int <Time tics - spent time>

* Containers
	Area
		Locations

	Agent
		Id int64
		Person
		messaging *MessagingService
		drinkingCoffee bool

	MessagingService
		InitService *InitService

	InitService
		Agents []Agent
		Containers []Container
		AgentCount int
}
*/
