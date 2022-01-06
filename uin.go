package lemi025

type Uin uint8

func (uin Uin) Volts() float32 {
	return (float32(uin) / 10)
}
