package com

type CtxKey string
type ErrorType string
type ErrorCode string
type IDPrefix string
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
	IDPrefixAccount                    IDPrefix = "acc"
	IDPrefixApiKey                     IDPrefix = "api_key"
	IDPrefixGroup                      IDPrefix = "grp"
	IDPrefixInventory                  IDPrefix = "inv"
	IDPrefixInventoryMovement          IDPrefix = "imv"
	IDPrefixInventoryMovementReference IDPrefix = "imv_ref"
	IDPrefixItem                       IDPrefix = "itm"
	IDPrefixItemIdentifier             IDPrefix = "itm_idf"
	IDPrefixItemVariantAttribute       IDPrefix = "iva"
	IDPrefixItemVariantAttributeOption IDPrefix = "iva_opt"
	IDPrefixRequest                    IDPrefix = "req"
	IDPrefixUser                       IDPrefix = "usr"
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
