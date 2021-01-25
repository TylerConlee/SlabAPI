// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CustomField struct {
	ID    *int    `json:"ID"`
	Value *string `json:"Value"`
}

type OrgFields struct {
	SLALevel string `json:"SLALevel"`
}

type Organization struct {
	URL                string       `json:"URL"`
	ID                 int          `json:"ID"`
	Name               string       `json:"Name"`
	CreatedAt          string       `json:"CreatedAt"`
	UpdatedAt          string       `json:"UpdatedAt"`
	DomainNames        []string     `json:"DomainNames"`
	Tags               []string     `json:"Tags"`
	OrganizationFields []*OrgFields `json:"OrganizationFields"`
}

type Ticket struct {
	URL            string         `json:"URL"`
	ID             int            `json:"ID"`
	CreatedAt      string         `json:"CreatedAt"`
	UpdatedAt      string         `json:"UpdatedAt"`
	Subject        string         `json:"Subject"`
	Description    string         `json:"Description"`
	Priority       string         `json:"Priority"`
	Status         string         `json:"Status"`
	RequesterID    int            `json:"RequesterID"`
	OrganizationID int            `json:"OrganizationID"`
	GroupID        int            `json:"GroupID"`
	Tags           []string       `json:"Tags"`
	CustomFields   []*CustomField `json:"CustomFields"`
	SLA            string         `json:"SLA"`
}

type Tickets struct {
	Tickets []*Ticket `json:"Tickets"`
	Count   int       `json:"Count"`
}

type Trigger struct {
	URL         string             `json:"URL"`
	ID          int                `json:"ID"`
	Title       string             `json:"Title"`
	RawTitle    string             `json:"RawTitle"`
	Position    int                `json:"Position"`
	Active      bool               `json:"Active"`
	Conditions  *TriggerConditions `json:"Conditions"`
	Actions     []*TriggerAction   `json:"Actions"`
	Description string             `json:"Description"`
	UpdatedAt   string             `json:"UpdatedAt"`
	CreatedAt   string             `json:"CreatedAt"`
}

type TriggerAction struct {
	Field string `json:"Field"`
	Value string `json:"Value"`
}

type TriggerCondition struct {
	Field    string `json:"Field"`
	Operator string `json:"Operator"`
	Value    string `json:"Value"`
}

type TriggerConditions struct {
	Any []*TriggerCondition `json:"Any"`
	All []*TriggerCondition `json:"All"`
}

type Triggers struct {
	Triggers []*Trigger `json:"Triggers"`
	Count    int        `json:"Count"`
}

type View struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Active      bool   `json:"Active"`
	Description string `json:"Description"`
	CreatedAt   string `json:"CreatedAt"`
	UpdatedAt   string `json:"UpdatedAt"`
}

type ViewCount struct {
	ViewID int    `json:"ViewID"`
	URL    string `json:"URL"`
	Value  int    `json:"Value"`
	Pretty string `json:"Pretty"`
	Fresh  bool   `json:"Fresh"`
}

type Views struct {
	Views []*View `json:"Views"`
	Count int     `json:"Count"`
}

type ZendeskConfig struct {
	User   string `json:"user"`
	Apikey string `json:"apikey"`
	URL    string `json:"url"`
}

type ZendeskConfigInput struct {
	User   string `json:"user"`
	Apikey string `json:"apikey"`
	URL    string `json:"url"`
}
