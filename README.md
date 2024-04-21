# Introduction

The Solar Inverter Data Logger project aims to provide a solution for retrieving real-time data from solar inverters and logging it into a MySQL database. This project utilizes Go programming language to interact with the solar inverter's API, fetch real-time data in JSON format, and parse it. The parsed data is then stored in a MySQL database for further analysis and monitoring.

## Features
Fetch real-time data from solar inverters using HTTP requests.
- Parse JSON data received from the solar inverter API.
- Store parsed data into a MySQL database for long-term storage and analysis.

## Installation


### Prerequisites
Before installing and running the Fronius Inverter Data Logger, ensure you have the following prerequisites installed and configured:

- Go programming language (version 1.16 or later)
- MySQL database server

#### Create a .env file: Create a .env file in the project directory and configure it with the following environment variables:

```
DBUSER=user
DBPASS=pass
DBNAME=dbname
DBHOST=dbhost
DBPORT=3306

DEVICE_ID=1
INVERTER_IP=#
REFRESH_INTERVAL=10 #in seconds
```

Make sure to replace the placeholder values (user, pass, dbname, etc.) with your actual database credentials and Fronius inverter IP address.
Build the binary: Build the Go application to create the executable binary:

```
go build -o fronius-inverter-data-logger .
```

### Running the Application
Once you have configured the .env file properly, you can run the binary:


```./fronius-inverter-data-logger```

This command will start the Fronius Inverter Data Logger application, which will fetch real-time data from the Fronius inverter, parse it, and store it in the configured MySQL database at the specified intervals.


