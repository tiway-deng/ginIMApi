package validators

type TalkRecord struct {
	Source    int `form:"source",valid:"Required;range(1,2)"`
	RecordId  int `form:"record_id",valid:"Required;Numeric; Min(0)"`
	ReceiveId int `form:"receive_id",valid:"Required;Numeric; Min(1)"`
}
