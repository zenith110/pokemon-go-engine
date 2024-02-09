import tomllib
import toml

with open("pokemon.toml", "rb") as f:
    data = tomllib.load(f)


pokemons = sorted(data['pokemon'], key=lambda x: x['id'])
pokemon = {
    "pokemon": pokemons
}

with open('pokemon-copy.toml', 'w', encoding="utf-8") as f:
    toml.dump(pokemon, f)