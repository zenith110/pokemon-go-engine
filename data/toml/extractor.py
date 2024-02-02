from pokebase import cache
import pokebase as pb
import toml
import os
cache.API_CACHE


MAXGEN = 8
data = []

# Get API data associated with that particular generation.
data.append(pb.generation(MAXGEN))
# Iterate through the list of Pokemon introduced in that generation.
for pokemon_gen in data:
    try:
        current_pokemon_gen = MAXGEN
        pokemon = []
        print(f"Going through gen {current_pokemon_gen}")
        for pokemon_species in pokemon_gen.pokemon_species:
            print(f"Going through {pokemon_species.name}: {pokemon_species.id}")
            pokemon_lookup = pb.pokemon(pokemon_species.id)
            type = []
            move = []
            ability = []
            print("Going through types")
            for pokemon_types in pokemon_lookup.types:
                type.append(pokemon_types.type.name)
            print("Going through moves")
            for pokemon_moves in pokemon_lookup.moves:
                pokemon_move = {
                    "name": pokemon_moves.move.name,
                    "level": pokemon_moves.version_group_details[-1].level_learned_at,
                    "method": pokemon_moves.version_group_details[-1].move_learn_method.name
                }
                move.append(pokemon_move)
            print("Going through abilities")
            for pokemon_abilities in pokemon_lookup.abilities:
                pokemon_ability = {
                    "name": pokemon_abilities.ability.name,
                    "isHidden": pokemon_abilities.is_hidden
                }
                ability.append(pokemon_ability)
            id = ""
            if(pokemon_lookup.id < 10):
                id = f"00{pokemon_lookup.id}"
            elif(pokemon_lookup.id < 100 and pokemon_lookup.id >= 10):
                id = f"0{pokemon_lookup.id}"
            else:
                id = pokemon_lookup.id

            pokemon_evolutions = []
            print("Going through evolutions")
            for evolution in pb.pokemon_species(pokemon_species.id).evolution_chain.chain.evolves_to:
                evolution_data_list = []
                for evolution_details in evolution.evolution_details:
                    match evolution_details.trigger.name:
                        case "level-up":
                            evolution_data_list.append(evolution_details.trigger.name)
                            evolution_data_list.append(str(evolution_details.min_level))
                        case "trade":
                            evolution_data_list.append(evolution_details.trigger.name)
                            try:
                                if(evolution_details.held_item.name == None):
                                    evolution_data_list.append("No item")
                                evolution_data_list.append(evolution_details.held_item.name)
                            except:
                                print("Skipping, no evos!")
                        case "Use item":
                            evolution_data_list.append(evolution_details.trigger.name)
                            evolution_data_list.append(evolution_details.item.name)
                        case _:
                            print("Could not find a trigger!")
                evolution_id = evolution.species.url.split("/")[-1 - 1]
                if(pokemon_lookup.id < 10):
                    evolution_id = f"00{pokemon_lookup.id}"
                elif(pokemon_lookup.id < 100 and pokemon_lookup.id >= 10):
                    evolution_id = f"0{pokemon_lookup.id}"
                else:
                    evolution_id = pokemon_lookup.id
                if(int(evolution_id) < 10):
                    evolution_id = f"00{evolution_id}"
                evolution_data = {
                    "name": evolution.species.name,
                    "methods": evolution_data_list,
                    "id": str(evolution_id)
                }
                pokemon_evolutions.append(evolution_data)
            pokemon_stats = {}
            print("Going through stats")
            for stats in pokemon_lookup.stats:
                pokemon_stats[stats.stat.name] = stats.base_stat
            dex_entries = []
            for dex_data in pb.pokemon_species(pokemon_species.id).flavor_text_entries:
                if(dex_data.language.name == "en"):
                    dex_entries.append(dex_data.flavor_text)
                continue

            dex_entry = dex_entries[-1].replace("\n", " ").replace("\x0c", " ")
            asset_data = {
                "front": f"assets/pokemon/front/{id}_front.png",
                "back": f"assets/pokemon/back/{id}_back.png",
                "shiny_front": f"assets/pokemon/shinyfront/{id}_front_shiny.png",
                "shiny_back": f"assets/pokemon/shinyback/{id}_shiny_back.png",
                "icon": f"assets/pokemon/icon/{id}_icon.png"
            }
            pokemon_data = {
                "id": str(id),
                "species": pokemon_lookup.species.name,
                "abilities": ability,
                "moves": move,
                "types": type,
                "evolutions": pokemon_evolutions,
                "dexEntry": dex_entry,
                "stats": pokemon_stats,
                "assetData": asset_data

            }
            print("Done, moving on!")
            print(pokemon_data)
            pokemon.append(pokemon_data)
            pokemon_result = {
                "pokemon": pokemon
            }
            
            output_file_name = f"pokemon_gen_{current_pokemon_gen}.toml"
            with open(output_file_name, "w") as toml_file:
                toml.dump(pokemon_result, toml_file)
    except:
        continue    
