#resource "stress_cpu" "my-server" {
  #duration = 30
#}

resource "stress_memory" "memory" {
  duration = 300
  size = 256
}
