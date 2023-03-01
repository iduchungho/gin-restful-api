package controller

import (
	"cse-foodcourt/src/database"
	"cse-foodcourt/src/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
)

func Login(c *gin.Context) {
	//TODO
}

func LogOut(c *gin.Context) {
	//TODO
}

func Signup(c *gin.Context) {
	// Get body off req body
	var customer model.Customer
	if c.BindJSON(&customer) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Get Cost from env file
	cost, errCost := strconv.Atoi(os.Getenv("COST"))
	if errCost != nil {
		panic("COST env not found")
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(customer.Password), cost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	// Create the customer
	newCustomer := model.Customer{
		CustomerID: customer.CustomerID,
		FirstName:  customer.FirstName,
		LastName:   customer.LastName,
		EmailID:    customer.EmailID,
		Password:   string(hash),
		PhoneNo:    customer.PhoneNo,
		State:      customer.State,
		Landmark:   customer.Landmark,
		Role:       "user",
	}
	result := database.GetConnection().Create(&newCustomer)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create new customer",
		})
		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{
		"message": "create customer successfully",
	})
}

func GetInfo(c *gin.Context) {

	var customer model.Customer
	result := database.GetConnection().Where("customer_id = ?", c.Param("id")).Find(&customer)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "customer ID not found",
		})
		return
	}

	// Get Customer info
	customerInfo := gin.H{
		"data": gin.H{
			"customer_id": customer.CustomerID,
			"first_name":  customer.FirstName,
			"last_name":   customer.LastName,
			"email_id":    customer.EmailID,
			"phone_no":    customer.PhoneNo,
			"state":       customer.State,
			"landmark":    customer.Landmark,
		},
	}

	// respond
	c.JSON(http.StatusOK, &customerInfo)
}
