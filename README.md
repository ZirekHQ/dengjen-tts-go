# dengjen-tts-go

Go bindings for [dengjen-tts](https://github.com/ZirekHQ/dengjen-tts),
published with prebuilt native binaries for linux/amd64,
linux/arm64, windows/amd64, and darwin/arm64.

```bash
go get github.com/ZirekHQ/dengjen-tts-go@v0.1.0
```

This repository is generated on every release from
[`bindings/go`](https://github.com/ZirekHQ/dengjen-tts/tree/main/bindings/go)
in the parent monorepo — do not hand-edit files here; changes go
through the parent repo's `bindings/go` instead, and a new
tagged release regenerates this repo from it.

On Windows, cgo has no rpath equivalent: copy
`lib/windows_amd64/libdengjen.dll` next to your built executable
(or put that directory on `PATH`) before running it. The name
matters — that is the import name baked into the linked binary.
