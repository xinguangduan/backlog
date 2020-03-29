package rest

import (
	"github.com/gin-gonic/gin"
	_db "microservice/backlog/db"
	"net/http"
	"strconv"
)

const (
	JSON_SUCCESS int = 1
	JSON_ERROR   int = 0
)

var db = _db.DBUtils

// 创建TODO条目
func Add(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := _db.BacklogModel{Title: c.PostForm("title"), Completed: completed}
	db.Save(&todo)
	c.JSON(http.StatusOK, gin.H{
		"status":     JSON_SUCCESS,
		"message":    "创建成功",
		"resourceId": todo.ID,
	})
}

// 获取所有条目
func All(c *gin.Context) {
	var backlog [] _db.BacklogModel
	var _backlog []_db.TransformedBacklog
	db.Find(&backlog)

	// 没有数据
	if len(backlog) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  JSON_ERROR,
			"message": "没有数据",
		})
		return
	}

	// 格式化
	for _, item := range backlog {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_backlog = append(_backlog, _db.TransformedBacklog{
			ID:        item.ID,
			Title:     item.Title,
			Completed: completed,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "ok",
		"data":    _backlog,
	})

}

// 根据id获取一个条目
func Take(c *gin.Context) {
	var backlog _db.BacklogModel
	bID := c.Param("id")

	db.First(&backlog, bID)
	if backlog.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  JSON_ERROR,
			"message": "条目不存在",
		})
		return
	}
	completed := false
	if backlog.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_backlog := _db.TransformedBacklog{
		ID:        backlog.ID,
		Title:     backlog.Title,
		Completed: completed,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "ok",
		"data":    _backlog,
	})
}

// 更新一个条目
func Update(c *gin.Context) {
	var backlog _db.BacklogModel
	bID := c.Param("id")
	db.First(&backlog, bID)
	if backlog.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  JSON_ERROR,
			"message": "条目不存在",
		})
		return
	}

	db.Model(&backlog).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&backlog).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "更新成功",
	})
}

// 删除条目
func Del(c *gin.Context) {
	var backlog _db.BacklogModel
	bID := c.Param("id")
	db.First(&backlog, bID)
	if backlog.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  JSON_ERROR,
			"message": "条目不存在",
		})
		return
	}
	db.Delete(&backlog)
	c.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "删除成功!",
	})
}
