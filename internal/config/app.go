package config

type App struct {
	Name string `json:"name"`
	Port int    `json:"port"`
	UID  string `json:"uid"`
}
