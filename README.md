# gowest22 material

Get presentation [here](https://docs.google.com/presentation/d/1u-2E9gS3hhb9_mWwDRsjyU9POVvfvC6svIk3NBkSxRQ/edit?usp=sharing).

## Compile protos
```bash
$ docker run -v $PWD:/defs namely/protoc-all \
  -f cli.proto \
  -l go \
  -o build/go/opts \
  --go-source-relative
```

## Install `protoc-go-inject-tag`
```bash
$ go install github.com/favadi/protoc-go-inject-tag@latest
```

## Inject tags into compiled protos
```bash
$ protoc-go-inject-tag -input="build/go/opts/*.go"
```
