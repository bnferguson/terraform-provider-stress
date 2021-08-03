package main

import (
	"math/rand"
	"runtime"
	"time"

	petname "github.com/dustinkirkland/golang-petname"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCpu() *schema.Resource {
	return &schema.Resource{
		Create: resourceCpuCreate,
		Read:   resourceCpuRead,
		Update: resourceCpuUpdate,
		Delete: resourceCpuDelete,

		Schema: map[string]*schema.Schema{
			"duration": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceCpuCreate(d *schema.ResourceData, m interface{}) error {
	duration := d.Get("duration").(int)
	done := make(chan int)
	rand.Seed(time.Now().UTC().UnixNano())
	id := petname.Generate(3, "-")
	d.SetId(id)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * time.Duration(duration))
	close(done)
	return resourceCpuRead(d, m)
}

func resourceCpuRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCpuUpdate(d *schema.ResourceData, m interface{}) error {
	duration := d.Get("duration").(int)
	done := make(chan int)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * time.Duration(duration))
	close(done)

	return resourceCpuRead(d, m)
}

func resourceCpuDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
