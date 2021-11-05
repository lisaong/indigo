# Compiling Audacity for Linux with mod-script-pipe

Tested with Ubuntu 18.04

## Build

```shell
sh build.sh
sudo chown -R $USER build/bin/Release
```

## Configure
Launch Audacity and verify that mod-script-pipe is enabled

```shell
cd build/bin/Release
LD_LIBRARY_PATH=lib/audacity ./audacity
```

1. Edit > Preferences > Modules

2. Set mod-script-pipe to "Enabled"

3. Close and re-launch audacity to verify settings are saved

## Test
```shell
cd audacity/scripts/piped-work
python3 pipe_test.py 
```