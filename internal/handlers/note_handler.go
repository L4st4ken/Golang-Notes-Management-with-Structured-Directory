package handlers

import (
	"encoding/json"
	"net/http"
	"notesmanagement/internal/models"
	"strconv"
	"strings"
	"notesmanagement/internal/services"
)

type NoteHandler struct{
	service *services.NoteService
}

func NewNoteHandler(service *services.NoteService) *NoteHandler{
	return &NoteHandler{service: service}
}

//Handling untuk GET api/note
func (h *NoteHandler) HandleNotes(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case http.MethodGet:
			h.GetAll(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *NoteHandler) GetAll(w http.ResponseWriter, r *http.Request){
	notes, err := h.service.GetAll()
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)

}

func (h *NoteHandler) Create(w http.ResponseWriter, r *http.Request){
	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//Simpan Ke DB
	if err := h.service.Create(&note); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Tampilkan Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)

	defer r.Body.Close()
}

//Handle Note yang menggunakan ID
func (h *NoteHandler) HandleNoteByID(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case http.MethodGet:
			h.GetByID(w, r)
		case http.MethodPut:
			h.Update(w, r)
		case http.MethodDelete:
			h.Delete(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *NoteHandler) GetByID(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/api/note/")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "Invalid Note ID", http.StatusBadRequest)
		return
	}

	note, err := h.service.GetByID(id)
	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func(h *NoteHandler) Update(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/api/note/")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "Invalid Note ID", http.StatusBadRequest)
		return
	}

	var note models.Note
	err = json.NewDecoder(r.Body).Decode(&note)
	if err != nil{
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	note.ID = id
	err = h.service.Update(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func (h *NoteHandler) Delete(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/api/note/")
	id,err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}
	
	err = h.service.Delete(id)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Note deleted successfully",
	})
}