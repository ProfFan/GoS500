/*
	Author: tjCFeng(LiuYang)
	EMail: tjCFeng@163.com
	[2016.02.02]
*/

package S500

import (
	
)

	const BaseGPIO = 0xB01B0000
	
	const (PA, PB, PC, PD, PE = 0, 1, 2, 3, 4)
	const (FunIN, FunOUT = 0, 1)
	
	type PORT struct {
		Port	uint8
		hMem	[]uint8
		
		OUTEN	*uint32
		INEN	*uint32
		DAT		*uint32
	}
	
func CreatePort(PORTx uint8) (*PORT, bool) {
	var Result bool = false
	
	if (PORTx > PE) { return nil, Result }
	
	port := &PORT{Port: PORTx}
	port.hMem, Result = IS500().GetMMap(BaseGPIO)
	if !Result { return nil, Result }
	
	Reg := (BaseGPIO & 0x00000FFF) + uint32(PORTx) * 0x0C
	port.OUTEN, Result = IS500().Register(port.hMem, Reg + 0x00)
	port.INEN, Result = IS500().Register(port.hMem, Reg + 0x04)
	port.DAT, Result = IS500().Register(port.hMem, Reg + 0x08)

	return port, Result
}

func FreePort(port *PORT) {
	if (port.hMem != nil) { IS500().FreeMMap(port.hMem) }
}

/******************************************************************************/
	type GPIO struct {
		Port	*PORT
		Pin		uint8
		Bit		uint32
	}
	
func CreateGPIO(PORTx uint8, PINx uint8) (*GPIO, bool) {
	var Result bool = false

	gpio := &GPIO{}
	gpio.Port, Result = CreatePort(PORTx)
	if !Result { return nil, Result }

	gpio.Pin = PINx
	gpio.Bit = (0x1 << PINx)
	
	return gpio, Result
}

func FreeGPIO(gpio *GPIO) {
	FreePort(gpio.Port)
}

func (this *GPIO) SetFun(Fun uint8) {
	IMFP().SetGPIO(this, Fun)
}

func (this *GPIO) SetData(Data bool) {
	switch (Data) {
		case true: *this.Port.DAT |= this.Bit
		case false: *this.Port.DAT &^= this.Bit
	}
}

func (this *GPIO) GetData() bool {
	return (*this.Port.DAT & this.Bit) == this.Bit
}

func (this *GPIO) Flip() {
	*this.Port.DAT ^= this.Bit
}

