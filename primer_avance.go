package main
/*	BUFIO se usa para la introduccion de datos por consola :)
	STRCONV se usa para las validaciones de tipo buhhhh :( 
	FMT es el formateador de texto, me permite imprimir texto en consola
	MATH creo que su nombre lo dice solo xd
*/
import (
	"fmt" 
	"math" 
	"bufio" 
	"strconv" 
	"os" 
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

func float_validation(input string) float64{
	temp, err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil || math.IsNaN(temp) || math.IsInf(temp,0){
		return 0
	}
	return temp
}

func integer_validation(input string) int{
	temp, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Entrada inválida: digite un número entero.")
		return 0;
	}
	return temp;
}

func main(){
	scanner := bufio.NewScanner(os.Stdin);
	//Declaracion de ciclo infinito se hace con for {} aura
	for{
		var op_string string;
		fmt.Println("Figura que desea calcular:\nA = Rectangulo\nB = Circunferencia\nC = Triangle\nD = Exit")
		fmt.Print("Opcion: ")
		if !scanner.Scan() {
			break
		}
		op_string = strings.ToLower(scanner.Text());
		switch op_string{
		case "a":
			var h float64
			var w float64
			fmt.Println("Haz elegido el Rectangulo a continuacion ingrese sus datos: ")
			fmt.Print("Altura : ")
			fmt.Scan(&h)
			fmt.Print("Anchura : ")
			fmt.Scan(&w)
			fig := Rectangle{nam: "Rectangulo", height: h, width: w}
			print_figure(fig);
		case "b":
			var r float64
			fmt.Println("Haz elegido la Circunferencia a continuacion ingrese sus datos: ")
			fmt.Print("Radio : ")
			fmt.Scan(&r)
			fig := Circle{nam: "Circunferencia", radious: r}
			print_figure(fig);
		case "c":
			var lado_a float64;
			var lado_b float64;
			var lado_c float64;
			fmt.Println("Haz elegido el Triangulo a continuacion ingrese sus datos: ")
			fmt.Print("Largo lado A: ")
			fmt.Scan(&lado_a)
			fmt.Print("Largo lado B: ")
			fmt.Scan(&lado_b)
			fmt.Print("Largo lado C: ")
			fmt.Scan(&lado_c)
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
