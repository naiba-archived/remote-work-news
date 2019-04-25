workflow "Build master and deploy on push" {
  on = "push"
  resolves = ["deploy"]
}

action "filter-master-branch" {
  uses = "actions/bin/filter@4227a6636cb419f91a0d1afb1216ecfab99e433a"
  args = "branch master"
}

action "docker-build" {
  uses = "actions/docker/cli@8cdf801b322af5f369e00d85e9cf3a7122f49108"
  needs = [
    "filter-master-branch",
  ]
  args = "build -t naiba/rwn ."
}

action "docker-login" {
  uses = "actions/docker/login@8cdf801b322af5f369e00d85e9cf3a7122f49108"
  needs = ["docker-build"]
  secrets = ["DOCKER_PASSWORD", "DOCKER_USERNAME"]
}

action "docker-push" {
  uses = "actions/docker/cli@8cdf801b322af5f369e00d85e9cf3a7122f49108"
  args = "push naiba/rwn"
  needs = ["docker-login"]
}

action "deploy" {
  uses = "maddox/actions/ssh@master"
  needs = ["docker-push"]
  secrets = ["PRIVATE_KEY", "PUBLIC_KEY", "HOST", "USER", "PORT"]
  args = "/NAIBA/script/rwn.sh"
}
