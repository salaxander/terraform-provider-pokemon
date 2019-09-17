package main

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mtslzr/pokeapi-go"
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
				Required: true,
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
	resp, err := pokeapi.Pokemon(name)
	log.Printf("Response: %v", resp.ID)
	if err != nil {
		d.SetId("")
		return err
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
