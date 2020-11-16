package seeder

import (
	"fmt"
	"math/rand"

	"github.com/Kamva/mgm/v2"
	"github.com/jaswdr/faker"
	"github.com/oleg-koval/go-service-jogadores/models"
)

type FakeGeoPoint struct {
	lat float64
	lon float64
}

// GoDrop this
func GoDrop() {
	mgm.Coll(&models.Player{}).Drop(mgm.Ctx())
	mgm.Coll(&models.Request{}).Drop(mgm.Ctx())
}

// GoFake this
func GoFake() {
	faker := faker.New()
	faker.Person().Name()
	faker.Internet().Email()

	person := faker.Person()
	internet := faker.Internet()

	for i := 0; i < 10; i++ {
		fmt.Println("<============")
		fmt.Println(person.Name())
		fmt.Println(internet.Email())

		interests := [6]string{"football", "basketball", "curling", "cs:go", "life", "quidditch"}
		randomIndex := rand.Intn(len(interests))
		pickInterest := interests[randomIndex]

		geoPoints := [4]FakeGeoPoint{
			{lat: 52.0841868, lon: 5.0824915},
			{lat: 52.3420674, lon: 4.8672316},
			{lat: 52.3643823, lon: 4.9022256},
			{lat: 52.0985526, lon: 5.1041158},
		}

		randomIndexGeoPoint := rand.Intn(len(geoPoints))
		pickGeoPoint := geoPoints[randomIndexGeoPoint]

		fmt.Println(pickInterest)
		fmt.Println(pickGeoPoint)

		player := models.CreatePlayer(person.Name(), internet.Email())

		mgm.Coll(player).Create(player)

		request := models.CreateRequest(player.ID.Hex(), pickInterest, pickGeoPoint.lat, pickGeoPoint.lon)
		mgm.Coll(request).Create(request)
		fmt.Println("============>")
	}
}
