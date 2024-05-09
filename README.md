# gophix

## Run

```shell
gophix fix <path>
```

Optionally, you can save the output logs to a file:

```shell
gophix fix <path> | tee logs.txt
```

## Build

```shell
go build -o gophix
```

## To do
- [ ] Write a better README
- [ ] Make it idempotent (right now it cannot find the associated json files after fixing file extensions)
  - Could be fixed by keeping state of the renamed files in an additional json file
- [ ] Add timezone offset to the metadata
  - Could be computed using the GPS coordinates if present