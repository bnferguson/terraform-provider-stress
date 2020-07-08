variable "cpu_stress_duration" {
  type = number
  default = "1"
}

variable "memory_stress_duration" {
  type = number
  default = "301"
}

variable "memory_stress_size" {
  type = number
  default = "400"
}

#resource "stress_cpu" "my-server" {
  #duration = var.cpu_stress_duration
#}

resource "stress_memory" "memory" {
  duration = var.memory_stress_duration
  size = var.memory_stress_size
}

#resource "stress_statefile" "statefile" {
  #count = 5000
  #size = 1 # size is currently not used
#}
