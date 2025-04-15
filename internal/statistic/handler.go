package statistic

import (
	"http_server/configs"
	"http_server/pkg/middleware"
	"http_server/pkg/response"
	"net/http"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type HandlerDeps struct {
	Config              *configs.Config
	StatisticRepository *Repository
}
type Handler struct {
	StatisticRepository *Repository
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		StatisticRepository: deps.StatisticRepository,
	}

	router.Handle("GET /statistic", middleware.IsAuthed(handler.GetStatistic(), deps.Config))
}

func (handler *Handler) GetStatistic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid from parameter", http.StatusBadRequest)
		}

		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid to parameter", http.StatusBadRequest)
		}

		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Invalid by parameter", http.StatusBadRequest)
		}

		statistic := handler.StatisticRepository.GetStatistic(by, from, to)
		
		response.Json(w, statistic, http.StatusOK)
	}
}
