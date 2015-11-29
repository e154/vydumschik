package vydumschik

import (
	"github.com/ghodss/yaml"
	"strconv"
	"fmt"
)

type street struct {
	Type			string		`yaml:"type"`
	Count			int			`yaml:"count"`
	Name			string		`yaml:"name"`
}

type addresses struct  {
	Streets 		[]*street	`yaml:"streets"`
}

var (
	Addresses *addresses
	MAX_APT = 100
	BUILDING_LIMITS = map[string]int{
		":street": 50,
		":alley": 10,
		":avenue": 120,
		":square": 10,
		":other": 20,
	}
	HOUSE_PROBABILITY = 20
	MODIFIER_PROBABILITY = 5
	BUILDING_MODIFIERS = []string{
		"а",
		"б",
		"в",
		"г",
		"д",
	}
)

type Address struct {}

func (a *Address) Street() string {
	a.getAddressFile()
	return Addresses.Streets[random(0, len(Addresses.Streets))].Name
}

func (a *Address) Street_address() string {

	a.getAddressFile()

	str := Addresses.Streets[random(0, len(Addresses.Streets))]
	bld := a.building(BUILDING_LIMITS[str.Type])

	apt := ""
	if random(0, 100) > 50 {
		if random(0, 100) < HOUSE_PROBABILITY {
			apt = bld
		} else {
			apt = bld + "/" + a.apartment()
		}
	} else {
		if random(0, 100) < HOUSE_PROBABILITY {
			apt = "д. " + bld
		} else {
			apt = "д. " + bld + ", кв. " + a.apartment()
		}
	}

	return str.Name + ", " + apt
}

func (a *Address) building(limit int) string {

	if limit == 0 {
		limit = BUILDING_LIMITS[":other"]
	}

	bld := (random(0, limit))
	if random(0, 100) < MODIFIER_PROBABILITY {
		return fmt.Sprintf("%d-%s", bld, BUILDING_MODIFIERS[random(0, len(BUILDING_MODIFIERS))])
	}

	return strconv.Itoa(bld)
}

func (a *Address) apartment() string {
	return strconv.Itoa(1+random(0, MAX_APT))
}

func (a *Address) getAddressFile() error {

	if Addresses != nil {
		return nil
	}

	Addresses = new(addresses)
	err := yaml.Unmarshal([]byte(dataAddresses), &Addresses)
	return err
}