
# maps

## Index

- [type BrocadeMaps](#type-brocademaps)
- [type DashboardMisc](#type-dashboardmisc)
- [type DashboardRule](#type-dashboardrule)
- [type Group](#type-group)
- [type MAPSConfig](#type-mapsconfig)
- [type MAPSPolicy](#type-mapspolicy)
- [type MonitoringSystemMatrix](#type-monitoringsystemmatrix)
- [type PausedCfg](#type-pausedcfg)
- [type Rule](#type-rule)
- [type SwitchStatusPolicyReport](#type-switchstatuspolicyreport)
- [type SystemResources](#type-systemresources)
- [type TestEmail](#type-testemail)


## type [BrocadeMaps](<brocadeMaps.go#L5>)
```go
type BrocadeMaps struct {
	XMLName	xml.Name	`json:"-" xml:"brocade-maps"`
	// The Switch Status Policy report. The SSP report
	// provides the overall health status of the switch and
	// includes enough information for you to investigate
	// further, if necessary.
	SwitchStatusPolicyReport	SwitchStatusPolicyReport	`json:"switch-status-policy-report" xml:"switch-status-policy-report"`
	// The system resources (such as CPU, RAM, and flash
	// memory usage). Note that usage is not real time and
	// may be delayed up to 2 minutes.
	SystemResources	SystemResources	`json:"system-resources,omitempty" xml:"system-resources"`
	// A list of elements or element groups that you want
	// to pause or resume monitoring. You can only pause or
	// resume monitoring of ports, FCIP circuits, or SFPs.
	PausedCfg	[]PausedCfg	`json:"paused-cfg,omitempty" xml:"paused-cfg"`
	// A list of groups to be monitored
	Group	[]Group	`json:"group" xml:"group"`
	// The MAPS configuration for the switch. You can
	// perform the following MAPS configurations using
	// this container:
	//  • View the current MAPS configuration.
	//  • Define the actions to take on the switch when a threshold is triggered.
	//  • Specify the e-mail addresses to which the alerts are sent.
	//  • Delete all user-defined MAPS configurations related to rules, groups, policies, and so on.
	MAPSConfig	MAPSConfig	`json:"maps-config" xml:"maps-config"`
	// A list of dashboards. The dashboard enables you to
	// view the events or rules triggered and the objects
	// on which the rules were triggered over a specified
	// period of time
	DashboardRule	[]DashboardRule	`json:"dashboard-rule" xml:"dashboard-rule"`
	// The dashboard miscellaeous infomation
	// (such as start time and operation).
	DashboardMisc	DashboardMisc	`json:"dashboard-misc" xml:"dashboard-misc"`
	// A list of rules. You can use the rules container to
	// configure and manage MAPS monitoring rules and to
	// display configured rules. A rule associates a
	// condition (threshold) with actions that are
	// triggered when the specified condition is reached. A
	// rule must be in the enabled MAPS policy to be
	// active. When you modify a rule in the enabled MAPS
	// policy, the rule does not take effect until you
	// re-enable the policy.
	Rule	[]Rule	`json:"rule" xml:"rule"`
	// The MAPS policy.
	// This container enables you to create and manage
	// monitoring policies. A MAPS policy is a set of rules
	// that define thresholds for measures and actions to
	// take when a threshold is triggered. When you enable
	// a policy, all of the rules in the policy are in
	// effect. A switch can have multiple policies.
	MAPSPolicy	[]MAPSPolicy	`json:"maps-policy" xml:"maps-policy"`
	// A list of monitoring systems.
	// Each monitoring system can support different time
	// bases, actions, and thresholds. Some monitoring
	// systems are supported only on specific systems.
	// For example, circuit or tunnel monitoring systems
	// are only supported on extension platforms.
	MonitoringSystemMatrix	[]MonitoringSystemMatrix	`json:"monitoring-system-matrix" xml:"monitoring-system-matrix"`
}
```

## type [DashboardMisc](<brocadeMaps.go#L248>)

The dashboard miscellaeous infomation
(such as start time and operation).
```go
type DashboardMisc struct {
	XMLName	xml.Name	`json:"-" xml:"dashboard-misc"`
	// The MAPS start time. MAPS is restartable service
	// which means the start time could be different than
	// system up time.
	MapsStartTime	string	`json:"maps-start-time,omitempty" xml:"maps-start-time"`
	// Whether to clear the dashboard data.
	ClearData	bool	`json:"clear-data,omitempty" xml:"clear-data"`
}
```

## type [DashboardRule](<brocadeMaps.go#L223>)

The dashboard enables you to view the events or rules
triggered and the objects on which the rules were
triggered over a specified period of time You can also
clear the dashboard data. You can view a triggered rules
list for the last 7 days.
```go
type DashboardRule struct {
	XMLName	xml.Name	`json:"-" xml:"dashboard-rule"`
	// The dashboard category. Each rule can belong to only
	// one category.
	Category	string	`json:"category,omitempty" xml:"category"`
	// The rule name.
	Name	string	`json:"name,omitempty" xml:"name"`
	// The number of times the rule was triggered for the category.
	TriggeredCount	uint32	`json:"triggered-count,omitempty" xml:"triggered-count"`
	// The date and time that the rule was last triggered.
	Timestamp	string	`json:"timestamp,omitempty" xml:"timestamp"`
	// The number of times a rule was triggered. The same
	// rule can be triggered multiple times for the same or
	// different objects. For example, if the
	// defALL_D_PORTSCRC_10 rule is triggered 20 times in
	// an hour for different objects, then the
	// repetition-count is 20.
	RepetitionCount	uint32	`json:"repetition-count,omitempty" xml:"repetition-count"`
	// The objects that violated the rule. For example,
	// port, circuit, and so on.
	Objects	[]string	`json:"objects,omitempty" xml:"objects>object"`
}
```

## type [Group](<brocadeMaps.go#L142>)

A group to be monitored
```go
type Group struct {
	XMLName	xml.Name	`json:"-" xml:"group"`
	// The group name. The group name must be unique; it is
	// not case sensitive and can contain up to 32
	// characters.
	Name	string	`json:"name,omitempty" xml:"name"`
	// The type of elements (such as, fc-port, sfp, fan,
	// and so on) contained in the group.
	GroupType	string	`json:"group-type,omitempty" xml:"group-type"`
	// The existing group feature name (such as node-wwn,
	// port-name, or unknown).
	GroupFeature	string	`json:"group-feature,omitempty" xml:"group-feature"`
	// The feature pattern. Specifies the wildcard
	// characters while defining the feature
	// characteristics. The wildcard characters "*" for any
	// string, "?" for any single character, "[expr]" for
	// one character from the set specified in the
	// expression, or '!' for negation of the string, are
	// supported. If '!' is specified in the pattern, the
	// pattern must be in single quotes. For example, if
	// you create a group where group feature is
	// "port-name" and the feature pattern is "brcdhost*",
	// the group has a membership defined as ports with a
	// port name that begins with "brcdhost."
	FeaturePattern	string	`json:"feature-pattern,omitempty" xml:"feature-pattern"`
	// Whether the group is a system-defined or
	// user-defined group. You cannot delete system defined
	// groups; however, you can augment the ports managed
	// by the group.
	IsPredefined	bool	`json:"is-predefined,omitempty" xml:"is-predefined"`
	// Whether you can modify the group. You can modify all
	// user-defined groups, and predefined port type
	// groups, except for the ALL_QUARANTINED_PORTS group
	// and any flow group (although the parameter is true).
	IsModifiable	bool	`json:"is-modifiable,omitempty" xml:"is-modifiable"`
	// A list of members (groups)
	// (such as, node-wwn or port-name).
	// This parameter is available only when the group
	// feature parameter is node WWN or port name
	// (group-feature=node-wwn or port-name). There must
	// be at least one member in the list.
	Members	[]string	`json:"members,omitempty" xml:"members>member"`
}
```

## type [MAPSConfig](<brocadeMaps.go#L188>)

The MAPS actions. You can enable one or more actions
globally at the switch level or per rule.
```go
type MAPSConfig struct {
	XMLName	xml.Name	`json:"-" xml:"maps-config"`
	// The MAPS actions. You can enable one or more actions
	// globally at the switch level or per rule.
	Actions	[]string	`json:"actions,omitempty" xml:"actions>action"`
	// The decommission behavior (with-disable or impair).
	// The default is with-disable.
	DecommisionCfg	string	`json:"decommision-cfg,omitempty" xml:"decommision-cfg"`
	// The recipient list for e-mail alerts.
	RecipientAddressList	[]string	`json:"recipient-address-list,omitempty" xml:"recipient-address-list>recipient-address"`
	// The e-mail address of the sender.
	SenderAddress	string	`json:"sender-address,omitempty" xml:"sender-address"`
	// The domain name. Enter none to clear the name.
	DomainName	string	`json:"domain-name,omitempty" xml:"domain-name"`
	// The relay IP address. Enter none to clear the IP
	// address.
	RelayIpAddress	string	`json:"relay-ip-address,omitempty" xml:"relay-ip-address"`
	// The test e-mail container.
	TestEmail	TestEmail	`json:"test-email,omitempty" xml:"test-email"`
}
```

## type [MAPSPolicy](<brocadeMaps.go#L330>)

The MAPS policy.
This container enables you to create and manage
monitoring policies. A MAPS policy is a set of rules
that define thresholds for measures and actions to take
when a threshold is triggered. When you enable a policy,
all of the rules in the policy are in effect. A switch
can have multiple policies.
```go
type MAPSPolicy struct {
	XMLName	xml.Name	`json:"-" xml:"maps-policy"`
	// The MAPS policy name.
	// The name for the policy must be unique; it is
	// case-sensitive and can contain up to 32 characters.
	Name	string	`json:"name,omitempty" xml:"name"`
	// A list of rules in the policy
	RuleList	[]string	`json:"rule-list,omitempty" xml:"rule-list>rule"`
	// Whether the policy is active.
	// You can configure multiple policies; however, only
	// one policycan be active at a time.
	IsActivePolicy	bool	`json:"is-active-policy,omitempty" xml:"is-active-policy"`
	// Whether the policy is predefined or user-defined.
	// Fabric OS ships with 4 predefined policies:
	// dflt_conservative_policy, dflt_aggressive_policy,
	// dflt_moderate_policy, and dflt_base_policy.
	IsPredefinedPolicy	bool	`json:"is-predefined-policy,omitempty" xml:"is-predefined-policy"`
}
```

## type [MonitoringSystemMatrix](<brocadeMaps.go#L355>)

Monitoring system.
Each monitoring system can support different time bases,
actions, and thresholds. Some monitoring systems are
supported only on specific systems. For example, circuit
or tunnel monitoring systems are only supported on
extension platforms.
```go
type MonitoringSystemMatrix struct {
	XMLName	xml.Name	`json:"-" xml:"monitoring-system-matrix"`
	// The monitoring system name
	// (CRC, BLADE_STATE, ITW, and so on).
	MonitoringSystem	string	`json:"monitoring-system,omitempty" xml:"monitoring-system"`
	// The dashboard category of the monitoring system.
	DashboardCategory	string	`json:"dashboard-category,omitempty" xml:"dashboard-category"`
	// The group type.
	GroupType	string	`json:"group-type,omitempty" xml:"group-type"`
	// A list of the supported time bases for the
	// monitoring system.
	BaseTimeBases	[]string	`json:"base-time-bases,omitempty" xml:"base-time-bases>time-base"`
	// The rule on rules time bases for the monitoring
	// system. This parameter is available only when the
	// is-rule-on-rule-supported parameter is active
	// (is-rule-on-rule-supported =true)
	RuleOnRuleTimeBases	[]string	`json:"rule-on-rule-time-bases,omitempty" xml:"rule-on-rule-time-bases>rule-on-rule-time-base"`
	// Whether the monitoring system is read only.
	IsReadOnly	bool	`json:"is-read-only,omitempty" xml:"is-read-only"`
	// Whether the monitoring system is in all logical switches or only the default logical switch.
	MonitoredLogicalSwitch	string	`json:"monitored-logical-switch,omitempty" xml:"monitored-logical-switch"`
	// Whether rule on rule is supported.
	IsRuleOnRuleSupported	bool	`json:"is-rule-on-rule-supported,omitempty" xml:"is-rule-on-rule-supported"`
	// Whether quiet time is supported.
	IsQuietTimeSupported	bool	`json:"is-quiet-time-supported,omitempty" xml:"is-quiet-time-supported"`
	// The minimum quiet time in seconds.
	MinimumQuietTime	uint32	`json:"minimum-quiet-time,omitempty" xml:"minimum-quiet-time"`
	// The monitoring type (event based or poll based)
	// for the monitoring system.
	MonitoringType	string	`json:"monitoring-type,omitempty" xml:"monitoring-type"`
	// The data type support. Each monitoring system
	// supports different data types for thresholds.
	DataType	string	`json:"data-type,omitempty" xml:"data-type"`
	// A description of the monitoring system.
	Description	string	`json:"description,omitempty" xml:"description"`
	// A list of the MAPS actions defined for this
	// monitoring system.
	Actions	[]string	`json:"actions,omitempty" xml:"actions>action"`
	// The data unit of the monitoring system.
	Unit	string	`json:"unit,omitempty" xml:"unit"`
	// The data range of the monitoring system.
	DataRange	string	`json:"data-range,omitempty" xml:"data-range"`
	// The supported operations.
	LogicalOperators	[]string	`json:"logical-operators,omitempty" xml:"logical-operators>logical-operator"`
}
```

## type [PausedCfg](<brocadeMaps.go#L131>)

A list of elements or element groups that you want to
pause or resume monitoring. You can only pause or resume
monitoring of ports, FCIP circuits, or SFPs.
```go
type PausedCfg struct {
	XMLName	xml.Name	`json:"-" xml:"paused-cfg"`
	// The element group of which you want to pause or
	// resume monitoring.
	GroupType	string	`json:"group-type,omitempty" xml:"group-type"`
	// A list of elements (ports, FCIP circuits, or SFPs).
	// There must be at least one member in the list.
	Members	[]string	`json:"members,omitempty" xml:"members>member"`
}
```

## type [Rule](<brocadeMaps.go#L264>)

A rule associates a condition (threshold) with actions
that are triggered when the specified condition is
reached. A rule must be in the enabled MAPS policy to be
active. When you modify a rule in the enabled MAPS
policy, the rule does not take effect until you
re-enable the policy.
```go
type Rule struct {
	XMLName	xml.Name	`json:"-" xml:"rule"`
	// The rule name.
	Name	string	`json:"name,omitempty" xml:"name"`
	// The rule name. A rule can be one of two types: base
	// or rule-on-rule. A base rule monitors the statistics
	// or FRUs whereas a rule-on-rule monitors the base
	// rule.
	IsRuleOnRule	bool	`json:"is-rule-on-rule,omitempty" xml:"is-rule-on-rule"`
	// The statistic or error to be monitored
	// (such as CRC, ITW, PS_STATE, and so on).
	MonitoringSystem	string	`json:"monitoring-system" xml:"monitoring-system"`
	// The time interval between two samples to be compared.
	TimeBase	string	`json:"time-base,omitempty" xml:"time-base"`
	// The relational operation to be used in evaluating
	// the condition.
	LogicalOperator	string	`json:"logical-operator,omitempty" xml:"logical-operator"`
	// The threshold value.
	// Thresholds are the values at which potential
	// problems might occur. For example, in configuring a
	// port threshold, you can select a specific value at
	// which an action is triggered because of too many
	// threshold violations.
	//  • For numerical values: 0-999999999. The upper limit may vary depending on the monitoring system
	//  category.
	//  • For percentage values: 0-100.
	//  • For FRU states: ON, OFF, IN, OUT, or FAULTY.
	//  • For temperature monitoring: IN_RANGE or OUT_OF_RANGE.
	//  • For FPI states: IO_FRAME_LOSS, IO_PERF_IMPACT, IO_LATENCY_CLEAR.
	//  • For Ethernet port state: UP or DOWN
	ThresholdValue	string	`json:"threshold-value,omitempty" xml:"threshold-value"`
	// The group name.
	// The group name must be unique; it is not case
	// sensitive and can contain up to 32 characters.
	GroupName	string	`json:"group-name,omitempty" xml:"group-name"`
	// The MAPS actions.
	Actions	[]string	`json:"actions,omitempty" xml:"actions>action"`
	// Whether the group is a system-defined or
	// user-defined group. You cannot delete system-defined
	// groups; however, you can augment the ports managed
	// by the group.
	IsPredefined	bool	`json:"is-predefined,omitempty" xml:"is-predefined"`
	// The user-configured severity
	// (such as warning, error, critical, info, or default).
	EventSeverity	string	`json:"event-severity,omitempty" xml:"event-severity"`
	// The port's toggle time in seconds.
	ToggleTime	uint32	`json:"toggle-time,omitempty" xml:"toggle-time"`
	// The quiet time in seconds.
	// The rule is not triggered until quiet time has
	// expired.
	QuietTime	uint32	`json:"quiet-time,omitempty" xml:"quiet-time"`
	// Whether to clear quiet time.
	QuietTimeClear	bool	`json:"quiet-time-clear,omitempty" xml:"quiet-time-clear"`
	// The unquarantine timeout in seconds.
	UnQuarantineTimeout	uint32	`json:"un-quarantine-timeout,omitempty" xml:"un-quarantine-timeout"`
	// Whether to clear the unquarantine timeout.
	UnQuarantineClear	bool	`json:"un-quarantine-clear,omitempty" xml:"un-quarantine-clear"`
}
```

## type [SwitchStatusPolicyReport](<brocadeMaps.go#L67>)

The Switch Status Policy report provides the overall
health status of the switch
```go
type SwitchStatusPolicyReport struct {
	XMLName	xml.Name	`json:"-" xml:"switch-status-policy-report"`
	// The overall status of the switch. The switch state
	// is determined by the state of one or more of its
	// components (power supply, fan, wwn, and so on).
	//
	// (unknown | down | marginal | healthy)
	SwitchHealth	string	`json:"switch-health,omitempty" xml:"switch-health"`
	// The state of the power supplies.
	PowerSupplyHealth	string	`json:"power-supply-health,omitempty" xml:"power-supply-health"`
	// The state of the fans.
	FanHealth	string	`json:"fan-health,omitempty" xml:"fan-health"`
	// The state of the WWN cards.
	WwnHealth	string	`json:"wwn-health,omitempty" xml:"wwn-health"`
	// The state of the temperature sensors.
	TemperatureSensorHealth	string	`json:"temperature-sensor-health,omitempty" xml:"temperature-sensor-health"`
	// The state of the high availability
	// (both control processors are in sync).
	HaHealth	string	`json:"ha-health,omitempty" xml:"ha-health"`
	// The state of the control processors.
	ControlProcessorHealth	string	`json:"control-processor-health,omitempty" xml:"control-processor-health"`
	// The state of the core blades.
	CoreBladeHealth	string	`json:"core-blade-health,omitempty" xml:"core-blade-health"`
	// The state of the application blades.
	BladeHealth	string	`json:"blade-health,omitempty" xml:"blade-health"`
	// The state of the flash usage.
	FlashHealth	string	`json:"flash-health,omitempty" xml:"flash-health"`
	// The state of the marginal ports
	MarginalPortHealth	string	`json:"marginal-port-health,omitempty" xml:"marginal-port-health"`
	// The state of the faulty ports
	FaultyPortHealth	string	`json:"faulty-port-health,omitempty" xml:"faulty-port-health"`
	// The state of the SFP.
	MissingSfpHealth	string	`json:"missing-sfp-health,omitempty" xml:"missing-sfp-health"`
	// The state of the error ports.
	ErrorPortHealth	string	`json:"error-port-health,omitempty" xml:"error-port-health"`
	// The state of the certificate.
	ExpiredCertificateHealth	string	`json:"expired-certificate-health,omitempty" xml:"expired-certificate-health"`
	// The state of the airflow. This parameter monitors
	// the air flow direction of the power supply fan FRUs
	// and blower FRUs and generates an alert if there is a
	// mismatch in the air flow direction of any two power
	// supply fans or any two blowers.
	AirflowMismatchHealth	string	`json:"airflow-mismatch-health,omitempty" xml:"airflow-mismatch-health"`
}
```

## type [SystemResources](<brocadeMaps.go#L116>)

The system resources
(such as CPU, RAM, and flash memory usage).
Note that usage is not real time and may be delayed up
to 2 minutes.
```go
type SystemResources struct {
	XMLName	xml.Name	`json:"-" xml:"system-resources"`
	// The percentage of CPU usage.
	CpuUsage	uint32	`json:"cpu-usage,omitempty" xml:"cpu-usage"`
	// The percentage of memory usage.
	MemoryUsage	uint32	`json:"memory-usage,omitempty" xml:"memory-usage"`
	// The total memory usage in kilobytes.
	TotalMemory	uint32	`json:"total-memory,omitempty" xml:"total-memory"`
	// The percentage of flash usage.
	FlashUsage	uint32	`json:"flash-usage,omitempty" xml:"flash-usage"`
}
```

## type [TestEmail](<brocadeMaps.go#L210>)

The test e-mail container.
```go
type TestEmail struct {
	XMLName	xml.Name	`json:"-" xml:"test-email"`
	// The subject for the e-mail alert.
	Subject	string	`json:"subject,omitempty" xml:"subject"`
	// The message for the e-mail alert.
	Body	string	`json:"body,omitempty" xml:"body"`
}
```

