#resource "stress_cpu" "my-server" {
  #duration = 60
#}

resource "stress_memory" "memory" {
  duration = 60
  size = 64
}
