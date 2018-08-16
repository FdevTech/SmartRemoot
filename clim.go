package service
import (
"log"
"github.com/paypal/gatt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"os/exec"
	"time"
	//"debug/elf"
)
var mOde int
var chaleur  int
var Temp int
var moDe int
func NewclimService() *gatt.Service {
	s := gatt.NewService(gatt.MustParseUUID("19fc95c0-c111-11e3-9904-0002a5d5c51b"))
	//Temp
	s.AddCharacteristic(gatt.MustParseUUID("41fac9e0-c111-11e3-9246-0002a5d5c51b")).HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			db, _ := sql.Open("sqlite3", "/home/pi/gopath/src/github.com/paypal/gatt/examples/service/cerist.db")
			row, _ := db.Query("SELECT value FROM values_ WHERE id=4")
			for row.Next() {
				var value int
				row.Scan(&value)
				row.Close()
				Temp = value
			}
			println(chaleur)
			log.Println(data[0])
			chaleur=int(data[0])

			if chaleur == 16 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_0").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_0").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_0").Output()

				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_0").Output()
				}

				println("temp is 16")
			} else if chaleur == 17 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_1").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_1").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_1").Output()

				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_1").Output()
				}
				println("temp is 17")

			} else if chaleur == 18 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_2").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_2").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_2").Output()
				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_2").Output()

				}
				println("temp is 18")

			} else if chaleur == 19 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_102ND").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_102ND").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_102ND").Output()
				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_102ND").Output()
				}
				println("temp is 19")

			} else if chaleur == 20 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_3").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_3").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_3").Output()
				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_3").Output()
				}
				println("temp is 20")

			} else if chaleur == 21 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_4").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_4").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_4").Output()
				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_4").Output()
				}
				println("temp is 21")

			} else if chaleur == 22 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_5").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_5").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_5").Output()
				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_5").Output()
				}
				println("temp is 22")

			} else if chaleur == 23 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_6").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_6").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_6").Output()
				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_6").Output()
				}
				println("temp is 23")

			} else if chaleur == 24 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_7").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_7").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_7").Output()
				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_7").Output()
				}
				println("temp is 24")

			} else if chaleur == 25 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_ALTERASE").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_8").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_8").Output()
				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_8").Output()
				}
				println("temp is 25")

			} else if chaleur == 26 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_9").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_9").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_9").Output()
				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_9").Output()
				}
				println("temp is 26")

			} else if chaleur == 27 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_A").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_A").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_A").Output()
				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_A").Output()
				}
				println("temp is 27")

			} else if chaleur == 28 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_AB").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_AB").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_AB").Output()
				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_AB").Output()
				}
				println("temp is 28")

			} else if chaleur == 29 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_ADDRESSBOOK").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_ADDRESSBOOK").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_ADDRESSBOOK").Output()
				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_ADDRESSBOOK").Output()
				}
				println("temp is 29")

			} else if chaleur == 30 {
				if mOde == 0 {
					println("mode is cold")
					exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_AGAIN").Output()
				} else if mOde == 1 {
					println("mode is dry")
					exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_AGAIN").Output()
				} else if mOde == 2 {
					println("mode is hot")
					exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_AGAIN").Output()
				} else if mOde == 3 {
					println("mode is Wind")
					exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_AGAIN").Output()
				}
				println("temp is 30")

			}
			println("temperature changed")
			db7, _ := sql.Open("sqlite3", "/home/pi/gopath/src/github.com/paypal/gatt/examples/service/cerist.db")
			row7, _ := db7.Prepare("update values_ 	SET value = " + strconv.Itoa(chaleur) + "  WHERE id=4")
			row7.Exec()

			return gatt.StatusSuccess
		})
	//
	//s.AddCharacteristic(gatt.MustParseUUID("42fac9e0-c111-11e3-9246-0002a5d5c51b")).HandleWriteFunc(
	//func(r gatt.Request, data []byte) (status byte) {
	//log.Println("temperature down")
	//log.Println(data)
	//return gatt.StatusSuccess
	//})
	//MOde
	s.AddCharacteristic(gatt.MustParseUUID("43fac9e0-c111-11e3-9246-0002a5d5c51b")).HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			db, _ := sql.Open("sqlite3", "/home/pi/gopath/src/github.com/paypal/gatt/examples/service/cerist.db")
			row, _ := db.Query("SELECT value FROM values_ WHERE id=3")
			for row.Next() {
				var value int
				row.Scan(&value)
				row.Close()
				mOde = value
			}
			println(mOde)
			log.Println(data[0])
			moDe=int(data[0])
			db3, _ := sql.Open("sqlite3", "/home/pi/gopath/src/github.com/paypal/gatt/examples/service/cerist.db")
			row3, _ := db3.Prepare("update values_ 	SET value = " + strconv.Itoa(moDe) + "  WHERE id=3")
			row3.Exec()
			if moDe == 0 {
				println("mode is cold")
				exec.Command("sh", "-c", "irsend SEND_ONCE condor KEY_ALS_TOGGLE").Output()
			} else if moDe == 1 {
				println("mode is dry")
				exec.Command("sh", "-c", "irsend SEND_ONCE modedRy KEY_ALS_TOGGLE").Output()
			} else if moDe == 2 {
				println("mode is hot")
				exec.Command("sh", "-c", "irsend SEND_ONCE modehot KEY_ALS_TOGGLE").Output()
			} else if moDe == 3 {
				println("mode is Wind")
				exec.Command("sh", "-c", "irsend SEND_ONCE modeWind KEY_ALS_TOGGLE").Output()
			}
			log.Println("mode changed")

			return gatt.StatusSuccess
		})
	s.AddCharacteristic(gatt.MustParseUUID("44fac9e0-c111-11e3-9246-0002a5d5c51b")).HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			if data[0] == 0 {
				exec.Command("sh", "-c", "irsend SEND_ONCE power KEY_POWER").Output()
				time.Sleep(500 * time.Millisecond)

				log.Println("state changed")
			} else {
				exec.Command("sh", "-c", "irsend SEND_ONCE power KEY_0").Output()
				time.Sleep(500 * time.Millisecond)
			}

			log.Println("state changed")

			log.Println("44fac9e0-c111-11e3-9246-0002a5d5c51b on/off")
			log.Println(data)
			return gatt.StatusSuccess
		})

	return s

}
