package main
import(
	"library_management/controllers"
	"library_management/services"
)

func main(){
	library := services.NewLibrary()
	controllers.StartConsole(library)
}