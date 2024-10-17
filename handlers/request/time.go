package request

type TimeReq struct {
	Timezone string `query:"timezone" validate:"required,timezone"`
}
