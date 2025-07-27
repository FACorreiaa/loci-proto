package common

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	c "github.com/FACorreiaa/loci-proto/modules/common/generated"
)

// Helper functions for common types and operations

// NewResponse creates a new Response with the given parameters
func NewResponse(success bool, message string) *c.Response {
	return &c.Response{
		Success: success,
		Message: message,
	}
}

// NewSuccessResponse creates a successful response
func NewSuccessResponse(message string) *c.Response {
	return NewResponse(true, message)
}

// NewErrorResponse creates an error response with error details
func NewErrorResponse(message string, errorDetails *c.ErrorDetails) *c.Response {
	return &c.Response{
		Success: false,
		Message: message,
		Error:   errorDetails,
	}
}

// NewErrorDetails creates error details with the given code and message
func NewErrorDetails(code, message string) *c.ErrorDetails {
	return &c.ErrorDetails{
		Code:    code,
		Message: message,
	}
}

// NewFieldError creates a field-specific validation error
func NewFieldError(field, message, code string) *c.FieldError {
	return &c.FieldError{
		Field:   field,
		Message: message,
		Code:    code,
	}
}

// NewPaginationRequest creates a pagination request with page and page size
func NewPaginationRequest(page, pageSize int32) *c.PaginationRequest {
	return &c.PaginationRequest{
		Page:     page,
		PageSize: pageSize,
	}
}

// NewPaginationResponse creates pagination response metadata
func NewPaginationResponse(currentPage, pageSize, totalItems, totalPages int32, hasNext, hasPrev bool) *c.PaginationResponse {
	return &c.PaginationResponse{
		CurrentPage:     currentPage,
		PageSize:        pageSize,
		TotalItems:      totalItems,
		TotalPages:      totalPages,
		HasNextPage:     hasNext,
		HasPreviousPage: hasPrev,
	}
}

// NewCoordinates creates a new Coordinates object
func NewCoordinates(lat, lng float64) *c.Coordinates {
	return &c.Coordinates{
		Latitude:  lat,
		Longitude: lng,
	}
}

// NewCoordinatesWithAltitude creates coordinates with altitude and accuracy
func NewCoordinatesWithAltitude(lat, lng, altitude, accuracy float64) *c.Coordinates {
	return &c.Coordinates{
		Latitude:  lat,
		Longitude: lng,
		Altitude:  altitude,
		Accuracy:  accuracy,
	}
}

// NewGeoBounds creates a geographic bounding box
func NewGeoBounds(swLat, swLng, neLat, neLng float64) *c.GeoBounds {
	return &c.GeoBounds{
		Southwest: NewCoordinates(swLat, swLng),
		Northeast: NewCoordinates(neLat, neLng),
	}
}

// NewAddress creates a new Address object
func NewAddress(street, city, state, postalCode, country, countryCode string) *c.Address {
	return &c.Address{
		Street:     street,
		City:       city,
		State:      state,
		PostalCode: postalCode,
		Country:    country,
		CountryCode: countryCode,
	}
}

// NewContactInfo creates contact information
func NewContactInfo(phone, email, website string) *c.ContactInfo {
	return &c.ContactInfo{
		Phone:   phone,
		Email:   email,
		Website: website,
	}
}

// NewSocialMedia creates a social media link
func NewSocialMedia(platform, url, handle string) *c.SocialMedia {
	return &c.SocialMedia{
		Platform: platform,
		Url:      url,
		Handle:   handle,
	}
}

// NewTimeSlot creates a time slot with open and close times
func NewTimeSlot(openTime, closeTime string) *c.TimeSlot {
	return &c.TimeSlot{
		OpenTime:  openTime,
		CloseTime: closeTime,
	}
}

// NewDaySchedule creates a day schedule
func NewDaySchedule(day c.DayOfWeek, timeSlots []*c.TimeSlot, isClosed bool) *c.DaySchedule {
	return &c.DaySchedule{
		Day:       day,
		TimeSlots: timeSlots,
		IsClosed:  isClosed,
	}
}

// NewRating creates a rating with average and count
func NewRating(average float64, count int32) *c.Rating {
	return &c.Rating{
		Average: average,
		Count:   count,
	}
}

// NewRatingBreakdown creates a rating breakdown
func NewRatingBreakdown(five, four, three, two, one int32) *c.RatingBreakdown {
	return &c.RatingBreakdown{
		FiveStar:  five,
		FourStar:  four,
		ThreeStar: three,
		TwoStar:   two,
		OneStar:   one,
	}
}

// NewPhoto creates a photo object
func NewPhoto(id, url, thumbnailUrl, caption, altText string) *c.Photo {
	return &c.Photo{
		Id:           id,
		Url:          url,
		ThumbnailUrl: thumbnailUrl,
		Caption:      caption,
		AltText:      altText,
	}
}

