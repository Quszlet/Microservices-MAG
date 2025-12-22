module github.com/Quszlet/doctors_service

go 1.25.1

require (
	github.com/Quszlet/libs/json v0.0.0
	github.com/Quszlet/libs/kafka v0.0.0
	github.com/Quszlet/libs/utils_sql v0.0.0
	github.com/fatih/structs v1.1.0
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/gorilla/mux v1.8.1
	github.com/jmoiron/sqlx v1.4.0
	github.com/lib/pq v1.10.9
	github.com/prometheus/client_golang v1.19.1
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.48.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/segmentio/kafka-go v0.4.47 // indirect
	golang.org/x/sys v0.17.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/Quszlet/libs/json => ../libs/json

replace github.com/Quszlet/libs/kafka => ../libs/kafka

replace github.com/Quszlet/libs/utils_sql => ../libs/utils_sql
