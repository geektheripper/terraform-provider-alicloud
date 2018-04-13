package cms

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AlarmHistoryInDescribeAlarmHistory is a nested struct in cms response
type AlarmHistoryInDescribeAlarmHistory struct {
	Id              string                              `json:"Id" xml:"Id"`
	AlertName       string                              `json:"AlertName" xml:"AlertName"`
	GroupId         string                              `json:"GroupId" xml:"GroupId"`
	Namespace       string                              `json:"Namespace" xml:"Namespace"`
	MetricName      string                              `json:"MetricName" xml:"MetricName"`
	Dimensions      string                              `json:"Dimensions" xml:"Dimensions"`
	Expression      string                              `json:"Expression" xml:"Expression"`
	EvaluationCount int                                 `json:"EvaluationCount" xml:"EvaluationCount"`
	Value           string                              `json:"Value" xml:"Value"`
	AlertTime       int                                 `json:"AlertTime" xml:"AlertTime"`
	LastTime        int                                 `json:"LastTime" xml:"LastTime"`
	Level           string                              `json:"Level" xml:"Level"`
	PreLevel        string                              `json:"PreLevel" xml:"PreLevel"`
	RuleName        string                              `json:"ruleName" xml:"ruleName"`
	State           string                              `json:"State" xml:"State"`
	Status          int                                 `json:"Status" xml:"Status"`
	UserId          string                              `json:"UserId" xml:"UserId"`
	Webhooks        string                              `json:"Webhooks" xml:"Webhooks"`
	ContactGroups   ContactGroupsInDescribeAlarmHistory `json:"ContactGroups" xml:"ContactGroups"`
	Contacts        Contacts                            `json:"Contacts" xml:"Contacts"`
	ContactALIIMs   ContactALIIMs                       `json:"ContactALIIMs" xml:"ContactALIIMs"`
	ContactSmses    ContactSmses                        `json:"ContactSmses" xml:"ContactSmses"`
	ContactMails    ContactMails                        `json:"ContactMails" xml:"ContactMails"`
}
