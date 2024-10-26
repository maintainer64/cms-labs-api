package lti_connector

import (
	"gitlab.com/a10869/api-modules/backend/app/queries"
	datastore "gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
	"gitlab.com/a10869/api-modules/backend/app/usecases/lti_connector/connector"
)

// LTIConnectorAPI упрощает взаимодействие с бекендом LTI протокола
type LTIConnectorAPI struct {
	LTIFormQueries             *datastore.LTIFormQueries
	LTILaunchDataQueries       *datastore.LTILaunchDataQueries
	UserQueries                *queries.UserQueries
	LTIProtocolDatastoreConfig *datastore.Config
}

func (c *LTIConnectorAPI) ConnectorByUserID(userID uint) (*connector.Connector, error) {
	userEntity, err := c.UserQueries.Get(userID)
	if err != nil {
		return nil, err
	}
	return c.ConnectorByLaunchID(userEntity.LastLaunchID)

}

func (c *LTIConnectorAPI) ConnectorByLaunchID(launchID string) (*connector.Connector, error) {
	launchEntity, err := c.LTILaunchDataQueries.Get(launchID)
	if err != nil {
		return nil, err
	}
	form, err := c.LTIFormQueries.Get(launchEntity.LTIFormID)
	if err != nil {
		return nil, err
	}
	conn, err := connector.New(
		c.LTIProtocolDatastoreConfig,
		launchEntity.ID,
		form.LTIClientID,
	)
	if err != nil {
		return nil, err
	}
	err = conn.SetSigningKey(form.PrivateKey)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
