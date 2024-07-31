// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * (title)
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.0
 */

package openapi




// MqttConnectEventSuccessResponse - Represents a successful response to an MQTT connection event.
type MqttConnectEventSuccessResponse struct {

	// The property provides a way for the upstream Webhook to authorize the client. There are different roles to grant initial permissions for PubSub WebSocket clients. See details in https://learn.microsoft.com/azure/azure-web-pubsub/concept-client-protocols#permissions
	Roles []string `json:"roles,omitempty"`

	// It should be \"mqtt\" or omitted.
	Subprotocol string `json:"subprotocol,omitempty"`

	// As the service allows anonymous connections, it's the connect event's responsibility to tell the service the user ID of the client connection. The service reads the user ID from the response payload userId if it exists.
	UserId string `json:"userId,omitempty"`

	// Initial groups the client joins. The property provides a convenient way for user to add the client to one or multiple groups. In this way, there's no need to have another call to add this connection to some groups.
	Groups []string `json:"groups,omitempty"`

	Mqtt MqttConnectEventSuccessResponseProperties `json:"mqtt,omitempty"`
}

// AssertMqttConnectEventSuccessResponseRequired checks if the required fields are not zero-ed
func AssertMqttConnectEventSuccessResponseRequired(obj MqttConnectEventSuccessResponse) error {
	if err := AssertMqttConnectEventSuccessResponsePropertiesRequired(obj.Mqtt); err != nil {
		return err
	}
	return nil
}

// AssertMqttConnectEventSuccessResponseConstraints checks if the values respects the defined constraints
func AssertMqttConnectEventSuccessResponseConstraints(obj MqttConnectEventSuccessResponse) error {
	if err := AssertMqttConnectEventSuccessResponsePropertiesConstraints(obj.Mqtt); err != nil {
		return err
	}
	return nil
}