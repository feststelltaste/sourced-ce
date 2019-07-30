:: we can't use makefile for windows because it depends on CI makefile which depends on shell

:: compile sourced-ce
if not exist build mkdir build
if not exist build/sourced-ce_windows_amd64 mkdir build/sourced-ce_windows_amd64
go build -tags=forceposix -o build/sourced-ce_windows_amd64/sourced.exe ./cmd/sourced

:: run tests
go test -timeout 20m -parallel 1 -count 1 -tags="forceposix integration" github.com/src-d/sourced-ce/test/ -v