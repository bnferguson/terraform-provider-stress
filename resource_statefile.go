package main

import (
	"fmt"
	"math/rand"
	"time"

	petname "github.com/dustinkirkland/golang-petname"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceStatefile() *schema.Resource {
	return &schema.Resource{
		Create: resourceStatefileCreate,
		Read:   resourceStatefileRead,
		Update: resourceStatefileUpdate,
		Delete: resourceStatefileDelete,

		Schema: map[string]*schema.Schema{
			"size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceStatefileCreate(d *schema.ResourceData, m interface{}) error {
	//size := d.Get("size").(int)
	//sizeInBytes := size * (1024 * 1024)
	//numberOfStrings := sizeInBytes / 4096
	const fourKString = "All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy All work and no play makes Jack a dull boy "

	// taken from https://golang.org/src/bytes/buffer_test.go
	var stateFileStrings []string //[2097152]string // 8gb / 4k
	//for i := 0; i < numberOfStrings; i++ {
	stateFileStrings = append(stateFileStrings, fourKString)
	//log.Printf("%v", stateFileStrings)
	//}

	if err := d.Set("data", stateFileStrings); err != nil {
		return fmt.Errorf("Error setting data: %s", err)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	id := petname.Generate(3, "-")
	d.SetId(id)

	return resourceStatefileRead(d, m)
}

func resourceStatefileRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceStatefileUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceStatefileDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
