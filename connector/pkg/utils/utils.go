package utils

import (
	"GoJira/pkg/structure"
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, projects structure.RestProjects) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func RespondError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}
