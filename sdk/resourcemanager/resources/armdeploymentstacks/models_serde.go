//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentstacks

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type DebugSetting.
func (d DebugSetting) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "detailLevel", d.DetailLevel)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DebugSetting.
func (d *DebugSetting) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "detailLevel":
			err = unpopulate(val, "DetailLevel", &d.DetailLevel)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DenySettings.
func (d DenySettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "applyToChildScopes", d.ApplyToChildScopes)
	populate(objectMap, "excludedActions", d.ExcludedActions)
	populate(objectMap, "excludedPrincipals", d.ExcludedPrincipals)
	populate(objectMap, "mode", d.Mode)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DenySettings.
func (d *DenySettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "applyToChildScopes":
			err = unpopulate(val, "ApplyToChildScopes", &d.ApplyToChildScopes)
			delete(rawMsg, key)
		case "excludedActions":
			err = unpopulate(val, "ExcludedActions", &d.ExcludedActions)
			delete(rawMsg, key)
		case "excludedPrincipals":
			err = unpopulate(val, "ExcludedPrincipals", &d.ExcludedPrincipals)
			delete(rawMsg, key)
		case "mode":
			err = unpopulate(val, "Mode", &d.Mode)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeploymentStack.
func (d DeploymentStack) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", d.ID)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "systemData", d.SystemData)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeploymentStack.
func (d *DeploymentStack) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &d.ID)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &d.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &d.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &d.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &d.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &d.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &d.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeploymentStackListResult.
func (d DeploymentStackListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeploymentStackListResult.
func (d *DeploymentStackListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &d.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &d.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeploymentStackProperties.
func (d DeploymentStackProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "actionOnUnmanage", d.ActionOnUnmanage)
	populate(objectMap, "debugSetting", d.DebugSetting)
	populate(objectMap, "deletedResources", d.DeletedResources)
	populate(objectMap, "denySettings", d.DenySettings)
	populate(objectMap, "deploymentId", d.DeploymentID)
	populate(objectMap, "deploymentScope", d.DeploymentScope)
	populate(objectMap, "description", d.Description)
	populate(objectMap, "detachedResources", d.DetachedResources)
	populate(objectMap, "duration", d.Duration)
	populate(objectMap, "error", d.Error)
	populate(objectMap, "failedResources", d.FailedResources)
	populateAny(objectMap, "outputs", d.Outputs)
	populateAny(objectMap, "parameters", d.Parameters)
	populate(objectMap, "parametersLink", d.ParametersLink)
	populate(objectMap, "provisioningState", d.ProvisioningState)
	populate(objectMap, "resources", d.Resources)
	populateAny(objectMap, "template", d.Template)
	populate(objectMap, "templateLink", d.TemplateLink)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeploymentStackProperties.
func (d *DeploymentStackProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionOnUnmanage":
			err = unpopulate(val, "ActionOnUnmanage", &d.ActionOnUnmanage)
			delete(rawMsg, key)
		case "debugSetting":
			err = unpopulate(val, "DebugSetting", &d.DebugSetting)
			delete(rawMsg, key)
		case "deletedResources":
			err = unpopulate(val, "DeletedResources", &d.DeletedResources)
			delete(rawMsg, key)
		case "denySettings":
			err = unpopulate(val, "DenySettings", &d.DenySettings)
			delete(rawMsg, key)
		case "deploymentId":
			err = unpopulate(val, "DeploymentID", &d.DeploymentID)
			delete(rawMsg, key)
		case "deploymentScope":
			err = unpopulate(val, "DeploymentScope", &d.DeploymentScope)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &d.Description)
			delete(rawMsg, key)
		case "detachedResources":
			err = unpopulate(val, "DetachedResources", &d.DetachedResources)
			delete(rawMsg, key)
		case "duration":
			err = unpopulate(val, "Duration", &d.Duration)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, "Error", &d.Error)
			delete(rawMsg, key)
		case "failedResources":
			err = unpopulate(val, "FailedResources", &d.FailedResources)
			delete(rawMsg, key)
		case "outputs":
			err = unpopulate(val, "Outputs", &d.Outputs)
			delete(rawMsg, key)
		case "parameters":
			err = unpopulate(val, "Parameters", &d.Parameters)
			delete(rawMsg, key)
		case "parametersLink":
			err = unpopulate(val, "ParametersLink", &d.ParametersLink)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &d.ProvisioningState)
			delete(rawMsg, key)
		case "resources":
			err = unpopulate(val, "Resources", &d.Resources)
			delete(rawMsg, key)
		case "template":
			err = unpopulate(val, "Template", &d.Template)
			delete(rawMsg, key)
		case "templateLink":
			err = unpopulate(val, "TemplateLink", &d.TemplateLink)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeploymentStackPropertiesActionOnUnmanage.
