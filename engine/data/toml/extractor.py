from pokebase import cache
import pokebase as pb
cache.API_CACHE

pokemon_data = []
bulba = pb.pokemon(1)

types = []
moves = []
abilities = []
for pokemon_types in bulba.types:
    types.append(pokemon_types.type.name)
for pokemon_moves in bulba.moves:
    pokemon_move = {
        "name": pokemon_moves.move.name,
        "level": pokemon_moves.version_group_details[-1].level_learned_at,
        "method": pokemon_moves.version_group_details[-1].move_learn_method.name
    }
    moves.append(pokemon_move)
for pokemon_abilities in bulba.abilities:
    pokemon_ability = {
        "name": pokemon_abilities.ability.name,
        "isHidden": pokemon_abilities.is_hidden
    }
    abilities.append(pokemon_ability)
id = ""
if(bulba.id < 10):
    id = f"00{bulba.id}"

pokemon_evolutions = []

for evolution in pb.pokemon_species(1).evolution_chain.chain.evolves_to:
    evolution_data_list = []
    for evolution_details in evolution.evolution_details:
        match evolution_details.trigger.name:
            case "level-up":
                evolution_data_list.append(evolution_details.trigger.name)
                evolution_data_list.append(evolution_details.min_level)
            case _:
                print("Could not find a trigger!")
    evolution_id = evolution.species.url.split("/")[-1 - 1]
    if(int(evolution_id) < 10):
        evolution_id = f"00{evolution_id}"
    evolution_data = {
        "name": evolution.species.name,
        "methods": evolution_data_list,
        "id": evolution_id
    }
    pokemon_evolutions.append(evolution_data)
pokemon_stats = {}
for stats in bulba.stats:
    pokemon_stats[stats.stat.name] = stats.base_stat
dex_entry = pb.pokemon_species(1).flavor_text_entries[0].flavor_text.replace("\n", " ").replace("\x0c", "")
pokemon = {
    "id": id, 
    "species": bulba.species.name,
    "abilities": abilities,
    "moves": moves,
    "types": types,
    "evolutions": pokemon_evolutions,
    "dexEntry": dex_entry
}



# commands of interest: abilities, held_items, id, moves, name, species, stats, types
# # MAX GEN
# MAXGEN = 3
# # View Pokemon from this generation number.
# GENERATION = 2
# data = []
# for i in range(1, MAXGEN + 1):
#     # Get API data associated with that particular generation.
#     data.append(pb.generation(i))

# pokemondata = []
# # Iterate through the list of Pokemon introduced in that generation.
# for pokemon in data[0].pokemon_species:
#     print(f"Now looking at {pokemon.name.title()} data!")
#     pokemondata.append(pb.pokemon(pokemon.name.title()))
# print(pokemondata[0])