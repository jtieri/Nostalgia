package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Recording struct {
	gorm.Model
	Source   string
	WOC      bool
	Year     uint
	Month    uint
	Day      uint
	Networks string
	Contents []Content `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments string
}

type RecordingInput struct {
	Source   string         `json:"source"`
	WOC      bool           `json:"woc"`
	Date     Date           `json:"date"`
	Networks []string       `json:"networks"`
	Contents []ContentInput `json:"contents"`
	Comments string         `json:"comments"`
	Archived bool           `json:"archived"`
}

type Date struct {
	Year  uint `json:"year"`
	Month uint `json:"month"`
	Day   uint `json:"day"`
}

type Content struct {
	gorm.Model
	RecordingID   uint
	Type          uint
	Title         string
	EpisodeTitles string
}

type ContentInput struct {
	Type          contentType `json:"type"`
	Title         string      `json:"title"`
	EpisodeTitles []string    `json:"episodes"`
}

type contentType int

const (
	Show contentType = iota
	Movie
)

func (t contentType) String() string {
	switch t {
	case Show:
		return "Show"
	case Movie:
		return "Movie"
	default:
		return fmt.Sprintf("%d", int(t))
	}
}

/*
	{
    "recording": {
        "source": "Where the files came from - URL is best",
        "woc": "Does it contain commercials or not? This should be true or false",
        "date": {
			"year": "Year, if unknown set to 0",
			"month": "Month, if unknown set to 0",
			"day": "Day, if unknown set to 0"
		}
        "networks": [
			"Network name, if more than one add an entry for each network",
			"Network name, if more than one add an entry for each network"
		],
        "contents": [
            {
                "type": "0 - Show, 1 - Movie. These are the only valid inputs",
                "title": "Show or movie title",
                "episodes": [
                    "Names of episodes of show, leave blank for movies",
                    "Names of episodes of show, leave blank for movies"
                ]
            },
            {
                "type": "0 - Show, 1 - Movie. These are the only valid inputs",
                "title": "Show or movie title",
                "episodes": [
                    "Names of episodes of show, leave blank for movies",
                    "Names of episodes of show, leave blank for movies"
                ]
            }
        ],
        "comments": "Enter any additional comments here"
		"archived": "true if this recording has already been parsed and put in the database"
		}
	}
*/
