data "pokemon" "ditto" {
    name = "ditto"
}

output "ditto" {
    value = data.pokemon.ditto
}