package core

import (
	"os"
	"strconv"
	"time"

	"github.com/hugolgst/rich-go/client"
)

func (drp DiscordRichPresence) SetupRichPresence() {
	err := client.Login(os.Getenv("DISCORD_RICH_PRESENCE_APP_ID"))
	if err != nil {
		panic(err)
	}

}

func (drp DiscordRichPresence) UpdateRichPresence(richPresenceData map[string]string) {
	var err error
	now := time.Now()

	switch richPresenceData["inBattle"] {
	case "false":
		err = client.SetActivity(client.Activity{
			State:      richPresenceData["state"],
			Details:    richPresenceData["details"],
			LargeImage: richPresenceData["largeImage"],
			LargeText:  richPresenceData["largeText"],
			SmallText:  "test",
			Timestamps: &client.Timestamps{
				Start: &now,
			},
			Buttons: []*client.Button{
				&client.Button{
					Label: "Source",
					Url:   "https://github.com/zenith110/pokemon-go-engine",
				},
			},
		})
	case "true":
		currentHP, _ := strconv.Atoi(richPresenceData["pokemonCurrentHP"])
		maxHP, _ := strconv.Atoi(richPresenceData["pokemonMaxHP"])
		err = client.SetActivity(client.Activity{
			State:      richPresenceData["state"],
			Details:    richPresenceData["details"],
			LargeImage: richPresenceData["largeImage"],
			LargeText:  richPresenceData["largeText"],
			Timestamps: &client.Timestamps{
				Start: &now,
			},
			Party: &client.Party{
				ID:         "-1",
				Players:    currentHP,
				MaxPlayers: maxHP,
			},
			Buttons: []*client.Button{
				&client.Button{
					Label: "Source",
					Url:   "https://github.com/zenith110/pokemon-go-engine",
				},
			},
		})
	}

	if err != nil {
		panic(err)
	}
}
