// Code generated by go-swagger; DO NOT EDIT.

package staking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new staking API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for staking API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetStakingDelegatorsDelegatorAddrDelegations(params *GetStakingDelegatorsDelegatorAddrDelegationsParams) (*GetStakingDelegatorsDelegatorAddrDelegationsOK, error)

	GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddr(params *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrParams) (*GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK, error)

	GetStakingDelegatorsDelegatorAddrTxs(params *GetStakingDelegatorsDelegatorAddrTxsParams) (*GetStakingDelegatorsDelegatorAddrTxsOK, *GetStakingDelegatorsDelegatorAddrTxsNoContent, error)

	GetStakingDelegatorsDelegatorAddrUnbondingDelegations(params *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) (*GetStakingDelegatorsDelegatorAddrUnbondingDelegationsOK, error)

	GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddr(params *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrParams) (*GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK, error)

	GetStakingDelegatorsDelegatorAddrValidators(params *GetStakingDelegatorsDelegatorAddrValidatorsParams) (*GetStakingDelegatorsDelegatorAddrValidatorsOK, error)

	GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddr(params *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams) (*GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK, error)

	GetStakingParameters(params *GetStakingParametersParams) (*GetStakingParametersOK, error)

	GetStakingPool(params *GetStakingPoolParams) (*GetStakingPoolOK, error)

	GetStakingRedelegations(params *GetStakingRedelegationsParams) (*GetStakingRedelegationsOK, error)

	GetStakingValidators(params *GetStakingValidatorsParams) (*GetStakingValidatorsOK, error)

	GetStakingValidatorsValidatorAddr(params *GetStakingValidatorsValidatorAddrParams) (*GetStakingValidatorsValidatorAddrOK, error)

	GetStakingValidatorsValidatorAddrDelegations(params *GetStakingValidatorsValidatorAddrDelegationsParams) (*GetStakingValidatorsValidatorAddrDelegationsOK, error)

	GetStakingValidatorsValidatorAddrUnbondingDelegations(params *GetStakingValidatorsValidatorAddrUnbondingDelegationsParams) (*GetStakingValidatorsValidatorAddrUnbondingDelegationsOK, error)

	PostStakingDelegatorsDelegatorAddrDelegations(params *PostStakingDelegatorsDelegatorAddrDelegationsParams) (*PostStakingDelegatorsDelegatorAddrDelegationsOK, error)

	PostStakingDelegatorsDelegatorAddrUnbondingDelegations(params *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) (*PostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetStakingDelegatorsDelegatorAddrDelegations gets all delegations from a delegator
*/
func (a *Client) GetStakingDelegatorsDelegatorAddrDelegations(params *GetStakingDelegatorsDelegatorAddrDelegationsParams) (*GetStakingDelegatorsDelegatorAddrDelegationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStakingDelegatorsDelegatorAddrDelegationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStakingDelegatorsDelegatorAddrDelegations",
		Method:             "GET",
		PathPattern:        "/staking/delegators/{delegatorAddr}/delegations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStakingDelegatorsDelegatorAddrDelegationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStakingDelegatorsDelegatorAddrDelegationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStakingDelegatorsDelegatorAddrDelegations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddr queries the current delegation between a delegator and a validator
*/
func (a *Client) GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddr(params *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrParams) (*GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddr",
		Method:             "GET",
		PathPattern:        "/staking/delegators/{delegatorAddr}/delegations/{validatorAddr}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddr: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStakingDelegatorsDelegatorAddrTxs gets all staking txs i e msgs from a delegator
*/
func (a *Client) GetStakingDelegatorsDelegatorAddrTxs(params *GetStakingDelegatorsDelegatorAddrTxsParams) (*GetStakingDelegatorsDelegatorAddrTxsOK, *GetStakingDelegatorsDelegatorAddrTxsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStakingDelegatorsDelegatorAddrTxsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStakingDelegatorsDelegatorAddrTxs",
		Method:             "GET",
		PathPattern:        "/staking/delegators/{delegatorAddr}/txs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStakingDelegatorsDelegatorAddrTxsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetStakingDelegatorsDelegatorAddrTxsOK:
		return value, nil, nil
	case *GetStakingDelegatorsDelegatorAddrTxsNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for staking: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStakingDelegatorsDelegatorAddrUnbondingDelegations gets all unbonding delegations from a delegator
*/
func (a *Client) GetStakingDelegatorsDelegatorAddrUnbondingDelegations(params *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) (*GetStakingDelegatorsDelegatorAddrUnbondingDelegationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStakingDelegatorsDelegatorAddrUnbondingDelegations",
		Method:             "GET",
		PathPattern:        "/staking/delegators/{delegatorAddr}/unbonding_delegations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStakingDelegatorsDelegatorAddrUnbondingDelegationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStakingDelegatorsDelegatorAddrUnbondingDelegationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStakingDelegatorsDelegatorAddrUnbondingDelegations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddr queries all unbonding delegations between a delegator and a validator
*/
func (a *Client) GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddr(params *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrParams) (*GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddr",
		Method:             "GET",
		PathPattern:        "/staking/delegators/{delegatorAddr}/unbonding_delegations/{validatorAddr}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddr: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStakingDelegatorsDelegatorAddrValidators queries all validators that a delegator is bonded to
*/
func (a *Client) GetStakingDelegatorsDelegatorAddrValidators(params *GetStakingDelegatorsDelegatorAddrValidatorsParams) (*GetStakingDelegatorsDelegatorAddrValidatorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStakingDelegatorsDelegatorAddrValidatorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStakingDelegatorsDelegatorAddrValidators",
		Method:             "GET",
		PathPattern:        "/staking/delegators/{delegatorAddr}/validators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStakingDelegatorsDelegatorAddrValidatorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStakingDelegatorsDelegatorAddrValidatorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStakingDelegatorsDelegatorAddrValidators: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddr queries a validator that a delegator is bonded to
*/
func (a *Client) GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddr(params *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams) (*GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddr",
		Method:             "GET",
		PathPattern:        "/staking/delegators/{delegatorAddr}/validators/{validatorAddr}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddr: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStakingParameters gets the current staking parameter values
*/
func (a *Client) GetStakingParameters(params *GetStakingParametersParams) (*GetStakingParametersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStakingParametersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStakingParameters",
		Method:             "GET",
		PathPattern:        "/staking/parameters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStakingParametersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStakingParametersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStakingParameters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStakingPool gets the current state of the staking pool
*/
func (a *Client) GetStakingPool(params *GetStakingPoolParams) (*GetStakingPoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStakingPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStakingPool",
		Method:             "GET",
		PathPattern:        "/staking/pool",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStakingPoolReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStakingPoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStakingPool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStakingRedelegations gets all redelegations filter by query params
*/
func (a *Client) GetStakingRedelegations(params *GetStakingRedelegationsParams) (*GetStakingRedelegationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStakingRedelegationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStakingRedelegations",
		Method:             "GET",
		PathPattern:        "/staking/redelegations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStakingRedelegationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStakingRedelegationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStakingRedelegations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStakingValidators gets all validator candidates by default it returns only the bonded validators
*/
func (a *Client) GetStakingValidators(params *GetStakingValidatorsParams) (*GetStakingValidatorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStakingValidatorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStakingValidators",
		Method:             "GET",
		PathPattern:        "/staking/validators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStakingValidatorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStakingValidatorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStakingValidators: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStakingValidatorsValidatorAddr queries the information from a single validator
*/
func (a *Client) GetStakingValidatorsValidatorAddr(params *GetStakingValidatorsValidatorAddrParams) (*GetStakingValidatorsValidatorAddrOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStakingValidatorsValidatorAddrParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStakingValidatorsValidatorAddr",
		Method:             "GET",
		PathPattern:        "/staking/validators/{validatorAddr}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStakingValidatorsValidatorAddrReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStakingValidatorsValidatorAddrOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStakingValidatorsValidatorAddr: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStakingValidatorsValidatorAddrDelegations gets all delegations from a validator
*/
func (a *Client) GetStakingValidatorsValidatorAddrDelegations(params *GetStakingValidatorsValidatorAddrDelegationsParams) (*GetStakingValidatorsValidatorAddrDelegationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStakingValidatorsValidatorAddrDelegationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStakingValidatorsValidatorAddrDelegations",
		Method:             "GET",
		PathPattern:        "/staking/validators/{validatorAddr}/delegations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStakingValidatorsValidatorAddrDelegationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStakingValidatorsValidatorAddrDelegationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStakingValidatorsValidatorAddrDelegations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStakingValidatorsValidatorAddrUnbondingDelegations gets all unbonding delegations from a validator
*/
func (a *Client) GetStakingValidatorsValidatorAddrUnbondingDelegations(params *GetStakingValidatorsValidatorAddrUnbondingDelegationsParams) (*GetStakingValidatorsValidatorAddrUnbondingDelegationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStakingValidatorsValidatorAddrUnbondingDelegationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStakingValidatorsValidatorAddrUnbondingDelegations",
		Method:             "GET",
		PathPattern:        "/staking/validators/{validatorAddr}/unbonding_delegations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStakingValidatorsValidatorAddrUnbondingDelegationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStakingValidatorsValidatorAddrUnbondingDelegationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStakingValidatorsValidatorAddrUnbondingDelegations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostStakingDelegatorsDelegatorAddrDelegations submits delegation
*/
func (a *Client) PostStakingDelegatorsDelegatorAddrDelegations(params *PostStakingDelegatorsDelegatorAddrDelegationsParams) (*PostStakingDelegatorsDelegatorAddrDelegationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStakingDelegatorsDelegatorAddrDelegationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostStakingDelegatorsDelegatorAddrDelegations",
		Method:             "POST",
		PathPattern:        "/staking/delegators/{delegatorAddr}/delegations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostStakingDelegatorsDelegatorAddrDelegationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostStakingDelegatorsDelegatorAddrDelegationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostStakingDelegatorsDelegatorAddrDelegations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostStakingDelegatorsDelegatorAddrUnbondingDelegations submits an unbonding delegation
*/
func (a *Client) PostStakingDelegatorsDelegatorAddrUnbondingDelegations(params *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) (*PostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStakingDelegatorsDelegatorAddrUnbondingDelegationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostStakingDelegatorsDelegatorAddrUnbondingDelegations",
		Method:             "POST",
		PathPattern:        "/staking/delegators/{delegatorAddr}/unbonding_delegations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostStakingDelegatorsDelegatorAddrUnbondingDelegationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostStakingDelegatorsDelegatorAddrUnbondingDelegations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
