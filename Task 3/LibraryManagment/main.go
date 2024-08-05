package main 

import (
	"library_managment/services"
	"library_managment/controllers"
)

func main(){
	new_lib := services.NewLibrary()
	controllers.LibraryController(new_lib)
}