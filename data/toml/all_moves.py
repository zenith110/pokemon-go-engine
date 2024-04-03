import pokebase as pb
from pokebase import cache
import toml
import re
cache.API_CACHE

data = []
all_moves = pb.APIResourceList("move")
id = 0
for move_data in all_moves:
    # Get API data for this move.
    move = pb.move(move_data["name"])
    effect_entry = []
    descriptions = []
    try:
        for description in move.flavor_text_entries:
            if(description.language.name == "en"):
                descriptions.append({
                    "description": description.flavor_text}
                    )
        for effect_data in move.effect_entries:
            if(effect_data.language.name == "en"):
                effect_description = effect_data.effect
                effect_entry.append({
                    "effect_text": effect_description
                })
    except Exception as e:
        print(f"Skipping, error is {e}")
    move_data = {
        "name": move.name,
        "Power": move.power,
        "pp": move.pp,
        "accuracy": move.accuracy,
        "descriptions": descriptions,
        "type": move.type,
        "kind_of_move": move.damage_class.name,
        "effects": effect_entry, 
        "id": id
    }
    data.append(move_data)
    id += 1
    pokemon_result = {
        "move": data
    }
    print(pokemon_result)
    output_file_name = f"moves.toml"
    with open(output_file_name, "w", encoding="utf-8") as toml_file:
        toml.dump(pokemon_result, toml_file)