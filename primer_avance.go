package main
/*	BUFIO se usa para la introduccion de datos por consola :)
	STRCONV se usa para las validaciones de tipo buhhhh :( 
	FMT es el formateador de texto, me permite imprimir texto en consola
	MATH creo que su nombre lo dice solo xd
*/
import ("fmt", "math", "bufio","strconv","os","strings")
//Estaba viendo traductores mientras estaba haciendo esto por eso es que es tan largo estaba ladillado 
// DECLARACION DE INTERFAZ Y ESTRUCTURAS
type Figure interface {
	area() float64;
	perimeter() float64;
	name() string;
}

type Rectangle struct {
	nombre string;
	height float64;
	witdh float64;
}

type Circle struct {
	nombre string;
	radious float64;
}

type Triangle struct {
	nombre string;
	a float64;
	b float64;
	c float64;
}
// METODOS HEREDADDS DE LA CLASE PADRE FIGURA
//Los de perimetro los hice por declaracion corta y area por la otra
func (r Rectangle) name() string {
	return r.name;
}

func (r Rectangle) area() float64{
	var area float64 = r.height * r.witdh;
	return area;
}
func (r Rectangle) perimeter() float64{
	perimeter := 2*(r.height + r.witdh);
	return perimeter;
}
func (c Circle) name() string {
	return c.name;
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
	return t.name;
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
	fmt.Printf("Hola soy un: %s\nArea: %.2f cm²\nPerimetro: %.2f" , f.name, f.area(),f.perimeter());
	return;
}

func float_validation(input string) (float64, bool){
	temp, err := strconv.ParseFloat(input,64)
	if err != nil || math.isNaN(temp) || math.isInf(tem,0){
		return 0, false;
	}
	return temp, true;
}

func integer_validation(input string) (float64, bool){
	temp, err := strconv.Atoi
}

func main(){
	scanner := bufio.NewScanner(os.Stdin);
	//Declaracion de ciclo infinito se hace con for {} aura
	for {
		fmt.Print("Ingrese que figura quiere calcular (1: Rectangulo / 2: Circulo / 3: Triangulo)")
		_, err := fmt
		//ENTRADA DE TEXTO POR CONSOLA
		input := strings.TrimSpace(scanner.Text());
		//VALIDACION DE TIPO DE DATOS DE ENTRADA
		var op int;
		op, err := strconv.ParseFloat(input,64);
		if err != nil || math.isNaN(op)
		
	}
}