package device

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"compumed/auth"
	"compumed/logging"
)

type StorageHandler struct {
	dbUrl     string
	dbpool    *pgxpool.Pool
	tableName string
}

var storageHandler *StorageHandler

func initialize(fullUrl string, tableName string) *StorageHandler {
	storageHandler = &StorageHandler{dbUrl: fullUrl}
	dbpool, err := pgxpool.New(context.Background(), fullUrl)
	storageHandler.dbpool = dbpool
	storageHandler.tableName = tableName
	if err != nil {
		logging.Log("Database connection error: ", err)
	} else {
		logging.Log("Database connection established")
	}
	return storageHandler
}

func GetStorageHandlerInstance() *StorageHandler {
	if storageHandler != nil {
		return storageHandler
	}

	return initialize(os.Getenv("POSTGRES_URL"), os.Getenv("POSTGRES_DEVICES_TABLE_NAME"))
}

func (sh *StorageHandler) Cleanup() error {
	sh.dbpool.Close()
	return nil
}

func (sh *StorageHandler) QueryUserDevices(user *auth.User) ([]Device, error) {
	ctx := context.Background()
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=@userID", sh.tableName)
	args := pgx.NamedArgs{
		"userID": user.ID,
	}
	rows, err := sh.dbpool.Query(ctx, query, args)
	if err != nil {
		logging.Log("Fetch Err: ", err)
		return nil, err
	}
	defer rows.Close()

	logging.Log("ROWS: ", rows)
	var devices []Device

	for rows.Next() {
		var device Device
		err := rows.Scan(&device.ID, &device.UserID, &device.Name)
		if err != nil {
			logging.Log("SCAN ERR: ", err)
		}
		devices = append(devices, device)
	}
	return devices, nil
}

// func (sh *StorageHandler) GetDevice(user *auth.User, id string) (*Device, error){
// 	rows, err := sh.db.Query(context.Background(), fmt.Sprintf("SELECT * FROM %s WHERE id=$1", sh.tableName), id)
// 	if err != nil {
// 		logging.Log("Fetch Err: ", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	device := &Device{}
// 	for rows.Next() {
// 		var hm HealthMeasurement
// 		err := rows.Scan(&hm.Id)
// 		if err != nil {
// 			logging.Log("SCAN ERR: ", err)
// 		}
// 		device.Data = append(device.Data, hm)
// 	}

// 	return device, nil
// }

func (sh *StorageHandler) InsertDevice(user *auth.User, name string) (*Device, error) {
	ctx := context.Background()
	query := fmt.Sprintf("INSERT INTO %s (user_id, name) VALUES (@userID, @deviceName) RETURNING *", sh.tableName)
	args := pgx.NamedArgs{
		"userID":     user.ID,
		"deviceName": name,
	}
	var device Device
	err := sh.dbpool.QueryRow(ctx, query, args).Scan(&device.ID, &device.UserID, &device.Name)
	if err != nil {
		logging.Log("CREATE FAILED: ", err)
	}

	return &device, err
}

// func (sh *StorageHandler) CreateMeasurement(user *auth.User, data *HealthData) (*HealthMeasurement, error){
// 	var hm HealthMeasurement
// 	err := sh.db.QueryRow(
// 		context.Background(),
// 		fmt.Sprintf("INSERT INTO %s (user_id, data) VALUES($1, $2) RETURNING *", sh.tableName), user.Id, data).Scan(&hm.Id, &hm.UserId, &hm.CreatedTime, &hm.Data)
// 	if err != nil {
// 		logging.Log("CREATE FAILED: ", err)
// 	}

// 	return &hm, err
// }

// func (sh *StorageHandler) UpdateMeasurement(user *auth.User, hm *HealthMeasurement) (*HealthMeasurement, error){
// 	var hm2 HealthMeasurement
// 	err := sh.db.QueryRow(
// 		context.Background(),
// 		fmt.Sprintf("UPDATE %s SET data=$1 WHERE user_id=$2 AND id=$3 RETURNING *", sh.tableName), hm.Data, user.Id, hm.Id).Scan(&hm2.Id, &hm2.UserId, &hm2.CreatedTime, &hm2.Data)
// 	if err != nil {
// 		logging.Log("UPDATE FAILED: ", err)
// 	}

// 	return &hm2, err
// }

// func (sh *StorageHandler) DeleteMeasurement(user *auth.User, id string) (error){

// 	_, err := sh.db.Exec(
// 		context.Background(),
// 		fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND id=$2", sh.tableName), user.Id, id)
// 	if err != nil {
// 		logging.Log("DELETE FAILED: ", err)
// 	}

// 	return err
// }
