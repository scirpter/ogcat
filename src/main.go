package main

import (
	"ogcat/src/base"
	"ogcat/src/common"
	"ogcat/src/enums"
	"ogcat/src/survey"
	"ogcat/src/utils/cnsl"
	"ogcat/src/utils/gen"
	"os"
	"time"
)

func main() {
	cnsl.ClearConsole()

	app := base.NewApp()

	if app.ProgramMode == enums.Sniping {
		qs := []*survey.Question{
			{
				Name:     "targetUsername",
				Prompt:   &survey.Input{Message: "Target Username:"},
				Validate: survey.Required,
			},
		}
		answers := struct {
			TargetUsername string
		}{}
		err := survey.Ask(qs, &answers)
		if err != nil {
			cnsl.Error("survey interrupted")
			time.Sleep(3 * time.Second)
			os.Exit(1)
		}
		go base.RunSniper(&answers.TargetUsername)
	} else {
		qs := []*survey.Question{
			{
				Name:     "includeABC",
				Prompt:   &survey.Confirm{Message: "Include ABC?"},
				Validate: survey.Required,
			},
			{
				Name:     "includeSpecial",
				Prompt:   &survey.Confirm{Message: "Include Special?"},
				Validate: survey.Required,
			},
			{
				Name:     "includeDigits",
				Prompt:   &survey.Confirm{Message: "Include Digits?"},
				Validate: survey.Required,
			},
			{
				Name:     "maxLen",
				Prompt:   &survey.Input{Message: "Max Length:"},
				Validate: survey.Required,
			},
			{
				Name:     "lengthOffset",
				Prompt:   &survey.Input{Message: "Minimum Length:"},
				Validate: survey.Required,
			},
			{
				Name:     "idleTime",
				Prompt:   &survey.Input{Message: "Idle Time (time in s to wait after check, stable with 9):"},
				Validate: survey.Required,
			},
		}
		answers := struct {
			IncludeABC     bool
			IncludeSpecial bool
			IncludeDigits  bool
			MaxLen         uint8
			LengthOffset   uint8
			IdleTime       uint8
		}{}
		err := survey.Ask(qs, &answers)
		if err != nil {
			cnsl.Error("survey interrupted")
			time.Sleep(3 * time.Second)
			os.Exit(1)
		}

		usernames := gen.GenUsernames(answers.LengthOffset, answers.MaxLen, answers.IncludeABC, answers.IncludeSpecial, answers.IncludeDigits)
		gen.FilterUsernamesFrom(&usernames)
		equalSplit := common.SplitStringListEqual(&usernames, uint8(len(*app.DiscordTokens)))

		for i, token := range *app.DiscordTokens {
			go base.RunChecker(&token, time.Duration(answers.IdleTime), &equalSplit[i])
		}
	}

	ch := make(chan bool)
	<-ch
}
