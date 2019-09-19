# Terraform Provider for Pokemon!
This is a **read-only** Terraform provider for interfacing with the public Pokemon API. Docs for the API can be found at [https://pokeapi.co/](https://pokeapi.co/).

## Usage Example
```hcl
# You can get a Pokemon by name, this will output all the other data about the given Pokemon. 
data "pokemon" "ditto" {
    name = "ditto"
}

# You can also get a Pokemon by number from 1 - 807
data "pokemon" "bulbasaur" {
    id = 1
}

# Want to get a random Pokemon? Use the random provider!
resource "random_integer" "pokemon_id" {
    min = 1
    max = 807
}

data "pokemon" "random" {
    id = random_integer.pokemon_id.result
}
```