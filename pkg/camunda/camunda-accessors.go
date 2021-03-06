// Copyright 2018 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-accessors; DO NOT EDIT.

package camunda

import (
	"time"
)

// GetCategory returns the Category field if it's non-nil, zero value otherwise.
func (c *CaseDefinition) GetCategory() string {
	if c == nil || c.Category == nil {
		return ""
	}
	return *c.Category
}

// GetDeploymentId returns the DeploymentId field if it's non-nil, zero value otherwise.
func (c *CaseDefinition) GetDeploymentId() string {
	if c == nil || c.DeploymentId == nil {
		return ""
	}
	return *c.DeploymentId
}

// GetHistoryTimeToLive returns the HistoryTimeToLive field if it's non-nil, zero value otherwise.
func (c *CaseDefinition) GetHistoryTimeToLive() int {
	if c == nil || c.HistoryTimeToLive == nil {
		return 0
	}
	return *c.HistoryTimeToLive
}

// GetId returns the Id field if it's non-nil, zero value otherwise.
func (c *CaseDefinition) GetId() string {
	if c == nil || c.Id == nil {
		return ""
	}
	return *c.Id
}

// GetKey returns the Key field if it's non-nil, zero value otherwise.
func (c *CaseDefinition) GetKey() string {
	if c == nil || c.Key == nil {
		return ""
	}
	return *c.Key
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (c *CaseDefinition) GetName() string {
	if c == nil || c.Name == nil {
		return ""
	}
	return *c.Name
}

// GetResource returns the Resource field if it's non-nil, zero value otherwise.
func (c *CaseDefinition) GetResource() string {
	if c == nil || c.Resource == nil {
		return ""
	}
	return *c.Resource
}

// GetTenantId returns the TenantId field if it's non-nil, zero value otherwise.
func (c *CaseDefinition) GetTenantId() string {
	if c == nil || c.TenantId == nil {
		return ""
	}
	return *c.TenantId
}

// GetVersion returns the Version field if it's non-nil, zero value otherwise.
func (c *CaseDefinition) GetVersion() int {
	if c == nil || c.Version == nil {
		return 0
	}
	return *c.Version
}

// GetCategory returns the Category field if it's non-nil, zero value otherwise.
func (d *DecisionDefinition) GetCategory() string {
	if d == nil || d.Category == nil {
		return ""
	}
	return *d.Category
}

// GetDecisionRequirementsDefinitionId returns the DecisionRequirementsDefinitionId field if it's non-nil, zero value otherwise.
func (d *DecisionDefinition) GetDecisionRequirementsDefinitionId() string {
	if d == nil || d.DecisionRequirementsDefinitionId == nil {
		return ""
	}
	return *d.DecisionRequirementsDefinitionId
}

// GetDecisionRequirementsDefinitionKey returns the DecisionRequirementsDefinitionKey field if it's non-nil, zero value otherwise.
func (d *DecisionDefinition) GetDecisionRequirementsDefinitionKey() string {
	if d == nil || d.DecisionRequirementsDefinitionKey == nil {
		return ""
	}
	return *d.DecisionRequirementsDefinitionKey
}

// GetDeploymentId returns the DeploymentId field if it's non-nil, zero value otherwise.
func (d *DecisionDefinition) GetDeploymentId() string {
	if d == nil || d.DeploymentId == nil {
		return ""
	}
	return *d.DeploymentId
}

// GetHistoryTimeToLive returns the HistoryTimeToLive field if it's non-nil, zero value otherwise.
func (d *DecisionDefinition) GetHistoryTimeToLive() int {
	if d == nil || d.HistoryTimeToLive == nil {
		return 0
	}
	return *d.HistoryTimeToLive
}

// GetId returns the Id field if it's non-nil, zero value otherwise.
func (d *DecisionDefinition) GetId() string {
	if d == nil || d.Id == nil {
		return ""
	}
	return *d.Id
}

// GetKey returns the Key field if it's non-nil, zero value otherwise.
func (d *DecisionDefinition) GetKey() string {
	if d == nil || d.Key == nil {
		return ""
	}
	return *d.Key
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (d *DecisionDefinition) GetName() string {
	if d == nil || d.Name == nil {
		return ""
	}
	return *d.Name
}

// GetResource returns the Resource field if it's non-nil, zero value otherwise.
func (d *DecisionDefinition) GetResource() string {
	if d == nil || d.Resource == nil {
		return ""
	}
	return *d.Resource
}

// GetTenantId returns the TenantId field if it's non-nil, zero value otherwise.
func (d *DecisionDefinition) GetTenantId() string {
	if d == nil || d.TenantId == nil {
		return ""
	}
	return *d.TenantId
}

// GetVersion returns the Version field if it's non-nil, zero value otherwise.
func (d *DecisionDefinition) GetVersion() int {
	if d == nil || d.Version == nil {
		return 0
	}
	return *d.Version
}

// GetVersionTag returns the VersionTag field if it's non-nil, zero value otherwise.
func (d *DecisionDefinition) GetVersionTag() string {
	if d == nil || d.VersionTag == nil {
		return ""
	}
	return *d.VersionTag
}

// GetCategory returns the Category field if it's non-nil, zero value otherwise.
func (d *DecisionRequirementsDefinition) GetCategory() string {
	if d == nil || d.Category == nil {
		return ""
	}
	return *d.Category
}

// GetDeploymentId returns the DeploymentId field if it's non-nil, zero value otherwise.
func (d *DecisionRequirementsDefinition) GetDeploymentId() string {
	if d == nil || d.DeploymentId == nil {
		return ""
	}
	return *d.DeploymentId
}

// GetId returns the Id field if it's non-nil, zero value otherwise.
func (d *DecisionRequirementsDefinition) GetId() string {
	if d == nil || d.Id == nil {
		return ""
	}
	return *d.Id
}

// GetKey returns the Key field if it's non-nil, zero value otherwise.
func (d *DecisionRequirementsDefinition) GetKey() string {
	if d == nil || d.Key == nil {
		return ""
	}
	return *d.Key
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (d *DecisionRequirementsDefinition) GetName() string {
	if d == nil || d.Name == nil {
		return ""
	}
	return *d.Name
}

// GetResource returns the Resource field if it's non-nil, zero value otherwise.
func (d *DecisionRequirementsDefinition) GetResource() string {
	if d == nil || d.Resource == nil {
		return ""
	}
	return *d.Resource
}

// GetTenantId returns the TenantId field if it's non-nil, zero value otherwise.
func (d *DecisionRequirementsDefinition) GetTenantId() string {
	if d == nil || d.TenantId == nil {
		return ""
	}
	return *d.TenantId
}

// GetVersion returns the Version field if it's non-nil, zero value otherwise.
func (d *DecisionRequirementsDefinition) GetVersion() int {
	if d == nil || d.Version == nil {
		return 0
	}
	return *d.Version
}

// GetDeploymentTime returns the DeploymentTime field if it's non-nil, zero value otherwise.
func (d *Deployment) GetDeploymentTime() time.Time {
	if d == nil || d.DeploymentTime == nil {
		return time.Time{}
	}
	return *d.DeploymentTime
}

// GetId returns the Id field if it's non-nil, zero value otherwise.
func (d *Deployment) GetId() string {
	if d == nil || d.Id == nil {
		return ""
	}
	return *d.Id
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (d *Deployment) GetName() string {
	if d == nil || d.Name == nil {
		return ""
	}
	return *d.Name
}

// GetSource returns the Source field if it's non-nil, zero value otherwise.
func (d *Deployment) GetSource() string {
	if d == nil || d.Source == nil {
		return ""
	}
	return *d.Source
}

// GetTenantId returns the TenantId field if it's non-nil, zero value otherwise.
func (d *Deployment) GetTenantId() string {
	if d == nil || d.TenantId == nil {
		return ""
	}
	return *d.TenantId
}

// GetDeployChangedOnly returns the DeployChangedOnly field if it's non-nil, zero value otherwise.
func (d *DeploymentCreateRequest) GetDeployChangedOnly() bool {
	if d == nil || d.DeployChangedOnly == nil {
		return false
	}
	return *d.DeployChangedOnly
}

// GetDeploymentName returns the DeploymentName field if it's non-nil, zero value otherwise.
func (d *DeploymentCreateRequest) GetDeploymentName() string {
	if d == nil || d.DeploymentName == nil {
		return ""
	}
	return *d.DeploymentName
}

// GetDeploymentSource returns the DeploymentSource field if it's non-nil, zero value otherwise.
func (d *DeploymentCreateRequest) GetDeploymentSource() string {
	if d == nil || d.DeploymentSource == nil {
		return ""
	}
	return *d.DeploymentSource
}

// GetEnableDuplicateFiltering returns the EnableDuplicateFiltering field if it's non-nil, zero value otherwise.
func (d *DeploymentCreateRequest) GetEnableDuplicateFiltering() bool {
	if d == nil || d.EnableDuplicateFiltering == nil {
		return false
	}
	return *d.EnableDuplicateFiltering
}

// GetTenantId returns the TenantId field if it's non-nil, zero value otherwise.
func (d *DeploymentCreateRequest) GetTenantId() string {
	if d == nil || d.TenantId == nil {
		return ""
	}
	return *d.TenantId
}

// GetAfter returns the After field if it's non-nil, zero value otherwise.
func (d *DeploymentGetRequest) GetAfter() time.Time {
	if d == nil || d.After == nil {
		return time.Time{}
	}
	return *d.After
}

// GetBefore returns the Before field if it's non-nil, zero value otherwise.
func (d *DeploymentGetRequest) GetBefore() time.Time {
	if d == nil || d.Before == nil {
		return time.Time{}
	}
	return *d.Before
}

// GetFirstResult returns the FirstResult field if it's non-nil, zero value otherwise.
func (d *DeploymentGetRequest) GetFirstResult() int {
	if d == nil || d.FirstResult == nil {
		return 0
	}
	return *d.FirstResult
}

// GetId returns the Id field if it's non-nil, zero value otherwise.
func (d *DeploymentGetRequest) GetId() string {
	if d == nil || d.Id == nil {
		return ""
	}
	return *d.Id
}

// GetIncludeDeploymentsWithoutTenantId returns the IncludeDeploymentsWithoutTenantId field if it's non-nil, zero value otherwise.
func (d *DeploymentGetRequest) GetIncludeDeploymentsWithoutTenantId() bool {
	if d == nil || d.IncludeDeploymentsWithoutTenantId == nil {
		return false
	}
	return *d.IncludeDeploymentsWithoutTenantId
}

// GetMaxResults returns the MaxResults field if it's non-nil, zero value otherwise.
func (d *DeploymentGetRequest) GetMaxResults() int {
	if d == nil || d.MaxResults == nil {
		return 0
	}
	return *d.MaxResults
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (d *DeploymentGetRequest) GetName() string {
	if d == nil || d.Name == nil {
		return ""
	}
	return *d.Name
}

// GetNameLike returns the NameLike field if it's non-nil, zero value otherwise.
func (d *DeploymentGetRequest) GetNameLike() string {
	if d == nil || d.NameLike == nil {
		return ""
	}
	return *d.NameLike
}

// GetSortBy returns the SortBy field if it's non-nil, zero value otherwise.
func (d *DeploymentGetRequest) GetSortBy() string {
	if d == nil || d.SortBy == nil {
		return ""
	}
	return *d.SortBy
}

// GetSortOrder returns the SortOrder field if it's non-nil, zero value otherwise.
func (d *DeploymentGetRequest) GetSortOrder() string {
	if d == nil || d.SortOrder == nil {
		return ""
	}
	return *d.SortOrder
}

// GetSource returns the Source field if it's non-nil, zero value otherwise.
func (d *DeploymentGetRequest) GetSource() string {
	if d == nil || d.Source == nil {
		return ""
	}
	return *d.Source
}

// GetWithoutSource returns the WithoutSource field if it's non-nil, zero value otherwise.
func (d *DeploymentGetRequest) GetWithoutSource() string {
	if d == nil || d.WithoutSource == nil {
		return ""
	}
	return *d.WithoutSource
}

// GetWithoutTentantId returns the WithoutTentantId field if it's non-nil, zero value otherwise.
func (d *DeploymentGetRequest) GetWithoutTentantId() bool {
	if d == nil || d.WithoutTentantId == nil {
		return false
	}
	return *d.WithoutTentantId
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (e *Engine) GetName() string {
	if e == nil || e.Name == nil {
		return ""
	}
	return *e.Name
}

// GetActivityId returns the ActivityId field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetActivityId() string {
	if e == nil || e.ActivityId == nil {
		return ""
	}
	return *e.ActivityId
}

// GetActivityInstanceId returns the ActivityInstanceId field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetActivityInstanceId() string {
	if e == nil || e.ActivityInstanceId == nil {
		return ""
	}
	return *e.ActivityInstanceId
}

// GetErrorMessage returns the ErrorMessage field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetErrorMessage() string {
	if e == nil || e.ErrorMessage == nil {
		return ""
	}
	return *e.ErrorMessage
}

// GetExecutionId returns the ExecutionId field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetExecutionId() string {
	if e == nil || e.ExecutionId == nil {
		return ""
	}
	return *e.ExecutionId
}

// GetId returns the Id field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetId() string {
	if e == nil || e.Id == nil {
		return ""
	}
	return *e.Id
}

// GetLockExpirationTime returns the LockExpirationTime field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetLockExpirationTime() string {
	if e == nil || e.LockExpirationTime == nil {
		return ""
	}
	return *e.LockExpirationTime
}

