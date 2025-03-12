package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"

	"golang.org/x/sys/unix"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/logos/ecs/ecs_systems"
	"github.com/logos/global"
	"github.com/logos/pkg/logger"
	"github.com/logos/pkg/setting"
	"github.com/logos/running/main/initial"
)

var (
	runMode  string
	cfg      string
	limitCfg string
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}

	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupLimitSetting()
	if err != nil {
		log.Fatalf("init.setupLimitSetting err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	global.Logger.Infof(ctx, "db_lock")
	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, os.Interrupt, unix.SIGTERM)

	world := initial.InitialWorld()
	biologies := initial.InitialBiology(world)

	ecs_systems.SaveUnitsState(biologies, "test_unit_before_")

	for i := 0; i < 10; i++ {
		ecs_systems.UpdateWorld(world, biologies)
	}

	ecs_systems.SaveUnitsState(biologies, "test_unit_after_")

	cancel()
}

func setupFlag() error {
	flag.StringVar(&runMode, "mode", "", "running level (info, debug)")
	flag.StringVar(&cfg, "config", "etc/", "assgin the path of config file")
	flag.StringVar(&limitCfg, "limitCfg", "etc/", "assgin the path of limit config file")
	flag.Parse()

	return nil
}

func setupLimitSetting() error {
	s, err := setting.NewLimitSetting(strings.Split(limitCfg, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("CharacterAttributeRanges", &global.CARSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("PracticeAttributes", &global.ParSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("WorldAttributeSetting", &global.WorldAttributeSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("GlobalAttributeSetting", &global.GlobalAttributeSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupSetting() error {
	s, err := setting.NewSetting(strings.Split(cfg, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("SaveLoad", &global.SaveLoadSetting)
	if err != nil {
		return err
	}

	// TODO: run mode

	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
