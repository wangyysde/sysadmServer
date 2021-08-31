package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"github.com/wangyysde/sysadmServer"
)

var roomManager *Manager

func main() {
	roomManager = NewRoomManager()
	router := sysadmServer.Default()
	router.SetHTMLTemplate(html)

	router.GET("/room/:roomid", roomGET)
	router.POST("/room/:roomid", roomPOST)
	router.DELETE("/room/:roomid", roomDELETE)
	router.GET("/stream/:roomid", stream)

	router.Run(":8080")
}

func stream(c *sysadmServer.Context) {
	roomid := c.Param("roomid")
	listener := roomManager.OpenListener(roomid)
	defer roomManager.CloseListener(roomid, listener)

	clientGone := c.Writer.CloseNotify()
	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case message := <-listener:
			c.SSEvent("message", message)
			return true
		}
	})
}

func roomGET(c *sysadmServer.Context) {
	roomid := c.Param("roomid")
	userid := fmt.Sprint(rand.Int31())
	c.HTML(http.StatusOK, "chat_room", sysadmServer.H{
		"roomid": roomid,
		"userid": userid,
	})
}

func roomPOST(c *sysadmServer.Context) {
	roomid := c.Param("roomid")
	userid := c.PostForm("user")
	message := c.PostForm("message")
	roomManager.Submit(userid, roomid, message)

	c.JSON(http.StatusOK, sysadmServer.H{
		"status":  "success",
		"message": message,
	})
}

func roomDELETE(c *sysadmServer.Context) {
	roomid := c.Param("roomid")
	roomManager.DeleteBroadcast(roomid)
}
