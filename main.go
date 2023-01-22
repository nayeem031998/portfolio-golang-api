package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("welcome to Nayeem PortoFolio")
	r := gin.New()
	///[Endpoints]
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"/about":       "View details about me",
			"/objective":   "View my carrer objective",
			"/academic":    "View my academic details",
			"/experience":  "View my work experience",
			"/hobbies":     "View my hobbies",
			"/socialmedia": "View my social media links",
			"/skills":      "View my skills",
			"/contact":     "View my contact details",
			"/resume":      "View my resume",
		})
	})

	r.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"about-me": "Myself Nayeem Akhtor from Golaghat,Assam,India.I am person with good morals and have a vey positive attitudes towards life. I am quite carrier oriented as well as quite adjustable in nature . I am a self motivated person and quite efficient in maintain the value of time. Expert in taking quick effective deccisions in stressfull situations.",
		})
	})
	r.GET("/objective", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Carrer-Objective ": "Desire to  work with such an organization where i can get an opportunity to prove my skills, along with dynamic growth and where every day will be a new challenge for me,as well as I can create a space for me to maintaining my individuality and to contribute my best to the organization. ",
		})
	})
	r.GET("/academic", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Academic-Details": "1: Mechanical Engineering From Board of Technical Education Uttar Pradesh Lucknow in 2021, 2: Higher Secondary  In SCIENCE Stream from Assam Higher Secondary Education Council In  2017, 3: High School from Board of Secondary Education Assam in 2015,4: Six Month Diploma IN Customer Application From NCH INNSTITUTIONS GOLAGHAT IN 2016, 5: diploma in AUTOCAD from Tool Room and Training Centre MSME GUWAHATI,6: Diploma in C++ Training Program from TCA PVT LTD NOIDA ",
		})
	})
	r.GET("/experience", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Work-Experience": "1: GOLANG Developer at FITPASS PVT LTD NEW DELHI,2: API Developer,3: PostgreSQL Developer,4: THREE MONTH C++ TRAING PROGRAM IN TCA PVT LTD NOIDA, Summer internship AT NUMALIGARH REFINARY LTD ASSAM IN Heat Exchanger Inspection",
		})
	})
	r.GET("/hobbies", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hobbies": "1: Travel,2: Blogging,3: Codding,3:Gaming",
		})
	})
	r.GET("/socialmedia", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"SOCIAL-MEDIA": "1: FACEBOOK- https://www.facebook.com/profile.php?id=100009394144204,2:INSTAGRAM - _nayeem_akhtar",
		})
	})

	r.GET("/skills", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"my-skills": "C++, GOLANG, POSTGRES, SQL, FLUTTER, GORM,GO.GIN,FIREBASE",
		})

	})
	r.GET("/contact", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"my-contact-info": "Email:nayeemakhtar371@gmail.com, Phone: 9954627635, Address: SHADIPUR ,NEW DELHI India",
		})
	})
	r.GET("/resume", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"my-resume": "https://firebasestorage.googleapis.com/v0/b/my-resume-a0f38.appspot.com/o/resume%20(1).pdf?alt=media&token=3791f063-468a-4aa5-9836-f060961f048a",
		})
	})
	///[Endpoints]
	r.Run(":4567")
}
