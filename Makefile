default:
	touch bin/{t,t.big}
	rm bin/{t,t.big}
	go build -a -installsuffix cgo -ldflags "-s -w" -o bin/t.big
	upx bin/t.big -o bin/t
	rm bin/t.big
