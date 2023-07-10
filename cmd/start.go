package cmd

import (
  "fmt"
  "log"
  "net/http"
  "strconv"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "github.com/gin-gonic/gin"
)

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

func startHttpServer() {
    fmt.Println("HTTP API - pentagon v1")
    router := gin.Default()
    listen_port := viper.GetString("server_port")
    loginRouter(router)
    locationRouter(router)
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

func locationRouter(router *gin.Engine) {
    //deals with locations
    router.GET("/locations/:id", func(c *gin.Context) {
        fmt.Println("Endpoint hit: returnLocation")
        calledId := c.Params.ByName("id")
        idInt, err := strconv.Atoi(calledId)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(idInt)
        c.String(http.StatusOK, "thanks for calling!")
    })
}
