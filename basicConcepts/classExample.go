package main

import "fmt"

type Employee struct {
	id int
	name string
	vacation bool
}

// implementacion del patron constructor en go
// CUANDO TRABAJAMOS CON STRUCTS Y LOS PASAMOS A FUNCIONES, GO LOS TRATA COMO SI FUERAN COPIAS, asi que es muy importante devolver la referencia
func NewEmployee (id int, name string, vacation bool) *Employee{
	return &Employee{
		id: id,
		name: name,
		vacation: vacation,
	}
}


// receiver functions se usan para implementar metodos dentro de structs como si fueran clases
func (employee *Employee) SetId(id int){
	employee.id = id
}

func (employee *Employee) SetName(name string){
	employee.name = name
}

func (employee *Employee) GetId() int {
	return employee.id
}

func (employee *Employee) GetName() string {
	return employee.name
}


func main(){
	employeeInstance := Employee{}
	fmt.Printf("%v", employeeInstance)
	fmt.Printf("%+v", employeeInstance)

	employeeInstance.id = 1
	employeeInstance.name = "Jose"
	fmt.Printf("%+v", employeeInstance)
	fmt.Printf("%v", employeeInstance)


	employeeInstance.SetId(1000)
	employeeInstance.SetName("Marco")
	fmt.Println(employeeInstance.GetId(), employeeInstance.GetName())

	// OTRAS FORMAS DE INSTANCIAS UN STRUCT
	secondEmployee := Employee{
		id: 2,
		name: "Alberto",
	}
	thirdEmployee := new(Employee) // cuando usas new TE VA A DEVOLVER UNA INSTANCIA DEL APUNTADOR DE EMPLOYEE
	fmt.Println(secondEmployee,thirdEmployee)


	// usando el constructor
	fourthEmploye := NewEmployee(4,"Alicia", false)
	fmt.Printf("%v\n", *fourthEmploye)
}





