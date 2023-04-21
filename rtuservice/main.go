package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/anonIot/srvgw/handler"
	"github.com/anonIot/srvgw/logs"
	"github.com/anonIot/srvgw/repository"
	"github.com/anonIot/srvgw/services"
	"github.com/goburrow/modbus"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func isAuthrized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "No Auth")
		endpoint(w, r)
	})
}

/*
**

	func getToken(name string) (string, error) {
		signingKey := []byte("keymaker")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name": name,
			"role": "redpill",
		})
		tokenString, err := token.SignedString(signingKey)
		return tokenString, err
	}

	func verifyToken(tokenString string) (jwt.Claims, error) {
		signingKey := []byte("keymaker")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
		})
		if err != nil {
			return nil, err
		}
		return token.Claims, err
	}

	func authMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			tokenString := r.Header.Get("Authorization")
			if len(tokenString) == 0 {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Missing Authorization Header"))
				return
			}
			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
			claims, err := verifyToken(tokenString)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Error verifying JWT token: " + err.Error()))
				return
			}
			name := claims.(jwt.MapClaims)["name"].(string)
			role := claims.(jwt.MapClaims)["role"].(string)

			r.Header.Set("name", name)
			r.Header.Set("role", role)

			next.ServeHTTP(w, r)
		})
	}
*/
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super hooks")
}

func main() {

	initTimeZone()

	rtuCon := initRtuConfig()

	client := repository.NewRtuBridgeDevice(rtuCon)
	rtuSev := services.NewRtuBridgeServiceDevice(client)
	acHandler := handler.NewRtuBridgeHandler(rtuSev)

	router := mux.NewRouter()

	// router.HandleFunc("/api/v2/indoor/{slaveID:[0-9]+}/{bmsID:[0-9]+}", acHandler.GetAcIndoor).Methods(http.MethodGet)
	router.Handle("/api/v2/indoor/{slaveID:[0-9]+}/{bmsID:[0-9]+}", isAuthrized(acHandler.GetAcIndoor)).Methods(http.MethodGet)
	router.HandleFunc("/api/v2/indoor/cmd/{slaveID:[0-9]+}/{bmsID:[0-9]+}/{cmd:[aA-zZ]+}/{val:[0-9]+}", acHandler.GetAcCmd).Methods(http.MethodGet)
	router.HandleFunc("/api/v2/indoor/{slaveID:[0-9]+}/{bmsID:[0-9]+}/power/{val:[0-1]+}", acHandler.GetAcPower).Methods(http.MethodGet)
	router.HandleFunc("/api/v2/indoor/{slaveID:[0-9]+}/{bmsID:[0-9]+}/temp/{val}", acHandler.GetAcTemp).Methods(http.MethodGet)
	router.HandleFunc("/api/v2/indoor/{slaveID:[0-9]+}/{bmsID:[0-9]+}/mode/{val:[0-9]+}", acHandler.GetAcMode).Methods(http.MethodGet)
	router.HandleFunc("/api/v2/indoor/{slaveID:[0-9]+}/{bmsID:[0-9]+}/fanspeed/{val:[0-9]+}", acHandler.GetFanSpeed).Methods(http.MethodGet)
	router.HandleFunc("/api/v2/indoor/{slaveID:[0-9]+}/{bmsID:[0-9]+}/swing/{val:[0-9]+}", acHandler.GetSwing).Methods(http.MethodGet)

	router.Handle("/auth", isAuthrized(homePage))

	logs.Log.Info("SRV RTU Local run on port: " + viper.GetString("srvlocal.port"))
	err := http.ListenAndServe(":3002", router)
	if err != nil {
		log.Fatalf("HTTP Server : %v", err)
	}

}

func initRtuConfig() *modbus.RTUClientHandler {

	initConfig()

	rtuPort := fmt.Sprintf("%v", viper.Get("rtu.uartport"))
	rtuBaudrate := viper.GetInt("rtu.baudrate")
	rtuParity := viper.GetString("rtu.parity")

	handler := modbus.NewRTUClientHandler(rtuPort)
	handler.BaudRate = rtuBaudrate
	handler.DataBits = 8
	handler.Parity = rtuParity
	handler.StopBits = viper.GetInt("rtu.stopbits")
	handler.SlaveId = 1
	handler.Timeout = 3 * time.Second

	err := handler.Connect()
	defer handler.Close()

	if err != nil {
		log.Fatalf(" No Connect : %v", err)
		return nil
	}
	return handler
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}
