TITLE BUILD-BAT

ECHO GO BUILD CLI : GO-MOD GET PACKAGE TO VENDOR 
go mod vendor

ECHO GO BUILD CLI : ROOM-STATUS SERVICE PROGRAM
go build -o bin\_roomstatus.exe build_cli\room_status.go

ECHO GO BUILD CLI : AUTH-RELATED SERVICE PROGRAM
go build -o bin\_auth.exe build_cli\auth_server.go


