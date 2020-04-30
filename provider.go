package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"stress_cpu":    resourceCpu(),
			"stress_memory": resourceMemory(),
		},
	}
}