// NewVideo creates a video object
func NewVideo(id, url, thumbnailUrl, title, description, provider string, duration int32) *c.Video {
	return &c.Video{
		Id:              id,
		Url:             url,
		ThumbnailUrl:    thumbnailUrl,
		Title:           title,
		Description:     description,
		DurationSeconds: duration,
		Provider:        provider,
	}
}

// NewSortOptions creates sort options
func NewSortOptions(field string, direction c.SortDirection) *c.SortOptions {
	return &c.SortOptions{
		Field:     field,
		Direction: direction,
	}
}

// NewFilterOptions creates filter options
func NewFilterOptions() *c.FilterOptions {
	return &c.FilterOptions{
		CustomFilters: make(map[string]string),
	}
}

// NewLocalizedString creates a localized string
func NewLocalizedString(languageCode, text string) *c.LocalizedString {
	return &c.LocalizedString{
		LanguageCode: languageCode,
		Text:         text,
	}
}

// NewMultilingualText creates multilingual text
func NewMultilingualText(defaultLanguage string, translations []*c.LocalizedString) *c.MultilingualText {
	return &c.MultilingualText{
		DefaultLanguage: defaultLanguage,
		Translations:    translations,
	}
}

// NewAuditInfo creates audit information with current timestamp
func NewAuditInfo(createdBy, updatedBy string) *c.AuditInfo {
	now := timestamppb.New(time.Now())
	return &c.AuditInfo{
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: createdBy,
		UpdatedBy: updatedBy,
		Version:   1,
	}
}

// UpdateAuditInfo updates the audit information for an update operation
func UpdateAuditInfo(auditInfo *c.AuditInfo, updatedBy string) *c.AuditInfo {
	if auditInfo == nil {
		return NewAuditInfo("", updatedBy)
	}
	
	auditInfo.UpdatedAt = timestamppb.New(time.Now())
	auditInfo.UpdatedBy = updatedBy
	auditInfo.Version++
	
	return auditInfo
}

// NewHealthCheckRequest creates a health check request
func NewHealthCheckRequest(service string) *c.HealthCheckRequest {
	return &c.HealthCheckRequest{
		Service: service,
		Request: NewBaseRequest("", ""),
	}
}

// NewHealthCheckResponse creates a health check response
func NewHealthCheckResponse(status, version string) *c.HealthCheckResponse {
	return &c.HealthCheckResponse{
		Status:    status,
		Version:   version,
		Timestamp: timestamppb.New(time.Now()),
		Components: make(map[string]*c.ComponentHealth),
		Response:  NewBaseResponse("", ""),
	}
}

// NewComponentHealth creates component health information
func NewComponentHealth(status, message string) *c.ComponentHealth {
	return &c.ComponentHealth{
		Status:  status,
		Message: message,
		Details: make(map[string]string),
	}
}

// NewFeatureFlag creates a feature flag
func NewFeatureFlag(name string, enabled bool, description string) *c.FeatureFlag {
	return &c.FeatureFlag{
		Name:        name,
		Enabled:     enabled,
		Description: description,
		Parameters:  make(map[string]string),
	}
}

// NewApiVersion creates API version information
func NewApiVersion(version string, deprecated bool) *c.ApiVersion {
	return &c.ApiVersion{
		Version:    version,
		Deprecated: deprecated,
	}
}

// NewRateLimitInfo creates rate limit information
func NewRateLimitInfo(remaining, limit int32, resetTime time.Time) *c.RateLimitInfo {
	return &c.RateLimitInfo{
		RequestsRemaining: remaining,
		RequestsLimit:     limit,
		ResetTime:         timestamppb.New(resetTime),
	}
}

// NewBaseRequest creates a base request
func NewBaseRequest(downstream, requestId string) *c.BaseRequest {
	return &c.BaseRequest{
		Downstream: downstream,
		RequestId:  requestId,
	}
}

// NewBaseResponse creates a base response
func NewBaseResponse(upstream, requestId string) *c.BaseResponse {
	return &c.BaseResponse{
		Upstream:  upstream,
		RequestId: requestId,
		Status:    "success",
	}
}

// Validation helpers

// ValidateCoordinates checks if coordinates are within valid ranges
func ValidateCoordinates(coords *c.Coordinates) bool {
	if coords == nil {
		return false
	}
	return coords.Latitude >= -90 && coords.Latitude <= 90 &&
		coords.Longitude >= -180 && coords.Longitude <= 180
}

// ValidatePagination checks if pagination parameters are valid
func ValidatePagination(pagination *c.PaginationRequest) bool {
	if pagination == nil {
		return true // Optional parameter
	}
	return pagination.Page > 0 && pagination.PageSize > 0 && pagination.PageSize <= 100
}

// ValidateRating checks if rating is within valid range
func ValidateRating(rating *c.Rating) bool {
	if rating == nil {
		return true // Optional parameter
	}
	return rating.Average >= 0 && rating.Average <= 5 && rating.Count >= 0
}