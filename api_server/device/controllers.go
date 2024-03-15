package device

import (
	"compumed/auth"
)

func GetDevices(user *auth.User) ([]Device, error) {
	storageHandler := GetStorageHandlerInstance()
	return storageHandler.QueryUserDevices(user)
}

func RegisterDevice(user *auth.User, name string) (*Device, error) {
	storageHandler := GetStorageHandlerInstance()
	return storageHandler.InsertDevice(user, name)
}

// func DeleteMeasurement(user *auth.User, id string) (error) {
// 	storageHandler := GetStorageHandlerInstance()
// 	return storageHandler.DeleteMeasurement(user, id)
// }
