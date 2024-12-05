package types


type PokemonAPIResponse struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Height int    `json:"height"`
}

type Pokemon struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	Cries struct{
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	Height                 int           `json:"height"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string      `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        string      `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"sprites"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}


type Sprites struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type TypeSlot struct {
	Slot int         `json:"slot"`
	Type TypeDetails `json:"type"`
}
type TypeDetails struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type GetAllPokemonsResponse struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Sprites Sprites    `json:"sprites"`
	Types   []TypeSlot `json:"types"`
}


type GenID string




type EvolutionChain struct {
    BabyTriggerItem interface{} `json:"baby_trigger_item"` // If this is always null, can keep as interface{} or change it to a specific type.
    Chain           Chain       `json:"chain"`
    ID              int         `json:"id"`
}

type Chain struct {
    EvolutionDetails []EvolutionDetails `json:"evolution_details"`
    EvolvesTo        []Chain            `json:"evolves_to"`
    IsBaby           bool               `json:"is_baby"`
    Species          Species            `json:"species"`
}

type EvolutionDetails struct {
    Gender                *int     `json:"gender"` // Nullable fields can use pointers.
    HeldItem              *string  `json:"held_item"`
    Item                  *string  `json:"item"`
    KnownMove             *string  `json:"known_move"`
    KnownMoveType         *string  `json:"known_move_type"`
    Location              *string  `json:"location"`
    MinAffection          *int     `json:"min_affection"`
    MinBeauty             *int     `json:"min_beauty"`
    MinHappiness          *int     `json:"min_happiness"`
    MinLevel              *int     `json:"min_level"`
    NeedsOverworldRain    bool     `json:"needs_overworld_rain"`
    PartySpecies          *string  `json:"party_species"`
    PartyType             *string  `json:"party_type"`
    RelativePhysicalStats *int     `json:"relative_physical_stats"`
    TimeOfDay             string   `json:"time_of_day"`
    TradeSpecies          *string  `json:"trade_species"`
    Trigger               Trigger  `json:"trigger"`
    TurnUpsideDown        bool     `json:"turn_upside_down"`
}

type Trigger struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type Species struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}




type SpeciesInfo struct {
	// BaseHappiness     int          `json:"base_happiness"`
	// CaptureRate       int          `json:"capture_rate"`
	// Color             NameAndURL   `json:"color"`
	// EggGroups         []NameAndURL `json:"egg_groups"`
	EvolutionChain    URL          `json:"evolution_chain"`
	// EvolvesFromSpecies *URL         `json:"evolves_from_species"`
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
	// FormDescriptions  []interface{} `json:"form_descriptions"` // If empty, you can change this to `[]string`.
	// FormsSwitchable   bool         `json:"forms_switchable"`
	// GenderRate        int          `json:"gender_rate"`
	// Genera            []Genus      `json:"genera"`
	// Generation        NameAndURL   `json:"generation"`
	// GrowthRate        NameAndURL   `json:"growth_rate"`
	// Habitat           NameAndURL   `json:"habitat"`
	// HasGenderDifferences bool      `json:"has_gender_differences"`
	// HatchCounter      int          `json:"hatch_counter"`
	// ID                int          `json:"id"`
	// IsBaby            bool         `json:"is_baby"`
	// IsLegendary       bool         `json:"is_legendary"`
	// IsMythical        bool         `json:"is_mythical"`
	// Name              string       `json:"name"`
	// Names             []LanguageName `json:"names"`
	// Order             int          `json:"order"`
	// PalParkEncounters []PalParkEncounter `json:"pal_park_encounters"`
	// PokedexNumbers    []PokedexNumber   `json:"pokedex_numbers"`
	// Shape             NameAndURL   `json:"shape"`
	// Varieties         []Variety    `json:"varieties"`
}

type NameAndURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type URL struct {
	URL string `json:"url"`
}

type FlavorText struct {
	FlavorText string     `json:"flavor_text"`
	Language   NameAndURL `json:"language"`
	Version    NameAndURL `json:"version"`
}

type Genus struct {
	Genus    string     `json:"genus"`
	Language NameAndURL `json:"language"`
}

type LanguageName struct {
	Language NameAndURL `json:"language"`
	Name     string     `json:"name"`
}

type PalParkEncounter struct {
	Area       NameAndURL `json:"area"`
	BaseScore  int        `json:"base_score"`
	Rate       int        `json:"rate"`
}

type PokedexNumber struct {
	EntryNumber int        `json:"entry_number"`
	Pokedex     NameAndURL `json:"pokedex"`
}

type Variety struct {
	IsDefault bool       `json:"is_default"`
	Pokemon   NameAndURL `json:"pokemon"`
}