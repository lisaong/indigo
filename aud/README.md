# Automating Audacity

Given an input audio file, use Audacity's scripting interface to apply audio effects, such as Noise Reduction.

Port the mod-script-pipe Python client to golang.

## Audacity Setup

Audacity on Linux does not include mod-script-pipe, so this uses a custom-built version of Audacity for Ubuntu 18.04. The [build instructions](docker/README.md) were adapted from Ubuntu 20.04, so should also work there with some minor adjustments to dependencies.sh.

### One-time setup

#### Import macros

1. Tools > Macros
2. Import ...
3. Select macros from [this](macros) folder

#### Get the noise profile

1. Load a representative audio file that contains noise in the first
2. Tools -> Apply Macro
3. Select the "Noise Profile" macro from the list

## Scripting

