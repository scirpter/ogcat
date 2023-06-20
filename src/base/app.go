package base

import (
	"fmt"
	"ogcat/src/common"
	"ogcat/src/enums"
	"ogcat/src/survey"
	"ogcat/src/utils/cnsl"
	"ogcat/src/utils/setup"
	"os"
	"time"
)

type App struct {
	DiscordTokens *[]string
	ProgramMode   enums.ProgramMode
}

func NewApp() *App {
	fmt.Print("\n\n")

	didConfigExist, _, err, _ := setup.CreateFilesIfNotExists()
	if err != nil {
		cnsl.Error(fmt.Sprintf("failed to create config file: %s", err.Error()))
		return nil
	}

	var discordTokens *[]string

	if !didConfigExist {
		qs := []*survey.Question{
			{
				Name:     "discordToken",
				Prompt:   &survey.Password{Message: "Discord Token:"},
				Validate: survey.Required,
			},
		}

		answers := struct {
			DiscordToken    string
			DiscordPassword string
		}{}

		err := survey.Ask(qs, &answers)
		if err != nil {
			cnsl.Error("survey interrupted")
			time.Sleep(3 * time.Second)
			os.Exit(1)
		}

		config := &common.Config{
			DiscordTokens:   []string{answers.DiscordToken},
			DiscordPassword: "unnecessary",
		}
		err = common.WriteToJSONFile(config)

		if err != nil {
			cnsl.Error(fmt.Sprintf("failed to write to config file: %s", err.Error()))
			time.Sleep(3 * time.Second)
			os.Exit(1)
		}
	}

	qs := []*survey.Question{
		{
			Name: "programMode",
			Prompt: &survey.Select{
				Message: "Select a program mode:",
				Options: enums.AllProgramModesStr(),
			},
			Validate: survey.Required,
		},
	}

	answers := struct {
		ProgramMode string
	}{}

	err = survey.Ask(qs, &answers)
	if err != nil {
		cnsl.Error("survey interrupted")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}

	discordTokens = &common.ReadJSONFile().DiscordTokens

	return &App{
		DiscordTokens: discordTokens,
		ProgramMode:   enums.ProgramMode(answers.ProgramMode),
	}
}
