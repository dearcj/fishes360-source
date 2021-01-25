package main

type StoryActions map[uint32]StoryAction

type StoryAction struct {
	Name     string
	IcoGfx   string
	StringID string
}

type StoryActionTypes struct {
	SA_FIGHT      uint32
	SA_GIFT       uint32
	SA_RUMOR      uint32
	SA_TRADE      uint32
	SA_CURSE      uint32
	SA_BLESS      uint32
	SA_DEMOTIVATE uint32
	SA_GOTO       uint32
	SA_SETFREE    uint32
	SA_GIVEOBJ    uint32
	SA_RETREAT    uint32
	SA_CUSTOM     uint32
}

type StoryConfig struct {
	ActionTypes StoryActionTypes
	Actions     StoryActions
}

var at = StoryActionTypes{
	SA_FIGHT:      1,
	SA_GIFT:       2,
	SA_RUMOR:      3,
	SA_TRADE:      4,
	SA_CURSE:      5,
	SA_BLESS:      6,
	SA_DEMOTIVATE: 8,
	SA_GOTO:       9,
	SA_SETFREE:    10,
	SA_GIVEOBJ:    11,
	SA_RETREAT:    12,
	SA_CUSTOM:     13,
}

var storyConfig = StoryConfig{
	ActionTypes: at,
	Actions: StoryActions{
		at.SA_RETREAT: StoryAction{
			StringID: "retreat",
			Name:     "Retreat",
			IcoGfx:   "retreat.png",
		},

		at.SA_FIGHT: StoryAction{
			StringID: "fight",
			Name:     "Start fight",
			IcoGfx:   "actionfight.png",
		},

		at.SA_TRADE: StoryAction{
			StringID: "trade",
			Name:     "Trade",
			IcoGfx:   "actiontrade.png",
		},

		at.SA_CURSE: StoryAction{
			StringID: "curse",
			Name:     "Curse somebody",
			IcoGfx:   "actionmagic.png",
		},

		at.SA_BLESS: StoryAction{
			StringID: "bless",
			Name:     "Blessing",
			IcoGfx:   "actionmagic.png",
		},

		at.SA_GIFT: StoryAction{
			StringID: "gift",
			Name:     "Take a gift",
			IcoGfx:   "actiongift.png",
		},

		at.SA_DEMOTIVATE: StoryAction{
			StringID: "demotivate",
			Name:     "Make a gift",
			IcoGfx:   "actiondemotivate.png",
		},

		at.SA_GOTO: StoryAction{
			StringID: "goto",
			Name:     "Open door",
			IcoGfx:   "actiondoor.png",
		},

		at.SA_RUMOR: StoryAction{
			StringID: "rumor",
			Name:     "Tell a rumor",
			IcoGfx:   "actionrumor.png",
		},

		at.SA_SETFREE: StoryAction{
			StringID: "setfree",
			Name:     "Set free",
			IcoGfx:   "actionfree.png",
		},

		at.SA_GIVEOBJ: StoryAction{
			StringID: "giveobj",
			Name:     "Give",
			IcoGfx:   "actiongive.png",
		},

		at.SA_CUSTOM: StoryAction{
			StringID: "custom",
			Name:     "Custom",
			IcoGfx:   "",
		},
	},
}
