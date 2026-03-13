package erratum

func Use(opener ResourceOpener, input string)  (b error) {
	var file Resource
    
    for{
        var err error
        file, err = opener()
        if err == nil{
            break
        }
        if _, ok:= err.(TransientError); ok{ //type assertion
            continue //keep on trying open the file if err is TransientError
        }
        return err
    }
    defer file.Close()
    
    defer func(){
        if r:= recover(); r!=nil{
            switch danny:= r.(type){
                case FrobError:
                	file.Defrob(danny.defrobTag)
                	b=danny
				case error:
                	b=danny
                default:
                	panic(r)                	
            }
        }
    }()
    
    file.Frob(input) // defer have to register before funtion call
    return nil
}
