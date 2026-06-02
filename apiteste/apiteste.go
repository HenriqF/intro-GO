package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/porra/:a", func(cont *gin.Context) {
		// receba := cont.Request.URL.Query()
		// for key, value := range receba {
		// 	fmt.Printf("%s -> %s\n", key, value)
		// }

		param := cont.Param("a")
		var reverse []byte

		for i := len(param) - 1; i >= 0; i-- {
			reverse = append(reverse, param[i])
		}

		cont.JSON(200, gin.H{"sigma": string(reverse)})
	})

	router.GET("bin/:string", func(cont *gin.Context) {
		var input []byte = []byte(cont.Param("string"))
		var output []byte

		for _, n := range input {
			conv := int(n)

			this_bin := make([]byte, 8)
			for i := range this_bin {
				this_bin[i] = 0
			}

			for i := 0; i < 8; i++ {
				this_bin[i] = byte(conv % 2)
				conv /= 2
			}

			for i, j := 0, len(this_bin)-1; i < j; i, j = i+1, j-1 {
				this_bin[i], this_bin[j] = this_bin[j]+48, this_bin[i]+48
			}
			output = append(output, this_bin...)
			output = append(output, ' ')
		}

		res := string(output)

		cont.JSON(200, res[:len(res)-1])
	})

	router.GET("/porra1/*a", func(cont *gin.Context) {
		param := cont.Param("a")
		fmt.Println(param)

		cont.JSON(200, gin.H{"sigma": "da bahia"})
	})

	router.Run("0.0.0.0:3050")
	fmt.Println("Hello")
}
