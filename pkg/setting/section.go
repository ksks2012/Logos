package setting

import (
	"fmt"
	"time"
)

type AppSettingS struct {
	RunMode      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	LogSavePath  string
	LogFileName  string
	LogFileExt   string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         []string
	SocketPath   string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type SaveLoadSettingS struct {
	SavePath     string
	SaveFileName string
	SaveFileExt  string
	LoadPath     string
	LoadFileName string
	LoadFileExt  string
}

// AttributeRange defines a range for generating random attributes
type AttributeRange struct {
	Min int
	Max int
}

// CharacterAttributeRanges stores the random range for all character attributes
type CARSettingS struct {
	VitalityRange     AttributeRange
	QiEnergyRange     AttributeRange
	StrengthRange     AttributeRange
	AgilityRange      AttributeRange
	IntelligenceRange AttributeRange
	AgeRange          AttributeRange
}

type PARSettingS struct {
	// TODO: spiritual root range
	BodyConstitutionRange AttributeRange
	ComprehensionRange    AttributeRange
	LuckFortuneRange      AttributeRange
	WillpowerRange        AttributeRange
	CultivationLevelRange AttributeRange
}

type WorldSettingS struct {
	BlockSizeRange AttributeRange
	AuraRange      AttributeRange
}

type GlobalSettingS struct {
	CharacterLimitRange AttributeRange
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}

	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func Display(setting interface{}) {
	fmt.Printf("%+v\n", setting)
}
