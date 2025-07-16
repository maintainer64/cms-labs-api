// Copyright (c) 2021 MacEwan University. All rights reserved.
//
// This source code is licensed under the MIT-style license found in
// the LICENSE file in the root directory of this source tree.

// Package lti_query implements a persistent SQL data store. It implements the RegistrationStorer interface.
package lti_query

import (
	"errors"
	"fmt"
	"net/url"

	"gitlab.com/a10869/api-modules/backend/app/models"
)

// FindRegistrationByIssuerAndClientID retrieves a registration from the SQL database.
func (q *LTIFormQueries) FindRegistrationByIssuerAndClientID(issuer, clientID string) (Registration, error) {
	q.Logger.Info().Msg(fmt.Sprintf(
		"LTIFormQueries: find registration by issuer: %+v and clientID: %+v",
		issuer,
		clientID,
	))
	if issuer == "" {
		return Registration{}, errors.New("received empty issuer argument")
	}
	entityDB := models.LTIForm{}
	query := q.DB.Where("base_uri = ?", issuer)
	if clientID != "" {
		// Use the client ID to disambiguate multiple registrations for an issuer.  The (optional) client ID
		// parameter can disambiguate between multiple registrations from a single issuer.
		//
		// Source: http://www.imsglobal.org/spec/lti/v1p3/#client_id-login-parameter
		query = query.Where("lti_client_id = ?", clientID)
	}
	query.Find(&entityDB)
	err := ErrRegistrationNotFound
	if entityDB.ID == 0 {
		return Registration{}, err
	}
	reg := Registration{}
	reg.ID = entityDB.ID
	reg.ClientID = entityDB.LTIClientID
	reg.Issuer = entityDB.BaseURI
	reg.AuthTokenURI, err = url.Parse(entityDB.LTIAuthTokenURI)
	if err != nil {
		return Registration{}, err
	}
	reg.AuthLoginURI, err = url.Parse(entityDB.LTIAuthLoginURI)
	if err != nil {
		return Registration{}, err
	}
	reg.KeysetURI, err = url.Parse(entityDB.KeySetURI)
	if err != nil {
		return Registration{}, err
	}
	reg.TargetLinkURI, err = url.Parse(entityDB.TargetLinkURI)
	if err != nil {
		return Registration{}, err
	}
	return reg, nil
}

// FindDeployment looks up and returns either a Deployment by the issuer and deployment ID or the datastore error
// ErrDeploymentNotFound.
func (q *LTIFormQueries) FindDeployment(issuer, deploymentID string) (Deployment, error) {
	q.Logger.Info().Msg(fmt.Sprintf(
		"LTIFormQueries: find deployment by issuer: %+v and deploymentID: %+v",
		issuer,
		deploymentID,
	))
	if issuer == "" {
		return Deployment{}, errors.New("received empty issuer argument")
	}
	if err := ValidateDeploymentID(deploymentID); err != nil {
		return Deployment{}, fmt.Errorf("received invalid deployment ID: %v", err)
	}
	entityDB := models.LTIForm{}
	result := q.DB.Where("base_uri = ?", issuer).Where("lti_deployment_id = ?", deploymentID).Find(&entityDB)
	if result.Error != nil {
		return Deployment{}, ErrRegistrationNotFound
	}
	return Deployment{DeploymentID: entityDB.LTIDeploymentID}, nil
}
