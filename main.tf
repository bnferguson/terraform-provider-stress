variable "cpu_stress_duration" {
  type = number
  default = "1"
}

variable "memory_stress_duration" {
  type = number
  default = "1"
}

variable "memory_stress_size" {
  type = number
  default = "750"
}

resource "stress_cpu" "my-server" {
  duration = var.cpu_stress_duration
}

resource "stress_memory" "memory" {
  duration = var.memory_stress_duration
  size = var.memory_stress_size
}

resource "stress_statefile" "statefile" {
  count = 64
  // Bigger than 3mb and you'll start seeing the RPC connection between tf and the provider break due to the 4mb limit
  size = 1
}
