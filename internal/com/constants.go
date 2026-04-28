package com

type CtxKey string
type ErrorType string
type ErrorCode string
type TypeIDPrefix string
type ResourceLabel string

const (
	CtxKeyAccountID CtxKey = "ctxKeyAccountID"
	CtxKeyRequestID CtxKey = "ctxKeyRequestID"
	CtxKeyUserID    CtxKey = "ctxKeyUserID"
)

const (
	ApiError            ErrorType = "api_error"
	InvalidRequestError ErrorType = "invalid_request_error"
)

const (
	ApiKeyExpired          ErrorCode = "api_key_expired"
	CountryCodeInvalid     ErrorCode = "country_code_invalid"
	CurrencyCodeInvalid    ErrorCode = "currency_code_invalid"
	FieldExpansionMaxDepth ErrorCode = "field_expansion_max_depth"
	FieldUnexpandable      ErrorCode = "field_unexpandable"
	IdentifierInvalid      ErrorCode = "identifier_invalid"
	ParameterInvalid       ErrorCode = "parameter_invalid"
	ParameterMissing       ErrorCode = "parameter_missing"
	ResourceNotFound       ErrorCode = "resource_not_found"
)

const (
	TypeIDPrefixAccount                    TypeIDPrefix = "acc"
	TypeIDPrefixApiKey                     TypeIDPrefix = "api_key"
	TypeIDPrefixGroup                      TypeIDPrefix = "grp"
	TypeIDPrefixInventory                  TypeIDPrefix = "inv"
	TypeIDPrefixInventoryMovement          TypeIDPrefix = "im"
	TypeIDPrefixInventoryMovementReference TypeIDPrefix = "im_ref"
	TypeIDPrefixItem                       TypeIDPrefix = "itm"
	TypeIDPrefixItemIdentifier             TypeIDPrefix = "itm_idr"
	TypeIDPrefixItemVariantAttribute       TypeIDPrefix = "iva"
	TypeIDPrefixItemVariantAttributeOption TypeIDPrefix = "iva_opt"
	TypeIDPrefixRequest                    TypeIDPrefix = "req"
	TypeIDPrefixUser                       TypeIDPrefix = "usr"
)

const (
	ResourceAccount                    ResourceLabel = "account"
	ResourceApiKey                     ResourceLabel = "api key"
	ResourceGroup                      ResourceLabel = "group"
	ResourceInventory                  ResourceLabel = "inventory"
	ResourceInventoryMovement          ResourceLabel = "inventory movement"
	ResourceInventoryMovementReference ResourceLabel = "inventory movement reference"
	ResourceItem                       ResourceLabel = "item"
	ResourceItemIdentifier             ResourceLabel = "item identifier"
	ResourceItemVariantAttribute       ResourceLabel = "item variant attribute"
	ResourceItemVariantAttributeOption ResourceLabel = "item variant attribute option"
	ResourceRequest                    ResourceLabel = "request"
	ResourceUser                       ResourceLabel = "user"
)
