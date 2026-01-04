package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nitesh111sinha/car-management/models"
	"github.com/nitesh111sinha/car-management/service"
	"go.opentelemetry.io/otel"
)

type EngineHandler struct {
	engineService service.EngineServiceInterface
}

func NewEngineHandler(engineService service.EngineServiceInterface) *EngineHandler {
	return &EngineHandler{
		engineService: engineService,
	}
}

func (h *EngineHandler) GetEngineById(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("engine-handler")
	ctx, span := tracer.Start(r.Context(), "GetEngineById-Handler")
	defer span.End()
	vars := mux.Vars(r)
	id := vars["id"]
	engine, err := h.engineService.GetEngineById(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(engine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("engine-handler")
	ctx, span := tracer.Start(r.Context(), "CreateEngine-Handler")
	defer span.End()
	var engine models.Engine
	err := json.NewDecoder(r.Body).Decode(&engine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	engine.EngineID = uuid.New()
	createdEngine, err := h.engineService.CreateEngine(ctx, engine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdEngine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *EngineHandler) GetEngines(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("engine-handler")
	ctx, span := tracer.Start(r.Context(), "GetEngines-Handler")
	defer span.End()
	engines, err := h.engineService.GetEngines(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(engines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *EngineHandler) UpdateEngine(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("engine-handler")
	ctx, span := tracer.Start(r.Context(), "UpdateEngine-Handler")
	defer span.End()
	vars := mux.Vars(r)
	id := vars["id"]
	var engine models.Engine
	err := json.NewDecoder(r.Body).Decode(&engine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	engineID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	engine.EngineID = engineID
	updatedEngine, err := h.engineService.UpdateEngine(ctx, engineID.String(), engine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(updatedEngine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *EngineHandler) DeleteEngine(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("engine-handler")
	ctx, span := tracer.Start(r.Context(), "DeleteEngine-Handler")
	defer span.End()
	vars := mux.Vars(r)
	id := vars["id"]
	if err := h.engineService.DeleteEngine(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
