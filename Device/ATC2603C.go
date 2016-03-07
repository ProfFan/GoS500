package Device

import (
	"github.com/tjCFeng/GoS500/S500"
)

	const ATC2603C_TWI	= S500.TWI_0
	const ATC2603C_ADDR	= 0x65 //(0xCA >> 1)
	
	/* PMU Register Address */
	const ATC2603C_PMU_BASE				= (0X00)
	const ATC2603C_PMU_SYS_CTL0			= (ATC2603C_PMU_BASE + 0x00)
	const ATC2603C_PMU_SYS_CTL1			= (ATC2603C_PMU_BASE + 0x01)
	const ATC2603C_PMU_SYS_CTL2			= (ATC2603C_PMU_BASE + 0x02)
	const ATC2603C_PMU_SYS_CTL3			= (ATC2603C_PMU_BASE + 0x03)
	const ATC2603C_PMU_SYS_CTL4			= (ATC2603C_PMU_BASE + 0x04)
	const ATC2603C_PMU_SYS_CTL5			= (ATC2603C_PMU_BASE + 0x05)
	const ATC2603C_PMU_SYS_CTL6			= (ATC2603C_PMU_BASE + 0x06)
	const ATC2603C_PMU_SYS_CTL7			= (ATC2603C_PMU_BASE + 0x07)
	const ATC2603C_PMU_SYS_CTL8			= (ATC2603C_PMU_BASE + 0x08)
	const ATC2603C_PMU_SYS_CTL9			= (ATC2603C_PMU_BASE + 0x09)
	const ATC2603C_PMU_BAT_CTL0			= (ATC2603C_PMU_BASE + 0x0A)
	const ATC2603C_PMU_BAT_CTL1			= (ATC2603C_PMU_BASE + 0x0B)
	const ATC2603C_PMU_VBUS_CTL0		= (ATC2603C_PMU_BASE + 0x0C)
	const ATC2603C_PMU_VBUS_CTL1		= (ATC2603C_PMU_BASE + 0x0D)
	const ATC2603C_PMU_WALL_CTL0		= (ATC2603C_PMU_BASE + 0x0E)
	const ATC2603C_PMU_WALL_CTL1		= (ATC2603C_PMU_BASE + 0x0F)
	const ATC2603C_PMU_SYS_PENDING		= (ATC2603C_PMU_BASE + 0x10)
	const ATC2603C_PMU_DC1_CTL0			= (ATC2603C_PMU_BASE + 0x11)
	const ATC2603C_PMU_DC1_CTL1			= (ATC2603C_PMU_BASE + 0x12)
	const ATC2603C_PMU_DC1_CTL2			= (ATC2603C_PMU_BASE + 0x13)
	const ATC2603C_PMU_DC2_CTL0			= (ATC2603C_PMU_BASE + 0x14)
	const ATC2603C_PMU_DC2_CTL1			= (ATC2603C_PMU_BASE + 0x15)
	const ATC2603C_PMU_DC2_CTL2			= (ATC2603C_PMU_BASE + 0x16)
	const ATC2603C_PMU_DC3_CTL0			= (ATC2603C_PMU_BASE + 0x17)
	const ATC2603C_PMU_DC3_CTL1			= (ATC2603C_PMU_BASE + 0x18)
	const ATC2603C_PMU_DC3_CTL2			= (ATC2603C_PMU_BASE + 0x19)
	const ATC2603C_PMU_DC4_CTL0			= (ATC2603C_PMU_BASE + 0x1A)
	const ATC2603C_PMU_DC4_CTL1			= (ATC2603C_PMU_BASE + 0x1B)
	const ATC2603C_PMU_DC5_CTL0			= (ATC2603C_PMU_BASE + 0x1C)
	const ATC2603C_PMU_DC5_CTL1			= (ATC2603C_PMU_BASE + 0x1D)
	const ATC2603C_PMU_LDO1_CTL			= (ATC2603C_PMU_BASE + 0x1E)
	const ATC2603C_PMU_LDO2_CTL			= (ATC2603C_PMU_BASE + 0x1F)
	const ATC2603C_PMU_LDO3_CTL			= (ATC2603C_PMU_BASE + 0x20)
	const ATC2603C_PMU_LDO4_CTL			= (ATC2603C_PMU_BASE + 0x21)
	const ATC2603C_PMU_LDO5_CTL			= (ATC2603C_PMU_BASE + 0x22)
	const ATC2603C_PMU_LDO6_CTL			= (ATC2603C_PMU_BASE + 0x23)
	const ATC2603C_PMU_LDO7_CTL			= (ATC2603C_PMU_BASE + 0x24)
	const ATC2603C_PMU_LDO8_CTL			= (ATC2603C_PMU_BASE + 0x25)
	const ATC2603C_PMU_LDO9_CTL			= (ATC2603C_PMU_BASE + 0x26)
	const ATC2603C_PMU_LDO10_CTL		= (ATC2603C_PMU_BASE + 0x27)
	const ATC2603C_PMU_LDO11_CTL		= (ATC2603C_PMU_BASE + 0x28)
	const ATC2603C_PMU_SWITCH_CTL		= (ATC2603C_PMU_BASE + 0x29)
	const ATC2603C_PMU_OV_CTL0			= (ATC2603C_PMU_BASE + 0x2A)
	const ATC2603C_PMU_OV_CTL1			= (ATC2603C_PMU_BASE + 0x2B)
	const ATC2603C_PMU_OV_STATUS		= (ATC2603C_PMU_BASE + 0x2C)
	const ATC2603C_PMU_OV_EN				= (ATC2603C_PMU_BASE + 0x2D)
	const ATC2603C_PMU_OV_INT_EN		= (ATC2603C_PMU_BASE + 0x2E)
	const ATC2603C_PMU_OC_CTL			= (ATC2603C_PMU_BASE + 0x2F)
	const ATC2603C_PMU_OC_STATUS		= (ATC2603C_PMU_BASE + 0x30)
	const ATC2603C_PMU_OC_EN				= (ATC2603C_PMU_BASE + 0x31)
	const ATC2603C_PMU_OC_INT_EN		= (ATC2603C_PMU_BASE + 0x32)
	const ATC2603C_PMU_UV_CTL0			= (ATC2603C_PMU_BASE + 0x33)
	const ATC2603C_PMU_UV_CTL1			= (ATC2603C_PMU_BASE + 0x34)
	const ATC2603C_PMU_UV_STATUS		= (ATC2603C_PMU_BASE + 0x35)
	const ATC2603C_PMU_UV_EN				= (ATC2603C_PMU_BASE + 0x36)
	const ATC2603C_PMU_UV_INT_EN		= (ATC2603C_PMU_BASE + 0x37)
	const ATC2603C_PMU_OT_CTL			= (ATC2603C_PMU_BASE + 0x38)
	const ATC2603C_PMU_CHARGER_CTL0		= (ATC2603C_PMU_BASE + 0x39)
	const ATC2603C_PMU_CHARGER_CTL1		= (ATC2603C_PMU_BASE + 0x3A)
	const ATC2603C_PMU_CHARGER_CTL2		= (ATC2603C_PMU_BASE + 0x3B)
	const ATC2603C_PMU_BAKCHARGER_CTL	= (ATC2603C_PMU_BASE + 0x3C)
	const ATC2603C_PMU_APDS_CTL			= (ATC2603C_PMU_BASE + 0x3D)
	const ATC2603C_PMU_AUXADC_CTL0		= (ATC2603C_PMU_BASE + 0x3E)
	const ATC2603C_PMU_AUXADC_CTL1		= (ATC2603C_PMU_BASE + 0x3F)
	const ATC2603C_PMU_BATVADC			= (ATC2603C_PMU_BASE + 0x40)
	const ATC2603C_PMU_BATIADC			= (ATC2603C_PMU_BASE + 0x41)
	const ATC2603C_PMU_WALLVADC			= (ATC2603C_PMU_BASE + 0x42)
	const ATC2603C_PMU_WALLIADC			= (ATC2603C_PMU_BASE + 0x43)
	const ATC2603C_PMU_VBUSVADC			= (ATC2603C_PMU_BASE + 0x44)
	const ATC2603C_PMU_VBUSIADC			= (ATC2603C_PMU_BASE + 0x45)
	const ATC2603C_PMU_SYSPWRADC		= (ATC2603C_PMU_BASE + 0x46)
	const ATC2603C_PMU_REMCONADC		= (ATC2603C_PMU_BASE + 0x47)
	const ATC2603C_PMU_SVCCADC			= (ATC2603C_PMU_BASE + 0x48)
	const ATC2603C_PMU_CHGIADC			= (ATC2603C_PMU_BASE + 0x49)
	const ATC2603C_PMU_IREFADC			= (ATC2603C_PMU_BASE + 0x4A)
	const ATC2603C_PMU_BAKBATADC		= (ATC2603C_PMU_BASE + 0x4B)
	const ATC2603C_PMU_ICTEMPADC		= (ATC2603C_PMU_BASE + 0x4C)
	const ATC2603C_PMU_AUXADC0			= (ATC2603C_PMU_BASE + 0x4D)
	const ATC2603C_PMU_AUXADC1			= (ATC2603C_PMU_BASE + 0x4E)
	const ATC2603C_PMU_AUXADC2			= (ATC2603C_PMU_BASE + 0x4F)
	const ATC2603C_PMU_ICMADC			= (ATC2603C_PMU_BASE + 0x50)
	const ATC2603C_PMU_BDG_CTL			= (ATC2603C_PMU_BASE + 0x51)
	const ATC2603C_RTC_CTL				= (ATC2603C_PMU_BASE + 0x52)
	const ATC2603C_RTC_MSALM				= (ATC2603C_PMU_BASE + 0x53)
	const ATC2603C_RTC_HALM				= (ATC2603C_PMU_BASE + 0x54)
	const ATC2603C_RTC_YMDALM			= (ATC2603C_PMU_BASE + 0x55)
	const ATC2603C_RTC_MS				= (ATC2603C_PMU_BASE + 0x56)
	const ATC2603C_RTC_H					= (ATC2603C_PMU_BASE + 0x57)
	const ATC2603C_RTC_DC				= (ATC2603C_PMU_BASE + 0x58)
	const ATC2603C_RTC_YMD				= (ATC2603C_PMU_BASE + 0x59)
	const ATC2603C_EFUSE_DAT				= (ATC2603C_PMU_BASE + 0x5A)
	const ATC2603C_EFUSECRTL1			= (ATC2603C_PMU_BASE + 0x5B)
	const ATC2603C_EFUSECRTL2			= (ATC2603C_PMU_BASE + 0x5C)
	const ATC2603C_PMU_FW_USE0			= (ATC2603C_PMU_BASE + 0x5D)
	const ATC2603C_PMU_FW_USE1			= (ATC2603C_PMU_BASE + 0x5E)
	const ATC2603C_PMU_FW_USE2			= (ATC2603C_PMU_BASE + 0x5F)
	const ATC2603C_PMU_FW_USE3			= (ATC2603C_PMU_BASE + 0x60)
	const ATC2603C_PMU_FW_USE4			= (ATC2603C_PMU_BASE + 0x61)
	const ATC2603C_PMU_ABNORMAL_STATUS	= (ATC2603C_PMU_BASE + 0x62)
	const ATC2603C_PMU_WALL_APDS_CTL	= (ATC2603C_PMU_BASE + 0x63)
	const ATC2603C_PMU_REMCON_CTL0		= (ATC2603C_PMU_BASE + 0x64)
	const ATC2603C_PMU_REMCON_CTL1		= (ATC2603C_PMU_BASE + 0x65)
	const ATC2603C_PMU_MUX_CTL0			= (ATC2603C_PMU_BASE + 0x66)
	const ATC2603C_PMU_SGPIO_CTL0		= (ATC2603C_PMU_BASE + 0x67)
	const ATC2603C_PMU_SGPIO_CTL1		= (ATC2603C_PMU_BASE + 0x68)
	const ATC2603C_PMU_SGPIO_CTL2		= (ATC2603C_PMU_BASE + 0x69)
	const ATC2603C_PMU_SGPIO_CTL3		= (ATC2603C_PMU_BASE + 0x6A)
	const ATC2603C_PMU_SGPIO_CTL4		= (ATC2603C_PMU_BASE + 0x6B)
	const ATC2603C_PWMCLK_CTL			= (ATC2603C_PMU_BASE + 0x6C)
	const ATC2603C_PWM0_CTL				= (ATC2603C_PMU_BASE + 0x6D)
	const ATC2603C_PWM1_CTL				= (ATC2603C_PMU_BASE + 0x6E)
	const ATC2603C_PMU_ADC_DBG0			= (ATC2603C_PMU_BASE + 0x70)
	const ATC2603C_PMU_ADC_DBG1			= (ATC2603C_PMU_BASE + 0x71)
	const ATC2603C_PMU_ADC_DBG2			= (ATC2603C_PMU_BASE + 0x72)
	const ATC2603C_PMU_ADC_DBG3			= (ATC2603C_PMU_BASE + 0x73)
	const ATC2603C_PMU_ADC_DBG4			= (ATC2603C_PMU_BASE + 0x74)
	const ATC2603C_IRC_CTL				= (ATC2603C_PMU_BASE + 0x80)
	const ATC2603C_IRC_STAT				= (ATC2603C_PMU_BASE + 0x81)
	const ATC2603C_IRC_CC				= (ATC2603C_PMU_BASE + 0x82)
	const ATC2603C_IRC_KDC				= (ATC2603C_PMU_BASE + 0x83)
	const ATC2603C_IRC_WK				= (ATC2603C_PMU_BASE + 0x84)
	const ATC2603C_IRC_RCC				= (ATC2603C_PMU_BASE + 0x85)
	const ATC2603C_IRC_FILTER			= (ATC2603C_PMU_BASE + 0x86)

	/* AUDIO_OUT Register Address */
	const ATC2603C_AUDIO_OUT_BASE		= (0xA0)
	const ATC2603C_AUDIOINOUT_CTL		= (ATC2603C_AUDIO_OUT_BASE + 0x0)
	const ATC2603C_AUDIO_DEBUGOUTCTL	= (ATC2603C_AUDIO_OUT_BASE + 0x1)
	const ATC2603C_DAC_DIGITALCTL		= (ATC2603C_AUDIO_OUT_BASE + 0x2)
	const ATC2603C_DAC_VOLUMECTL0		= (ATC2603C_AUDIO_OUT_BASE + 0x3)
	const ATC2603C_DAC_ANALOG0			= (ATC2603C_AUDIO_OUT_BASE + 0x4)
	const ATC2603C_DAC_ANALOG1			= (ATC2603C_AUDIO_OUT_BASE + 0x5)
	const ATC2603C_DAC_ANALOG2			= (ATC2603C_AUDIO_OUT_BASE + 0x6)
	const ATC2603C_DAC_ANALOG3			= (ATC2603C_AUDIO_OUT_BASE + 0x7)

	/* AUDIO_IN Register Address */
	const ATC2603C_AUDIO_IN_BASE	= (0xA0)
	const ATC2603C_ADC_DIGITALCTL	= (ATC2603C_AUDIO_IN_BASE + 0x8)
	const ATC2603C_ADC_HPFCTL		= (ATC2603C_AUDIO_IN_BASE + 0x9)
	const ATC2603C_ADC_CTL			= (ATC2603C_AUDIO_IN_BASE + 0xA)
	const ATC2603C_AGC_CTL0			= (ATC2603C_AUDIO_IN_BASE + 0xB)
	const ATC2603C_AGC_CTL1			= (ATC2603C_AUDIO_IN_BASE + 0xC)
	const ATC2603C_AGC_CTL2			= (ATC2603C_AUDIO_IN_BASE + 0xD)
	const ATC2603C_ADC_ANALOG0		= (ATC2603C_AUDIO_IN_BASE + 0xE)
	const ATC2603C_ADC_ANALOG1		= (ATC2603C_AUDIO_IN_BASE + 0xF)

	/* PCM_IF Register Address */
	const ATC2603C_PCM_IF_BASE	= (0xA0)
	const ATC2603C_PCM0_CTL		= (ATC2603C_PCM_IF_BASE + 0x10)
	const ATC2603C_PCM1_CTL		= (ATC2603C_PCM_IF_BASE + 0x11)
	const ATC2603C_PCM2_CTL		= (ATC2603C_PCM_IF_BASE + 0x12)
	const ATC2603C_PCMIF_CTL		= (ATC2603C_PCM_IF_BASE + 0x13)

	/* CMU_CONTROL Register Address */
	const ATC2603C_CMU_CONTROL_BASE	= (0xC0)
	const ATC2603C_CMU_DEVRST		= (ATC2603C_CMU_CONTROL_BASE + 0x1)

	/* INTS Register Address */
	const ATC2603C_INTS_BASE	= (0xC8)
	const ATC2603C_INTS_PD	= (ATC2603C_INTS_BASE + 0x0)
	const ATC2603C_INTS_MSK	= (ATC2603C_INTS_BASE + 0x1)

	/* MFP Register Address */
	const ATC2603C_MFP_BASE		= (0xD0)
	const ATC2603C_MFP_CTL		= (ATC2603C_MFP_BASE + 0x00)
	const ATC2603C_PAD_VSEL		= (ATC2603C_MFP_BASE + 0x01)
	const ATC2603C_GPIO_OUTEN	= (ATC2603C_MFP_BASE + 0x02)
	const ATC2603C_GPIO_INEN		= (ATC2603C_MFP_BASE + 0x03)
	const ATC2603C_GPIO_DAT		= (ATC2603C_MFP_BASE + 0x04)
	const ATC2603C_PAD_DRV		= (ATC2603C_MFP_BASE + 0x05)
	const ATC2603C_PAD_EN		= (ATC2603C_MFP_BASE + 0x06)
	const ATC2603C_DEBUG_SEL		= (ATC2603C_MFP_BASE + 0x07)
	const ATC2603C_DEBUG_IE		= (ATC2603C_MFP_BASE + 0x08)
	const ATC2603C_DEBUG_OE		= (ATC2603C_MFP_BASE + 0x09)
	const ATC2603C_BIST_START	= (ATC2603C_MFP_BASE + 0x0A)
	const ATC2603C_BIST_RESULT	= (ATC2603C_MFP_BASE + 0x0B)
	const ATC2603C_CHIP_VER		= (ATC2603C_MFP_BASE + 0x0C)

	/* TWSI Register Address */
	const ATC2603C_TWI_BASE	= (0xF8)
	const ATC2603C_SADDR		= (ATC2603C_TWI_BASE + 0x7)
	
	
	
	const (V_WALL, V_VBUS, V_SYSPWR = 0, 2, 4)
	const (A_WALL, A_VBUS = 0, 1)


	var iATC2603C *ATC2603C = nil

