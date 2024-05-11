# gophix

A simple tool to merge Google Photos Takeout json metadata with the media files themselves. This works with both photos
and videos.

This works with all the media formats compatible with [Phil Harvey's ExifTool](https://exiftool.org). If the format is
not supported, gophix will write the metadata in a xmp sidecar file.

## Features

- Merges media creation date, GPS coordinates, and description from json files with the media files
- Fixes media file extensions if necessary (for example png files with a jpg extension)
- Works with any media format (although some formats require a sidecar file)
- Cleans up the json files

## Motivation

Before I built this, I looked for an existing tool that could do this, but I couldn't find one that satisfied all my
needs. Some of the issues I had with existing tools were:

- They did not work with all the quirky json file names that Google Photos Takeout generates
- They were prioritizing the metadata from the media file itself over the metadata from the json file
- They were not merging the GPS metadata or the description from the json file with the media
- They did not work with videos

## Usage

First, download your [Google Photos Takeout](http://takeout.google.com) archive and extract it. You should have a folder
with the following structure:

```
Takeout/
  Google Photos/
    Photos of 2001/
      IMG_0001.jpg
      IMG_0001.json
      IMG_0002.jpg
      IMG_0002.json
    Photos of 2002/
      IMG_0003.jpg
      IMG_0003.json
      IMG_0004.jpg
      IMG_0004.json
```

gophix supports 2 operations: `fix` and `clean-json`.

### fix

To fix the files and merge the metadata, run the following command:

```shell
gophix fix <path>
```

Where `<path>` is the path to the `Takeout/Google Photos` folder. For example:

```shell
gophix fix ~/Downloads/Takeout/Google\ Photos
```

**Note** the backslash before the space in the path. This is necessary because the shell interprets the space as a
separator between arguments.

Optionally, you can save the output logs to a file:

```shell
gophix fix <path> | tee logs.txt
```

### clean-json

To clean up the json files (potentially after running `fix`), run the following command:

```shell
gophix clean-json <path>
```

## Build

```shell
go build -o gophix
```

## To do

- [ ] Make it idempotent (right now it cannot find the associated json files after fixing file extensions)
    - Could be fixed by keeping state of the renamed files in an additional json file
- [ ] Release binaries
- [ ] Add timezone offset to the metadata
    - Could be computed using the GPS coordinates if present