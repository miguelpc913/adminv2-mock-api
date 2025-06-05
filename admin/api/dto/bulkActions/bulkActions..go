package bulkactionsDto

type AffectedReservation struct {
	ReservationId      int    `json:"reservationId"`
	EventStartDatetime string `json:"eventStartDatetime"`
}

type CancelEventsValidationResponse struct {
	AffectedEvents       int                   `json:"affectedEvents"`
	AffectedReservations []AffectedReservation `json:"affectedReservations"`
}

type BasicValidationResponse struct {
	AffectedEvents int `json:"affectedEvents"`
}

type BulkActionType struct {
	Type string `json:"type"`
}

type CommonBulkFields struct {
	VenueCapacityId int      `json:"venueCapacityId"`
	SelectedDates   []string `json:"selectedDates"`
	StartDate       string   `json:"startDate"`
	EndDate         string   `json:"endDate"`
	WeekDays        []int    `json:"weekDays"`
	StartTime       string   `json:"startTime"`
	EndTime         string   `json:"endTime"`
}

type EnableEventsRequest struct {
	CommonBulkFields
	Type string `json:"type"`
}

type DisableEventsRequest struct {
	CommonBulkFields
	Type                      string `json:"type"`
	TargetAvailabilityGroupId int    `json:"targetAvailabilityGroupId"`
	Quantity                  int    `json:"quantity"`
}

type CancelEventsRequest struct {
	VenueCapacityId    int      `json:"venueCapacityId"`
	Type               string   `json:"type"`
	SelectedDates      []string `json:"selectedDates"`
	StartTime          string   `json:"startTime"`
	EndTime            string   `json:"endTime"`
	SendEmail          bool     `json:"sendEmail"`
	CancellationReason string   `json:"cancellationReason"`
	SendRescheduleLink bool     `json:"sendRescheduleLink"`
}

type RemoveAvailabilityRequest struct {
	CommonBulkFields
	Type string `json:"type"`
}

type MoveAvailabilityRequest struct {
	CommonBulkFields
	Type                      string `json:"type"`
	SourceAvailabilityGroupId int    `json:"sourceAvailabilityGroupId"`
	TargetAvailabilityGroupId int    `json:"targetAvailabilityGroupId"`
	Quantity                  int    `json:"quantity"`
}

type EventMaxGroupsRequest struct {
	CommonBulkFields
	Type            string `json:"type"`
	MaxReservations int    `json:"maxReservations"`
}
