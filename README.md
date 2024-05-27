# MEMERENAMER

**This is an exploratory project meant for myself; no warranty, no support.**

What I'm trying to do here is building a _meme renamer_. A program that will
help me easily go through my meme folder and give the million different jpegs a
new name.

A neat idea I've come up with for that is: circumvent *any* kind of OS specific
gui shenanigans by making it a http server. Sacrilege, I know. I hold the
opinion that if anyone has figured out how to do good GUIs, it's
web-developers. Not by competence but by brownian motion and the invisible hand
of the market. They had to figure out how to render interactive widgets on
screen for decades and, as much as a "real" programmer might dislike this, are
doing an alright job of it (framework bloat notwithstanding). So rather than
fiddling with nasty maybe-maybe-not crossplatform tools, I offload that work
onto Google and Mozilla.

From there the core idea is simple. Start the server in a directory (your meme
folder), and start serving one image after another, along with a form for
things like name and description (maybe fiddle with exif data? idk).

## TODO

- server recognizes properly formatted image names and won't offer them for renaming.
- skip button
- all lowercase renames? -> command line option/ checkboxes on the site?
- suport video formats

## misc

chores that have cropped up

No one here but us chickens

## DONE

- add css to handle screen width
- basics:
  - shows images that are in current directory
  - has an input form
  - renames files
  - shows next image after rename
- name can be written like a sentence and get reformatted to `snake_case` without any weird chars.
- support more file types than just jpg -> now also png!


# Feature Creep

If it's gonna happen anyway, might as well plan for it.

## TODO

1. navbar to the side, that shows images to be renamed.
2. neat htmx auto-check if a filename is taken.
3. preview of what the name will be reformatted to.
4. maybe put extra metadata into the files?
5. cmdline args to control various things (e.g. port to run on)


# Build

This project is dependency free and compiles to a standalone, placement independent binary.

```{bash}
    go build
```

or using make

```{bash}
    make
    sudo make install
```


# Usage

```{bash}
    # if not installed
    <path to binary>/memerenamer
    # else
    memerenamer
```

