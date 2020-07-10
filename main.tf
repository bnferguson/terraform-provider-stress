variable "cpu_stress_duration" {
  type = number
  default = "60"
}

variable "memory_stress_duration" {
  type = number
  default = "900"
}

variable "memory_stress_size" {
  type = number
  default = "401"
}

#resource "stress_cpu" "my-server" {
  #duration = var.cpu_stress_duration
#}

resource "stress_memory" "memory" {
  duration = var.memory_stress_duration
  size = var.memory_stress_size
}

resource "stress_statefile" "statefile" {
  count = 500
  size = 1 # size is currently not used
}
