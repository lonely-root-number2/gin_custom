module gin-demo

go 1.16

require github.com/gin-gonic/gin v1.8.1 
replace (
	github.com/gin-gonic/gin v1.8.1 => ./mgin@v1.0
)
