# Automating Audacity

Given an input audio file, use Audacity's scripting interface to apply audio effects, such as Noise Reduction.

Port the mod-script-pipe Python client to golang.

## Audacity version

Audacity on Linux does not include mod-script-pipe, so this uses a custom-built version of Audacity for Ubuntu 18.04. The [build instructions](docker/README.md) were adapted from Ubuntu 20.04, so should also work there with some minor adjustments to dependencies.sh.