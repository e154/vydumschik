package vydumschik

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"fmt"
)

type male struct {
	Male_middle		string		`yaml:"male_middle"`
	Female_middle	string		`yaml:"female_middle"`
	Name			string		`yaml:"name"`
}

type names struct {
	Male			[]*male		`yaml:"male"`
	Female			[]string	`yaml:"female"`
	Surnames		[]string	`yaml:"surnames"`
}

var (
	Names		*names
)

type Name struct {}

func (n *Name) First_name(gender string) string {

	err := n.getNameFile()
	if err != nil {
		return err.Error()
	}

	if gender == "" {
		gender = randomGender()
	}

	switch gender {
	case "male":
		return Names.Male[random(0, len(Names.Male))].Name
	case "female":
		return Names.Female[random(0, len(Names.Female))]
	default:
	}

	return ""
}

func (n *Name) Middle_name(gender string) string {

	err := n.getNameFile()
	if err != nil {
		return err.Error()
	}

	if gender == "" {
		gender = randomGender()
	}

	switch gender {
	case "male":
		return Names.Male[random(0, len(Names.Male))].Male_middle
	case "female":
		return Names.Male[random(0, len(Names.Male))].Female_middle
	default:
	}

	return ""
}

func (n *Name) Sur_name(gender string) string {

	err := n.getNameFile()
	if err != nil {
		return err.Error()
	}

	if gender == "" {
		gender = randomGender()
	}

	surname := Names.Surnames[random(0, len(Names.Surnames))]

	if gender == "female" {
		surname += "а"
	}

	return surname
}

func (n *Name) Full_name(gender string) string {

	err := n.getNameFile()
	if err != nil {
		return err.Error()
	}

	if gender == "" {
		gender = randomGender()
	}

	return fmt.Sprintf("%s %s %s", n.Sur_name(gender), n.First_name(gender), n.Middle_name(gender))
}

func (n *Name) getNameFile() error {

	if Names != nil {
		return nil
	}

	file, err := ioutil.ReadFile("data/names.yml")
	if err != nil {
		return err
	}

	Names = new(names)
	err = yaml.Unmarshal(file, &Names)
	return err
}