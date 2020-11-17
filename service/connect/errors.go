// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeContactFlowNotPublishedException for service response error code
	// "ContactFlowNotPublishedException".
	//
	// The contact flow has not been published.
	ErrCodeContactFlowNotPublishedException = "ContactFlowNotPublishedException"

	// ErrCodeContactNotFoundException for service response error code
	// "ContactNotFoundException".
	//
	// The contact with the specified ID is not active or does not exist.
	ErrCodeContactNotFoundException = "ContactNotFoundException"

	// ErrCodeDestinationNotAllowedException for service response error code
	// "DestinationNotAllowedException".
	//
	// Outbound calls to the destination number are not allowed.
	ErrCodeDestinationNotAllowedException = "DestinationNotAllowedException"

	// ErrCodeDuplicateResourceException for service response error code
	// "DuplicateResourceException".
	//
	// A resource with the specified name already exists.
	ErrCodeDuplicateResourceException = "DuplicateResourceException"

	// ErrCodeInternalServiceException for service response error code
	// "InternalServiceException".
	//
	// Request processing failed due to an error or failure with the service.
	ErrCodeInternalServiceException = "InternalServiceException"

	// ErrCodeInvalidContactFlowException for service response error code
	// "InvalidContactFlowException".
	//
	// The contact flow is not valid.
	ErrCodeInvalidContactFlowException = "InvalidContactFlowException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// One or more of the specified parameters are not valid.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The request is not valid.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The allowed limit for the resource has been exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeOutboundContactNotPermittedException for service response error code
	// "OutboundContactNotPermittedException".
	//
	// The contact is not permitted.
	ErrCodeOutboundContactNotPermittedException = "OutboundContactNotPermittedException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// That resource is already in use. Please try another.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource was not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The throttling limit has been exceeded.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeUserNotFoundException for service response error code
	// "UserNotFoundException".
	//
	// No user with the specified credentials was found in the Amazon Connect instance.
	ErrCodeUserNotFoundException = "UserNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ContactFlowNotPublishedException":     newErrorContactFlowNotPublishedException,
	"ContactNotFoundException":             newErrorContactNotFoundException,
	"DestinationNotAllowedException":       newErrorDestinationNotAllowedException,
	"DuplicateResourceException":           newErrorDuplicateResourceException,
	"InternalServiceException":             newErrorInternalServiceException,
	"InvalidContactFlowException":          newErrorInvalidContactFlowException,
	"InvalidParameterException":            newErrorInvalidParameterException,
	"InvalidRequestException":              newErrorInvalidRequestException,
	"LimitExceededException":               newErrorLimitExceededException,
	"OutboundContactNotPermittedException": newErrorOutboundContactNotPermittedException,
	"ResourceInUseException":               newErrorResourceInUseException,
	"ResourceNotFoundException":            newErrorResourceNotFoundException,
	"ThrottlingException":                  newErrorThrottlingException,
	"UserNotFoundException":                newErrorUserNotFoundException,
}
