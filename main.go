package main

import (
	"github.com/jtieri/Nostalgia/config"
	"github.com/jtieri/Nostalgia/webserver"
	"github.com/jtieri/Nostalgia/webserver/app"
	"github.com/jtieri/Nostalgia/webserver/controllers"
	"github.com/jtieri/Nostalgia/webserver/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"strconv"
)

// Media Server runs filescan of all new directories relative to the last scan.
// JSON files are parsed, recording structs are created and checked, structs get sent to webserver endpoint
// Web Server will listen at http://hostname/api/v1/tv/ for incoming post requests and try to prepare entries for DB
// -------------------------------------------------------------------------------------------------------------------
// For some reason I was getting the error below when trying to build the project on Windows 10 x64
// 		Go\pkg\tool\windows_amd64\link.exe: running gcc failed: exec: "gcc": executable file not found in %PATH%
// The issue was with sqlite3 and the solution I found on GitHub was to build from the path with the command below
// 		go build -ldflags="-linkmode=internal -extld=none" -x
// The issue doesn't seem to exist using CLI
func main() {
	c := config.Load()

	db, err := gorm.Open(sqlite.Open(c.Database.Name), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: " + err.Error())
	}
	app.WebApp = app.New(db)

	err = db.AutoMigrate(&models.Post{}, &models.Recording{}, &models.Content{})
	if err != nil {
		log.Fatal("Failed to migrate the database schemas: " + err.Error())
	}

	/*
			// Start the Media Server
			ms := mediaserver.New(c.MediaServer.Port, c.MediaServer.Host)
			err := ms.Start()
			if err != nil {
				log.Fatal("Failed to read the media directories contents due to: " + err.Error())
			}

			// Start the Stream
			err = ms.StreamQueue.Start(c.MediaServer.Host, c.MediaServer.Port)
			if err != nil {
				log.Fatal("Failed to start the stream due to: " + err.Error())
			}

		// Open our jsonFile
		jsonFile, err := os.Open("test_recording.json")

		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened users.json")

		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		// read our opened jsonFile as a byte array.
		byteValue, _ := ioutil.ReadAll(jsonFile)

		// we initialize our Users array
		var recording models.RecordingInput

		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		err = json.Unmarshal(byteValue, &recording)
		if err != nil {
			panic(err.Error())
		}
		printRecording(&recording)

		recording.Archived = true
		updateRec, _ := json.Marshal(recording)
		ioutil.WriteFile("update_recording.json", updateRec, 0644)
	*/

	// Start the Web Server
	ws := webserver.New()
	err = ws.Start(c.WebServer.Host, c.WebServer.Port)
	if err != nil {
		log.Fatal("Failed to start the web server due to:  " + err.Error())
	}
}

func testSerialize() {
	networks := []string{"CN", "AS"}
	serializedNetworks := controllers.Serialize(networks)
	log.Println(serializedNetworks)
	log.Println(controllers.Deserialize(serializedNetworks))
	log.Println("-----------------------------------------------------------")

	networks2 := []string{"Nick"}
	serializedNetworks2 := controllers.Serialize(networks2)
	log.Println(serializedNetworks2)
	log.Println(controllers.Deserialize(serializedNetworks2))
	log.Println("-----------------------------------------------------------")

	networks3 := []string{"AS", "MTV", "ABC", "CNN", "BBC"}
	serializedNetworks3 := controllers.Serialize(networks3)
	log.Println(serializedNetworks3)
	log.Println(controllers.Deserialize(serializedNetworks3))
	log.Println("-----------------------------------------------------------")

	var networks4 []string
	serializedNetworks4 := controllers.Serialize(networks4)
	log.Println(serializedNetworks4)
	log.Println(controllers.Deserialize(serializedNetworks4))
	log.Println("-----------------------------------------------------------")
}

func printBlogPost(post *models.Post) {
	log.Println("-------------------------------------------------------")
	log.Println("Created at: " + post.CreatedAt.String())
	log.Println("Updated at: " + post.UpdatedAt.String())
	//log.Println("Deleted at: " + post.DeletedAt.Time.String())
	log.Println("Post ID: " + strconv.Itoa(int(post.ID)))
	log.Println("Title: " + post.Title)
	log.Println("Body: " + post.Body)
}

func printRecording(recording *models.RecordingInput) {
	log.Println("---------------------------------------------------")
	log.Println("Source: " + recording.Source)
	log.Printf("WOC: " + strconv.FormatBool(recording.WOC))
	log.Printf("Year: %s", strconv.Itoa(int(recording.Date.Year)))
	log.Printf("Month: %s", strconv.Itoa(int(recording.Date.Month)))
	log.Printf("Day: %s", strconv.Itoa(int(recording.Date.Day)))

	for i, network := range recording.Networks {
		log.Println("Network" + strconv.Itoa(i) + ": " + network)
	}

	for i, content := range recording.Contents {
		log.Println("ContentInput" + strconv.Itoa(i) + ": " + content.Title)
		for c, episode := range content.EpisodeTitles {
			log.Println("Episode" + strconv.Itoa(c) + ": " + episode)
		}
		log.Println("ContentInput Type: " + content.Type.String())
	}

	log.Println("Comments: " + recording.Comments)
	log.Println("Archived: " + strconv.FormatBool(recording.Archived))
}
