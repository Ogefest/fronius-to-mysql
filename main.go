package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't load env file")
	}
	openDb()

	sleepInterval, err := strconv.Atoi(os.Getenv("REFRESH_INTERVAL"))
	if err != nil {
		fmt.Printf("Invalid interval: %s set default 10\n", err)
		sleepInterval = 10
	}

	deviceId, err := strconv.Atoi(os.Getenv("DEVICE_ID"))
	if err != nil {
		fmt.Printf("Invalid DEVICE_ID, %s\n", err)
	}

	for {

		data, err := getData(deviceId)
		if err != nil || data.PAC.Value == 0 {
			fmt.Printf("Error getting data: %s, PAC: %f\n", err, data.PAC.Value)
		} else {
			insertStmt := "INSERT INTO inverter_data (day_energy, fac, iac, idc, pac, total_energy, uac, udc, year_energy, device_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
			_, err = db.Exec(insertStmt, data.DAY_ENERGY.Value, data.FAC.Value, data.IAC.Value, data.IDC.Value, data.PAC.Value, data.TOTAL_ENERGY.Value, data.UAC.Value, data.UDC.Value, data.YEAR_ENERGY.Value, deviceId)
			if err != nil {
				log.Fatal(err)
			}
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			fmt.Printf("%s PAC: %.2f %s\n", currentTime, data.PAC.Value, data.PAC.Unit)
		}

		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}

}

func openDb() {
	// var err error

	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DBHOST") + ":" + os.Getenv("DBPORT"),
		DBName:               os.Getenv("DBNAME"),
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// db, err = sql.Open("duckdb", os.Getenv("DB_PATH")+"?access_mode=READ_WRITE")
	if err != nil {
		log.Fatal("Can't open database", err)
	}
}

func getData(deviceId int) (Data, error) {

	url := fmt.Sprintf("http://%s/solar_api/v1/GetInverterRealtimeData.cgi?Scope=Device&DeviceId=%d&DataCollection=CommonInverterData", os.Getenv("INVERTER_IP"), deviceId)

	// Make HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return Data{}, err
	}
	defer resp.Body.Close()

	// Decode JSON response
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return Data{}, err
	}

	data := response.Body.Data

	return data, nil
}
