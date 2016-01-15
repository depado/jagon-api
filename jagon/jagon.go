package jagon

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Confession is a simple confession.
type Confession struct {
	Confession string `json:"confession,omitempty"`
}

// Confess alows you to confess to the almighty Jagon !
func Confess(c *gin.Context) {
	var err error
	var confession = Confession{}

	if err = c.BindJSON(&confession); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Jagon didn't understand your confession."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"confession": confession.Confession, "status": "Jagon forgives you."})
}

// Prophet gives some information about the current prophet.
func Prophet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"current": gin.H{
			"name":  "Nojag",
			"alive": true,
		},
		"previous": "",
	})
}

// Apostles gives some information about the current apostotles.
func Apostles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"current": []gin.H{
			gin.H{"name": "Hekmon", "alive": true, "location": "In a dark underground temple, accompanied by a dwarf, an elf and a Shiba Inu."},
			gin.H{"name": "Thonain", "alive": true, "location": "In the back of a seedy, dark and smelly alley, right near a drunk's vomit from last night."},
		},
	})
}

// Praise allows you to praise the almighty Jagon.
func Praise(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": "Jagon heard you praise him. And He is pleased."})
}

// Pray allows you to pray the almighty Jagon.
func Pray(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": "Jagon heard your prayer and will take it in consideration."})
}

// Sacrifice represents a single sacrifice
type Sacrifice struct {
	What string `json:"what,omitempty"`
}

// DoSacrifice allows you to sacrifice something in the name of Jagon
func DoSacrifice(c *gin.Context) {
	var err error
	var s = Sacrifice{}

	if err = c.BindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Jagon didn't understand what you were trying to do."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"what": s.What, "status": "Jagon accepts your sacrifice. And he is pleased."})
}
