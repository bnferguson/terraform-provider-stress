variable "cpu_stress_duration" {
  type = number
  default = "3"
}

variable "memory_stress_duration" {
  type = number
  default = "300"
}

variable "memory_stress_size" {
  type = number
  default = "1000"
}

resource "stress_cpu" "my-server" {
  duration = var.cpu_stress_duration
}

resource "stress_memory" "memory" {
  duration = var.memory_stress_duration
  size = var.memory_stress_size
}
