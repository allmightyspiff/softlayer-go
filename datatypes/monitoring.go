/**
 * Copyright 2016 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * AUTOMATICALLY GENERATED CODE - DO NOT MODIFY
 */

package datatypes

// A monitoring agent object contains information describing the agent.
type Monitoring_Agent struct {
	Entity

	// The current status of the corresponding agent
	AgentStatus *Monitoring_Agent_Status `json:"agentStatus,omitempty" xmlrpc:"agentStatus"`

	// A count of all custom configuration profiles associated with the corresponding agent
	ConfigurationProfileCount *uint `json:"configurationProfileCount,omitempty" xmlrpc:"configurationProfileCount"`

	// All custom configuration profiles associated with the corresponding agent
	ConfigurationProfiles []Configuration_Template_Section_Profile `json:"configurationProfiles,omitempty" xmlrpc:"configurationProfiles"`

	// A template of an agent's current configuration which contains information about the structure of the configuration values.
	ConfigurationTemplate *Configuration_Template `json:"configurationTemplate,omitempty" xmlrpc:"configurationTemplate"`

	// Internal identifier of a configuration template that is used to configure this agent
	ConfigurationTemplateId *int `json:"configurationTemplateId,omitempty" xmlrpc:"configurationTemplateId"`

	// A count of the values associated with the corresponding Agent configuration.
	ConfigurationValueCount *uint `json:"configurationValueCount,omitempty" xmlrpc:"configurationValueCount"`

	// The values associated with the corresponding Agent configuration.
	ConfigurationValues []Monitoring_Agent_Configuration_Value `json:"configurationValues,omitempty" xmlrpc:"configurationValues"`

	// Created date
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// SoftLayer hardware related to the agent.
	Hardware *Hardware `json:"hardware,omitempty" xmlrpc:"hardware"`

	// Internal identifier of a monitoring agent
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Last modified date
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// Monitoring agent name
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// Contains general information relating to a single SoftLayer product.
	ProductItem *Product_Item `json:"productItem,omitempty" xmlrpc:"productItem"`

	// Indicates if this monitoring agent resides on your local box or on a SoftLayer monitoring cluster.
	RemoteMonitoringAgentFlag *bool `json:"remoteMonitoringAgentFlag,omitempty" xmlrpc:"remoteMonitoringAgentFlag"`

	// Internal identifier of a monitoring robot that this agent belongs to
	RobotId *int `json:"robotId,omitempty" xmlrpc:"robotId"`

	// A description for a specific installation of a Software Component
	SoftwareDescription *Software_Description `json:"softwareDescription,omitempty" xmlrpc:"softwareDescription"`

	// Internal identifier of a monitoring agent status
	StatusId *int `json:"statusId,omitempty" xmlrpc:"statusId"`

	// Monitoring agent status name.
	StatusName *string `json:"statusName,omitempty" xmlrpc:"statusName"`

	// Softlayer_Virtual_Guest object related to the monitoring agent, which this virtual guest object and hardware is on the server of the running agent.
	VirtualGuest *Virtual_Guest `json:"virtualGuest,omitempty" xmlrpc:"virtualGuest"`
}

// The SoftLayer_Monitoring_Agent_Configuration_Template_Group class is consisted of configuration templates for agents in a monitoring package.
type Monitoring_Agent_Configuration_Template_Group struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// Internal identifier of a SoftLayer account that this configuration template belongs to
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// A count of
	ConfigurationTemplateCount *uint `json:"configurationTemplateCount,omitempty" xmlrpc:"configurationTemplateCount"`

	// A count of
	ConfigurationTemplateReferenceCount *uint `json:"configurationTemplateReferenceCount,omitempty" xmlrpc:"configurationTemplateReferenceCount"`

	// no documentation yet
	ConfigurationTemplateReferences []Monitoring_Agent_Configuration_Template_Group_Reference `json:"configurationTemplateReferences,omitempty" xmlrpc:"configurationTemplateReferences"`

	// no documentation yet
	ConfigurationTemplates []Configuration_Template `json:"configurationTemplates,omitempty" xmlrpc:"configurationTemplates"`

	// Created date
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// Description of a monitoring agent configuration group
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Internal identifier of a monitoring agent configuration group
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	Item *Product_Item `json:"item,omitempty" xmlrpc:"item"`

	// Internal identifier of a configuration template type
	ItemId *int `json:"itemId,omitempty" xmlrpc:"itemId"`

	// Last modified date
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// Name of a monitoring agent configuration group
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference class holds the reference information, essentially a SQL join, between a monitoring configuration group and agent configuration templates.
type Monitoring_Agent_Configuration_Template_Group_Reference struct {
	Entity

	// no documentation yet
	ConfigurationTemplate *Configuration_Template `json:"configurationTemplate,omitempty" xmlrpc:"configurationTemplate"`

	// Internal identifier of a configuration template
	ConfigurationTemplateId *int `json:"configurationTemplateId,omitempty" xmlrpc:"configurationTemplateId"`

	// Internal identifier of a configuration group reference record
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	TemplateGroup *Monitoring_Agent_Configuration_Template_Group `json:"templateGroup,omitempty" xmlrpc:"templateGroup"`

	// Internal identifier of a monitoring agent configuration group
	TemplateGroupId *int `json:"templateGroupId,omitempty" xmlrpc:"templateGroupId"`
}

