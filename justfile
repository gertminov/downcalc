build:
  go build -o downcalc ./main

crosscompile:
  just ensurebindir
  env GOOS=linux GOARCH=amd64 go build -o bin/downcalc-amd64-linux ./main; 
  env GOOS=darwin GOARCH=amd64 go build -o bin/downcalc-amd64-macos ./main; 
  env GOOS=windows GOARCH=amd64 go build -o bin/downcalc-amd64-windows.exe ./main; 
  env GOOS=darwin GOARCH=arm64 go build -o bin/downcalc-arm64-macos ./main; 

ensurebindir:
  mkdir -p bin
