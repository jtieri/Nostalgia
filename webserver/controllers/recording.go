package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jtieri/Nostalgia/webserver/models"
	"github.com/jtieri/Nostalgia/webserver/repo"
	"net/http"
	"strconv"
	"strings"
)

func CreateRecording(c *gin.Context) {
	var input models.RecordingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recording := &models.Recording{
		Source:   input.Source,
		WOC:      input.WOC,
		Year:     input.Date.Year,
		Month:    input.Date.Month,
		Day:      input.Date.Day,
		Networks: Serialize(input.Networks),
		Contents: createContents(input.Contents),
		Comments: input.Comments,
	}

	err := repo.CreateRecording(recording)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recording})
}

func GetAllRecordings(c *gin.Context) {
	recordings, err := repo.GetAllRecordings()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recordings)
}

func GetRecordingsByYear(c *gin.Context) {
	idString := c.Param("year")
	id, _ := strconv.Atoi(idString)

	recordings, err := repo.GetRecordingsByYear(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recordings)
}

func GetRecordingsByNetwork(c *gin.Context) {
	network := c.Param("network")

	recordings, err := repo.GetRecordingsByNetwork(network)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recordings)
}

func createContents(contents []models.ContentInput) []models.Content {
	var contentsDB []models.Content

	for _, content := range contents {
		var ctype uint

		if content.Type == models.Show {
			ctype = 0
		} else if content.Type == models.Movie {
			ctype = 1
		}

		contentsDB = append(contentsDB, models.Content{
			Type:          ctype,
			Title:         content.Title,
			EpisodeTitles: Serialize(content.EpisodeTitles),
		})
	}

	return contentsDB
}

func Serialize(strSlice []string) string {
	var serializedStrings strings.Builder

	for i, str := range strSlice {
		if i < len(strSlice)-1 {
			serializedStrings.WriteString(str + ",")
		} else {
			serializedStrings.WriteString(str)
		}
	}

	return serializedStrings.String()
}

func Deserialize(serializedStrings string) []string {
	var strSlice []string

	var str string
	for _, letter := range serializedStrings {
		if letter == ',' {
			strSlice = append(strSlice, str)
			str = ""
		} else {
			str += string(letter)
		}
	}
	strSlice = append(strSlice, str)

	return strSlice
}
