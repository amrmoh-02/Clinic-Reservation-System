package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func makeRequest(endpoint, apiKey, method, payloadStr string) (string, error) {
	baseurl := "https://data.mongodb-api.com/app/data-hbqst/endpoint/data/v1/action"
	url := fmt.Sprintf("%s/%s", baseurl, endpoint)

	payload := strings.NewReader(payloadStr)

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")
	req.Header.Add("api-key", apiKey)
	req.Header.Add("Accept", "application/ejson")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func signIn(apiKey, username, password string) {
	method := "POST"
	payload := fmt.Sprintf(`{
        "collection": "users",
        "database": "Hospital",
        "dataSource": "Cluster0",
        "username": "%s",
        "password": "%s"
    }`, username, password)

	response, err := makeRequest("signIn", apiKey, method, payload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if strings.Contains(response, "doctor") {
		fmt.Println("Signed in as a doctor.")
	} else if strings.Contains(response, "patient") {
		fmt.Println("Signed in as a patient.")
	} else {
		fmt.Println("Invalid credentials or user type.")
	}
}

func signUp(apiKey, username, password, userType string) {
	method := "POST"
	var payloadStr string
	if userType == "doctor" {
		payloadStr = fmt.Sprintf(`{
			"collection": "users",
			"database": "Hospital",
			"dataSource": "Cluster0",
			"username": "%s",
			"password": "%s",
			"DName": "%s"
		}`, username, password, username)

		makeRequest("signUp", apiKey, method, payloadStr)
		setSchedule(apiKey, username, "Monday: 9 AM - 5 PM")
	} else if userType == "patient" {
		payloadStr = fmt.Sprintf(`{
			"collection": "users",
			"database": "Hospital",
			"dataSource": "Cluster0",
			"username": "%s",
			"password": "%s",
			"Pname": "%s"
		}`, username, password, username)
		makeRequest("signUp", apiKey, method, payloadStr)
		providePatientInfo(apiKey, username, "25")
	} else {
		fmt.Println("Invalid user type")
	}
}

func setSchedule(apiKey, DName, schedule string) {
	method := "POST"
	payloadStr := fmt.Sprintf(`{
        "collection": "doctors",
        "database": "Hospital",
        "dataSource": "Cluster0",
		"projection": {"_id": 1},
        "DName": "%s",
        "schedule": "%s"
    }`, DName, schedule)

	makeRequest("setSchedule", apiKey, method, payloadStr)
}
func providePatientInfo(apiKey, Pname, age string) {
	method := "POST"
	payloadStr := fmt.Sprintf(`{
        "collection": "patients",
        "database": "Hospital",
        "dataSource": "Cluster0",
        "Pname": "%s",
        "age": "%s"
    }`, Pname, age)

	makeRequest("providePatientInfo", apiKey, method, payloadStr)
}

func viewAvailableSlots(apiKey, id string) {
	method := "GET"
	payload := fmt.Sprintf(`{
        "collection": "doctors",
        "database": "Hospital",
        "dataSource": "Cluster0",
        "id": "%s"
    }`, id)

	makeRequest("viewAvailableSlots", apiKey, method, payload)
}

func updateAppointment(apiKey, id, Pid, time, date, status string) {
	method := "PUT"
	payload := fmt.Sprintf(`{
        "collection": "appointments",
        "database": "Hospital",
        "dataSource": "Cluster0",
        "id": "%s",
        "Pid": "%s",
        "time": "%s",
        "date": "%s",
        "status": "%s"
    }`, id, Pid, time, date, status)

	makeRequest("updateAppointment", apiKey, method, payload)
}

func cancelAppointment(apiKey, id, Pid string) {
	method := "DELETE"
	payload := fmt.Sprintf(`{
        "collection": "appointments",
        "database": "Hospital",
        "dataSource": "Cluster0",
        "id": "%s",
        "Pid": "%s"
    }`, id, Pid)

	makeRequest("cancelAppointment", apiKey, method, payload)
}

func viewReservations(apiKey, Pid string) {
	method := "GET"
	payload := fmt.Sprintf(`{
        "collection": "appointments",
        "database": "Hospital",
        "dataSource": "Cluster0",
        "Pid": "%s"
    }`, Pid)

	makeRequest("viewReservations", apiKey, method, payload)
}

func main() {

	signInAPIKey := "xRDvZVwjeQNPcxfwbS2wOyr2NIfoTykQJ0ioBKj9KjWIqe2xcOFmXBT8QZR819T2"
	signUpAPIKey := "SrMu6gItDG0gKDPiGS3V7ZNz8dWqW26AYZRa8sF6Sj0KoYdOoQ1udi0yhQf1QuRS"
	setScheduleAPIKey := "09VB3bYkNHg373as6LUR7T0B6GNRRLnsBZ3QkO5xGfYRghY513g6ZhsVNvVGbIO2"
	viewAvailableSlotsAPIKey := "Ys0E5TIqynHQ1r62MTR27Ui2LMq8AETB84FbzUfqIX6srK8KpIVqaPmQQ0loNHUI"
	updateAppointmentAPIKey := "gTdiArgHXZps9UMxuWH6KAekN0I9ahyjmDPJMUvNKMbl8IIu85pYmJQQdQHwtIJe"
	cancelAppointmentAPIKey := "uHkV5Cxnrbjrpyx6ZUwcGyLfoWCjMmkqe6KEgbVgieSCR30vZPPI7PLoAS8DrXYE"
	viewReservationsAPIKey := "yfjFEG7UADJsIVU22oqMG3RfBFcLOS60Isk7HSP7lTo7tDW7X1KOAGMwMOW217oV"

	signUp(signUpAPIKey, "doctor_john", "password123", "doctor")
	signUp(signUpAPIKey, "patient_jane", "password456", "patient")
	signIn(signInAPIKey, "doctor_john", "password123")
	setSchedule(setScheduleAPIKey, "Dr.Ali", "Monday: 9 AM - 5 PM")
	viewAvailableSlots(viewAvailableSlotsAPIKey, "Dr.Ali")
	updateAppointment(updateAppointmentAPIKey, "1", "20206100", "1:45PM", "Monday", "Confirmed")
	cancelAppointment(cancelAppointmentAPIKey, "1", "20206100")
	viewReservations(viewReservationsAPIKey, "20206100")

}
