package models

type RoverResponse struct {
	Photos []Photo `json:"photos"`
}

type Photo struct {
	ID        int    `json:"id"`
	Sol       int    `json:"sol"`
	ImgSrc    string `json:"img_src"`
	EarthDate string `json:"earth_date"`
	Camera    Camera `json:"camera"`
	Rover     Rover  `json:"rover"`
}

type Camera struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	RoverID  int    `json:"rover_id"`
}

type Rover struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	LandingDate string `json:"landing_date"`
	LaunchDate  string `json:"launch_date"`
	Status      string `json:"status"`
}
