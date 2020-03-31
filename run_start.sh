go run build_cli/auth_server.go start -c=config.yaml >> log/auth.log 2>&1 &
go run build_cli/room_status.go start -c=config.yaml >> log/room_status.log 2>&1 &