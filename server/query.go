package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/koblas/impalathing"
)

func Query(c *gin.Context) {
	logrus.Info("query")
	c.HTML(200, "query", gin.H{})
}

func PostQuery(c *gin.Context) {
	q := c.PostForm("query")
	logrus.Info(q)
	res, err := impalaQuery(q)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(res)
	c.HTML(200, "query", gin.H{
		"query": q,
		"res":   q,
	})
}

func impalaQuery(query string) ([]map[string]interface{}, error) {
	conn, err := impalathing.Connect("192.168.1.112", 21000, impalathing.DefaultOptions)
	if err != nil {
		return nil, err
	}
	_, err = conn.Query("USE gameadmindw")
	if err != nil {
		return nil, err
	}
	var res impalathing.RowSet
	res, err = conn.Query(query)
	if err != nil {
		return nil, err
	}

	return res.FetchAll(), nil
}
