package models

import (
	"encoding/json"
	"testing"
)

func TestPhotoResponse_Unmarshal(t *testing.T) {
	sampleJSON := `
	{
	  "photos": [
		{
		  "id": 102693,
		  "sol": 1004,
		  "camera": {
			"id": 20,
			"name": "FHAZ",
			"rover_id": 5,
			"full_name": "Front Hazard Avoidance Camera"
		  },
		  "img_src": "http://mars.nasa.gov/msl-raw-images/image.jpg",
		  "earth_date": "2015-06-03",
		  "rover": {
			"id": 5,
			"name": "Curiosity",
			"landing_date": "2012-08-06",
			"launch_date": "2011-11-26",
			"status": "active"
		  }
		}
	  ]
	}`

	var resp PhotoResponse
	err := json.Unmarshal([]byte(sampleJSON), &resp)
	if err != nil {
		t.Fatalf("failed to unmarshal sample JSON: %v", err)
	}

	if len(resp.Photos) != 1 {
		t.Errorf("expected 1 photo, got %d", len(resp.Photos))
	}

	if resp.Photos[0].Rover.Name != "Curiosity" {
		t.Errorf("expected rover name 'Curiosity', got '%s'", resp.Photos[0].Rover.Name)
	}
}