// GetPriority returns the Priority field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetPriority() int {
	if e == nil || e.Priority == nil {
		return 0
	}
	return *e.Priority
}

// GetProcessDefinitionId returns the ProcessDefinitionId field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetProcessDefinitionId() string {
	if e == nil || e.ProcessDefinitionId == nil {
		return ""
	}
	return *e.ProcessDefinitionId
}

// GetProcessDefinitionKey returns the ProcessDefinitionKey field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetProcessDefinitionKey() string {
	if e == nil || e.ProcessDefinitionKey == nil {
		return ""
	}
	return *e.ProcessDefinitionKey
}

// GetProcessInstanceId returns the ProcessInstanceId field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetProcessInstanceId() string {
	if e == nil || e.ProcessInstanceId == nil {
		return ""
	}
	return *e.ProcessInstanceId
}

// GetRetries returns the Retries field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetRetries() int {
	if e == nil || e.Retries == nil {
		return 0
	}
	return *e.Retries
}

// GetSuspended returns the Suspended field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetSuspended() bool {
	if e == nil || e.Suspended == nil {
		return false
	}
	return *e.Suspended
}

// GetTenantId returns the TenantId field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetTenantId() string {
	if e == nil || e.TenantId == nil {
		return ""
	}
	return *e.TenantId
}

