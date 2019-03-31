package main

//ErrorMsg estructura del Mensaje de ErrorMsg
type ErrorMsg struct {
	TrackID        int          `json:"trackId"`
	CpnID          int          `json:"cpnId"`
	TrackType      string       `json:"trackType"`
	CreateDate     int          `json:"createDate"`
	ResolutionDate int          `json:"resolutionDate"`
	TrackState     int          `json:"trackState"`
	UsrID          int          `json:"usrId"`
	UsrName        string       `json:"usrName"`
	TicketID       int          `json:"ticketId"`
	TrackData      TrackData    `json:"trackData"`
	InstanceData   InstanceData `json:"instanceData"`
}

//TrackData estructura
type TrackData struct{
  Msg string `json:"msg"`
}

//InstanceData estructura
type InstanceData struct{
  CpnName         string `json:"cpnName"`
  LegalName       string `json:"legalName"`
  CpnCode         string `json:"cpnCode"`
  ContactName     string `json:"contactName"`
  ContactEmail    string `json:"contactEmail"`
  ContactPhone    string `json:"contactPhone"`
  LegalAgentName  string `json:"legalAgentName"`
  LegalAgentCode  string `json:"legalAgentCode"`
  LegalAgentEmail string `json:"legalAgentEmail"`
}
