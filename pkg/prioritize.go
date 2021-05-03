package pkg

import (
	"github.com/gin-gonic/gin"
	"k8s.io/kubernetes/pkg/scheduler/api"
	"log"
	"math/rand"
	"net/http"
)

const (
	// lucky priority gives a random [0, schedulerapi.MaxPriority] score
	// currently schedulerapi.MaxPriority is 10
	luckyPrioMsg = "node %v/%v is lucky to get score %v\n"
)


func Prioritize(c *gin.Context)  {
	var args api.ExtenderArgs
	if err :=c.Bind(&args);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":"error"})
		return
	}


	log.Printf("pod name:%v,nodelist: %v,node Name: %v",args.Pod.Name,args.Nodes.Items,args.NodeNames)


	c.JSON(http.StatusOK,prioritize(args))
}

func prioritize(args api.ExtenderArgs) *api.HostPriorityList {
	//pod := args.Pod
	nodes := args.Nodes.Items

	hostPriorityList := make(api.HostPriorityList, len(nodes))
	for i, node := range nodes {
		score := rand.Intn(api.MaxPriority + 1)
		log.Printf(luckyPrioMsg, node.Name, node.Namespace, score)
		hostPriorityList[i] = api.HostPriority{
			Host:  node.Name,
			Score: score,
		}
	}

	return &hostPriorityList
}