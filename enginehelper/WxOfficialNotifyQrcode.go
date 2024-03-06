package bus

import (
	"e.coding.net/zhechat/magic/taihao/bootstrap"
	"e.coding.net/zhechat/magic/taihao/library/eventutil"
	"e.coding.net/zhechat/magic/taihao/library/logutil"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill/message"
)

const (
	wxOfficialNotifyQrcode = "wx-official-notify-qrcode"
)

type wxOfficialNotifyQrcodePayload struct {
	Facade     *bootstrap.Facade
	FromOpenid string
	QrcodeID   uint
	Domain     string
}

func WxOfficialNotifyQrcodeSub(f func(facade *bootstrap.Facade, fromOpenid string, qrcodeID uint, domain string)) {
	eventutil.Sub(wxOfficialNotifyQrcode+"-dose", wxOfficialNotifyQrcode, func(msg *message.Message) error {
		msg.Ack()
		var payload wxOfficialNotifyQrcodePayload
		if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			return err
		}
		f(payload.Facade, payload.FromOpenid, payload.QrcodeID, payload.Domain)
		return nil
	})
}
func WxOfficialNotifyQrcodePub(facade *bootstrap.Facade, fromOpenid string, qrcodeID uint, domain string) {
	err := eventutil.Pub(wxOfficialNotifyQrcode, wxOfficialNotifyQrcodePayload{
		Facade:     facade,
		FromOpenid: fromOpenid,
		QrcodeID:   qrcodeID,
		Domain:     domain,
	})
	logutil.Info("WxOfficialNotifyQrcode", fromOpenid, " ", qrcodeID, err)
}
