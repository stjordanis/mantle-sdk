// Code generated by go-swagger; DO NOT EDIT.

package oracle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new oracle API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for oracle API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetOracleDenomsActives(params *GetOracleDenomsActivesParams) (*GetOracleDenomsActivesOK, error)

	GetOracleDenomsDenomExchangeRate(params *GetOracleDenomsDenomExchangeRateParams) (*GetOracleDenomsDenomExchangeRateOK, error)

	GetOracleDenomsDenomPrevotes(params *GetOracleDenomsDenomPrevotesParams) (*GetOracleDenomsDenomPrevotesOK, error)

	GetOracleDenomsDenomPrevotesValidator(params *GetOracleDenomsDenomPrevotesValidatorParams) (*GetOracleDenomsDenomPrevotesValidatorOK, error)

	GetOracleDenomsDenomVotes(params *GetOracleDenomsDenomVotesParams) (*GetOracleDenomsDenomVotesOK, error)

	GetOracleDenomsDenomVotesValidator(params *GetOracleDenomsDenomVotesValidatorParams) (*GetOracleDenomsDenomVotesValidatorOK, error)

	GetOracleDenomsExchangeRates(params *GetOracleDenomsExchangeRatesParams) (*GetOracleDenomsExchangeRatesOK, error)

	GetOracleParameters(params *GetOracleParametersParams) (*GetOracleParametersOK, error)

	GetOracleVotersValidatorAggregatePrevote(params *GetOracleVotersValidatorAggregatePrevoteParams) (*GetOracleVotersValidatorAggregatePrevoteOK, error)

	GetOracleVotersValidatorAggregateVote(params *GetOracleVotersValidatorAggregateVoteParams) (*GetOracleVotersValidatorAggregateVoteOK, error)

	GetOracleVotersValidatorFeeder(params *GetOracleVotersValidatorFeederParams) (*GetOracleVotersValidatorFeederOK, error)

	GetOracleVotersValidatorMiss(params *GetOracleVotersValidatorMissParams) (*GetOracleVotersValidatorMissOK, error)

	GetOracleVotersValidatorPrevotes(params *GetOracleVotersValidatorPrevotesParams) (*GetOracleVotersValidatorPrevotesOK, error)

	GetOracleVotersValidatorVotes(params *GetOracleVotersValidatorVotesParams) (*GetOracleVotersValidatorVotesOK, error)

	PostOracleDenomsDenomPrevotes(params *PostOracleDenomsDenomPrevotesParams) (*PostOracleDenomsDenomPrevotesOK, error)

	PostOracleDenomsDenomVotes(params *PostOracleDenomsDenomVotesParams) (*PostOracleDenomsDenomVotesOK, error)

	PostOracleVotersValidatorAggregatePrevote(params *PostOracleVotersValidatorAggregatePrevoteParams) (*PostOracleVotersValidatorAggregatePrevoteOK, error)

	PostOracleVotersValidatorAggregateVote(params *PostOracleVotersValidatorAggregateVoteParams) (*PostOracleVotersValidatorAggregateVoteOK, error)

	PostOracleVotersValidatorFeeder(params *PostOracleVotersValidatorFeederParams) (*PostOracleVotersValidatorFeederOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetOracleDenomsActives gets all activated denoms
*/
func (a *Client) GetOracleDenomsActives(params *GetOracleDenomsActivesParams) (*GetOracleDenomsActivesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleDenomsActivesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOracleDenomsActives",
		Method:             "GET",
		PathPattern:        "/oracle/denoms/actives",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOracleDenomsActivesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleDenomsActivesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleDenomsActives: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleDenomsDenomExchangeRate gets the current effective exchange rate in luna for the asset
*/
func (a *Client) GetOracleDenomsDenomExchangeRate(params *GetOracleDenomsDenomExchangeRateParams) (*GetOracleDenomsDenomExchangeRateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleDenomsDenomExchangeRateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOracleDenomsDenomExchangeRate",
		Method:             "GET",
		PathPattern:        "/oracle/denoms/{denom}/exchange_rate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOracleDenomsDenomExchangeRateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleDenomsDenomExchangeRateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleDenomsDenomExchangeRate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleDenomsDenomPrevotes requests to get the currently outstanding exchange rate oracle prevote
*/
func (a *Client) GetOracleDenomsDenomPrevotes(params *GetOracleDenomsDenomPrevotesParams) (*GetOracleDenomsDenomPrevotesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleDenomsDenomPrevotesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOracleDenomsDenomPrevotes",
		Method:             "GET",
		PathPattern:        "/oracle/denoms/{denom}/prevotes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOracleDenomsDenomPrevotesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleDenomsDenomPrevotesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleDenomsDenomPrevotes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleDenomsDenomPrevotesValidator requests to get the currently outstanding exchange rate oracle prevote
*/
func (a *Client) GetOracleDenomsDenomPrevotesValidator(params *GetOracleDenomsDenomPrevotesValidatorParams) (*GetOracleDenomsDenomPrevotesValidatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleDenomsDenomPrevotesValidatorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOracleDenomsDenomPrevotesValidator",
		Method:             "GET",
		PathPattern:        "/oracle/denoms/{denom}/prevotes/{validator}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOracleDenomsDenomPrevotesValidatorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleDenomsDenomPrevotesValidatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleDenomsDenomPrevotesValidator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleDenomsDenomVotes requests to get the currently unelected outstanding exchange rate oracle vote
*/
func (a *Client) GetOracleDenomsDenomVotes(params *GetOracleDenomsDenomVotesParams) (*GetOracleDenomsDenomVotesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleDenomsDenomVotesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOracleDenomsDenomVotes",
		Method:             "GET",
		PathPattern:        "/oracle/denoms/{denom}/votes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOracleDenomsDenomVotesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleDenomsDenomVotesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleDenomsDenomVotes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleDenomsDenomVotesValidator requests to get the currently unelected outstanding exchange rate oracle vote
*/
func (a *Client) GetOracleDenomsDenomVotesValidator(params *GetOracleDenomsDenomVotesValidatorParams) (*GetOracleDenomsDenomVotesValidatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleDenomsDenomVotesValidatorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOracleDenomsDenomVotesValidator",
		Method:             "GET",
		PathPattern:        "/oracle/denoms/{denom}/votes/{validator}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOracleDenomsDenomVotesValidatorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleDenomsDenomVotesValidatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleDenomsDenomVotesValidator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleDenomsExchangeRates gets all activated exchange rates
*/
func (a *Client) GetOracleDenomsExchangeRates(params *GetOracleDenomsExchangeRatesParams) (*GetOracleDenomsExchangeRatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleDenomsExchangeRatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOracleDenomsExchangeRates",
		Method:             "GET",
		PathPattern:        "/oracle/denoms/exchange_rates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOracleDenomsExchangeRatesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleDenomsExchangeRatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleDenomsExchangeRates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleParameters gets oracle params
*/
func (a *Client) GetOracleParameters(params *GetOracleParametersParams) (*GetOracleParametersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleParametersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOracleParameters",
		Method:             "GET",
		PathPattern:        "/oracle/parameters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOracleParametersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleParametersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleParameters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleVotersValidatorAggregatePrevote gets the currently outstanding aggregate exchange rate oracle prevote of a validator
*/
func (a *Client) GetOracleVotersValidatorAggregatePrevote(params *GetOracleVotersValidatorAggregatePrevoteParams) (*GetOracleVotersValidatorAggregatePrevoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleVotersValidatorAggregatePrevoteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOracleVotersValidatorAggregatePrevote",
		Method:             "GET",
		PathPattern:        "/oracle/voters/{validator}/aggregate_prevote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOracleVotersValidatorAggregatePrevoteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleVotersValidatorAggregatePrevoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleVotersValidatorAggregatePrevote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleVotersValidatorAggregateVote gets the currently outstanding aggregate exchange rate oracle vote of a validator
*/
func (a *Client) GetOracleVotersValidatorAggregateVote(params *GetOracleVotersValidatorAggregateVoteParams) (*GetOracleVotersValidatorAggregateVoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleVotersValidatorAggregateVoteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOracleVotersValidatorAggregateVote",
		Method:             "GET",
		PathPattern:        "/oracle/voters/{validator}/aggregate_vote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOracleVotersValidatorAggregateVoteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleVotersValidatorAggregateVoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleVotersValidatorAggregateVote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleVotersValidatorFeeder gets delegated oracle feeder of a validator
*/
func (a *Client) GetOracleVotersValidatorFeeder(params *GetOracleVotersValidatorFeederParams) (*GetOracleVotersValidatorFeederOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleVotersValidatorFeederParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOracleVotersValidatorFeeder",
		Method:             "GET",
		PathPattern:        "/oracle/voters/{validator}/feeder",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOracleVotersValidatorFeederReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleVotersValidatorFeederOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleVotersValidatorFeeder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleVotersValidatorMiss gets the number of vote periods missed in this oracle slash window
*/
func (a *Client) GetOracleVotersValidatorMiss(params *GetOracleVotersValidatorMissParams) (*GetOracleVotersValidatorMissOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleVotersValidatorMissParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOracleVotersValidatorMiss",
		Method:             "GET",
		PathPattern:        "/oracle/voters/{validator}/miss",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOracleVotersValidatorMissReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleVotersValidatorMissOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleVotersValidatorMiss: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleVotersValidatorPrevotes requests to get the currently outstanding exchange rate oracle prevotes of a validator
*/
func (a *Client) GetOracleVotersValidatorPrevotes(params *GetOracleVotersValidatorPrevotesParams) (*GetOracleVotersValidatorPrevotesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleVotersValidatorPrevotesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOracleVotersValidatorPrevotes",
		Method:             "GET",
		PathPattern:        "/oracle/voters/{validator}/prevotes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOracleVotersValidatorPrevotesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleVotersValidatorPrevotesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleVotersValidatorPrevotes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleVotersValidatorVotes requests to get the currently outstanding exchange rate oracle votes of a validator
*/
func (a *Client) GetOracleVotersValidatorVotes(params *GetOracleVotersValidatorVotesParams) (*GetOracleVotersValidatorVotesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleVotersValidatorVotesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOracleVotersValidatorVotes",
		Method:             "GET",
		PathPattern:        "/oracle/voters/{validator}/votes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOracleVotersValidatorVotesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleVotersValidatorVotesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleVotersValidatorVotes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostOracleDenomsDenomPrevotes generates oracle exchange rate prevote message containing hash of an vote
*/
func (a *Client) PostOracleDenomsDenomPrevotes(params *PostOracleDenomsDenomPrevotesParams) (*PostOracleDenomsDenomPrevotesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOracleDenomsDenomPrevotesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOracleDenomsDenomPrevotes",
		Method:             "POST",
		PathPattern:        "/oracle/denoms/{denom}/prevotes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOracleDenomsDenomPrevotesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostOracleDenomsDenomPrevotesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostOracleDenomsDenomPrevotes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostOracleDenomsDenomVotes generates oracle exchange rate vote message containing exchange rate and salt for an prevote
*/
func (a *Client) PostOracleDenomsDenomVotes(params *PostOracleDenomsDenomVotesParams) (*PostOracleDenomsDenomVotesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOracleDenomsDenomVotesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOracleDenomsDenomVotes",
		Method:             "POST",
		PathPattern:        "/oracle/denoms/{denom}/votes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOracleDenomsDenomVotesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostOracleDenomsDenomVotesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostOracleDenomsDenomVotes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostOracleVotersValidatorAggregatePrevote generates oracle aggregate exchange rate prevote message containing a hash
*/
func (a *Client) PostOracleVotersValidatorAggregatePrevote(params *PostOracleVotersValidatorAggregatePrevoteParams) (*PostOracleVotersValidatorAggregatePrevoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOracleVotersValidatorAggregatePrevoteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOracleVotersValidatorAggregatePrevote",
		Method:             "POST",
		PathPattern:        "/oracle/voters/{validator}/aggregate_prevote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOracleVotersValidatorAggregatePrevoteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostOracleVotersValidatorAggregatePrevoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostOracleVotersValidatorAggregatePrevote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostOracleVotersValidatorAggregateVote generates oracle aggregate exchange rate vote message containing exchange rates and salt to prove the aggregate prevote
*/
func (a *Client) PostOracleVotersValidatorAggregateVote(params *PostOracleVotersValidatorAggregateVoteParams) (*PostOracleVotersValidatorAggregateVoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOracleVotersValidatorAggregateVoteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOracleVotersValidatorAggregateVote",
		Method:             "POST",
		PathPattern:        "/oracle/voters/{validator}/aggregate_vote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOracleVotersValidatorAggregateVoteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostOracleVotersValidatorAggregateVoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostOracleVotersValidatorAggregateVote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostOracleVotersValidatorFeeder generates oracle feeder right delegation message
*/
func (a *Client) PostOracleVotersValidatorFeeder(params *PostOracleVotersValidatorFeederParams) (*PostOracleVotersValidatorFeederOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOracleVotersValidatorFeederParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOracleVotersValidatorFeeder",
		Method:             "POST",
		PathPattern:        "/oracle/voters/{validator}/feeder",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOracleVotersValidatorFeederReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostOracleVotersValidatorFeederOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostOracleVotersValidatorFeeder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
