package main

/*	BUFIO se usa para la introduccion de datos por consola :)
	STRCONV se usa para las validaciones de tipo buhhhh :(
	FMT es el formateador de texto, me permite imprimir texto en consola
	MATH creo que su nombre lo dice solo xd
*/
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//Estaba viendo traductores mientras estaba haciendo esto por eso es que es tan largo estaba ladillado
// DECLARACION DE INTERFAZ Y ESTRUCTURAS
type Figure interface {
	area() float64;
	perimeter() float64;
	name() string;
}

type Rectangle struct {
	nam string;
	height float64;
	width float64;
}

type Circle struct {
	nam string;
	radious float64;
}

type Triangle struct {
	nam string;
	a float64;
	b float64;
	c float64;
}
// METODOS HEREDADDS DE LA CLASE PADRE FIGURA
//Los de perimetro los hice por declaracion corta y area por la otra
func (r Rectangle) name() string {
	return r.nam;
}

func (r Rectangle) area() float64{
	var area float64 = r.height * r.width;
	return area;
}
func (r Rectangle) perimeter() float64{
	perimeter := 2*(r.height + r.width);
	return perimeter;
}
func (c Circle) name() string {
	return c.nam;
}

func (c Circle) area() float64{
	var area float64 = math.Pi * c.radious * c.radious;
	return area;
}
func(c Circle) perimeter() float64{
	perimeter := math.Pi * c.radious * 2;
	return perimeter;
}

func (t Triangle) name() string {
	return t.nam;
}

func (t Triangle) area() float64{
	var s float64 = (t.a + t.b + t.c)/2; 
	var area float64 = math.Sqrt(s*(s-t.a)*(s-t.b)*(s-t.c));
	return area;
}
func (t Triangle) perimeter() float64{
	perimeter := t.a + t.b + t.c;
	return perimeter;
}

func print_figure(f Figure){
	fmt.Printf("Hola soy un: %s\nArea: %.2f cm²\nPerimetro: %.2f\n" , f.name(), f.area(),f.perimeter());
	return;
}

//FUNCIONES HELPERS

func float_input(text string, r *bufio.Reader) float64{
	fmt.Print(text)
	input_string, _ := r.ReadString('\n');
	input, err := strconv.ParseFloat(strings.TrimSpace(input_string),64)
	if err != nil || math.IsNaN(input) || math.IsInf(input,0){
		fmt.Println("Se le asigno el valor de 0 por mala escritura de datos")
		return 0
	}
	return input
}




func main(){
	scanner := bufio.NewReader(os.Stdin);
	//Declaracion de ciclo infinito se hace con for {} aura
	for{
		var op_string string;
		fmt.Println("Figura que desea calcular:\nA = Rectangulo\nB = Circunferencia\nC = Triangle\nD = Exit")
		fmt.Print("Opcion: ")
		
		op_string, e := scanner.ReadString('\n');
		if e!= nil{
			fmt.Print("Como pudo dar error estooo")
			break
		}
		op_p := strings.ToLower(op_string)
		op := strings.TrimSpace(op_p);
		switch op{
		case "a":
			fmt.Println("Haz elegido el Rectangulo a continuacion ingrese sus datos: ")
			var h float64 = float_input("Altura : ", scanner)
			var w float64 = float_input("Anchura : ", scanner)
			fig := Rectangle{nam: "Rectangulo", height: h, width: w}
			print_figure(fig);
		case "b":
			fmt.Println("Haz elegido la Circunferencia a continuacion ingrese sus datos: ")
			var r float64 = float_input("Radio : ",scanner)
			fig := Circle{nam: "Circunferencia", radious: r}
			print_figure(fig);
		case "c":
			fmt.Println("Haz elegido el Triangulo a continuacion ingrese sus datos: ")
			var lado_a float64 = float_input("Largo lado A: ", scanner);
			var lado_b float64 = float_input("Largo lado B: ", scanner);
			var lado_c float64 = float_input("Largo lado C: ", scanner);
			fig := Triangle{nam: "Triangulo", a: lado_a, b: lado_b, c: lado_c}
			print_figure(fig);
		case "d":
			fmt.Println("Gracias por usar este sitema")
			os.Exit(0); 
		default:
			fmt.Println("Error: Campo no valido")
		}
		
	}
}
