package main

//ErrorMsg estructura del Mensaje de ErrorMsg
type ErrorMsg struct{
  TrackId         int     `json:"trackId"`
  CpnId           int     `json:"cpnId"`
  TrackType       string  `json:"trackType"`
  CreateDate      int     `json:"createDate"`
  ResolutionDate  int     `json:"resolutionDate"`
  TrackState      int     `json:"trackState"`
  UsrId           int     `json:"usrId"`
  UsrName         string  `json:"usrName"`
  TicketId        int     `json:"ticketId"`
  // TrackData
  // InstanceData
}

//TrackData estructura
type TrackData struct{
  Msg string `json: "msg"`
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
