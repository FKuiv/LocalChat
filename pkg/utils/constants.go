package utils

import "time"

const (
	KB                        = 1 << 10
	MB                        = 1 << 20
	GB                        = 1 << 30
	URL_expiration_time       = time.Second * 60
	MINIO_bucket              = "localchat"
	MULTIPART_FORM_MAX_MEMORY = 5 * MB
)
