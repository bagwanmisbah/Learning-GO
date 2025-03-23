1.Download and Install GO from [https://go.dev/dl/](https://go.dev/dl/) according to OS

2.Complete Installation WIZARD

3.Create a folder to store go programs 

4.Remember  Package contains GO files

Modules contain Number of Packages 

5.Initialize New Module using command

        go mod init foldername
        go mod init go-prac
        
6 go.mod will contain something like
        
        module go-prac

        go 1.24.1

7. Create a main.go file
8. write basic code as
   
        package main
        
        import "fmt"
        
        func main() {
        	fmt.Println("Hello world ! My name is Misbah ! ")
        }

9.To Execute it
9.1 First build a binary like 

    go build main.go
    
then execute binary as 

    ./main

9.2 Simply execute without creating binary 

    go run main.go

![image](https://github.com/user-attachments/assets/f844729c-8edd-4507-b35e-06e1298309e0)

VARIABLES and CONSTANTS :

    var varname datatype = value 

anything we declare we must use 

    package main
    
    import "fmt"
    
    func main() {
    	// fmt.Println("Hello world ! My name is Misbah ! ")
    	var intnum int16 = 3242
    	fmt.Println(intnum)
    }


NOTE : int16 has its own limitations therefore 

    package main
    
    import "fmt"
    
    func main() {
    	// fmt.Println("Hello world ! My name is Misbah ! ")
    	var intnum int16 = 32767
    	intnum += 1
    	fmt.Println(intnum)
    }

program will cause integer overflow because int16 can store values 

only in the range -32,768 to 32,767.

intnum is declared as int16 and assigned 32767, the maximum possible value.

When you do intnum += 1, it overflows, and Go wraps it around to -32,768 (minimum value of int16).

thus intnum must be assigned int32 or uint16 will give result 32768 

    package main
    
    import "fmt"
    
    func main() {
      // fmt.Println("Hello world ! My name is Misbah ! ")
      var intnum uint16 = 32767
      intnum += 1
      fmt.Println(intnum)
    
      var fnum float32 = 56.789
      fmt.Println(fnum)
    
      // cannot perform operations with mixed datatypes
      //must perform typecasting to do that
      var fnum1 float32 = 15.0
      var inum1 int32 = 2
      var res = fnum1 * float32(inum1)
      fmt.Println(res)
    
      //integer division results in integer without the presence of decimal points
      var i1 int32 = 100
      var i2 int32 = 3
      fmt.Println(i1 / i2)
      // fmt.Println(i1 % i2)
    
      var s1 string = `Hello
      world`
      fmt.Println(s1)
      s1 = s1 + " " + " its me"
      fmt.Println(s1)
    
      var b bool = true
      fmt.Println(!b)
    
      //val are default set for diff integers
      var i int
      fmt.Println(i)
      // for all int def val is 0
    
      // shorthand declaration
      inum2 := 5
      fmt.Println(inum2)
    
      // multiple declaration
    
      v1, v2 := 7, 8
      fmt.Println(v1, v2)
    
      s := "hi there"
      fmt.Println(s)
    
      //adding types is good practice
    
      const myconst int = 12
      //must always be initialised
    
    }







    

   
