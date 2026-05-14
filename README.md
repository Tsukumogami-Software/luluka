# 📓 Luluka: Ebiten shader tool

Luluka lets you view and tweak shaders made with Kage for [Ebiten](https://github.com/hajimehoshi/ebiten).

<img width="928" height="480" alt="luluka" src="https://github.com/user-attachments/assets/cde37724-a1e6-4d84-ba9c-517c821f92e1" />


## Installation

Install luluka using the following command:

```sh
go install github.com/Tsukumogami-Software/luluka@v1.0.0
```

## Usage

Launch a shader file by passing it as the first argument to luluka:

```sh
luluka sample/transition.kage
```

Pass textures using `-i` (max 4):

```sh
luluka sample/transition.kage -i image.png -i image2.png
```

Pass uniform variables using `-u`:

```sh
luluka sample/transition.kage -i image.png -i image2.png -u Steepness:80 -u Speed:0.08
```

Vectors, array and matrices are supported using their flattened index:

```sh
luluka sample/transition.kage -i image2.png -i image.png -u Steepness:80 -u Seed.0:15 -u Seed.1:100 -u Seed.2:5000 -u Seed.3:5000 -u Speed:0.08
```

The following uniforms are automatically given passed by Luluka:
* Time: the number of seconds since the program start
* Cursor: the relative cursor position
* MouseButtons: the mouse buttons status encoded on an int

For convenience, you can use a YAML file instead of passing every uniform to the command line.

```yaml
Steepness: 80
Seed: [15.0, 100.0, 5000.0]
Speed: 0.08
```

```sh
luluka sample/transition.kage -i image2.png -i image.png -v values.yaml
```

## Build from source

Clone and compile this repository:

```sh
git clone git@github.com:Tsukumogami-Software/luluka.git
cd luluka
go build -o luluka .
```

(Requires go 1.26)

## Learn more about Kage

Check out the amazing [Kage's desk](https://github.com/tinne26/kage-desk/).
