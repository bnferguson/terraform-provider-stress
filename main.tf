variable "cpu_stress_duration" {
  type = number
  default = "1"
}

variable "memory_stress_duration" {
  type = number
  default = "900"
}

variable "memory_stress_size" {
  type = number
  default = "300"
}

#resource "stress_cpu" "my-server" {
  #duration = var.cpu_stress_duration
#}

resource "stress_memory" "memory" {
  duration = var.memory_stress_duration
  size = var.memory_stress_size
}

resource "stress_statefile" "statefile" {
  count = 2500
  size = 1 # size is currently not used
}
