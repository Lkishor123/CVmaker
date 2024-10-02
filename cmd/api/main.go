package main

import (
	ltex "CVmaker/internal/ltexmod"
	"fmt"
)

func main() {
	// router := gin.Default()

	// router.POST("/generate", func(c *gin.Context) {
	// var resume Resume
	// if err := c.ShouldBindJSON(&resume); err != nil {
	// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// return
	// }
	rusme, err := ltex.ParseJSON("data.json")
	if err != nil {
		fmt.Println(err)
	}
	err = ltex.GenerateLaTeX(rusme, "template.tex", "output.tex")
	if err != nil {
		fmt.Println(err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	return
	}

	err = ltex.CompilePDF("output.tex")
	if err != nil {
		fmt.Println(err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	return
	}

	// c.FileAttachment("output.pdf", "resume.pdf")
	// })

	// router.Run(":8080")
}
