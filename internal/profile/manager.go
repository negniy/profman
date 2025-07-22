package profile

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func CreateProfile(name, user, project string) error {
	profile := Profile{
		User:    user,
		Project: project,
	}
	data, err := yaml.Marshal(&profile)
	if err != nil {
		return err
	}
	err = os.WriteFile(name+".yaml", data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func GetProfile(name string) error {
	data, err := os.ReadFile(name + ".yaml")
	if err != nil {
		return err
	}
	var profile Profile
	err = yaml.Unmarshal(data, &profile)
	if err != nil {
		return err
	}
	if profile.Project == "" || profile.User == "" {
		return fmt.Errorf("Error: cannot unmarshal")
	}
	fmt.Println("Profile name: " + name + "\n" + "\tUser: " + profile.User + "\n" + "\tProject: " + profile.Project)
	return nil
}

func ListProfile() error {
	files, err := os.ReadDir(".")
	if err != nil {
		return err
	}
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			err = GetProfile(strings.TrimSuffix(file.Name(), ".yaml"))
			if err != nil {
				if strings.Contains(err.Error(), "cannot unmarshal") {
					continue
				}
				return err
			}
		}
	}
	return nil
}

func DeleteProfile(name string) error {
	files, err := os.ReadDir(".")
	if err != nil {
		return err
	}
	for _, file := range files {
		if strings.Compare(name+".yaml", file.Name()) == 0 {
			err = os.Remove(file.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
