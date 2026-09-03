package lasagnamaster
// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int{
    if time==0{
        time=2
    }
    return time*len(layers)
}
// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int,float64){
    c1:=0
    c2:=0
    for i:=0;i<len(layers);i++{
        if layers[i]=="sauce"{
            c1++
        }else if layers[i]=="noodles"{
            c2++
        }
    }
    return c2*50,float64(c1)*0.2
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string,myList []string){
    if len(myList)==0||len(friendsList)==0{
        return
    }
    myList[len(myList)-1]=friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, num int) []float64{
    newArray:=make([]float64,len(quantities))
    portion:=float64(num)/2.0
    for i:=0;i<len(newArray);i++{
        newArray[i]=portion*quantities[i]
    }
    return newArray
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
