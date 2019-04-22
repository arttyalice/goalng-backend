package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gitlab.com/?/?/middlewares"
	"gitlab.com/?/?/models"
	"gitlab.com/?/?/utils"
)

// RegisterStaffEndpoints - to let main register these endpoints
func RegisterStaffEndpoints(router *gin.RouterGroup) {
	staffRouter := router.Group("")
	// Annonymous Routes
	staffRouter.POST("/login", staffLogin)

	// Admin Routes
	staffRouter.Use(middlewares.LoginRequire())
	{
		staffRouter.GET("", getStaffList)
		staffRouter.GET("/admins", getAdminList)
		staffRouter.GET("/check/:uname/station/:sID", checkIfUsernameExist)
		staffRouter.GET("/roles", getStaffRoleList)

	}

	staffRouter.Use(middlewares.SuperAdminRequired())
	{
		staffRouter.POST("/create", createStaff)
		staffRouter.POST("/edit/:staffID", editStaff)
		staffRouter.GET("/info/:staffID", getStaffInfomation)
	}
}

func getStaffList(c *gin.Context) {
	var staffs []models.Staff
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("bad request", http.StatusBadRequest))
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("bad request", http.StatusBadRequest))
		return
	}
	hID := c.Query("hid")
	rID := c.Query("rid")

	if err := models.GetStaffList(&staffs, size, page, hID, rID); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("internal error", http.StatusInternalServerError))
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(staffs))
}

func getAdminList (c *gin.Context) {
	var staffs []models.Staff
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("bad request", http.StatusBadRequest))
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("bad request", http.StatusBadRequest))
		return
	}
	hID := c.Query("hid")

	if err := models.GetAdminStaffList(&staffs, size, page, hID); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("internal error", http.StatusInternalServerError))
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(staffs))
}

func getStaffRoleList (c *gin.Context) {
	var typelist []models.StaffType
	if err := models.GetStaffTypeList(&typelist); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("internal error", http.StatusInternalServerError))
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(typelist))
}

func getStaffInfomation (c *gin.Context) {
	staffID := c.Param("staffID")
	var staffInfo models.StaffInfomation
	if err := models.GetStaffInfomation(&staffInfo, staffID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("internal error", http.StatusInternalServerError))
	}
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(staffInfo))
}

func checkIfUsernameExist (c *gin.Context) {
	loginName := c.Param("uname")
	var stationID uint64
	var err error
	if stationID, err = strconv.ParseUint(c.Param("sID"), 10, 32); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("bad request", http.StatusBadRequest))
		return
	}
	if isAvailable := models.CheckUserExist(loginName, stationID); isAvailable {
		c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(gin.H{
			"username":  loginName,
			"available": false,
		}))
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(gin.H{
		"username":  loginName,
		"available": true,
	}))
}

func createStaff (c *gin.Context) {
	var staff models.Staff
	c.BindJSON(&staff)
	if isDuplicated := models.FindDuplicateStaff(&staff); isDuplicated {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("duplicated user", http.StatusInternalServerError))
		return
	}
	if err := models.CreateStaff(&staff); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("internal error", http.StatusInternalServerError))
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(gin.H{
		"status": true,
	}))
}

func editStaff (c *gin.Context) {
	var staff models.Staff
	c.BindJSON(&staff)
}

func staffLogin (c *gin.Context) {
	var data models.StaffInfomation
	username := c.PostForm("login_name")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("bad request", http.StatusBadRequest))
		return
	}
	staff := models.Staff{
		LoginName:         &username,
		LoginHashPassword: &password,
	}
	isErr, isFound, err := models.StaffLogin(&staff, &data)
	if isErr {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("internal error", http.StatusInternalServerError))
		return
	}
	if !isFound {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid username or password", http.StatusInternalServerError))
		return
	}
	if data.User.UserToken == nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Fail generating token", http.StatusInternalServerError))
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(data))
}
