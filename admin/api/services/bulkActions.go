package services

import (
	"encoding/json"
	"io"
	"net/http"

	bulkactionsDto "github.com/tiqueteo/adminv2-mock-api/api/dto/bulkActions"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
)

func (sm *ServiceManager) PostBulkActionsValidate(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	var typeHolder bulkactionsDto.BulkActionType
	if err := json.Unmarshal(body, &typeHolder); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "could not determine action type"})
		return
	}

	switch typeHolder.Type {
	case "cancel_events":
		var req bulkactionsDto.CancelEventsRequest
		if err := json.Unmarshal(body, &req); err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		helpers.WriteJSON(w, http.StatusOK, generateCancelEventsResponse(req.SelectedDates, req.StartTime))

	case "enable_events":
		var req bulkactionsDto.EnableEventsRequest
		json.Unmarshal(body, &req)
		helpers.WriteJSON(w, http.StatusOK, bulkactionsDto.BasicValidationResponse{
			AffectedEvents: len(req.SelectedDates),
		})

	case "disable_events":
		var req bulkactionsDto.DisableEventsRequest
		json.Unmarshal(body, &req)
		helpers.WriteJSON(w, http.StatusOK, bulkactionsDto.BasicValidationResponse{
			AffectedEvents: len(req.SelectedDates),
		})

	case "remove_availability":
		var req bulkactionsDto.RemoveAvailabilityRequest
		json.Unmarshal(body, &req)
		helpers.WriteJSON(w, http.StatusOK, bulkactionsDto.BasicValidationResponse{
			AffectedEvents: len(req.SelectedDates),
		})

	case "move_availability":
		var req bulkactionsDto.MoveAvailabilityRequest
		json.Unmarshal(body, &req)
		helpers.WriteJSON(w, http.StatusOK, bulkactionsDto.BasicValidationResponse{
			AffectedEvents: len(req.SelectedDates),
		})

	case "event_max_groups":
		var req bulkactionsDto.EventMaxGroupsRequest
		json.Unmarshal(body, &req)
		helpers.WriteJSON(w, http.StatusOK, bulkactionsDto.BasicValidationResponse{
			AffectedEvents: len(req.SelectedDates),
		})

	default:
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "unsupported bulk action type"})
	}
}

func generateCancelEventsResponse(dates []string, startTime string) bulkactionsDto.CancelEventsValidationResponse {
	reservations := make([]bulkactionsDto.AffectedReservation, 0)
	for i, date := range dates {
		reservations = append(reservations, bulkactionsDto.AffectedReservation{
			ReservationId:      1000 + i,
			EventStartDatetime: date + "T" + startTime,
		})
	}

	return bulkactionsDto.CancelEventsValidationResponse{
		AffectedEvents:       len(dates),
		AffectedReservations: reservations,
	}
}

func (sm *ServiceManager) PostBulkActionsExecute(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	var typeHolder bulkactionsDto.BulkActionType
	if err := json.Unmarshal(body, &typeHolder); err != nil || typeHolder.Type == "" {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "missing or invalid type field"})
		return
	}

	switch typeHolder.Type {
	case "enable_events":
		var req bulkactionsDto.EnableEventsRequest
		json.Unmarshal(body, &req)

	case "disable_events":
		var req bulkactionsDto.DisableEventsRequest
		json.Unmarshal(body, &req)

	case "cancel_events":
		var req bulkactionsDto.CancelEventsRequest
		json.Unmarshal(body, &req)

	case "remove_availability":
		var req bulkactionsDto.RemoveAvailabilityRequest
		json.Unmarshal(body, &req)

	case "move_availability":
		var req bulkactionsDto.MoveAvailabilityRequest
		json.Unmarshal(body, &req)

	case "event_max_groups":
		var req bulkactionsDto.EventMaxGroupsRequest
		json.Unmarshal(body, &req)

	default:
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "unsupported bulk action type"})
		return
	}

	w.WriteHeader(http.StatusOK)
}
