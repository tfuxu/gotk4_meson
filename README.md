# Gotk4 Meson Example
This repository showcases an example of compiling Gotk4 projects using Meson build system.

## How to build:

### Prerequisites:

The following packages are required to build this project:

- Golang `go`
- Gtk 4 `gtk4`
- Meson `meson`
- Ninja `ninja-build`

### Building Instructions:

#### Global installation

```shell
git clone https://github.com/tfuxu/gotk4_meson.git
cd gotk4_meson
meson setup builddir
meson configure builddir -Dprefix=/usr/local
ninja -C builddir install
```

#### Local build (for testing and development purposes)

```shell
git clone https://github.com/tfuxu/gotk4_meson.git
cd gotk4_meson
meson setup builddir
meson configure builddir -Dprefix="$(pwd)/builddir" -Dbuildtype=debug
ninja -C builddir install
meson devenv -C builddir ./bin/gotk4_meson
```

> **Note** 
> During testing and development, as a convenience, you can use the [`local_run.sh`](./local_run.sh) script to quickly rebuild local builds.

## TODO:
- [ ] Add example Flatpak manifest
- [ ] Get translations to work
