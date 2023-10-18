package payload

func New() []byte {
	uuid := []byte("628aace1-ad98-4f44-8a5c-cb73862c46c4")
	payload := []byte{123, 34, 116, 97, 115, 107, 34, 58, 34, 116, 101, 115, 116, 34, 125}

	return append(uuid, payload...)
}
