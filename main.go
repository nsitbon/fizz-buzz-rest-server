package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"log"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type ApiParams struct {
	String1 string `form:"string1" binding:"required"`
	String2 string `form:"string2" binding:"required"`
	Int1 uint `form:"int1" binding:"required"`
	Int2 uint `form:"int2" binding:"required"`
	Limit uint `form:"limit" binding:"required,lte=1024"`
}

func main() {
	createHttpServer().Run()
}

func createHttpServer() *gin.Engine {
	r := gin.New()
	r.GET("/generate", generate)
	return r
}

func generate(ctx *gin.Context) {
	p, err := extractApiParamsFromRequest(ctx)

	if err == nil {
		ctx.JSON(200, generateElements(p.Limit, p.Int1, p.Int2, p.String1, p.String2))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":  err.Error()})
	}
}

func extractApiParamsFromRequest(ctx *gin.Context) (*ApiParams, error) {
	apiParams := &ApiParams{}
	return apiParams, ctx.ShouldBindWith(apiParams, binding.Form)
}

func generateElements(limit uint, int1 uint, int2 uint, string1 string, string2 string) []string {
	queue := make([]string, limit)

	for i := uint(1); i <= limit; i++ {
		queue[i - 1] = generateElement(i, int1, int2, string1, string2)
	}

	return queue
}

func generateElement(value uint, int1 uint, int2 uint, string1 string, string2 string) (elem string) {
	bitMap := valueIfTrue(isMultipleOf(value, int1), 1) + valueIfTrue(isMultipleOf(value, int2), 2)

	switch bitMap {
	case 0: elem = strconv.FormatUint(uint64(value), 10)
	case 1: elem = string1
	case 2: elem = string2
	case 3: elem = fmt.Sprintf("%s%s", string1, string2)
	default: log.Fatal("should never happen")
	}

	return
}

func valueIfTrue(condition bool, left uint) uint {
	if condition { return left } else { return 0 }
}

func isMultipleOf(value uint, multiple uint) bool {
	return (value % multiple) == 0
}