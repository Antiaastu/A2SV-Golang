package main
import(
	"context"
	"log"
	"task_manager_5/data"
	"task_manager_5/router"
	"time"
)

func main(){
	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	err := data.InitMongoDB(ctx, "mongodb://localhost:27017")
	if err != nil{
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer data.CloseMongoDB()
	r := router.SetUpRouter()
	r.Run(":8080")
}