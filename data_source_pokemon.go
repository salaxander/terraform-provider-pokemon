package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func dataSourcePokemon() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePokemonRead,
		Schema: map[string]*schema.Schema{
			"base_experience": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"height": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"is_default": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"order": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func dataSourcePokemonRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	var resp structs.Pokemon
	if name == "" {
		resp = randomPokemon()
	} else {
		resp, _ = pokeapi.Pokemon(name)
	}
	d.SetId(strconv.Itoa(resp.ID))
	d.Set("base_experience", resp.BaseExperience)
	d.Set("height", resp.Height)
	d.Set("is_default", resp.IsDefault)
	d.Set("name", resp.Name)
	d.Set("order", resp.Order)
	d.Set("weight", resp.Weight)

	return nil
}

func randomPokemon() structs.Pokemon {
	id := rand.Intn(807)
	if id == 0 {
		randomPokemon()
	}
	resp, err := pokeapi.Pokemon(strconv.Itoa(id))
	if err != nil {
		log.Fatalf("Error getting random Pokemon: %v", err)
	}
	return resp
}
