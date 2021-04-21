// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package private

import (
	externalRef0 "playbook-dispatcher/internal/api/controllers/public"
)

// RunCreated defines model for RunCreated.
type RunCreated struct {

	// status code of the request
	Code int `json:"code"`

	// Unique identifier of a Playbook run
	Id *externalRef0.RunId `json:"id,omitempty"`
}

// RunInput defines model for RunInput.
type RunInput struct {

	// Identifier of the tenant
	Account externalRef0.Account `json:"account"`

	// Optionally, information about hosts involved in the Playbook run can be provided.
	// This information is used to pre-allocate run_host resources.
	// Moreover, it can be used to create a connection between a run_host resource and host inventory.
	Hosts *RunInputHosts `json:"hosts,omitempty"`

	// Additional metadata about the Playbook run. Can be used for filtering purposes.
	Labels *externalRef0.Labels `json:"labels,omitempty"`

	// Identifier of the host to which a given Playbook is addressed
	Recipient externalRef0.RunRecipient `json:"recipient"`

	// Amount of seconds after which the run is considered failed due to timeout
	Timeout *externalRef0.RunTimeout `json:"timeout,omitempty"`

	// URL hosting the Playbook
	Url externalRef0.Url `json:"url"`
}

// RunInputHosts defines model for RunInputHosts.
type RunInputHosts []struct {

	// Host name as known to Ansible inventory.
	// Used to identify the host in status reports.
	AnsibleHost string `json:"ansible_host"`

	// Inventory id of the given host
	InventoryId *string `json:"inventory_id,omitempty"`
}

// RunsCreated defines model for RunsCreated.
type RunsCreated []RunCreated

// ApiInternalRunsCreateJSONBody defines parameters for ApiInternalRunsCreate.
type ApiInternalRunsCreateJSONBody []RunInput

// ApiInternalRunsCreateRequestBody defines body for ApiInternalRunsCreate for application/json ContentType.
type ApiInternalRunsCreateJSONRequestBody ApiInternalRunsCreateJSONBody
