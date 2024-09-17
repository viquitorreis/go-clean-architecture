provider "docker" {}

resource "docker_image" "nginx" {
  name         = "nginx:latest"
  keep_locally = true
}

resource "docker_container" "nginx" {
  image = docker_image.nginx.image_id
  name  = "nginx2"
  ports {
    internal = 80
    external = 8000
  }
}

resource "docker_image" "postgres" {
  name         = "postgres:latest"
  keep_locally = true
}

resource "docker_container" "postgres" {
  image = docker_image.postgres.image_id
  name  = "psqlterraform"
  ports {
    internal = 5432
    external = 5555
  }

  env = [
    "POSTGRES_USER=admin",
    "POSTGRES_PASSWORD=adminpassword",
    "POSTGRES_DB=mydatabase"
  ]
}