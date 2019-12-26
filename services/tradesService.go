package services

type Filter struct {
	UID int
}

type TradesService struct{}

func (hm *TradesService) getCountTrades(in *Filter, out *string) error {
	fmt.Println("call getCountTrades", in)

	*out = "Hello " + in.Name + ", your age is " + strconv.Itoa(in.Age) + "."
	return nil
}