// GetTopicName returns the TopicName field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetTopicName() string {
	if e == nil || e.TopicName == nil {
		return ""
	}
	return *e.TopicName
}

// GetWorkerId returns the WorkerId field if it's non-nil, zero value otherwise.
func (e *ExternalTask) GetWorkerId() string {
	if e == nil || e.WorkerId == nil {
		return ""
	}
	return *e.WorkerId
}

// GetActivityId returns the ActivityId field if it's non-nil, zero value otherwise.
func (i *Incident) GetActivityId() string {
	if i == nil || i.ActivityId == nil {
		return ""
	}
	return *i.ActivityId
}

// GetCauseIncidentId returns the CauseIncidentId field if it's non-nil, zero value otherwise.
func (i *Incident) GetCauseIncidentId() string {
	if i == nil || i.CauseIncidentId == nil {
		return ""
	}
	return *i.CauseIncidentId
}

// GetConfiguration returns the Configuration field if it's non-nil, zero value otherwise.
func (i *Incident) GetConfiguration() string {
	if i == nil || i.Configuration == nil {
		return ""
	}
	return *i.Configuration
}

// GetExecutionId returns the ExecutionId field if it's non-nil, zero value otherwise.
func (i *Incident) GetExecutionId() string {
	if i == nil || i.ExecutionId == nil {
		return ""
	}
	return *i.ExecutionId
}

// GetId returns the Id field if it's non-nil, zero value otherwise.
func (i *Incident) GetId() string {
	if i == nil || i.Id == nil {
		return ""
	}
	return *i.Id
}

// GetIncidentMessage returns the IncidentMessage field if it's non-nil, zero value otherwise.
func (i *Incident) GetIncidentMessage() string {
	if i == nil || i.IncidentMessage == nil {
		return ""
	}
	return *i.IncidentMessage
}

// GetIncidentTimestamp returns the IncidentTimestamp field if it's non-nil, zero value otherwise.
func (i *Incident) GetIncidentTimestamp() time.Time {
	if i == nil || i.IncidentTimestamp == nil {
		return time.Time{}
	}
	return *i.IncidentTimestamp
}

// GetIncidentType returns the IncidentType field if it's non-nil, zero value otherwise.
func (i *Incident) GetIncidentType() string {
	if i == nil || i.IncidentType == nil {
		return ""
	}
	return *i.IncidentType
}

// GetJobDefinitionId returns the JobDefinitionId field if it's non-nil, zero value otherwise.
func (i *Incident) GetJobDefinitionId() string {
	if i == nil || i.JobDefinitionId == nil {
		return ""
	}
	return *i.JobDefinitionId
}

// GetProcessDefinitionId returns the ProcessDefinitionId field if it's non-nil, zero value otherwise.
func (i *Incident) GetProcessDefinitionId() string {
	if i == nil || i.ProcessDefinitionId == nil {
		return ""
	}
	return *i.ProcessDefinitionId
}

// GetProcessInstanceId returns the ProcessInstanceId field if it's non-nil, zero value otherwise.
func (i *Incident) GetProcessInstanceId() string {
	if i == nil || i.ProcessInstanceId == nil {
		return ""
	}
	return *i.ProcessInstanceId
}

// GetRootCauseIncidentId returns the RootCauseIncidentId field if it's non-nil, zero value otherwise.
func (i *Incident) GetRootCauseIncidentId() string {
	if i == nil || i.RootCauseIncidentId == nil {
		return ""
	}
	return *i.RootCauseIncidentId
}

// GetTenantId returns the TenantId field if it's non-nil, zero value otherwise.
func (i *Incident) GetTenantId() string {
	if i == nil || i.TenantId == nil {
		return ""
	}
	return *i.TenantId
}

// GetHref returns the Href field if it's non-nil, zero value otherwise.
func (l *Link) GetHref() string {
	if l == nil || l.Href == nil {
		return ""
	}
	return *l.Href
}

// GetMethod returns the Method field if it's non-nil, zero value otherwise.
func (l *Link) GetMethod() string {
	if l == nil || l.Method == nil {
		return ""
	}
	return *l.Method
}

// GetRel returns the Rel field if it's non-nil, zero value otherwise.
func (l *Link) GetRel() string {
	if l == nil || l.Rel == nil {
		return ""
	}
	return *l.Rel
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (m *Metrics) GetName() string {
	if m == nil || m.Name == nil {
		return ""
	}
	return *m.Name
}

// GetReporter returns the Reporter field if it's non-nil, zero value otherwise.
func (m *Metrics) GetReporter() string {
	if m == nil || m.Reporter == nil {
		return ""
	}
	return *m.Reporter
}

// GetTimestamp returns the Timestamp field if it's non-nil, zero value otherwise.
func (m *Metrics) GetTimestamp() time.Time {
	if m == nil || m.Timestamp == nil {
		return time.Time{}
	}
	return *m.Timestamp
}

// GetValue returns the Value field if it's non-nil, zero value otherwise.
func (m *Metrics) GetValue() int {
	if m == nil || m.Value == nil {
		return 0
	}
	return *m.Value
}

// GetCategory returns the Category field if it's non-nil, zero value otherwise.
func (p *ProcessDefinition) GetCategory() string {
	if p == nil || p.Category == nil {
		return ""
	}
	return *p.Category
}

// GetDeploymentId returns the DeploymentId field if it's non-nil, zero value otherwise.
func (p *ProcessDefinition) GetDeploymentId() string {
	if p == nil || p.DeploymentId == nil {
		return ""
	}
	return *p.DeploymentId
}

// GetDescription returns the Description field if it's non-nil, zero value otherwise.
func (p *ProcessDefinition) GetDescription() string {
	if p == nil || p.Description == nil {
		return ""
	}
	return *p.Description
}

// GetDiagram returns the Diagram field if it's non-nil, zero value otherwise.
func (p *ProcessDefinition) GetDiagram() string {
	if p == nil || p.Diagram == nil {
		return ""
	}
	return *p.Diagram
}

// GetHistoryTimeToLive returns the HistoryTimeToLive field if it's non-nil, zero value otherwise.
func (p *ProcessDefinition) GetHistoryTimeToLive() int {
	if p == nil || p.HistoryTimeToLive == nil {
		return 0
	}
	return *p.HistoryTimeToLive
}

// GetId returns the Id field if it's non-nil, zero value otherwise.
func (p *ProcessDefinition) GetId() string {
	if p == nil || p.Id == nil {
		return ""
	}
	return *p.Id
}

// GetKey returns the Key field if it's non-nil, zero value otherwise.
func (p *ProcessDefinition) GetKey() string {
	if p == nil || p.Key == nil {
		return ""
	}
	return *p.Key
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (p *ProcessDefinition) GetName() string {
	if p == nil || p.Name == nil {
		return ""
	}
	return *p.Name
}

// GetResource returns the Resource field if it's non-nil, zero value otherwise.
func (p *ProcessDefinition) GetResource() string {
	if p == nil || p.Resource == nil {
		return ""
	}
	return *p.Resource
}

// GetSuspended returns the Suspended field if it's non-nil, zero value otherwise.
func (p *ProcessDefinition) GetSuspended() bool {
	if p == nil || p.Suspended == nil {
		return false
	}
	return *p.Suspended
}

// GetTenantId returns the TenantId field if it's non-nil, zero value otherwise.
func (p *ProcessDefinition) GetTenantId() string {
	if p == nil || p.TenantId == nil {
		return ""
	}
	return *p.TenantId
}

// GetVersion returns the Version field if it's non-nil, zero value otherwise.
func (p *ProcessDefinition) GetVersion() int {
	if p == nil || p.Version == nil {
		return 0
	}
	return *p.Version
}

// GetVersionTag returns the VersionTag field if it's non-nil, zero value otherwise.
func (p *ProcessDefinition) GetVersionTag() string {
	if p == nil || p.VersionTag == nil {
		return ""
	}
	return *p.VersionTag
}

// GetBusinessKey returns the BusinessKey field if it's non-nil, zero value otherwise.
func (p *ProcessInstance) GetBusinessKey() string {
	if p == nil || p.BusinessKey == nil {
		return ""
	}
	return *p.BusinessKey
}

// GetCaseInstanceId returns the CaseInstanceId field if it's non-nil, zero value otherwise.
func (p *ProcessInstance) GetCaseInstanceId() string {
	if p == nil || p.CaseInstanceId == nil {
		return ""
	}
	return *p.CaseInstanceId
}

// GetDefinitionId returns the DefinitionId field if it's non-nil, zero value otherwise.
func (p *ProcessInstance) GetDefinitionId() string {
	if p == nil || p.DefinitionId == nil {
		return ""
	}
	return *p.DefinitionId
}

// GetEnded returns the Ended field if it's non-nil, zero value otherwise.
func (p *ProcessInstance) GetEnded() bool {
	if p == nil || p.Ended == nil {
		return false
	}
	return *p.Ended
}

// GetId returns the Id field if it's non-nil, zero value otherwise.
func (p *ProcessInstance) GetId() string {
	if p == nil || p.Id == nil {
		return ""
	}
	return *p.Id
}

// GetSuspended returns the Suspended field if it's non-nil, zero value otherwise.
func (p *ProcessInstance) GetSuspended() bool {
	if p == nil || p.Suspended == nil {
		return false
	}
	return *p.Suspended
}

// GetTenantId returns the TenantId field if it's non-nil, zero value otherwise.
func (p *ProcessInstance) GetTenantId() string {
	if p == nil || p.TenantId == nil {
		return ""
	}
	return *p.TenantId
}

// GetEncoding returns the Encoding field if it's non-nil, zero value otherwise.
func (v *ValueInfo) GetEncoding() string {
	if v == nil || v.Encoding == nil {
		return ""
	}
	return *v.Encoding
}

// GetFilename returns the Filename field if it's non-nil, zero value otherwise.
func (v *ValueInfo) GetFilename() string {
	if v == nil || v.Filename == nil {
		return ""
	}
	return *v.Filename
}

// GetMimeType returns the MimeType field if it's non-nil, zero value otherwise.
func (v *ValueInfo) GetMimeType() string {
	if v == nil || v.MimeType == nil {
		return ""
	}
	return *v.MimeType
}

// GetObjectTypeName returns the ObjectTypeName field if it's non-nil, zero value otherwise.
func (v *ValueInfo) GetObjectTypeName() string {
	if v == nil || v.ObjectTypeName == nil {
		return ""
	}
	return *v.ObjectTypeName
}

// GetSerializationDataFormat returns the SerializationDataFormat field if it's non-nil, zero value otherwise.
func (v *ValueInfo) GetSerializationDataFormat() string {
	if v == nil || v.SerializationDataFormat == nil {
		return ""
	}
	return *v.SerializationDataFormat
}
