#resource "stress_cpu" "my-server" {
  #duration = 30
#}

resource "stress_memory" "memory" {
  duration = 30
  size = 128
}
