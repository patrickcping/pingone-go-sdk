# EntityArrayEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServers** | Pointer to [**[]APIServer**](APIServer.md) |  | [optional] 
**Assignments** | Pointer to [**[]ApplicationRoleAssignment**](ApplicationRoleAssignment.md) |  | [optional] 
**AuthorizationAttributes** | Pointer to [**[]AuthorizeEditorDataDefinitionsAttributeDefinitionDTO**](AuthorizeEditorDataDefinitionsAttributeDefinitionDTO.md) |  | [optional] 
**AuthorizationConditions** | Pointer to [**[]AuthorizeEditorDataDefinitionsConditionDefinitionDTO**](AuthorizeEditorDataDefinitionsConditionDefinitionDTO.md) |  | [optional] 
**AuthorizationConnectorTemplates** | Pointer to [**[]AuthorizeEditorDataConnectorsConnectorTemplateDTO**](AuthorizeEditorDataConnectorsConnectorTemplateDTO.md) |  | [optional] 
**AuthorizationPolicies** | Pointer to [**[]AuthorizeEditorDataPoliciesReferenceablePolicyDTO**](AuthorizeEditorDataPoliciesReferenceablePolicyDTO.md) |  | [optional] 
**AuthorizationProcessors** | Pointer to [**[]AuthorizeEditorDataDefinitionsProcessorDefinitionDTO**](AuthorizeEditorDataDefinitionsProcessorDefinitionDTO.md) |  | [optional] 
**AuthorizationRules** | Pointer to [**[]AuthorizeEditorDataRulesReferenceableRuleDTO**](AuthorizeEditorDataRulesReferenceableRuleDTO.md) |  | [optional] 
**AuthorizationServices** | Pointer to [**[]AuthorizeEditorDataDefinitionsServiceDefinitionDTO**](AuthorizeEditorDataDefinitionsServiceDefinitionDTO.md) |  | [optional] 
**AuthorizationStatements** | Pointer to [**[]AuthorizeEditorDataStatementsReferenceableStatementDTO**](AuthorizeEditorDataStatementsReferenceableStatementDTO.md) |  | [optional] 
**AuthorizationVersions** | Pointer to [**[]AuthorizeEditorDataAuthorizationVersionDTO**](AuthorizeEditorDataAuthorizationVersionDTO.md) |  | [optional] 
**DecisionEndpoints** | Pointer to [**[]DecisionEndpoint**](DecisionEndpoint.md) |  | [optional] 
**Operations** | Pointer to [**[]APIServerOperation**](APIServerOperation.md) |  | [optional] 
**Permissions** | Pointer to [**[]EntityArrayEmbeddedPermissionsInner**](EntityArrayEmbeddedPermissionsInner.md) |  | [optional] 
**Resources** | Pointer to [**[]ApplicationResource**](ApplicationResource.md) |  | [optional] 
**Roles** | Pointer to [**[]ApplicationRole**](ApplicationRole.md) |  | [optional] 

## Methods

### NewEntityArrayEmbedded

`func NewEntityArrayEmbedded() *EntityArrayEmbedded`

NewEntityArrayEmbedded instantiates a new EntityArrayEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedWithDefaults

`func NewEntityArrayEmbeddedWithDefaults() *EntityArrayEmbedded`

NewEntityArrayEmbeddedWithDefaults instantiates a new EntityArrayEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiServers

`func (o *EntityArrayEmbedded) GetApiServers() []APIServer`

GetApiServers returns the ApiServers field if non-nil, zero value otherwise.

### GetApiServersOk

`func (o *EntityArrayEmbedded) GetApiServersOk() (*[]APIServer, bool)`

GetApiServersOk returns a tuple with the ApiServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServers

`func (o *EntityArrayEmbedded) SetApiServers(v []APIServer)`

SetApiServers sets ApiServers field to given value.

### HasApiServers

`func (o *EntityArrayEmbedded) HasApiServers() bool`

HasApiServers returns a boolean if a field has been set.

### GetAssignments

