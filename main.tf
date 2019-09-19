### DATA ###

data "pokemon" "ditto" {
    name = "ditto"
}

data "pokemon" "bulbasaur" {
    number = 1
}

resource "random_integer" "pokemon_number" {
    min = 1
    max = 807
}

data "pokemon" "random" {
    number = random_integer.pokemon_number.result
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