package vkBot

import "vkBot/events"

var FirstKeyboard = events.ViewKeyBoard{
	OneTime: true,
	Buttons: [][]events.Button{
		{
			{
				Action: events.Action{
					TypeAction: "text",
					Label:      filmsCmd,
				},
			},
		},
		{
			{
				Action: events.Action{
					TypeAction: "text",
					Label:      serialCmd,
				},
			},
		},
		{
			{
				Action: events.Action{
					TypeAction: "text",
					Label:      animeCmd,
				},
			},
		},
		{
			{
				Action: events.Action{
					TypeAction: "text",
					Label:      descriptionCmd,
				},
			},
		},
	},
}

var SecondKeyBoard = events.ViewKeyBoard{
	OneTime: true,
	Buttons: [][]events.Button{
		{
			{
				Action: events.Action{
					TypeAction: "text",
					Label:      comedyCmd,
				},
			},
		},
		{
			{
				Action: events.Action{
					TypeAction: "text",
					Label:      dramaCmd,
				},
			},
		},
		{
			{
				Action: events.Action{
					TypeAction: "text",
					Label:      actionCmd,
				},
			},
		},
		{
			{
				Action: events.Action{
					TypeAction: "text",
					Label:      horrorCmd,
				},
			},
		},
	},
}
