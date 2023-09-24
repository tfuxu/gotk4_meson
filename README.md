# Gotk4 Meson Example
This repository showcases an example of compiling Gotk4 projects using Meson build system.

## How to build:

### Flatpak Builder:

#### Prerequisites:

- Flatpak Builder `flatpak-builder`
- GNOME SDK runtime `org.gnome.Sdk//45`
- GNOME Platform runtime `org.gnome.Platform//45`

Install required runtimes:
```shell
flatpak install org.gnome.Sdk//45 org.gnome.Platform//45
```

#### Building Instructions:

##### User installation
```shell
git clone https://github.com/tfuxu/gotk4_meson.git
cd gotk4_meson
flatpak-builder --install --user --force-clean repo/ build-aux/flatpak/io.github.tfuxu.gotk4_meson.json
```

##### System installation
```shell
git clone https://github.com/tfuxu/gotk4_meson.git
cd gotk4_meson
flatpak-builder --install --system --force-clean repo/ build-aux/flatpak/io.github.tfuxu.gotk4_meson.json
```

### Meson Build System:

#### Prerequisites:

The following packages are required to build this project:

- Golang `go`
- Gtk 4 `gtk4`
- Meson `meson`
- Ninja `ninja-build`

#### Building Instructions:

##### Global installation

```shell
git clone https://github.com/tfuxu/gotk4_meson.git
cd gotk4_meson
meson setup builddir
meson configure builddir -Dprefix=/usr/local
ninja -C builddir install
```

##### Local build (for testing and development purposes)

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
- [x] Add example Flatpak manifest
- [ ] Get translations to work

## Notes:
- It takes really long (around 40 minutes) to finish building Flatpak package without any cached libraries (see [#66](https://github.com/diamondburned/gotk4/issues/66) in gotk4)
