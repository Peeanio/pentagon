package cmd

import (
  "fmt"
  "encoding/json"
  "log"
  "io/ioutil"
  "net/http"
  "strconv"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "github.com/gin-gonic/gin"
)

type Grid struct {
	Width int `json:"width"`
	Height int `json:"height"`
	Locations []Location `json:"locations"`
	Units []Unit `json:"units"`
}

type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
	Name string `json:"name"`
	Notes []string `json:"notes"`
}

type Unit struct {
	X int `json:"x"`
	Y int `json:"y"`
	Type string `json:"type"` // infantry, armor, artillery, etc.
	HP int `json:"hp"`      // hit points
	Moves int `json:"moves"` // number of moves left this turn
}


func init() {
  rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
  Use:   "server",
  Short: "server, the Pentagon http endpoint",
  Long:  `server, the Pentagon http endpoint, using configuration values defined in config file`,
  Run: func(cmd *cobra.Command, args []string) {
    startHttpServer()
  },
}

func loadData() Grid{
    data, err := ioutil.ReadFile("pentagon_source.json")
    if err != nil {
        fmt.Println(err)
    }
    var grid Grid
    err = json.Unmarshal(data, &grid)
    if err != nil {
	fmt.Printf("unmarshal error: %+v\n", err)
    }
    return grid
}

func startHttpServer() {
    fmt.Println("HTTP API - pentagon v1")
    router := gin.Default()
    listen_port := viper.GetString("server_port")
    grid := loadData()
    loginRouter(router)
    locationRouter(router, grid.Locations)
    unitRouter(router, grid.Units)
    log.Fatal(router.Run(listen_port))
}

func loginRouter(router *gin.Engine) {
    //deals with login cookies

    router.GET("/login", func(c *gin.Context){
        cookie, err := c.Cookie("gin_cookie")

        if err != nil {
            cookie = "NotSet"
            c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
        }
        fmt.Printf("Cookie value: %s \n", cookie)
    })
}

func unitRouter(router *gin.Engine, units []Unit) {
	router.GET("/units/:id", func(c *gin.Context) {
		fmt.Println("Endpoint Hit: returnUnit")
		calledId := c.Params.ByName("id")
		idInt, err := strconv.Atoi(calledId)
		if err != nil {
			fmt.Println(err)
		}
		unit, err := json.Marshal(units[idInt])
		c.String(http.StatusOK, string(unit))
	})
}

func locationRouter(router *gin.Engine, locations []Location) {
    //deals with locations
    router.GET("/locations/:id", func(c *gin.Context) {
        fmt.Println("Endpoint hit: returnLocation")
        calledId := c.Params.ByName("id")
        idInt, err := strconv.Atoi(calledId)
        if err != nil {
            fmt.Println(err)
        }
	location, err := json.Marshal(locations[idInt])
	c.String(http.StatusOK, string(location))
    })

    // router.PUT("/locations/:id", func(c *gin.Context) {
    //     fmt.Println("Endpoint hit: updateLocation")
    //     var json Location
    //     calledId := c.Params.ByName("id")
    //     idInt, err := strconv.Atoi(calledId)
    //     if err != nil {
    //         fmt.Println(err)
    //     }
    //     if err := c.ShouldBindJSON(&json); err != nil {
    //         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    //         return
    //     }
    //     Locations[idInt] = Location{ID: json.ID, Title: json.Title, Agent: Agent{CodeName: json.Agent.CodeName, RealName: json.Agent.RealName}, Status: json.Status}
    //     c.String(http.StatusOK, "thanks for calling!")
    // })

}
