package pkg

import (
	"github.com/gin-gonic/gin"
	"k8s.io/kubernetes/pkg/scheduler/api"
	"log"
	"net/http"
)

func Filter(c *gin.Context)  {
	var args api.ExtenderArgs
	if err :=c.Bind(&args);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":"error"})
		return
	}


	log.Printf("pod name:%v,nodelist: %v,node Name: %v\n",args.Pod.Name,args.Nodes.Items,args.NodeNames)

	result := api.ExtenderFilterResult{
		Nodes:args.Nodes,
		NodeNames:args.NodeNames,
	}
	c.JSON(http.StatusOK,result)
}
