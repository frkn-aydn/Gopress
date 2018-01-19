out:
	GOOS=windows GOARCH=386 go build -o gopress-windows-V0.0.1.exe ./server/main.go
	GOOS=darwin  GOARCH=386 go build -o gopress-mac-V0.0.1         ./server/main.go
	GOOS=linux   GOARCH=386 go build -o gopress-linux-V0.0.1       ./server/main.go
