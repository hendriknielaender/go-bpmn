// Package bpmn implements definitions and utilities
// for the Business Process Model and Notation 2.0
package bpmn

import (
	"encoding/xml"
)

type Definitions struct {
	XMLName            xml.Name `xml:"definitions"`
	Id                 string   `xml:"id,attr"`
	Name               string   `xml:"name,attr"`
	TargetNamespace    string   `xml:"targetNamespace,attr"`
	ExpressionLanguage string   `xml:"expressionLanguage,attr"`
	TypeLanguage       string   `xml:"typeLanguage,attr"`
	Exporter           string   `xml:"exporter,attr"`
	ExporterVersion    string   `xml:"exporterVersion,attr"`
	Process            Process  `xml:"process"`
}

type Process struct {
	XMLName                      xml.Name           `xml:"process"`
	Id                           string             `xml:"id,attr"`
	Name                         string             `xml:"name,attr"`
	ProcessType                  string             `xml:"processType,attr"`
	IsClosed                     bool               `xml:"isClosed,attr"`
	IsExecutable                 bool               `xml:"isExecutable,attr"`
	DefinitionalCollaborationRef string             `xml:"definitionalCollaborationRef,attr"`
	StartEvents                  []StartEvent       `xml:"startEvent"`
	EndEvents                    []EndEvent         `xml:"endEvent"`
	SequenceFlows                []SequenceFlow     `xml:"sequenceFlow"`
	ServiceTasks                 []ServiceTask      `xml:"serviceTask"`
	ExclusiveGateway             []ExclusiveGateway `xml:"exclusiveGateway"`
}

type StartEvent struct {
	XMLName             xml.Name `xml:"startEvent"`
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IsInterrupting      bool     `xml:"isInterrupting,attr"`
	ParallelMultiple    bool     `xml:"parallelMultiple,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

type EndEvent struct {
	XMLName             xml.Name `xml:"endEvent"`
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

type SequenceFlow struct {
	XMLName             xml.Name     `xml:"sequenceFlow"`
	Id                  string       `xml:"id,attr"`
	Name                string       `xml:"name,attr"`
	SourceRef           string       `xml:"sourceRef,attr"`
	TargetRef           string       `xml:"targetRef,attr"`
	ConditionExpression []Expression `xml:"conditionExpression"`
}

type Expression struct {
	XMLName xml.Name `xml:"conditionExpression"`
	Text    string   `xml:",innerxml"`
}

type ServiceTask struct {
	XMLName             xml.Name `xml:"serviceTask"`
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	Default             string   `xml:"default,attr"`
	CompletionQuantity  int      `xml:"completionQuantity,attr"`
	IsForCompensation   bool     `xml:"isForCompensation,attr"`
	OperationRef        string   `xml:"operationRef,attr"`
	Implementation      string   `xml:"implementation,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

type ExclusiveGateway struct {
	XMLName             xml.Name `xml:"exclusiveGateway"`
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}
