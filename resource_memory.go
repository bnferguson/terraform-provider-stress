package main

import (
	"fmt"
	"github.com/dustinkirkland/golang-petname"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"math/rand"
	"runtime"
	"time"
)

func resourceMemory() *schema.Resource {
	return &schema.Resource{
		Create: resourceMemoryCreate,
		Read:   resourceMemoryRead,
		Update: resourceMemoryUpdate,
		Delete: resourceMemoryDelete,

		Schema: map[string]*schema.Schema{
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceMemoryCreate(d *schema.ResourceData, m interface{}) error {

	// Below is an example of using our PrintMemUsage() function
	// Print our starting memory usage (should be around 0mb)
	PrintMemUsage()

	size := d.Get("size").(int)
	sizeInBytes := size * (1024 * 1024)

	// taken from https://golang.org/src/bytes/buffer_test.go
	var testBytes []byte
	testBytes = make([]byte, sizeInBytes)
	for i := 0; i < sizeInBytes; i++ {
		testBytes[i] = 'a' + byte(i%26)
	}

	PrintMemUsage()

	duration := d.Get("duration").(int)
	time.Sleep(time.Second * time.Duration(duration))

	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()

	rand.Seed(time.Now().UTC().UnixNano())
	id := petname.Generate(3, "-")
	d.SetId(id)

	return resourceMemoryRead(d, m)
}

func resourceMemoryRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMemoryUpdate(d *schema.ResourceData, m interface{}) error {
	// Below is an example of using our PrintMemUsage() function
	// Print our starting memory usage (should be around 0mb)
	PrintMemUsage()

	size := d.Get("size").(int)
	sizeInBytes := size * (1024 * 1024)

	// taken from https://golang.org/src/bytes/buffer_test.go
	var testBytes []byte
	testBytes = make([]byte, sizeInBytes)
	for i := 0; i < sizeInBytes; i++ {
		testBytes[i] = 'a' + byte(i%26)
	}

	PrintMemUsage()

	duration := d.Get("duration").(int)
	time.Sleep(time.Second * time.Duration(duration))

	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()

	rand.Seed(time.Now().UTC().UnixNano())
	id := petname.Generate(3, "-")
	d.SetId(id)

	return resourceMemoryRead(d, m)
}

func resourceMemoryDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
