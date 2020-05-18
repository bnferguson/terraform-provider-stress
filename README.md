# terraform-provider-stress
Terrifying Terraforming Tricks To Test TFE's Tenacity

This is a very simple provider that exists to do synthetic stress testing against Terraform Enterprise installations. It is very WIP and very ill advised to actually use. It does simple things like allocate and hold on to memory and stress out CPUs for durations. This is an unpublished module and at the moment it is only used for [sideloading](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins) into local runs or using the `terraform.d/` dir to load a Linux version of it into TFE (other methods would be to use bundles etc).

At the moment only a couple things are supported, `resource_stress_memory` and `resource_stress_cpu`. In the future things like disk/state file stresses will be added as needed but this is a good start.

## Usage

Don't.

## But if you must

```hcl
resource "stress_cpu" "why-are-you-doing-this" {
  duration = 300 # How long to keep the load up in Seconds
}

resource "stress_memory" "sad-ram-dot-emoji" {
  duration = 300 # How long to hold the memory
  size = 128 # Memory to allocate in MB
}

resource "stress_statefile" "this-can-crash-tf" {
  count = 64 # adjust count to adjust bloat, but careful this is not linear for the RAM consumed
  size = 2 # 2MB, do not go over 3 as there is a size limit for setting attributes over grpc of ~4MB
}
```
