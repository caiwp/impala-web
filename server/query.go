package server

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/koblas/impalathing"
	"github.com/utrack/gin-csrf"
)

func Query(c *gin.Context) {
	logrus.Info("query")
	t := csrf.GetToken(c)

	c.HTML(200, "query", gin.H{
		"token": t,
	})
}

func PostQuery(c *gin.Context) {
	var s string
	q := c.PostForm("query")
	logrus.Info(q)

	res, err := impalaQuery(q)
	if err != nil {
		s = err.Error()
		logrus.Error(err)
	} else {
		s = processRes(res)
	}

	t := csrf.GetToken(c)

	c.HTML(200, "query", gin.H{
		"query": q,
		"res":   s,
		"token": t,
	})
}

func processRes(sl []map[string]interface{}) string {
	var s string
	for _, v := range sl {
		for key, value := range v {
			s = s + key + " : " + fmt.Sprint(value) + " | "
		}
		s += "\n"
	}
	return s
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
