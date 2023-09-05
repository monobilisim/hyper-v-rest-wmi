//go:build windows

package rest

import (
	"encoding/json"
	"errors"
	"hyper-v-rest/wmi"
	"net/http"

	"github.com/gorilla/mux"
)

func vms(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp response

	data, err := wmi.VMs()
	if err != nil {
		httpError(w, err, http.StatusInternalServerError, resp)
		return
	}

	if len(data) == 0 {
		httpError(w, errors.New("no VM found"), http.StatusNotFound, resp)
		return
	}

	resp.Result = "success"
	resp.Message = "VMs are listed in data field."
	resp.Data = data

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}

func memory(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp response

	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		httpError(w, errors.New("name is missing in parameters"), http.StatusBadRequest, resp)
		return
	}

	data, err := wmi.Memory(name)
	if err != nil {
		httpError(w, err, http.StatusInternalServerError, resp)
		return
	}

	if len(data) == 0 {
		httpError(w, errors.New("no memory info found"), http.StatusNotFound, resp)
		return
	}

	resp.Result = "success"
	resp.Message = "Memory info is displayed in data field."
	resp.Data = data

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}

func summary(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp response

	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		httpError(w, errors.New("name is missing in parameters"), http.StatusBadRequest, resp)
		return
	}

	data, err := wmi.Summary(name)
	if err != nil {
		httpError(w, err, http.StatusInternalServerError, resp)
		return
	}

	if len(data) == 0 {
		httpError(w, errors.New("no summary info found"), http.StatusNotFound, resp)
		return
	}

	resp.Result = "success"
	resp.Message = "Summary info is displayed in data field."
	resp.Data = data

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}

func vhd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp response

	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		httpError(w, errors.New("name is missing in parameters"), http.StatusBadRequest, resp)
		return
	}

	data, err := wmi.Vhd(name)
	if err != nil {
		httpError(w, err, http.StatusInternalServerError, resp)
		return
	}

	if len(data) == 0 {
		httpError(w, errors.New("no image info found"), http.StatusNotFound, resp)
		return
	}

	resp.Result = "success"
	resp.Message = "Image info is displayed in data field."
	resp.Data = json.RawMessage(data)

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}

func ip(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp response

	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		httpError(w, errors.New("Name is missing in parameters"), http.StatusBadRequest, resp)
		return
	}

	data, err := wmi.Ip(name)
	if err != nil {
		httpError(w, err, http.StatusInternalServerError, resp)
		return
	}

	if len(data) == 0 {
		httpError(w, errors.New("No network info found"), http.StatusNotFound, resp)
		return
	}

	resp.Result = "success"
	resp.Message = "Network info is displayed in data field."
	resp.Data = data

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}

func version(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp response

	resp.Result = "success"
	resp.Message = "Version is displayed in data field."
	resp.Data = "2.0"

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}
