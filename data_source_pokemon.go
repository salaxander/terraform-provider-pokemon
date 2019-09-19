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
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
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
	id := strconv.Itoa(d.Get("id").(int))

	var resp structs.Pokemon

	if name != "" {
		resp, _ = pokeapi.Pokemon(name)
	} else if id != "" {
		resp, _ = pokeapi.Pokemon(id)
	} else {
		d.SetId("")
		return nil
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
