package monitor

type Event struct {
	Product      string `json:"product"`
	ResourceID   string `json:"resourceId"`
	Level        string `json:"level"`
	InstanceName string `json:"instanceName"`
	RegionID     string `json:"regionId"`
	GroupID      string `json:"groupId"`
	Name         string `json:"name"`
	Content      struct {
		InstanceID string `json:"instanceId"`
		Action     string `json:"action"`
	} `json:"content"`
	Status string `json:"status"`
}