func (d DeploymentStackPropertiesActionOnUnmanage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "managementGroups", d.ManagementGroups)
	populate(objectMap, "resourceGroups", d.ResourceGroups)
	populate(objectMap, "resources", d.Resources)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeploymentStackPropertiesActionOnUnmanage.
func (d *DeploymentStackPropertiesActionOnUnmanage) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "managementGroups":
			err = unpopulate(val, "ManagementGroups", &d.ManagementGroups)
			delete(rawMsg, key)
		case "resourceGroups":
			err = unpopulate(val, "ResourceGroups", &d.ResourceGroups)
			delete(rawMsg, key)
		case "resources":
			err = unpopulate(val, "Resources", &d.Resources)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeploymentStackTemplateDefinition.
func (d DeploymentStackTemplateDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateAny(objectMap, "template", d.Template)
	populate(objectMap, "templateLink", d.TemplateLink)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeploymentStackTemplateDefinition.
func (d *DeploymentStackTemplateDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "template":
			err = unpopulate(val, "Template", &d.Template)
			delete(rawMsg, key)
		case "templateLink":
			err = unpopulate(val, "TemplateLink", &d.TemplateLink)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorAdditionalInfo.
func (e ErrorAdditionalInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateAny(objectMap, "info", e.Info)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorAdditionalInfo.
func (e *ErrorAdditionalInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "info":
			err = unpopulate(val, "Info", &e.Info)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &e.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetail.
func (e *ErrorDetail) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalInfo":
			err = unpopulate(val, "AdditionalInfo", &e.AdditionalInfo)
			delete(rawMsg, key)
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, "Details", &e.Details)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &e.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", e.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponse.
func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, "Error", &e.Error)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ManagedResourceReference.
func (m ManagedResourceReference) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "denyStatus", m.DenyStatus)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "status", m.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ManagedResourceReference.
func (m *ManagedResourceReference) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "denyStatus":
			err = unpopulate(val, "DenyStatus", &m.DenyStatus)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &m.ID)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &m.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ParametersLink.
func (p ParametersLink) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "contentVersion", p.ContentVersion)
	populate(objectMap, "uri", p.URI)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ParametersLink.
func (p *ParametersLink) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contentVersion":
			err = unpopulate(val, "ContentVersion", &p.ContentVersion)
			delete(rawMsg, key)
		case "uri":
			err = unpopulate(val, "URI", &p.URI)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourceReference.
func (r ResourceReference) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourceReference.
func (r *ResourceReference) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourceReferenceExtended.
func (r ResourceReferenceExtended) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", r.Error)
	populate(objectMap, "id", r.ID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourceReferenceExtended.
func (r *ResourceReferenceExtended) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, "Error", &r.Error)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateDateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateDateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateDateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateDateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TemplateLink.
func (t TemplateLink) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "contentVersion", t.ContentVersion)
	populate(objectMap, "id", t.ID)
	populate(objectMap, "queryString", t.QueryString)
	populate(objectMap, "relativePath", t.RelativePath)
	populate(objectMap, "uri", t.URI)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TemplateLink.
func (t *TemplateLink) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contentVersion":
			err = unpopulate(val, "ContentVersion", &t.ContentVersion)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &t.ID)
			delete(rawMsg, key)
		case "queryString":
			err = unpopulate(val, "QueryString", &t.QueryString)
			delete(rawMsg, key)
		case "relativePath":
			err = unpopulate(val, "RelativePath", &t.RelativePath)
			delete(rawMsg, key)
		case "uri":
			err = unpopulate(val, "URI", &t.URI)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateAny(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
