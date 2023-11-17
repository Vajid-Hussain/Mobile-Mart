package middlewire

import (
	"fmt"
	"net/http"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type TokenRequirement struct {
	JwtTokenUseCase interfaceUseCase.IJwtTokenUseCase
	securityKeys    config.Token
}

var token TokenRequirement

func NewJwtTokenMiddleWire(jwtUseCase interfaceUseCase.IJwtTokenUseCase, keys config.Token) {
	token = TokenRequirement{JwtTokenUseCase: jwtUseCase, securityKeys: keys}
}

func SellerAuthorization(c *gin.Context) {
	accessToken := c.Request.Header.Get("authorization")
	refreshToken := c.Request.Header.Get("refreshtoken")

	id, err := service.VerifyAcessToken(accessToken, token.securityKeys.SellerSecurityKey)
	if err != nil {
		err := service.VerifyRefreshToken(refreshToken, token.securityKeys.SellerSecurityKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
			c.Abort()
		} else {
			_, err := token.JwtTokenUseCase.GetDataForCreteAccessToken(id)
			if err != nil {
				c.JSON(http.StatusUnauthorized, err.Error())
				c.Abort()
			} else {
				token, err := service.GenerateAcessToken(token.securityKeys.SellerSecurityKey, id)
				if err != nil {
					c.JSON(http.StatusUnauthorized, err.Error())
					c.Abort()
				} else {
					c.JSON(http.StatusOK, gin.H{"token": token})
					c.Set("SellerID", id)
					c.Next()
				}
			}
		}
	} else {
		c.JSON(http.StatusOK, "all perfect, your access token is uptodate")
		c.Set("SellerID", id)
		fmt.Println("access token is upto date")
		c.Next()
	}
	c.Set("SellerID", id)
	c.Next()
}

func UserAuthorization(c *gin.Context) {
	accessToken := c.Request.Header.Get("authorization")
	refreshToken := c.Request.Header.Get("refreshtoken")

	id, err := service.VerifyAcessToken(accessToken, token.securityKeys.UserSecurityKey)
	if err != nil {
		err := service.VerifyRefreshToken(refreshToken, token.securityKeys.UserSecurityKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
			c.Abort()
		} else {
			_, err := token.JwtTokenUseCase.GetStatusOfUser(id)
			if err != nil {
				c.JSON(http.StatusUnauthorized, err.Error())
				c.Abort()
			} else {
				token, err := service.GenerateAcessToken(token.securityKeys.UserSecurityKey, id)
				if err != nil {
					c.JSON(http.StatusUnauthorized, err.Error())
					c.Abort()
				} else {
					c.JSON(http.StatusOK, gin.H{"token": token})
					c.Next()
				}
			}
		}
	} else {
		c.JSON(http.StatusOK, "all perfect, your access token is uptodate")
		fmt.Println("access token is upto date")
	}
	c.Next()
}

func AdminAuthorization(c *gin.Context) {
	adminToken := c.GetHeader("Authorization")

	err := service.VerifyRefreshToken(adminToken, token.securityKeys.AdminSecurityKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		c.Abort()
	}
	c.Next()
}
