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








    

   
