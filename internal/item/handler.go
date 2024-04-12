package item

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shynn12/biocad/internal/config"
	"github.com/shynn12/biocad/internal/handlers"
	"github.com/shynn12/biocad/pkg/logging"
)

type handler struct {
	logger  *logging.Logger
	service Service
	cfg     *config.Config
}

func NewHandler(logger *logging.Logger, service Service, cfg *config.Config) handlers.Handler {
	return &handler{
		logger:  logger,
		service: service,
		cfg:     cfg,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/biocad/element/{guid}", h.getOne).Methods(http.MethodGet)
	router.HandleFunc("/biocad/{page}", h.getTSV).Methods(http.MethodGet)
}

func (h *handler) getTSV(w http.ResponseWriter, r *http.Request) {
	v, ok := mux.Vars(r)["page"]
	if !ok {
		v = "1"
	}
	page, err := strconv.Atoi(v)
	if err != nil {
		h.logger.Error(err)
	}

	items := h.service.GetAllItems(context.Background(), page, h.cfg.PerPage)

	if err != nil {
		h.logger.Error(err)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		h.logger.Error(err)
	}
}

func (h *handler) getOne(w http.ResponseWriter, r *http.Request) {

	guid, ok := mux.Vars(r)["guid"]
	if !ok {
		h.logger.Error("error while parsing guid")
	}

	item := h.service.GetOneByGuid(context.Background(), guid)

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(item)
	if err != nil {
		h.logger.Error(err)
	}

}
