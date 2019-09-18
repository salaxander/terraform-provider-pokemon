data "pokemon" "ditto" {
    name = "ditto"
}

data "pokemon" "random" {

}

output "ditto" {
    value = data.pokemon.ditto
}

output "random" {
    value = data.pokemon.random
}