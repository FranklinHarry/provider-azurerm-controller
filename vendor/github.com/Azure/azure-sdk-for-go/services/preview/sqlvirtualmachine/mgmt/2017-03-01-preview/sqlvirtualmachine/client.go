// Package sqlvirtualmachine implements the Azure ARM Sqlvirtualmachine service API version 2017-03-01-preview.
//
// The SQL virtual machine management API provides a RESTful set of web APIs that interact with Azure Compute, Network
// & Storage services to manage your SQL Server virtual machine. The API enables users to create, delete and retrieve a
// SQL virtual machine, SQL virtual machine group or availability group listener.
package sqlvirtualmachine

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Sqlvirtualmachine
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Sqlvirtualmachine.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}
