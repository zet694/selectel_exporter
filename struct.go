package main

type Price struct {
	Status string `json:"status,omitempty"`
	Data   struct {
		Primary struct {
			Main int `json:"main,omitempty"`
		}
		Storage struct {
			Main int `json:"main,omitempty"`
		}
		Vpc struct {
			Main int `json:"main,omitempty"`
		}
		Vmware struct {
			Main int `json:"main,omitempty"`
		}
	}
}
