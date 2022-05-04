# mage-loot sample

![mage-boots](mage-boots.png)

This project serves as an example for the [mage-loot](https://github.com/aserto-dev/mage-loot) project.

The binary is supposed to print ASCII art of shoes.
To do this, at build time we download a random shoe image, then we use [ascii-image-converter](https://github.com/TheZoraiz/ascii-image-converter) to create ASCII art.
We define a binary dependency to this tool in our `Depfile` and then use it from the `magefile`.

It uses a common build target to compile this project and to version it using [sver](https://github.com/aserto-dev/sver).
