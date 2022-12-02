variable "TAG" {
  default = "latest"
}

group "release" {
    targets = ["amd64", "arm64", "386", "armv6", "armv7"]
}

target "amd64" {
    dockerfile = "Dockerfile"
    platforms = ["linux/amd64"]
    args = {
        LOCAL_FOLDER = "linux-amd64"
    }
    target = "ghactions"
    tags = [
        "f100024/syncthing_exporter:latest",
        "f100024/syncthing_exporter:${TAG}"
    ]
}

target "arm64" {
    dockerfile = "Dockerfile"
    platforms = ["linux/arm64"]
    args = {
        LOCAL_FOLDER = "linux-arm64"
    }
    target = "ghactions"
    tags = [
        "f100024/syncthing_exporter:latest",
        "f100024/syncthing_exporter:${TAG}"
    ]
}

target "386" {
    dockerfile = "Dockerfile"
    platforms = ["linux/386"]
    args = {
        LOCAL_FOLDER = "linux-386"
    }
    target = "ghactions"
    tags = [
        "f100024/syncthing_exporter:latest",
        "f100024/syncthing_exporter:${TAG}"
    ]
}

target "armv6" {
    dockerfile = "Dockerfile"
    platforms = ["linux/arm/v6"]
    args = {
        LOCAL_FOLDER = "linux-armv6"
    }
    target = "ghactions"
    tags = [
        "f100024/syncthing_exporter:latest",
        "f100024/syncthing_exporter:${TAG}"
    ]
}

target "armv7" {
    dockerfile = "Dockerfile"
    platforms = ["linux/arm/v7"]
    args = {
        LOCAL_FOLDER = "linux-armv7"
    }
    target = "ghactions"
    tags = [
        "f100024/syncthing_exporter:latest",
        "f100024/syncthing_exporter:${TAG}"
    ]
}
