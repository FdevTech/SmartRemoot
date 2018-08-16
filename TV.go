package service
import (
"strconv"
"log"
"os/exec"
"github.com/paypal/gatt"
"time"
"fmt"
"database/sql"
_ "github.com/mattn/go-sqlite3"
)


var son int
var sound =0
var Req int
var Req2 int
var n int
var m int
var x,reqS,temps,a,b int

func NewTVService() *gatt.Service {
s := gatt.NewService(gatt.MustParseUUID("18fc95c0-c111-11e3-9904-0002a5d5c51b"))
// vol up
s.AddCharacteristic(gatt.MustParseUUID("51fac9e0-c111-11e3-9246-0002a5d5c51b")).HandleWriteFunc(

func(r gatt.Request, data []byte) (status byte) {
db, _ := sql.Open("sqlite3", "/home/pi/gopath/src/github.com/paypal/gatt/examples/service/cerist.db")
row, _ := db.Query("SELECT value FROM values_ WHERE id=1")
for row.Next() {
		var value int
		row.Scan(&value)
		row.Close()
		son = value
	}
println(son)
log.Println(data[0])
log.Println("sound up")
log.Println(data[0])
Req=int(data[0])
m=Req-son
if m > 0{ 
for i := 0; i < m; i++ {
exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_VOLUMEUP").Output()		
time.Sleep(250* time.Millisecond)
log.Println(i)

}
db3,_ := sql.Open("sqlite3", "/home/pi/gopath/src/github.com/paypal/gatt/examples/service/cerist.db")
row3, _ := db3.Prepare("update values_ 	SET value = "+strconv.Itoa(Req)+"  WHERE remote_id=1")
row3.Exec()
}else{
println("rak tokhrote")
}
return gatt.StatusSuccess

})
//vol down

s.AddCharacteristic(gatt.MustParseUUID("52fac9e0-c111-11e3-9246-0002a5d5c51b")).HandleWriteFunc(

func(r gatt.Request, data []byte) (status byte) {
log.Println("sound down")
db, _ := sql.Open("sqlite3", "/home/pi/gopath/src/github.com/paypal/gatt/examples/service/cerist.db")
row, _ := db.Query("SELECT value FROM values_ WHERE id=1")
for row.Next() {
		var value int
		row.Scan(&value)
		row.Close()
		son = value
	}
println(son)
Req2=int(data[0])
log.Println(data[0])
n=son-Req2
println(n)
println(Req2)
if n > 0 {
for i := 0; i < n; i++ {
exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_VOLUMEDOWN").Output()
time.Sleep(250* time.Millisecond)
log.Println(i)


}

db2,_ := sql.Open("sqlite3", "/home/pi/gopath/src/github.com/paypal/gatt/examples/service/cerist.db")
row2, _ := db2.Prepare("update values_ 	SET value = "+strconv.Itoa(Req2)+"  WHERE remote_id=1")
row2.Exec()
}else {
println ("rak tokhrete")
}
return gatt.StatusSuccess
})


s.AddCharacteristic(gatt.MustParseUUID("53fac9e0-c111-11e3-9246-0002a5d5c51b")).HandleWriteFunc(
func(r gatt.Request, data []byte) (status byte) {
log.Println("53fac9e0-c111-11e3-9246-0002a5d5c51b on/off")
exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_POWER").Output()
log.Println(data[0])
if data[0]==0{

log.Println("salut Anis")

}
return gatt.StatusSuccess
})

s.AddCharacteristic(gatt.MustParseUUID("54fac9e0-c111-11e3-9246-0002a5d5c51b")).HandleReadFunc(
func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
log.Println(sound)
fmt.Fprintf(rsp, "%v", sound)
})

//timmer
s.AddCharacteristic(gatt.MustParseUUID("55fac9e0-c111-11e3-9246-0002a5d5c51b")).HandleWriteFunc(
func(r gatt.Request, data []byte) (status byte) {
reqS=int(data[0])
db, _ := sql.Open("sqlite3", "/home/pi/gopath/src/github.com/paypal/gatt/examples/service/cerist.db")
row, _ := db.Query("SELECT value FROM values_ WHERE id=2")
for row.Next() {
		var value int
		row.Scan(&value)
		row.Close()
		temps = value
	}
println(temps)
x=reqS-temps
println(x)
println("wache hsabt")

log.Println("55fac9e0-c111-11e3-9246-0002a5d5c51b timer ")
if reqS ==0 {
	println("timer is off")
}else if reqS==1{

	println("timer is 15 mns")
} else if reqS==2{

	println("timer is 30 mns")
} else if reqS==3 {

	println("timer is 60 mns")
} else {

	println("timer is 90 mns ")
}
	

if x > 0 {
		println("je suis la")
		a=x
		exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_MENU").Output()
		time.Sleep(500* time.Millisecond)

		exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_DOWN").Output()
		time.Sleep(500* time.Millisecond)

		exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_RIGHT").Output()
		time.Sleep(500* time.Millisecond)

		exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_OK").Output()
		time.Sleep(500* time.Millisecond)

		exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_DOWN").Output()
		time.Sleep(500* time.Millisecond)

		for i := 0; i < a; i++ {
			exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_RIGHT").Output()
			time.Sleep(500* time.Millisecond)
			println(i)

                                        }

	}else if x <0 {
		b=-x
		println("negatif")
		exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_MENU").Output()
		time.Sleep(500* time.Millisecond)

		exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_DOWN").Output()
		time.Sleep(500* time.Millisecond)

		exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_RIGHT").Output()
		time.Sleep(500* time.Millisecond)

		exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_OK").Output()
		time.Sleep(500* time.Millisecond)

		exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_DOWN").Output()
		time.Sleep(500* time.Millisecond)

		for i := 0; i < b; i++ {
			exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_LEFT").Output()
			time.Sleep(500* time.Millisecond)
			println(i)

		}


	}else {
		println("rak tokhrete")
	}

exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_BACK").Output()
time.Sleep(500* time.Millisecond)

exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_BACK").Output()
		time.Sleep(500* time.Millisecond)
exec.Command("sh","-c","irsend SEND_ONCE remoteC KEY_BACK").Output()
		time.Sleep(500* time.Millisecond)



db6,_ := sql.Open("sqlite3", "/home/pi/gopath/src/github.com/paypal/gatt/examples/service/cerist.db")
row6, _ := db6.Prepare("update values_ 	SET value = "+strconv.Itoa(reqS)+"  WHERE id=2")
row6.Exec()




return gatt.StatusSuccess
})

return s
}
