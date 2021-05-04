package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/kubernetes/pkg/scheduler/api"
	"log"
	"net/http"
)

const (
	// lucky priority gives a random [0, schedulerapi.MaxPriority] score
	// currently schedulerapi.MaxPriority is 10
	luckyPrioMsg = "node %v is lucky to get score %v\n"
)


func Prioritize(c *gin.Context)  {
	var args api.ExtenderArgs
	if err :=c.Bind(&args);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":"error"})
		return
	}


	log.Printf("pod name:%v\n",args.Pod.Name)


	c.JSON(http.StatusOK,prioritize(args))
}

func prioritize(args api.ExtenderArgs) *api.HostPriorityList {
	//pod := args.Pod
	nodes := *args.NodeNames
	hostPriorityList := make(api.HostPriorityList, len(*args.NodeNames))
	for i, node := range nodes {
		score :=0
		if node == "master"{
			fmt.Println("node is master")
			score =1000+score
		}
		//score := rand.Intn(api.MaxPriority + 1)
		log.Printf(luckyPrioMsg, node, score)
		hostPriorityList[i] = api.HostPriority{
			Host:  node,
			Score: score,
		}
	}

	return &hostPriorityList
}