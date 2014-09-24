logs
===
##The logs is a simple and easy to use the go source code, Basic data structure support.
##[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

#for example
    package main
    
    import "logs"
    
    func main() {
    	if err := logs.Init("./logs.conf"); err != nil {
    		panic(err)
    	}
    	logs.Debug("%s", "Debug")
    	logs.Info("%s", "Info")
    	logs.Warning("%s", "Warning")
    	logs.Error("%s", "Error")
    }