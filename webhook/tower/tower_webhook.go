package tower

type TWWebHook struct {

}

func (tw *TWWebHook) String() string{
	return "Tower Webhook messgae"
}

func (tw  *TWWebHook)ParseMessage() (TWMessage, error){
	return nil,nil
}