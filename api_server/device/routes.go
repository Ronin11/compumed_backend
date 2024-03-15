package device

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"compumed/logging"
	"compumed/auth"
)

func CreateRoutes(router *mux.Router) http.Handler {
	logging.Log("Creating device routes")

	router.HandleFunc("/user_devices", auth.BuildRouteWithUser(getDevices).Handle).Methods("GET")
	router.HandleFunc("/register", auth.BuildRouteWithUser(register).Handle).Methods("POST")
	// router.HandleFunc("/addMeasurement", BuildRouteWithUser(addMeasurement).Handle).Methods("POST")
	// router.HandleFunc("/getMeasurement/{id}", BuildRouteWithUser(getMeasurement).Handle).Methods("GET")
	// router.HandleFunc("/updateMeasurement", BuildRouteWithUser(updateMeasurement).Handle).Methods("PUT")
	// router.HandleFunc("/deleteMeasurement", BuildRouteWithUser(deleteMeasurement).Handle).Methods("DELETE")

	return router
}

func getDevices(w http.ResponseWriter, r *http.Request, user *auth.User) {
	devices, err := GetDevices(user)
	if err != nil {
		logging.Log("ERR: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(devices)
	if err != nil {
		logging.Log("ERR: ", err)
		return
	}

	w.Write(jsonResponse)
}

type registerDeviceReq struct {
	Name string `json:"name"`
}

func register(w http.ResponseWriter, r *http.Request, user *auth.User) {
	var data registerDeviceReq
	err := json.NewDecoder(r.Body).Decode(&data)
	if (registerDeviceReq{}) == data {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		logging.Log("ADD ERR: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	device, err := RegisterDevice(user, data.Name)

	if err != nil {
		logging.Log("ERR: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(device.Serialize())
}

// func updateMeasurement(w http.ResponseWriter, r *http.Request, user *auth.User) {
// 	var measurement health.HealthMeasurement
// 	err := json.NewDecoder(r.Body).Decode(&measurement)
// 	if (health.HealthMeasurement{}) == measurement {
// 		w.WriteHeader(http.StatusNoContent)
// 		return
// 	}
// 	if err != nil {
// 		logging.Log("UPDATE ERR: ", err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	upatedMeasurement, err := health.UpdateMeasurement(user, &measurement)

// 	if err != nil {
// 		logging.Log("ERR: ", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(upatedMeasurement.Serialize())
// }

// func deleteMeasurement(w http.ResponseWriter, r *http.Request, user *auth.User) {
// 	var idObj IdReqObj
// 	err := json.NewDecoder(r.Body).Decode(&idObj)
// 	if (IdReqObj{}) == idObj {
// 		w.WriteHeader(http.StatusNoContent)
// 		return
// 	}
// 	if err != nil {
// 		logging.Log("DELETE ERR: ", err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	err = health.DeleteMeasurement(user, idObj.Id)

// 	if err != nil {
// 		logging.Log("GET ERR: ", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }
