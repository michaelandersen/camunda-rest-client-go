package camunda

import "github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda/util"

// https://docs.camunda.org/manual/7.8/reference/rest/decision-definition/get/#result
// https://github.com/camunda/camunda-bpm-platform/blob/master/engine-rest/engine-rest/src/main/java/org/camunda/bpm/engine/rest/dto/repository/DecisionDefinitionDto.java
type DecisionDefinition struct {
	Id                                *string `json:"id,omitempty"`
	Key                               *string `json:"key,omitempty"`
	Category                          *string `json:"category,omitempty"`
	Name                              *string `json:"name,omitempty"`
	Version                           *int    `json:"Version,omitempty"`
	Resource                          *string `json:"resource,omitempty"`
	DeploymentId                      *string `json:"deploymentId,omitempty"`
	TenantId                          *string `json:"tenantId,omitempty"`
	DecisionRequirementsDefinitionId  *string `json:"decisionRequirementsDefinitionId,omitempty"`
	DecisionRequirementsDefinitionKey *string `json:"decisionRequirementsDefinitionKey,omitempty"`
	VersionTag                        *string `json:"versionTag,omitempty"`
	HistoryTimeToLive                 *int    `json:"historyTimeToLive,omitempty"` // days
}

func (d DecisionDefinition) String() string {
	return util.Stringify(d)
}
