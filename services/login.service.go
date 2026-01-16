package services
import(
    "context"
	// "log"
	"trading/config"
	"trading/db"
	
    // "github.com/gofiber/fiber/v2"
)
func SignUp(username string,email string,password string)error{
    dB:=config.OpenMysqlConnectionQuery()
	defer dB.Close()
    params:=db.CreateUserParams{
        Username   :username,
        Email      :email,
        PasswordHash :password,
    }
    if _,err:=dB.CreateUser(context.Background(),params);err!=nil{
        return err
    }
    return nil




}
func Login(email string,password string)(bool,string){
    dB:=config.OpenMysqlConnectionQuery()
	defer dB.Close()
    value,err:=dB.GetUserByEmail(context.Background(), email)
    if err!=nil{
        return false,"error"

    }
    if password==value.PasswordHash{
        return true,"success"
    }
    return false,"unsuccessfull"




}
func SendOtp(){

}