// Monitoring agent configuration value
type Monitoring_Agent_Configuration_Value struct {
	Entity

	// Internal identifier of a monitoring Agent that this configuration value
	AgentId *int `json:"agentId,omitempty" xmlrpc:"agentId"`

	// Internal identifier of a monitoring configuration definition by which
	ConfigurationDefinitionId *int `json:"configurationDefinitionId,omitempty" xmlrpc:"configurationDefinitionId"`

	// no documentation yet
	Definition *Configuration_Template_Section_Definition `json:"definition,omitempty" xmlrpc:"definition"`

	// User-friendly description of a configuration value.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Internal identifier of a monitoring configuration value.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The metric data type used to retrieve metric data currently being tracked.
	MetricDataType *Container_Metric_Data_Type `json:"metricDataType,omitempty" xmlrpc:"metricDataType"`

	// no documentation yet
	MonitoringAgent *Monitoring_Agent `json:"monitoringAgent,omitempty" xmlrpc:"monitoringAgent"`

	// no documentation yet
	Profile *Configuration_Template_Section_Profile `json:"profile,omitempty" xmlrpc:"profile"`

	// Internal identifier of a configuration profile. Configuration profile is associated with a configuration section type of "Template section".
	//
	// A "Template section" defines skeleton configuration definitions. For instance, if you want to monitor additional hard disks with "CPU/Memory/Disk Monitoring Agent", you will have to add a new configuration profiles.
	ProfileId *int `json:"profileId,omitempty" xmlrpc:"profileId"`

	// Configuration value
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// Monitoring agent status
type Monitoring_Agent_Status struct {
	Entity

	// Description of a monitoring agent status
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Internal identifier of a monitoring agent status
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Monitoring agent status name
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// The SoftLayer_Monitoring_Robot data type contains general information relating to a monitoring robot.
type Monitoring_Robot struct {
	Entity

	// The account associated with the corresponding robot.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// Internal identifier of a SoftLayer account that this robot belongs to
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// Internal identifier of a monitoring robot
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A count of the program (monitoring agent) that gets details of a system or application and reporting of the metric data and triggers alarms for predefined events.
	MonitoringAgentCount *uint `json:"monitoringAgentCount,omitempty" xmlrpc:"monitoringAgentCount"`

	// The program (monitoring agent) that gets details of a system or application and reporting of the metric data and triggers alarms for predefined events.
	MonitoringAgents []Monitoring_Agent `json:"monitoringAgents,omitempty" xmlrpc:"monitoringAgents"`

	// Robot name
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// The current status of the robot.
	RobotStatus *Monitoring_Robot_Status `json:"robotStatus,omitempty" xmlrpc:"robotStatus"`

	// The SoftLayer_Software_Component that corresponds to the robot installation on the server.
	SoftwareComponent *Software_Component `json:"softwareComponent,omitempty" xmlrpc:"softwareComponent"`

	// Internal identifier of a monitoring robot status
	StatusId *int `json:"statusId,omitempty" xmlrpc:"statusId"`
}

// Your monitoring robot will be in "Active" status under normal circumstances. If you perform an OS reload, your robot will be in "Reclaim" status until it's reloaded on your server or virtual server.
//
// Advanced monitoring system requires "Nimsoft Monitoring (Advanced)" service running and TCP ports 48000 - 48020 to be open on your server or virtual server. Monitoring agents cannot be managed nor can the usage data be updated if these ports are closed. Your monitoring robot will be in "Limited Connectivity" status if our monitoring management system cannot communicate with your system.
//
// See [[SoftLayer_Monitoring_Robot::resetStatus|resetStatus]] service for more details.
type Monitoring_Robot_Status struct {
	Entity

	// Monitoring robot status description
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Internal identifier of a monitoring robot status
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Monitoring robot status name
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}
