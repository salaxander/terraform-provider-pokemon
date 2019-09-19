package main

import (
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
				Computed: true,
			},
			"height": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_default": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"number": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"order": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourcePokemonRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	number := strconv.Itoa(d.Get("number").(int))

	var resp structs.Pokemon

	if name != "" {
		resp, _ = pokeapi.Pokemon(name)
	} else if number != "" {
		resp, _ = pokeapi.Pokemon(number)
	} else {
		d.SetId("")
		return nil
	}

	d.SetId(strconv.Itoa(resp.ID))
	d.Set("base_experience", resp.BaseExperience)
	d.Set("height", resp.Height)
	d.Set("is_default", resp.IsDefault)
	d.Set("name", resp.Name)
	d.Set("number", resp.ID)
	d.Set("order", resp.Order)
	d.Set("weight", resp.Weight)

	return nil
}