func IATC2603C() (*ATC2603C) {
	if (iATC2603C == nil) {
		iATC2603C = &ATC2603C{}

		iATC2603C.twi, _ = S500.CreateTWI(nil, nil, S500.TWI_0)
		iATC2603C.twi.SetSpeed(S500.TWI_SpeedLow)
	}

	return iATC2603C
}

func FreeATC2603C() {
	S500.FreeTWI(iATC2603C.twi)
}

	type ATC2603C struct {
		twi		*S500.TWI
	}

func (this *ATC2603C) Read(Reg uint8, Len uint8) ([]uint8, bool) {
	return this.twi.Read(ATC2603C_ADDR, Reg, Len)
}

func (this *ATC2603C) Write(Reg uint8, Data []uint8) bool {
	return this.twi.Write(ATC2603C_ADDR, Reg, Data)
}

/**********************************************************************/
func (this *ATC2603C) Version() uint8 {
	Ver, Result := this.Read(ATC2603C_CHIP_VER, 1)
	if !Result { return 255 }
	return Ver[0]
}

func (this *ATC2603C) Temperature() float32 {
	Data, Result := this.Read(ATC2603C_PMU_ICTEMPADC, 2)
	if !Result { return 0.0 }
	return float32((uint16(Data[0]) << 8) + uint16(Data[1])) * 0.1949 - 44.899
}

