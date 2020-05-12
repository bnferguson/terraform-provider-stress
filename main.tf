variable "cpu_stress_duration" {
  type = number
  default = "30"
}

variable "memory_stress_duration" {
  type = number
  default = "30"
}

variable "memory_stress_size" {
  type = number
  default = "128"
}

resource "stress_cpu" "my-server" {
  duration = var.cpu_stress_duration
}

resource "stress_memory" "memory" {
  duration = var.memory_stress_duration
  size = var.memory_stress_size
}
