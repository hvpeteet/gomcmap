# gomcmap

This is a fork of kch42's [gomcmap](https://github.com/kch42/gomcmap). I just wanted to play around with this library since it seemed awesome, but it did not yet support all the functionality that I wanted so I have added some functionality. For the original documentation see the [original repo](https://github.com/kch42/gomcmap).

Mainly I intend to just add more examples and code to make new kinds of maps. I only intend to edit the original code base when I need a new feature that is not yet supported.

## List of added features

1. Can lookup BlockID by the block name (see mcmap/block:BlockIDFromString)

## Playground rules

This is only meant to be a playground and not a project with full support or backwards compatibility as I will likely make some breaking changes in the future. For a dependable repo please use the [original repo](https://github.com/kch42/gomcmap).

## Install

	go get github.com/hvpeteet/gomcmap/mcmap

## Examples

See `mcmap/examples`

## Compatibility

Currently only the Anvil map format is supported. Tested with Minecraft 1.6.2, will probably work with older versions too, but I haven't tested it.

## WARNING

Although I tested the library with some maps, I can't guarantee that everything always works (especially if you use mods in your Minecraft installation). So make a backup of your maps, just in case!
