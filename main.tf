### DATA ###

data "pokemon" "ditto" {
    name = "ditto"
}

data "pokemon" "bulbasaur" {
    id = 1
}

resource "random_integer" "pokemon_id" {
    min = 1
    max = 807
}

data "pokemon" "random" {
    id = random_integer.pokemon_id.result
}

### OUTPUT ###

output "ditto" {
    value = data.pokemon.ditto
}

output "bulbasaur" {
    value = data.pokemon.bulbasaur
}

output "random" {
    value = data.pokemon.random
}