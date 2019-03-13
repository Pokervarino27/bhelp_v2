package main

//ErrorMsg estructura del Mensaje de ErrorMsg
type ErrorMsg struct{
  TrackId         int     `json:"trackId,omitempty"`
  CpnId           int     `json:"cpnId,omitempty"`
  TrackType       string  `json:"trackType,omitempty"`
  CreateDate      int     `json:"createDate,omitempty"`
  ResolutionDate  int     `json:"resolutionDate,omitempty"`
  TrackState      int     `json:"trackState,omitempty"`
  UsrId           int     `json:"usrId,omitempty"`
  UsrName         string  `json:"usrName,omitempty"`
  ticketId        int     `json:"ticketId,omitempty"`
  TrackData
  InstanceData
}

//TrackData estructura
type TrackData struct{
  Msg string `json: "msg,omitempty"`
}

//InstanceData estructura
type InstanceData struct{
  CpnName         string `json:"cpnName,omitempty"`
  LegalName       string `json:"legalName,omitempty"`
  CpnCode         string `json:"cpnCode,omitempty"`
  ContactName     string `json:"contactName,omitempty"`
  ContactEmail    string `json:"contactEmail,omitempty"`
  ContactPhone    string `json:"contactPhone,omitempty"`
  LegalAgentName  string `json:"legalAgentName,omitempty"`
  LegalAgentCode  string `json:"legalAgentCode,omitempty"`
  LegalAgentEmail string `json:"legalAgentEmail,omitempty"`
}
