package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	// limit the maximum size of JSON allowed to be read
	maxBytes := 1048576 // 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{}) // ensure body only has a single JSON file
	// EOF just a constant implying that you have reached the end of the file. EOF is the error
	// returned by Read when no more input is available. In this check, we just make sure that we have
	// read the entire JSON that we have received. If we get an error that is `io.EOF`,
	// then all is good --> there is nothing else to read. However, any other errors means that something
	// went wrong --> we have more data to read, and we shouldn't have (and don't want) any more.
	if err != io.EOF {
		return errors.New("body must have only a single json value")
	}

	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// check if we have any headers
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) {
	// default status code if not provided
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	app.writeJSON(w, statusCode, payload)
}