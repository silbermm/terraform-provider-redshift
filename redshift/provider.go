package redshift

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"log"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Description: "Redshift url",
			},
			"user": {
				Type:        schema.TypeString,
				Description: "master user",
			},
			"password": {
				Type:        schema.TypeString,
				Description: "master password",
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "port",
			},
			"database": {
				Type:        schema.TypeString,
				Description: "default database",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"redshift_user": redshiftUser(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	config := Config{
		url: d.Get("url").(string),
		user: d.Get("user").(string),
		password: d.Get("password").(string),
		port: d.Get("port").(int),
		database: d.Get("database").(string),
	}

	log.Println("[INFO] Initializing Redshift client")
	client := config.Client()

	//Test the connection
	err := client.Ping()

	if err != nil {
		log.Println("[ERROR] Could not establist Redshift connection with credentials")
		panic(err.Error()) // proper error handling instead of panic
	}

	return client, nil
}