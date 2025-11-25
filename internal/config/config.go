package config

import(
	"os"
	"encoding/json"
)

type Config struct{
	DBURL	string `json:"db_url"`
	Current_User_Name	string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"


func Read() (Config, error){
	home, err := os.UserHomeDir()
	if err != nil{
		return Config{}, err
	}
	filePath := home + "/" + configFileName
	data, err := os.ReadFile(filePath)
	if err != nil{
		return Config{}, err
	}
	
	var cfg Config

	if err := json.Unmarshal(data, &cfg); err !=nil{
		return Config{}, err
	}
	return cfg, nil
}


func (c *Config) SetUser(userName string) error {
	c.Current_User_Name = userName

	home, err := os.UserHomeDir()
	if err != nil{
		return err
	}
	filePath := home + "/" + configFileName
	jsonData, err := json.Marshal(c)
	if err != nil{
		return err
	}
	if err := os.WriteFile(filePath, jsonData, 0644); err != nil{
		return err
	}

	return nil
}