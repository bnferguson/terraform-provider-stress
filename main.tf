resource "stress_cpu" "my-server" {
  duration = 300
}

resource "stress_memory" "memory" {
  duration = 300
  size = 128
}