`func (o *EntityArrayEmbedded) GetAssignments() []ApplicationRoleAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *EntityArrayEmbedded) GetAssignmentsOk() (*[]ApplicationRoleAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *EntityArrayEmbedded) SetAssignments(v []ApplicationRoleAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *EntityArrayEmbedded) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetAuthorizationAttributes

`func (o *EntityArrayEmbedded) GetAuthorizationAttributes() []AuthorizeEditorDataDefinitionsAttributeDefinitionDTO`

GetAuthorizationAttributes returns the AuthorizationAttributes field if non-nil, zero value otherwise.

### GetAuthorizationAttributesOk

`func (o *EntityArrayEmbedded) GetAuthorizationAttributesOk() (*[]AuthorizeEditorDataDefinitionsAttributeDefinitionDTO, bool)`

GetAuthorizationAttributesOk returns a tuple with the AuthorizationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationAttributes

`func (o *EntityArrayEmbedded) SetAuthorizationAttributes(v []AuthorizeEditorDataDefinitionsAttributeDefinitionDTO)`

SetAuthorizationAttributes sets AuthorizationAttributes field to given value.

### HasAuthorizationAttributes

`func (o *EntityArrayEmbedded) HasAuthorizationAttributes() bool`

HasAuthorizationAttributes returns a boolean if a field has been set.

### GetAuthorizationConditions

`func (o *EntityArrayEmbedded) GetAuthorizationConditions() []AuthorizeEditorDataDefinitionsConditionDefinitionDTO`

GetAuthorizationConditions returns the AuthorizationConditions field if non-nil, zero value otherwise.

### GetAuthorizationConditionsOk

`func (o *EntityArrayEmbedded) GetAuthorizationConditionsOk() (*[]AuthorizeEditorDataDefinitionsConditionDefinitionDTO, bool)`

GetAuthorizationConditionsOk returns a tuple with the AuthorizationConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationConditions

`func (o *EntityArrayEmbedded) SetAuthorizationConditions(v []AuthorizeEditorDataDefinitionsConditionDefinitionDTO)`

SetAuthorizationConditions sets AuthorizationConditions field to given value.

### HasAuthorizationConditions

`func (o *EntityArrayEmbedded) HasAuthorizationConditions() bool`

HasAuthorizationConditions returns a boolean if a field has been set.

### GetAuthorizationConnectorTemplates

`func (o *EntityArrayEmbedded) GetAuthorizationConnectorTemplates() []AuthorizeEditorDataConnectorsConnectorTemplateDTO`

GetAuthorizationConnectorTemplates returns the AuthorizationConnectorTemplates field if non-nil, zero value otherwise.

### GetAuthorizationConnectorTemplatesOk

`func (o *EntityArrayEmbedded) GetAuthorizationConnectorTemplatesOk() (*[]AuthorizeEditorDataConnectorsConnectorTemplateDTO, bool)`

GetAuthorizationConnectorTemplatesOk returns a tuple with the AuthorizationConnectorTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationConnectorTemplates

`func (o *EntityArrayEmbedded) SetAuthorizationConnectorTemplates(v []AuthorizeEditorDataConnectorsConnectorTemplateDTO)`

SetAuthorizationConnectorTemplates sets AuthorizationConnectorTemplates field to given value.

### HasAuthorizationConnectorTemplates

`func (o *EntityArrayEmbedded) HasAuthorizationConnectorTemplates() bool`

HasAuthorizationConnectorTemplates returns a boolean if a field has been set.

### GetAuthorizationPolicies

`func (o *EntityArrayEmbedded) GetAuthorizationPolicies() []AuthorizeEditorDataPoliciesReferenceablePolicyDTO`

GetAuthorizationPolicies returns the AuthorizationPolicies field if non-nil, zero value otherwise.

### GetAuthorizationPoliciesOk

`func (o *EntityArrayEmbedded) GetAuthorizationPoliciesOk() (*[]AuthorizeEditorDataPoliciesReferenceablePolicyDTO, bool)`

GetAuthorizationPoliciesOk returns a tuple with the AuthorizationPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationPolicies

`func (o *EntityArrayEmbedded) SetAuthorizationPolicies(v []AuthorizeEditorDataPoliciesReferenceablePolicyDTO)`

SetAuthorizationPolicies sets AuthorizationPolicies field to given value.

### HasAuthorizationPolicies

`func (o *EntityArrayEmbedded) HasAuthorizationPolicies() bool`

HasAuthorizationPolicies returns a boolean if a field has been set.

### GetAuthorizationProcessors

`func (o *EntityArrayEmbedded) GetAuthorizationProcessors() []AuthorizeEditorDataDefinitionsProcessorDefinitionDTO`

GetAuthorizationProcessors returns the AuthorizationProcessors field if non-nil, zero value otherwise.

### GetAuthorizationProcessorsOk

`func (o *EntityArrayEmbedded) GetAuthorizationProcessorsOk() (*[]AuthorizeEditorDataDefinitionsProcessorDefinitionDTO, bool)`

GetAuthorizationProcessorsOk returns a tuple with the AuthorizationProcessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationProcessors

`func (o *EntityArrayEmbedded) SetAuthorizationProcessors(v []AuthorizeEditorDataDefinitionsProcessorDefinitionDTO)`

SetAuthorizationProcessors sets AuthorizationProcessors field to given value.

### HasAuthorizationProcessors

`func (o *EntityArrayEmbedded) HasAuthorizationProcessors() bool`

HasAuthorizationProcessors returns a boolean if a field has been set.

### GetAuthorizationRules

`func (o *EntityArrayEmbedded) GetAuthorizationRules() []AuthorizeEditorDataRulesReferenceableRuleDTO`

GetAuthorizationRules returns the AuthorizationRules field if non-nil, zero value otherwise.

### GetAuthorizationRulesOk

`func (o *EntityArrayEmbedded) GetAuthorizationRulesOk() (*[]AuthorizeEditorDataRulesReferenceableRuleDTO, bool)`

GetAuthorizationRulesOk returns a tuple with the AuthorizationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationRules

`func (o *EntityArrayEmbedded) SetAuthorizationRules(v []AuthorizeEditorDataRulesReferenceableRuleDTO)`

SetAuthorizationRules sets AuthorizationRules field to given value.

### HasAuthorizationRules

`func (o *EntityArrayEmbedded) HasAuthorizationRules() bool`

HasAuthorizationRules returns a boolean if a field has been set.

### GetAuthorizationServices

`func (o *EntityArrayEmbedded) GetAuthorizationServices() []AuthorizeEditorDataDefinitionsServiceDefinitionDTO`

GetAuthorizationServices returns the AuthorizationServices field if non-nil, zero value otherwise.

### GetAuthorizationServicesOk

`func (o *EntityArrayEmbedded) GetAuthorizationServicesOk() (*[]AuthorizeEditorDataDefinitionsServiceDefinitionDTO, bool)`

GetAuthorizationServicesOk returns a tuple with the AuthorizationServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServices

`func (o *EntityArrayEmbedded) SetAuthorizationServices(v []AuthorizeEditorDataDefinitionsServiceDefinitionDTO)`

SetAuthorizationServices sets AuthorizationServices field to given value.

### HasAuthorizationServices

`func (o *EntityArrayEmbedded) HasAuthorizationServices() bool`

HasAuthorizationServices returns a boolean if a field has been set.

### GetAuthorizationStatements

`func (o *EntityArrayEmbedded) GetAuthorizationStatements() []AuthorizeEditorDataStatementsReferenceableStatementDTO`

GetAuthorizationStatements returns the AuthorizationStatements field if non-nil, zero value otherwise.

### GetAuthorizationStatementsOk

`func (o *EntityArrayEmbedded) GetAuthorizationStatementsOk() (*[]AuthorizeEditorDataStatementsReferenceableStatementDTO, bool)`

GetAuthorizationStatementsOk returns a tuple with the AuthorizationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationStatements

`func (o *EntityArrayEmbedded) SetAuthorizationStatements(v []AuthorizeEditorDataStatementsReferenceableStatementDTO)`

SetAuthorizationStatements sets AuthorizationStatements field to given value.

### HasAuthorizationStatements

`func (o *EntityArrayEmbedded) HasAuthorizationStatements() bool`

HasAuthorizationStatements returns a boolean if a field has been set.

### GetAuthorizationVersions

`func (o *EntityArrayEmbedded) GetAuthorizationVersions() []AuthorizeEditorDataAuthorizationVersionDTO`

GetAuthorizationVersions returns the AuthorizationVersions field if non-nil, zero value otherwise.

### GetAuthorizationVersionsOk

`func (o *EntityArrayEmbedded) GetAuthorizationVersionsOk() (*[]AuthorizeEditorDataAuthorizationVersionDTO, bool)`

GetAuthorizationVersionsOk returns a tuple with the AuthorizationVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationVersions

`func (o *EntityArrayEmbedded) SetAuthorizationVersions(v []AuthorizeEditorDataAuthorizationVersionDTO)`

SetAuthorizationVersions sets AuthorizationVersions field to given value.

### HasAuthorizationVersions

`func (o *EntityArrayEmbedded) HasAuthorizationVersions() bool`

HasAuthorizationVersions returns a boolean if a field has been set.

### GetDecisionEndpoints

`func (o *EntityArrayEmbedded) GetDecisionEndpoints() []DecisionEndpoint`

GetDecisionEndpoints returns the DecisionEndpoints field if non-nil, zero value otherwise.

### GetDecisionEndpointsOk

`func (o *EntityArrayEmbedded) GetDecisionEndpointsOk() (*[]DecisionEndpoint, bool)`

GetDecisionEndpointsOk returns a tuple with the DecisionEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionEndpoints

`func (o *EntityArrayEmbedded) SetDecisionEndpoints(v []DecisionEndpoint)`

SetDecisionEndpoints sets DecisionEndpoints field to given value.

### HasDecisionEndpoints

`func (o *EntityArrayEmbedded) HasDecisionEndpoints() bool`

HasDecisionEndpoints returns a boolean if a field has been set.

### GetOperations

`func (o *EntityArrayEmbedded) GetOperations() []APIServerOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *EntityArrayEmbedded) GetOperationsOk() (*[]APIServerOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *EntityArrayEmbedded) SetOperations(v []APIServerOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *EntityArrayEmbedded) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetPermissions

`func (o *EntityArrayEmbedded) GetPermissions() []EntityArrayEmbeddedPermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *EntityArrayEmbedded) GetPermissionsOk() (*[]EntityArrayEmbeddedPermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *EntityArrayEmbedded) SetPermissions(v []EntityArrayEmbeddedPermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *EntityArrayEmbedded) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetResources

`func (o *EntityArrayEmbedded) GetResources() []ApplicationResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *EntityArrayEmbedded) GetResourcesOk() (*[]ApplicationResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *EntityArrayEmbedded) SetResources(v []ApplicationResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *EntityArrayEmbedded) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRoles

`func (o *EntityArrayEmbedded) GetRoles() []ApplicationRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *EntityArrayEmbedded) GetRolesOk() (*[]ApplicationRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *EntityArrayEmbedded) SetRoles(v []ApplicationRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *EntityArrayEmbedded) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


