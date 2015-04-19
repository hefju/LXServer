package main
import (
    "net/http"
    "github.com/gin-gonic/gin"
    "log"
    "time"
    "github.com/hefju/LXServer/tools"
    "strconv"
)
func main(){
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "hello world")
    })
    router.POST("/", func(c *gin.Context) {
        var json CopyFileArgs
        c.Bind(&json) // This will infer what binder to use depending on the content-type header.
        log.Println("jutest:",json)
    })

    router.POST("/copy",func(c *gin.Context){
        var json CopyFileArgs
        c.Bind(&json) // This will infer what binder to use depending on the content-type header.
        //log.Println("jutest:",json)
        file:= &tools.File {}

       count,err:= file.Copy(json.Oldfile,json.Newfile)
        if err!=nil{
            c.String(http.StatusOK,"faile:"+err.Error())
        }else{
            c.String(http.StatusOK,"success:"+strconv.FormatInt(count,10))
        }

    })
    //获取服务器时间
    router.GET("/time", func(c *gin.Context) {
        c.String(http.StatusOK,time.Now().Format("2006-01-02 15:04:05"))
    })

    router.Run(":8082")
}

type LXArgs struct {
    Tool string `json:"tool" `
    Hefju string `json:"hefju" `
}
type CopyFileArgs struct {
    Oldfile string `json:"oldfile" `
    Newfile string `json:"newfile" `
}

