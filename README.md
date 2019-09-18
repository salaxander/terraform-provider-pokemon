# Terraform Provider for Pokemon!
This is a **read-only** Terraform provider for interfacing with the public Pokemon API. Docs for the API can be found at [https://pokeapi.co/](https://pokeapi.co/).

## Usage Example
```hcl
# You can get a Pokemon by name, this will output all the other data about the given Pokemon. 
data "pokemon" "ditto" {
    name = "ditto"
}

# You can also get a random Pokemon from number 1 - 807, just leave the fields empty
data "pokemon" "random" {

}
```