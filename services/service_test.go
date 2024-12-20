package pokemonService_test

import (
	"fmt"
	"sort"
	"testing"

	pokemonServiceHelper "github.com/Tonipenyallop/pokedex-api/helpers"
	"github.com/Tonipenyallop/pokedex-api/types"
)

func TestGetEvolutionChainPokemonNames(t *testing.T) {
	fmt.Println("TestGetEvolutionChainPokemonNames was called")

	// case eeveee
	eeveee := types.EvolutionChain{
		BabyTriggerItem: nil,
		Chain: types.Chain{
			EvolutionDetails: []types.EvolutionDetail{},
			EvolvesTo: []types.Chain{
				// Vaporeon Evolution
				{
					EvolutionDetails: []types.EvolutionDetail{
						{
							Gender:   nil,
							HeldItem: nil,
							Item: &types.Item{
								Name: "water-stone",
								URL:  "https://pokeapi.co/api/v2/item/84/",
							},
							KnownMove:             nil,
							KnownMoveType:         nil,
							Location:              nil,
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "use-item",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/3/",
							},
							TurnUpsideDown: false,
						},
					},
					EvolvesTo: []types.Chain{},
					IsBaby:    false,
					Species: types.Species{
						Name: "vaporeon",
						URL:  "https://pokeapi.co/api/v2/pokemon-species/134/",
					},
				},
				// Jolteon Evolution
				{
					EvolutionDetails: []types.EvolutionDetail{
						{
							Gender:   nil,
							HeldItem: nil,
							Item: &types.Item{
								Name: "thunder-stone",
								URL:  "https://pokeapi.co/api/v2/item/83/",
							},
							KnownMove:             nil,
							KnownMoveType:         nil,
							Location:              nil,
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "use-item",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/3/",
							},
							TurnUpsideDown: false,
						},
					},
					EvolvesTo: []types.Chain{},
					IsBaby:    false,
					Species: types.Species{
						Name: "jolteon",
						URL:  "https://pokeapi.co/api/v2/pokemon-species/135/",
					},
				},
				// Flareon Evolution
				{
					EvolutionDetails: []types.EvolutionDetail{
						{
							Gender:   nil,
							HeldItem: nil,
							Item: &types.Item{
								Name: "fire-stone",
								URL:  "https://pokeapi.co/api/v2/item/82/",
							},
							KnownMove:             nil,
							KnownMoveType:         nil,
							Location:              nil,
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "use-item",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/3/",
							},
							TurnUpsideDown: false,
						},
					},
					EvolvesTo: []types.Chain{},
					IsBaby:    false,
					Species: types.Species{
						Name: "flareon",
						URL:  "https://pokeapi.co/api/v2/pokemon-species/136/",
					},
				},
				// Espeon Evolution
				{
					EvolutionDetails: []types.EvolutionDetail{
						{
							Gender:                nil,
							HeldItem:              nil,
							Item:                  nil,
							KnownMove:             nil,
							KnownMoveType:         nil,
							Location:              nil,
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "day",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "level-up",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/1/",
							},
							TurnUpsideDown: false,
						},
					},
					EvolvesTo: []types.Chain{},
					IsBaby:    false,
					Species: types.Species{
						Name: "espeon",
						URL:  "https://pokeapi.co/api/v2/pokemon-species/196/",
					},
				},
				// Umbreon Evolution
				{
					EvolutionDetails: []types.EvolutionDetail{
						{
							Gender:                nil,
							HeldItem:              nil,
							Item:                  nil,
							KnownMove:             nil,
							KnownMoveType:         nil,
							Location:              nil,
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "night",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "level-up",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/1/",
							},
							TurnUpsideDown: false,
						},
					},
					EvolvesTo: []types.Chain{},
					IsBaby:    false,
					Species: types.Species{
						Name: "umbreon",
						URL:  "https://pokeapi.co/api/v2/pokemon-species/197/",
					},
				},
				// Leafeon Evolution
				{
					EvolutionDetails: []types.EvolutionDetail{
						// First Evolution Detail
						{
							Gender:        nil,
							HeldItem:      nil,
							Item:          nil,
							KnownMove:     nil,
							KnownMoveType: nil,
							Location: &types.Location{
								Name: "eterna-forest",
								URL:  "https://pokeapi.co/api/v2/location/8/",
							},
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "level-up",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/1/",
							},
							TurnUpsideDown: false,
						},
						// Second Evolution Detail
						{
							Gender:        nil,
							HeldItem:      nil,
							Item:          nil,
							KnownMove:     nil,
							KnownMoveType: nil,
							Location: &types.Location{
								Name: "pinwheel-forest",
								URL:  "https://pokeapi.co/api/v2/location/375/",
							},
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "level-up",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/1/",
							},
							TurnUpsideDown: false,
						},
						// Third Evolution Detail
						{
							Gender:        nil,
							HeldItem:      nil,
							Item:          nil,
							KnownMove:     nil,
							KnownMoveType: nil,
							Location: &types.Location{
								Name: "kalos-route-20",
								URL:  "https://pokeapi.co/api/v2/location/650/",
							},
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "level-up",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/1/",
							},
							TurnUpsideDown: false,
						},
						// Fourth Evolution Detail
						{
							Gender:   nil,
							HeldItem: nil,
							Item: &types.Item{
								Name: "leaf-stone",
								URL:  "https://pokeapi.co/api/v2/item/85/",
							},
							KnownMove:             nil,
							KnownMoveType:         nil,
							Location:              nil,
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "use-item",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/3/",
							},
							TurnUpsideDown: false,
						},
					},
					EvolvesTo: []types.Chain{},
					IsBaby:    false,
					Species: types.Species{
						Name: "leafeon",
						URL:  "https://pokeapi.co/api/v2/pokemon-species/470/",
					},
				},
				// Glaceon Evolution
				{
					EvolutionDetails: []types.EvolutionDetail{
						// First Evolution Detail
						{
							Gender:        nil,
							HeldItem:      nil,
							Item:          nil,
							KnownMove:     nil,
							KnownMoveType: nil,
							Location: &types.Location{
								Name: "sinnoh-route-217",
								URL:  "https://pokeapi.co/api/v2/location/48/",
							},
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "level-up",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/1/",
							},
							TurnUpsideDown: false,
						},
						// Second Evolution Detail
						{
							Gender:        nil,
							HeldItem:      nil,
							Item:          nil,
							KnownMove:     nil,
							KnownMoveType: nil,
							Location: &types.Location{
								Name: "twist-mountain",
								URL:  "https://pokeapi.co/api/v2/location/380/",
							},
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "level-up",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/1/",
							},
							TurnUpsideDown: false,
						},
						// Third Evolution Detail
						{
							Gender:        nil,
							HeldItem:      nil,
							Item:          nil,
							KnownMove:     nil,
							KnownMoveType: nil,
							Location: &types.Location{
								Name: "frost-cavern",
								URL:  "https://pokeapi.co/api/v2/location/640/",
							},
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "level-up",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/1/",
							},
							TurnUpsideDown: false,
						},
						// Fourth Evolution Detail
						{
							Gender:   nil,
							HeldItem: nil,
							Item: &types.Item{
								Name: "ice-stone",
								URL:  "https://pokeapi.co/api/v2/item/885/",
							},
							KnownMove:             nil,
							KnownMoveType:         nil,
							Location:              nil,
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "use-item",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/3/",
							},
							TurnUpsideDown: false,
						},
					},
					EvolvesTo: []types.Chain{},
					IsBaby:    false,
					Species: types.Species{
						Name: "glaceon",
						URL:  "https://pokeapi.co/api/v2/pokemon-species/471/",
					},
				},
				// Sylveon Evolution
				{
					EvolutionDetails: []types.EvolutionDetail{
						// First Evolution Detail
						{
							Gender:    nil,
							HeldItem:  nil,
							Item:      nil,
							KnownMove: nil,
							KnownMoveType: &types.Type{
								Name: "fairy",
								URL:  "https://pokeapi.co/api/v2/type/18/",
							},
							Location:              nil,
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "level-up",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/1/",
							},
							TurnUpsideDown: false,
						},
						// Second Evolution Detail
						{
							Gender:    nil,
							HeldItem:  nil,
							Item:      nil,
							KnownMove: nil,
							KnownMoveType: &types.Type{
								Name: "fairy",
								URL:  "https://pokeapi.co/api/v2/type/18/",
							},
							Location:              nil,
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "level-up",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/1/",
							},
							TurnUpsideDown: false,
						},
					},
					EvolvesTo: []types.Chain{},
					IsBaby:    false,
					Species: types.Species{
						Name: "sylveon",
						URL:  "https://pokeapi.co/api/v2/pokemon-species/700/",
					},
				},
			},
			IsBaby: false,
			Species: types.Species{
				Name: "eevee",
				URL:  "https://pokeapi.co/api/v2/pokemon-species/133/",
			},
		},
		ID: 67,
	}

	tmpExpected := []struct{
		caseExpected []string
	}{

		{caseExpected: []string{"vaporeon", "jolteon", "flareon", "espeon", "umbreon", "leafeon", "glaceon", "sylveon", "eevee",},}
		{caseExpected: []string{"venusaur","ivysaur","bulbasaur"},}
	}
	

	expected := []string{
		"vaporeon", "jolteon", "flareon", "espeon", "umbreon", "leafeon", "glaceon", "sylveon", "eevee",
	}

	sort.Slice(expected, func(i, j int) bool {
		return expected[i] < expected[j]
	})

	res := pokemonServiceHelper.GetEvolutionChainPokemonNames(&eeveee)
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	if len(expected) != len(res) {
		t.Fatalf("expect %d is equal to %d", len(expected), len(res))
	}

	for index := 0; index < len(expected); index++ {
		fmt.Printf("expected[index]:%s, res[index]:%s", expected[index], res[index])
		if expected[index] != res[index] {
			t.Fatalf("expect %s is equal to %s", expected[index], res[index])
		}
	}

	// case bulbasaur
	bulbasaur := types.EvolutionChain{
		BabyTriggerItem: nil,
		Chain: types.Chain{
			EvolutionDetails: []types.EvolutionDetail{},
			EvolvesTo: []types.Chain{
				{
					EvolutionDetails: []types.EvolutionDetail{
						{
							Gender:                nil,
							HeldItem:              nil,
							Item:                  nil,
							KnownMove:             nil,
							KnownMoveType:         nil,
							Location:              nil,
							MinAffection:          nil,
							MinBeauty:             nil,
							MinHappiness:          nil,
							MinLevel:              nil,
							NeedsOverworldRain:    false,
							PartySpecies:          nil,
							PartyType:             nil,
							RelativePhysicalStats: nil,
							TimeOfDay:             "",
							TradeSpecies:          nil,
							Trigger: types.Trigger{
								Name: "level-up",
								URL:  "https://pokeapi.co/api/v2/evolution-trigger/1/",
							},
							TurnUpsideDown: false,
						},
					},
					EvolvesTo: []types.Chain{
						{
							EvolutionDetails: []types.EvolutionDetail{
								{
									Gender:                nil,
									HeldItem:              nil,
									Item:                  nil,
									KnownMove:             nil,
									KnownMoveType:         nil,
									Location:              nil,
									MinAffection:          nil,
									MinBeauty:             nil,
									MinHappiness:          nil,
									MinLevel:              nil,
									NeedsOverworldRain:    false,
									PartySpecies:          nil,
									PartyType:             nil,
									RelativePhysicalStats: nil,
									TimeOfDay:             "",
									TradeSpecies:          nil,
									Trigger: types.Trigger{
										Name: "level-up",
										URL:  "https://pokeapi.co/api/v2/evolution-trigger/1/",
									},
									TurnUpsideDown: false,
								},
							},
							EvolvesTo: []types.Chain{},
							IsBaby:    false,
							Species: types.Species{
								Name: "venusaur",
								URL:  "https://pokeapi.co/api/v2/pokemon-species/3/",
							},
						},
					},
					IsBaby: false,
					Species: types.Species{
						Name: "ivysaur",
						URL:  "https://pokeapi.co/api/v2/pokemon-species/2/",
					},
				},
			},
			IsBaby: false,
			Species: types.Species{
				Name: "bulbasaur",
				URL:  "https://pokeapi.co/api/v2/pokemon-species/1/",
			},
		},
		ID: 1,
	}
	expected2 := []string{"venusaur","ivysaur","bulbasaur"}

	sort.Slice(expected2,func(i, j int) bool {
		return expected2[i] < expected2[j]
	})
	
	res2 := pokemonServiceHelper.GetEvolutionChainPokemonNames(&bulbasaur)
	sort.Slice(res2,func(i, j int) bool {
		return res2[i] < res2[j]
	})

	if len(expected2) != len(res2) {
		t.Fatalf("expect %d is equal to %d", len(expected2), len(res2))
	}

	for index := 0; index < len(expected2); index++ {
		fmt.Printf("expected2[index]:%s, res2[index]:%s", expected2[index], res2[index])
		if expected2[index] != res2[index] {
			t.Fatalf("expect %s is equal to %s", expected2[index], res2[index])
		}
	}


	// case mewtwo

}