func (this *ATC2603C) VoltageMV(Index uint8) uint32 {
	switch (Index) {
		case V_WALL, V_VBUS, V_SYSPWR:
		default: return 0.0
	}
	
	Data, Result := this.Read(ATC2603C_PMU_WALLVADC + Index, 2)
	if !Result { return 0.0 }
	
	V := (uint32(Data[0]) << 8) + uint32(Data[1])
	V *= uint32(Data[0])
	return uint32(float32(V) * 2.5)
}

func (this *ATC2603C) CurrentMA(Index uint8) uint32 {
	var DBG, REF uint32
	
	switch (Index) {
		case A_WALL:
			Data, Result := this.Read(ATC2603C_PMU_ADC_DBG2, 2)
			if !Result { return 0.0 }
			DBG = (uint32(Data[0]) << 8) + uint32(Data[1])
			Data, Result = this.Read(ATC2603C_PMU_IREFADC, 2)
			if !Result { return 0.0 }
			REF = (uint32(Data[0]) << 8) + uint32(Data[1])
			return uint32(float32(DBG) / float32(REF) * 1527)
		case A_VBUS:
			Data, Result := this.Read(ATC2603C_PMU_ADC_DBG1, 2)
			if !Result { return 0.0 }
			DBG = (uint32(Data[0]) << 8) + uint32(Data[1])
			Data, Result = this.Read(ATC2603C_PMU_IREFADC, 2)
			if !Result { return 0.0 }
			REF = (uint32(Data[0]) << 8) + uint32(Data[1])
			return uint32(float32(DBG) / float32(REF) * 1509)
		default: return 0.0
	}
}

func (this *ATC2603C) ADC(Index uint8) uint32 {
	if (Index > 2) { return 0 }
	
	Data, Result := this.Read(ATC2603C_PMU_AUXADC0 + Index, 2)
	if !Result { return 0 }
	return ((uint32(Data[0]) << 8) + uint32(Data[1])) & 0x3FF
}

func (this *ATC2603C) ConvertADCmV(Data uint32) uint32 {
	return Data * 3000 / 1023
}
