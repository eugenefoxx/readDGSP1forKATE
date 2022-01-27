package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	//	/home/a20272/Code/github.com/eugenefoxx/readDGS/lotnamesdata/lotnamesdata.go
	"github.com/eugenefoxx/readDGSP1forKATE/lotnamesdata"
	"github.com/eugenefoxx/readDGSP1forKATE/recipesap"
	"github.com/eugenefoxx/readDGSP1forKATE/setting"
	"github.com/fatih/color"
	"github.com/skratchdot/open-golang/open"
	"gopkg.in/ini.v1"
)

var lineL1 = flag.Bool("L1", false, "Введите флаг L1 для поиска рецепта на линии 1")
var lineL2 = flag.Bool("L2", false, "Введите флаг L2 для поиска рецепта на линии 2")

func main() {
	start := time.Now()
	flag.Parse()
	if !*lineL1 && !*lineL2 {
		color.Set(color.FgRed)
		fmt.Println("Введите флаг")
		color.Unset()
		log.Fatal()
	}

	var s string

	if len(os.Args) < 3 {
		color.Set(color.FgRed)
		fmt.Println("Введите название рецепта после флага")
		color.Unset()
		log.Fatal()
	} else if len(os.Args) > 3 {
		color.Set(color.FgRed)
		fmt.Println("После флага должен быть один аргумент")
		color.Unset()
		log.Fatal()
	} else {
		for i := 2; i < len(os.Args); i++ {
			s += os.Args[i]
		}
	}

	fmt.Println(s)

	if *lineL1 {
		cmd := &exec.Cmd{
			//Path: "/home/a20272/Code/github.com/eugenefoxx/readDGS/fileL1.sh",
			Path: setting.Conf.Bash + "fileL1.sh",
			//Args: []string{"/home/a20272/Code/github.com/eugenefoxx/readDGS/fileL1.sh", s},
			Args:   []string{setting.Conf.Bash + "fileL1.sh", s},
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}
		cmd.Run()

		// Read entire file content, giving us little control but
		// making it very simple. No need to close the file.
		//content, err := ioutil.ReadFile("/home/a20272/Code/github.com/eugenefoxx/readDGS/DGS_files/ExtractL1/fileresult.txt")
		content, err := ioutil.ReadFile(setting.Conf.DGSExtractL1 + "fileresult.txt")
		if err != nil {
			log.Fatal(err)
		}

		// Convert []byte to string and print to screen
		text := string(content)
		// чищу от переноса строки
		text = strings.TrimSuffix(text, "\n")
		//	text2 := strconv.Quote(text)
		//	fmt.Println("text2 -", text2)
		//	cfg, err := ini.Load("Brain-1.crb")
		cfg, err := ini.LoadSources(ini.LoadOptions{
			UnparseableSections: []string{
				"LotNames",
				"MultiPattern<1>",
				"MultiPattern<2>",
				"MultiPattern<3>",
				"MultiPattern<4>",
				"MultiPattern<5>",
				"MultiPattern<6>",
				"MultiPattern<7>",
				"PatternAttribute<1>",
				"PatternAttribute<2>",
				"PatternAttribute<3>",
				"PatternAttribute<4>",
				"PatternAttribute<5>",
				"PatternAttribute<6>",
				"PatternAttribute<7>",
				"PatternGrouping<1>",
				"PatternGrouping<2>",
				"PatternGrouping<3>",
				"PatternGrouping<4>",
				"PatternGrouping<5>",
				"PatternGrouping<6>",
				"PatternGrouping<7>",
				"RecogMarkSet<1>",
				"RecogMarkSet<2>",
				"RecogMarkSet<3>",
				"RecogMarkSet<4>",
				"RecogMarkSet<5>",
				"RecogMarkSet<6>",
				"RecogMarkSet<7>",
				"PositionData<1>",
				"PositionData<2>",
				"PositionData<3>",
				"PositionData<4>",
				"PositionData<5>",
				"PositionData<6>",
				"PositionData<7>",
				"PatternData<1>",
				"PatternData<2>",
				"PatternData<3>",
				"PatternData<4>",
				"PatternData<5>",
				"PatternData<6>",
				"PatternData<7>",
				"FilterData<1>",
				"FilterData<2>",
				"FilterData<3>",
				"FilterData<4>",
				"FilterData<5>",
				"FilterData<6>",
				"FilterData<7>",
				"SupportPinLocation<1>",
				"SupportPinLocation<2>",
				"SupportPinLocation<3>",
				"SupportPinLocation<4>",
				"SupportPinLocation<5>",
				"SupportPinLocation<6>",
				"SupportPinLocation<7>",
				"SupportPinExtension<1>",
				"SupportPinExtension<2>",
				"SupportPinExtension<3>",
				"SupportPinExtension<4>",
				"SupportPinExtension<5>",
				"SupportPinExtension<6>",
				"SupportPinExtension<7>",
				"WarpageGroup<1>",
				"WarpageGroup<2>",
				"WarpageGroup<3>",
				"WarpageGroup<4>",
				"WarpageGroup<5>",
				"WarpageGroup<6>",
				"WarpageGroup<7>",
				"RecogSeqInfo<1>",
				"RecogSeqInfo<2>",
				"RecogSeqInfo<3>",
				"RecogSeqInfo<4>",
				"RecogSeqInfo<5>",
				"RecogSeqInfo<6>",
				"RecogSeqInfo<7>",
				"ObstacleData<1>",
				"ObstacleData<2>",
				"ObstacleData<3>",
				"ObstacleData<4>",
				"ObstacleData<5>",
				"ObstacleData<6>",
				"ObstacleData<7>",
				"PlacementGroup<1>",
				"PlacementGroup<2>",
				"PlacementGroup<3>",
				"PlacementGroup<4>",
				"PlacementGroup<5>",
				"PlacementGroup<6>",
				"PlacementGroup<7>",
				"L-R-Matrix<1>",
				"L-R-Matrix<2>",
				"L-R-Matrix<3>",
				"L-R-Matrix<4>",
				"L-R-Matrix<5>",
				"L-R-Matrix<6>",
				"L-R-Matrix<7>",
				"InsertPositionData<1>",
				"InsertPositionData<2>",
				"InsertPositionData<3>",
				"InsertPositionData<4>",
				"InsertPositionData<5>",
				"InsertPositionData<6>",
				"InsertPositionData<7>",
				"InsertPatternData<1>",
				"InsertPatternData<2>",
				"InsertPatternData<3>",
				"InsertPatternData<4>",
				"InsertPatternData<5>",
				"InsertPatternData<6>",
				"InsertPatternData<7>",
				"SoftSwitch<1>",
				"SoftSwitch<2>",
				"SoftSwitch<3>",
				"SoftSwitch<4>",
				"SoftSwitch<5>",
				"SoftSwitch<6>",
				"SoftSwitch<7>",
				"SoftSwitchPerBeam<1>",
				"SoftSwitchPerBeam<2>",
				"SoftSwitchPerBeam<3>",
				"SoftSwitchPerBeam<4>",
				"SoftSwitchPerBeam<5>",
				"SoftSwitchPerBeam<6>",
				"SoftSwitchPerBeam<7>",
				"ComSwitch<1>",
				"ComSwitch<2>",
				"ComSwitch<3>",
				"ComSwitch<4>",
				"ComSwitch<5>",
				"ComSwitch<6>",
				"ComSwitch<7>",
				"PCBConveyance<1>",
				"PCBConveyance<2>",
				"PCBConveyance<3>",
				"PCBConveyance<4>",
				"PCBConveyance<5>",
				"PCBConveyance<6>",
				"PCBConveyance<7>",
				"NozzlePosition<1>",
				"NozzlePosition<2>",
				"NozzlePosition<3>",
				"NozzlePosition<4>",
				"NozzlePosition<5>",
				"NozzlePosition<6>",
				"NozzlePosition<7>",
				"SimplePriority<1>",
				"SimplePriority<2>",
				"SimplePriority<3>",
				"SimplePriority<4>",
				"SimplePriority<5>",
				"SimplePriority<6>",
				"SimplePriority<7>",
				"SeqPriority<1>",
				"SeqPriority<2>",
				"SeqPriority<3>",
				"SeqPriority<4>",
				"SeqPriority<5>",
				"SeqPriority<6>",
				"SeqPriority<7>",
				"RelativePriority<1>",
				"RelativePriority<2>",
				"RelativePriority<3>",
				"RelativePriority<4>",
				"RelativePriority<5>",
				"RelativePriority<6>",
				"RelativePriority<7>",
				"VstockData<1>",
				"VstockData<2>",
				"VstockData<3>",
				"VstockData<4>",
				"VstockData<5>",
				"VstockData<6>",
				"VstockData<7>",
				"VnozzleStock<1>",
				"VnozzleStock<2>",
				"VnozzleStock<3>",
				"VnozzleStock<4>",
				"VnozzleStock<5>",
				"VnozzleStock<6>",
				"VnozzleStock<7>",
				"VpaletteData<1>",
				"VpaletteData<2>",
				"VpaletteData<3>",
				"VpaletteData<4>",
				"VpaletteData<5>",
				"VpaletteData<6>",
				"VpaletteData<7>",
				"FixedPlacementCount<1>",
				"FixedPlacementCount<2>",
				"FixedPlacementCount<3>",
				"FixedPlacementCount<4>",
				"FixedPlacementCount<5>",
				"FixedPlacementCount<6>",
				"FixedPlacementCount<7>",
				"TransUnitSetting<1>",
				"TransUnitSetting<2>",
				"TransUnitSetting<3>",
				"TransUnitSetting<4>",
				"TransUnitSetting<5>",
				"TransUnitSetting<6>",
				"TransUnitSetting<7>",
				"TrayTableMode<1>",
				"TrayTableMode<2>",
				"TrayTableMode<3>",
				"TrayTableMode<4>",
				"TrayTableMode<5>",
				"TrayTableMode<6>",
				"TrayTableMode<7>",
				"MachineTurnInfo<1>",
				"MachineTurnInfo<2>",
				"MachineTurnInfo<3>",
				"MachineTurnInfo<4>",
				"MachineTurnInfo<5>",
				"MachineTurnInfo<6>",
				"MachineTurnInfo<7>",
				"ShareRecogHeadControl<1>",
				"ShareRecogHeadControl<2>",
				"ShareRecogHeadControl<3>",
				"ShareRecogHeadControl<4>",
				"ShareRecogHeadControl<5>",
				"ShareRecogHeadControl<6>",
				"ShareRecogHeadControl<7>",
				"ShareMntTurnControl<1>",
				"ShareMntTurnControl<2>",
				"ShareMntTurnControl<3>",
				"ShareMntTurnControl<4>",
				"ShareMntTurnControl<5>",
				"ShareMntTurnControl<6>",
				"ShareMntTurnControl<7>",
				"GroundData<1>",
				"GroundData<2>",
				"GroundData<3>",
				"GroundData<4>",
				"GroundData<5>",
				"GroundData<6>",
				"GroundData<7>",
				"LandData<1>",
				"LandData<2>",
				"LandData<3>",
				"LandData<4>",
				"LandData<5>",
				"LandData<6>",
				"LandData<7>",
				"SolderData<1>",
				"SolderData<2>",
				"SolderData<3>",
				"SolderData<4>",
				"SolderData<5>",
				"SolderData<6>",
				"SolderData<7>",
				"SolderJdgData<1>",
				"SolderJdgData<2>",
				"SolderJdgData<3>",
				"SolderJdgData<4>",
				"SolderJdgData<5>",
				"SolderJdgData<6>",
				"SolderJdgData<7>",
				"SolderJdgParameter<1>",
				"SolderJdgParameter<2>",
				"SolderJdgParameter<3>",
				"SolderJdgParameter<4>",
				"SolderJdgParameter<5>",
				"SolderJdgParameter<6>",
				"SolderJdgParameter<7>",
				"SolderJdgLookUpTable<1>",
				"SolderJdgLookUpTable<2>",
				"SolderJdgLookUpTable<3>",
				"SolderJdgLookUpTable<4>",
				"SolderJdgLookUpTable<5>",
				"SolderJdgLookUpTable<6>",
				"SolderJdgLookUpTable<7>",
				"ShootingCondition<1>",
				"ShootingCondition<2>",
				"ShootingCondition<3>",
				"ShootingCondition<4>",
				"ShootingCondition<5>",
				"ShootingCondition<6>",
				"ShootingCondition<7>",
				"LookUpTableReference<1>",
				"LookUpTableReference<2>",
				"LookUpTableReference<3>",
				"LookUpTableReference<4>",
				"LookUpTableReference<5>",
				"LookUpTableReference<6>",
				"LookUpTableReference<7>",
				"LookUpTableReference2<1>",
				"LookUpTableReference2<2>",
				"LookUpTableReference2<3>",
				"LookUpTableReference2<4>",
				"LookUpTableReference2<5>",
				"LookUpTableReference2<6>",
				"LookUpTableReference2<7>",
				"InspectViewData<1>",
				"InspectViewData<2>",
				"InspectViewData<3>",
				"InspectViewData<4>",
				"InspectViewData<5>",
				"InspectViewData<6>",
				"InspectViewData<7>",
				"InspectTargetData<1>",
				"InspectTargetData<2>",
				"InspectTargetData<3>",
				"InspectTargetData<4>",
				"InspectTargetData<5>",
				"InspectTargetData<6>",
				"InspectTargetData<7>",
				"InspectViewRetryData<1>",
				"InspectViewRetryData<2>",
				"InspectViewRetryData<3>",
				"InspectViewRetryData<4>",
				"InspectViewRetryData<5>",
				"InspectViewRetryData<6>",
				"InspectViewRetryData<7>",
				"InspectTargetRetryData<1>",
				"InspectTargetRetryData<2>",
				"InspectTargetRetryData<3>",
				"InspectTargetRetryData<4>",
				"InspectTargetRetryData<5>",
				"InspectTargetRetryData<6>",
				"InspectTargetRetryData<7>",
				"ApcCtrlData<1>",
				"ApcCtrlData<2>",
				"ApcCtrlData<3>",
				"ApcCtrlData<4>",
				"ApcCtrlData<5>",
				"ApcCtrlData<6>",
				"ApcCtrlData<7>",
				"NozzleSelect<1>",
				"NozzleSelect<2>",
				"NozzleSelect<3>",
				"NozzleSelect<4>",
				"NozzleSelect<5>",
				"NozzleSelect<6>",
				"NozzleSelect<7>",
				"DischargeData<1>",
				"DischargeData<2>",
				"DischargeData<3>",
				"DischargeData<4>",
				"DischargeData<5>",
				"DischargeData<6>",
				"DischargeData<7>",
				"DispensePositionData<1>",
				"DispensePositionData<2>",
				"DispensePositionData<3>",
				"DispensePositionData<4>",
				"DispensePositionData<5>",
				"DispensePositionData<6>",
				"DispensePositionData<7>",
				"NozzleStock",
				"StockData",
				"NonArrangeParts",
				"PaletteData",
				"PartsListData",
				"FixedPairParts",
				"PairPartsArrange",
				"PartsData",
				"PartsGroup",
				"PartsGroupNames",
				"PartsLabel",
				"AlterSupplyPartGroup",
				"SameNameSubstitutePart",
				"VendorKeywordLink",
				"RecognitionMark",
				"ShapeBase",
				"ShapeBaseShape",
				"ShapeBasePosUp",
				"ShapeBasePosMethod",
				"ShapeBaseChk",
				"ShapeBaseEtc",
				"ShapePol",
				"ShapeMold",
				"ShapeLeadGroup",
				"ShapeBallGroup",
				"ShapeInsLead",
				"ShapeSquare",
				"ShapeCircle",
				"ShapeLine",
				"ShapeCorner",
				"ShapeRectangle",
				"ShapeUnevenBall",
				"ShapeUnevenBallB",
				"ShapeMntTeachPos",
				"ShapeCutLead",
				"ShapeAutoCutLead",
				"ShapeLackBall",
				"ShapeBlockMatch",
				"ShapeRandomBall",
				"ShapeBoxSide",
				"ShapeLeadPoint",
				"ShapeLeadBend",
				"ShapeLeadIrr",
				"ShapeReinforcedLead",
				"ShapeInsLeadDetail",
				"ShapeIrrRange",
				"ShapeTakeupRange",
				"ShapeAccRange",
				"ShapeBoxPos",
				"ShapeBoxAccPos",
				"ShapeAccSlant",
				"ShapePointCorrect",
				"ShapeAllBallInspect",
				"ShapeNoAngle",
				"ShapeNoCenter",
				"ShapeAngleCalcMethod",
				"ShapeLeadDetectPos",
				"ShapeLeadDetectWidth",
				"ShapeLeadHPos",
				"ShapeBallHPos",
				"ShapeCenterDef",
				"ShapeMountOffset",
				"ShapeCorrectionFlg",
				"ShapeSizeAllow",
				"ShapeLeadWAllow",
				"ShapeLeadBendAllow",
				"ShapeFloatAllow",
				"ShapeFloatAllowMaster",
				"ShapeLeadLackChk",
				"ShapeRegistPosChk",
				"ShapeSizeChk",
				"ShapeAngleCalcChk",
				"ShapeCenterRangeChk",
				"ShapeTakeUpChk",
				"ShapeOverChk",
				"ShapeAdjacentChk",
				"ShapeSupplyDirChk",
				"ShapeLeadDirChk",
				"ShapeSideStandChk",
				"ShapeStandChk",
				"ShapeRollChk",
				"ShapeBrightChk",
				"ShapeFloatChk",
				"ShapeFluxChk",
				"ShapeSolderChk",
				"ShapePOPRecChk",
				"ShapeDblPickUpChk",
				"ShapeLeadFlipFlopChk",
				"ShapeBrightDiffChk",
				"ShapePOPRecChkEx",
				"ShapeBumpBridgeChk",
				"ShapeInverse",
				"ShapeEasySupportFlag",
				"ShapeEdgePrm",
				"ShapeTriPrm",
				"ShapeSamplingPrm",
				"ShapeAccuPrm",
				"Shape2DPrm",
				"Shape3DPrm",
				"Shape3DSensor",
				"Shape3DSensorPSD",
				"ShapeOldFashion",
				"ShapeInsLeadDetailMaster",
				"ShapeCMRef",
				"ShapeRndShpExSize",
				"ShapeOutlineRecog",
				"ShapeAutoAngleAdjust",
				"ShapeSupportPin",
				"ChipData",
				"ChipExtensionQFP",
				"ChipExtensionBGA",
				"ChipExtensionCT",
				"ChipExtensionFC",
				"ChipExtensionBARE",
				"ChipExtensionGR",
				"GRDetailPINFRM",
				"GRDetailARYFORM",
				"GRDetailNONGRID",
				"GRDetailPINAD",
				"GRDetailBMPAD",
				"GRDetailPINGAD",
				"GRDetailBMPGAD",
				"GRDetailINSPGAD",
				"GRDetailRTAD",
				"GRDetailCIRAD",
				"GRDetailLNAD",
				"GRDetailCORAD",
				"GRDetailRTGAD",
				"GRDetailPINMS",
				"GRDetailPINBS",
				"GRDetailBMPARYBT",
				"GRDetailGRBINA",
				"GRDetailGRBINB",
				"GRDetailGRBINC",
				"MountCondition",
				"TrayData",
				"FeederData",
				"FeederMaxData",
				"NozzleDataNPM",
				"NozzleDataCM",
				"NozzleMaxData",
				"SupplyInspectConditionSet",
				"SupplyInspectPolarityCondition",
				"SupplyInspectVerifyCondition",
				"SupplyInspectTraceCondition",
				"SupplyInspectExistCondition",
				"SupplyInspectPosCondition",
				"SupplyInspectPolarityRecog",
				"SupplyInspectCharacterRecog",
				"SupplyInspect2DCodeRecog",
				"SupplyInspectExistRecog",
				"SupplyInspectPosRecog",
				"HeadCamRecMountCondition",
				"HeadCamInspectionCondition",
				"HeadCamRecData",
				"FeederExSupplyInfo",
				"LCRChkData",
				"LCRChkDataCmn",
				"LCRChkDataEQ",
				"ChipJdgData",
				"SpecialJdgData",
				"GrandLibData",
				"InspectConditionSet",
				"PartInspectCondition",
				"InspectBoxSet",
				"LookUpTable",
				"InspectAlgo",
				"InspectBoxParameter",
				"InspectPivotSet",
				"InspectPivotData",
				"LampStockData",
				"PartsInspectAreaCondition",
				"SimplePartsInspectCondition",
				"PartsInspectSkipKind",
				"DNozzleShapeData",
				"BondData",
				"DNozzleControlData",
				"DispenseCondition",
				"MultiDisData",
				"DrawData",
				"DrawPosData",
				"DischargeRecogData",
				"PrintCondition",
				"Cleaning",
				"PrintCounter",
				"MultiSpeedEditData",
				"InsertCondition",
				"InsertLeadTwist",
				"NozzleDataBodyChuck",
				"NozzleDataLeadChuck",
				"NozzleDataSwing",
				"FixedFeeder",
				"FixedNozzle",
				"FixedNozzleStock",
				"VagueFixFeeder",
				"VagueFixFeederSetupRules",
				"VagueFixFeederSetupRuleCondition",
				"VagueFixFeederSetupRuleAssign",
				"ForceArrangeBeam",
				"CarryoverStockData",
				"BeamData",
				"FdrTblMode",
				"LineInfoData",
				"InsertHeadType",
				"StageTactInfo<1>",
				"StageTactInfo<2>",
				"StageTactInfo<3>",
				"StageTactInfo<4>",
				"StageTactInfo<5>",
				"StageTactInfo<6>",
				"StageTactInfo<7>",
				"LaneStageTactInfo<1>",
				"LaneStageTactInfo<2>",
				"LaneStageTactInfo<3>",
				"LaneStageTactInfo<4>",
				"LaneStageTactInfo<5>",
				"LaneStageTactInfo<6>",
				"LaneStageTactInfo<7>",
				"HeadInfo<1>",
				"HeadInfo<2>",
				"HeadInfo<3>",
				"HeadInfo<4>",
				"HeadInfo<5>",
				"HeadInfo<6>",
				"HeadInfo<7>",
				"SimultaneousProduct",
				"ArrangementGroupProducts",
				"ArrangementGroupCondition",
				"OptimizeOptionsPerBeam"},
		}, text)
		// "Brain-1.crb"
		// text2
		// "/home/eugenearch/Code/github.com/eugenefoxx/test/readIni/DGS_files/Extract/Brain-1.crb"
		if err != nil {
			fmt.Printf("Fail to read file: %v\n", err)
			os.Exit(1)
		}
		//	ss := "NPM_brain-1_p"
		//	dd := "1"

		//	tt := "BoardData<" + dd + ">"
		//	var str string
		// Classic read of values, default section can be represented as empty string
		//	fmt.Println("App Mode:", cfg.Section(tt).Key("Z").String())
		//	fmt.Println("GGG:", cfg.Section("LotNames").Body())
		str := cfg.Section("LotNames").Body()
		fmt.Println(str)
		fmt.Println("strings.ToLower:", strings.Contains(strings.ToLower(str), strings.ToLower(s)))
		errContains := strings.Contains(strings.ToLower(str), strings.ToLower(s))
		if errContains != true {
			fmt.Println("Шаблон рецепта не найден.Попробуйте изменить регистр P/p или S/s.")
			log.Fatal(errContains)
		}
		d1 := []byte(str)
		//err1 := ioutil.WriteFile("/home/a20272/Code/github.com/eugenefoxx/readDGS/files/datLotNames", d1, 0644)
		err1 := ioutil.WriteFile(setting.Conf.Files+"datLotNames", d1, 0644)
		check(err1)

		// lotnamesdataSide - идентификатор для парсинга нужной стороны платы
		//lotnamesdataSide := lotnamesdata.LotNamesData("/home/a20272/Code/github.com/eugenefoxx/readDGS/files/datLotNames", s)
		lotnamesdataSide := lotnamesdata.LotNamesData(setting.Conf.Files+"datLotNames", s)
		//lotnamesdata.LotNamesData("data1")
		//lotnamesData.LotNamesData("dat1")
		fmt.Println("Ответ функции lotnamesdata.LotNamesData:", lotnamesdataSide)
		patternData := "PatternData<" + lotnamesdataSide + ">"
		str2 := cfg.Section(patternData).Body()
		d2 := []byte(str2)

		//err2 := ioutil.WriteFile("/home/a20272/Code/github.com/eugenefoxx/readDGS/files/datpatternData", d2, 0644)
		err2 := ioutil.WriteFile(setting.Conf.Files+"datpatternData", d2, 0644)
		check(err2)

		//	sed -i '1d' filename
		//delCmd := exec.Command("bash", "-c", "sed -i '1d' /home/a20272/Code/github.com/eugenefoxx/readDGS/files/datpatternData")
		delCmd := exec.Command("bash", "-c", "sed -i '1d' \""+setting.Conf.Files+"\"datpatternData")
		_, errdelCmd := delCmd.Output()
		if errdelCmd != nil {
			panic(errdelCmd)
		}

		//awk 'BEGIN {FS = OFS = " "} {s[$6] += 1} END {for (key in s) {print key, s[key]}}' datpatternData > outdatpatternData
		//countCmd := exec.Command("bash", "-c", "awk 'BEGIN {FS = OFS = " "} {s[$6] += 1} END {for (key in s) {print key, s[key]}}' /home/eugenearch/Code/github.com/eugenefoxx/test/readIni/datpatternData > /home/eugenearch/Code/github.com/eugenefoxx/test/readIni/outdatpatternData")
		//countCmd := exec.Command("/bin/sh", "-c", "/home/a20272/Code/github.com/eugenefoxx/readDGS/count.sh")
		countCmd := exec.Command("/bin/sh", "-c", setting.Conf.Bash+"count.sh")
		_, errdcountCmd := countCmd.Output()
		if errdcountCmd != nil {
			panic(errdcountCmd)
		}

		//	[PartsData]
		//	partsData := cfg.Section("PartsData").Body()
		str3 := cfg.Section("PartsData").Body()
		d3 := []byte(str3)

		//err3 := ioutil.WriteFile("/home/a20272/Code/github.com/eugenefoxx/readDGS/files/datpartsData", d3, 0644)
		err3 := ioutil.WriteFile(setting.Conf.Files+"datpartsData", d3, 0644)
		check(err3)
	}

	if *lineL2 {
		cmd := &exec.Cmd{
			//Path: "/home/a20272/Code/github.com/eugenefoxx/readDGS/fileL2.sh",
			Path: setting.Conf.Bash + "fileL2.sh",
			//	Args: []string{"/home/a20272/Code/github.com/eugenefoxx/readDGS/fileL2.sh", s},
			Args:   []string{setting.Conf.Bash + "fileL2.sh", s},
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}
		cmd.Run()

		// Read entire file content, giving us little control but
		// making it very simple. No need to close the file.
		//content, err := ioutil.ReadFile("/home/a20272/Code/github.com/eugenefoxx/readDGS/DGS_files/ExtractL2/fileresult.txt")
		content, err := ioutil.ReadFile(setting.Conf.DGSExtractL2 + "fileresult.txt")
		if err != nil {
			log.Fatal(err)
		}

		// Convert []byte to string and print to screen
		text := string(content)
		// чищу от переноса строки
		text = strings.TrimSuffix(text, "\n")
		//	text2 := strconv.Quote(text)
		//	fmt.Println("text2 -", text2)
		//	cfg, err := ini.Load("Brain-1.crb")
		cfg, err := ini.LoadSources(ini.LoadOptions{
			UnparseableSections: []string{"LotNames",
				"MultiPattern<1>",
				"MultiPattern<2>",
				"PatternAttribute<1>",
				"PatternAttribute<2>",
				"PatternGrouping<1>",
				"PatternGrouping<2>",
				"RecogMarkSet<1>",
				"RecogMarkSet<2>",
				"PositionData<1>",
				"PositionData<2>",
				"PatternData<1>",
				"PatternData<2>",
				"FilterData<1>",
				"FilterData<2>",
				"SupportPinLocation<1>",
				"SupportPinLocation<2>",
				"SupportPinExtension<1>",
				"SupportPinExtension<2>",
				"WarpageGroup<1>",
				"WarpageGroup<2>",
				"RecogSeqInfo<1>",
				"RecogSeqInfo<2>",
				"ObstacleData<1>",
				"ObstacleData<2>",
				"PlacementGroup<1>",
				"PlacementGroup<2>",
				"L-R-Matrix<1>",
				"L-R-Matrix<2>",
				"InsertPositionData<1>",
				"InsertPositionData<2>",
				"InsertPatternData<1>",
				"InsertPatternData<2>",
				"SoftSwitch<1>",
				"SoftSwitch<2>",
				"SoftSwitchPerBeam<1>",
				"SoftSwitchPerBeam<2>",
				"ComSwitch<1>",
				"ComSwitch<2>",
				"PCBConveyance<1>",
				"PCBConveyance<2>",
				"NozzlePosition<1>",
				"NozzlePosition<2>",
				"SimplePriority<1>",
				"SimplePriority<2>",
				"SeqPriority<1>",
				"SeqPriority<2>",
				"RelativePriority<1>",
				"RelativePriority<2>",
				"VstockData<1>",
				"VstockData<2>",
				"VnozzleStock<1>",
				"VnozzleStock<2>",
				"VpaletteData<1>",
				"VpaletteData<2>",
				"FixedPlacementCount<1>",
				"FixedPlacementCount<2>",
				"TransUnitSetting<1>",
				"TransUnitSetting<2>",
				"TrayTableMode<1>",
				"TrayTableMode<2>",
				"MachineTurnInfo<1>",
				"MachineTurnInfo<2>",
				"ShareRecogHeadControl<1>",
				"ShareRecogHeadControl<2>",
				"ShareMntTurnControl<1>",
				"ShareMntTurnControl<2>",
				"GroundData<1>",
				"GroundData<2>",
				"LandData<1>",
				"LandData<2>",
				"SolderData<1>",
				"SolderData<2>",
				"SolderJdgData<1>",
				"SolderJdgData<2>",
				"SolderJdgParameter<1>",
				"SolderJdgParameter<2>",
				"SolderJdgLookUpTable<1>",
				"SolderJdgLookUpTable<2>",
				"ShootingCondition<1>",
				"ShootingCondition<2>",
				"LookUpTableReference<1>",
				"LookUpTableReference<2>",
				"LookUpTableReference2<1>",
				"LookUpTableReference2<2>",
				"InspectViewData<1>",
				"InspectViewData<2>",
				"InspectTargetData<1>",
				"InspectTargetData<2>",
				"InspectViewRetryData<1>",
				"InspectViewRetryData<2>",
				"InspectTargetRetryData<1>",
				"InspectTargetRetryData<2>",
				"ApcCtrlData<1>",
				"ApcCtrlData<2>",
				"NozzleSelect<1>",
				"NozzleSelect<2>",
				"DischargeData<1>",
				"DischargeData<2>",
				"DispensePositionData<1>",
				"DispensePositionData<2>",
				"NozzleStock",
				"StockData",
				"NonArrangeParts",
				"PaletteData",
				"PartsListData",
				"FixedPairParts",
				"PairPartsArrange",
				"PartsData",
				"PartsGroup",
				"PartsGroupNames",
				"PartsLabel",
				"AlterSupplyPartGroup",
				"SameNameSubstitutePart",
				"VendorKeywordLink",
				"RecognitionMark",
				"ShapeBase",
				"ShapeBaseShape",
				"ShapeBasePosUp",
				"ShapeBasePosMethod",
				"ShapeBaseChk",
				"ShapeBaseEtc",
				"ShapePol",
				"ShapeMold",
				"ShapeLeadGroup",
				"ShapeBallGroup",
				"ShapeInsLead",
				"ShapeSquare",
				"ShapeCircle",
				"ShapeLine",
				"ShapeCorner",
				"ShapeRectangle",
				"ShapeUnevenBall",
				"ShapeUnevenBallB",
				"ShapeMntTeachPos",
				"ShapeCutLead",
				"ShapeAutoCutLead",
				"ShapeLackBall",
				"ShapeBlockMatch",
				"ShapeRandomBall",
				"ShapeBoxSide",
				"ShapeLeadPoint",
				"ShapeLeadBend",
				"ShapeLeadIrr",
				"ShapeReinforcedLead",
				"ShapeInsLeadDetail",
				"ShapeIrrRange",
				"ShapeTakeupRange",
				"ShapeAccRange",
				"ShapeBoxPos",
				"ShapeBoxAccPos",
				"ShapeAccSlant",
				"ShapePointCorrect",
				"ShapeAllBallInspect",
				"ShapeNoAngle",
				"ShapeNoCenter",
				"ShapeAngleCalcMethod",
				"ShapeLeadDetectPos",
				"ShapeLeadDetectWidth",
				"ShapeLeadHPos",
				"ShapeBallHPos",
				"ShapeCenterDef",
				"ShapeMountOffset",
				"ShapeCorrectionFlg",
				"ShapeSizeAllow",
				"ShapeLeadWAllow",
				"ShapeLeadBendAllow",
				"ShapeFloatAllow",
				"ShapeFloatAllowMaster",
				"ShapeLeadLackChk",
				"ShapeRegistPosChk",
				"ShapeSizeChk",
				"ShapeAngleCalcChk",
				"ShapeCenterRangeChk",
				"ShapeTakeUpChk",
				"ShapeOverChk",
				"ShapeAdjacentChk",
				"ShapeSupplyDirChk",
				"ShapeLeadDirChk",
				"ShapeSideStandChk",
				"ShapeStandChk",
				"ShapeRollChk",
				"ShapeBrightChk",
				"ShapeFloatChk",
				"ShapeFluxChk",
				"ShapeSolderChk",
				"ShapePOPRecChk",
				"ShapeDblPickUpChk",
				"ShapeLeadFlipFlopChk",
				"ShapeBrightDiffChk",
				"ShapePOPRecChkEx",
				"ShapeBumpBridgeChk",
				"ShapeInverse",
				"ShapeEasySupportFlag",
				"ShapeEdgePrm",
				"ShapeTriPrm",
				"ShapeSamplingPrm",
				"ShapeAccuPrm",
				"Shape2DPrm",
				"Shape3DPrm",
				"Shape3DSensor",
				"Shape3DSensorPSD",
				"ShapeOldFashion",
				"ShapeInsLeadDetailMaster",
				"ShapeCMRef",
				"ShapeRndShpExSize",
				"ShapeOutlineRecog",
				"ShapeAutoAngleAdjust",
				"ShapeSupportPin",
				"ChipData",
				"ChipExtensionQFP",
				"ChipExtensionBGA",
				"ChipExtensionCT",
				"ChipExtensionFC",
				"ChipExtensionBARE",
				"ChipExtensionGR",
				"GRDetailPINFRM",
				"GRDetailARYFORM",
				"GRDetailNONGRID",
				"GRDetailPINAD",
				"GRDetailBMPAD",
				"GRDetailPINGAD",
				"GRDetailBMPGAD",
				"GRDetailINSPGAD",
				"GRDetailRTAD",
				"GRDetailCIRAD",
				"GRDetailLNAD",
				"GRDetailCORAD",
				"GRDetailRTGAD",
				"GRDetailPINMS",
				"GRDetailPINBS",
				"GRDetailBMPARYBT",
				"GRDetailGRBINA",
				"GRDetailGRBINB",
				"GRDetailGRBINC",
				"MountCondition",
				"TrayData",
				"FeederData",
				"FeederMaxData",
				"NozzleDataNPM",
				"NozzleDataCM",
				"NozzleMaxData",
				"SupplyInspectConditionSet",
				"SupplyInspectPolarityCondition",
				"SupplyInspectVerifyCondition",
				"SupplyInspectTraceCondition",
				"SupplyInspectExistCondition",
				"SupplyInspectPosCondition",
				"SupplyInspectPolarityRecog",
				"SupplyInspectCharacterRecog",
				"SupplyInspect2DCodeRecog",
				"SupplyInspectExistRecog",
				"SupplyInspectPosRecog",
				"HeadCamRecMountCondition",
				"HeadCamInspectionCondition",
				"HeadCamRecData",
				"FeederExSupplyInfo",
				"LCRChkData",
				"LCRChkDataCmn",
				"LCRChkDataEQ",
				"ChipJdgData",
				"SpecialJdgData",
				"GrandLibData",
				"InspectConditionSet",
				"PartInspectCondition",
				"InspectBoxSet",
				"LookUpTable",
				"InspectAlgo",
				"InspectBoxParameter",
				"InspectPivotSet",
				"InspectPivotData",
				"LampStockData",
				"PartsInspectAreaCondition",
				"SimplePartsInspectCondition",
				"PartsInspectSkipKind",
				"DNozzleShapeData",
				"BondData",
				"DNozzleControlData",
				"DispenseCondition",
				"MultiDisData",
				"DrawData",
				"DrawPosData",
				"DischargeRecogData",
				"PrintCondition",
				"Cleaning",
				"PrintCounter",
				"MultiSpeedEditData",
				"InsertCondition",
				"InsertLeadTwist",
				"NozzleDataBodyChuck",
				"NozzleDataLeadChuck",
				"NozzleDataSwing",
				"FixedFeeder",
				"FixedNozzle",
				"FixedNozzleStock",
				"VagueFixFeeder",
				"VagueFixFeederSetupRules",
				"VagueFixFeederSetupRuleCondition",
				"VagueFixFeederSetupRuleAssign",
				"ForceArrangeBeam",
				"CarryoverStockData",
				"BeamData",
				"FdrTblMode",
				"LineInfoData",
				"InsertHeadType",
				"StageTactInfo<1>",
				"StageTactInfo<2>",
				"LaneStageTactInfo<1>",
				"LaneStageTactInfo<2>",
				"HeadInfo<1>",
				"HeadInfo<2>",
				"SimultaneousProduct",
				"ArrangementGroupProducts",
				"ArrangementGroupCondition",
				"OptimizeOptionsPerBeam"},
		}, text)
		// "Brain-1.crb"
		// text2
		// "/home/eugenearch/Code/github.com/eugenefoxx/test/readIni/DGS_files/Extract/Brain-1.crb"
		if err != nil {
			fmt.Printf("Fail to read file: %v\n", err)
			os.Exit(1)
		}
		//	ss := "NPM_brain-1_p"
		//	dd := "1"

		//	tt := "BoardData<" + dd + ">"
		//	var str string
		// Classic read of values, default section can be represented as empty string
		//	fmt.Println("App Mode:", cfg.Section(tt).Key("Z").String())
		//	fmt.Println("GGG:", cfg.Section("LotNames").Body())
		str := cfg.Section("LotNames").Body()
		fmt.Println(str)
		fmt.Println(strings.Contains(strings.ToLower(str), strings.ToLower(s)))
		errContains := strings.Contains(strings.ToLower(str), strings.ToLower(s))
		if errContains != true {
			fmt.Println("Шаблон рецепта не найден.Попробуйте изменить регистр P/p или S/s.")
			log.Fatal(errContains)
		}
		d1 := []byte(str)
		//err1 := ioutil.WriteFile("/home/a20272/Code/github.com/eugenefoxx/readDGS/files/datLotNames", d1, 0644)
		err1 := ioutil.WriteFile(setting.Conf.Files+"datLotNames", d1, 0644)
		check(err1)

		// lotnamesdataSide - идентификатор для парсинга нужной стороны платы
		//lotnamesdataSide := lotnamesdata.LotNamesData("/home/a20272/Code/github.com/eugenefoxx/readDGS/files/datLotNames", s)
		lotnamesdataSide := lotnamesdata.LotNamesData(setting.Conf.Files+"datLotNames", s)
		//lotnamesdata.LotNamesData("data1")
		//lotnamesData.LotNamesData("dat1")
		fmt.Println("Ответ функции lotnamesdata.LotNamesData:", lotnamesdataSide)
		patternData := "PatternData<" + lotnamesdataSide + ">"
		str2 := cfg.Section(patternData).Body()
		d2 := []byte(str2)

		//err2 := ioutil.WriteFile("/home/a20272/Code/github.com/eugenefoxx/readDGS/files/datpatternData", d2, 0644)
		err2 := ioutil.WriteFile(setting.Conf.Files+"datpatternData", d2, 0644)
		check(err2)

		//	sed -i '1d' filename
		//delCmd := exec.Command("bash", "-c", "sed -i '1d' /home/a20272/Code/github.com/eugenefoxx/readDGS/files/datpatternData")
		delCmd := exec.Command("bash", "-c", "sed -i '1d' \""+setting.Conf.Files+"\"datpatternData")
		_, errdelCmd := delCmd.Output()
		if errdelCmd != nil {
			panic(errdelCmd)
		}

		//awk 'BEGIN {FS = OFS = " "} {s[$6] += 1} END {for (key in s) {print key, s[key]}}' datpatternData > outdatpatternData
		//countCmd := exec.Command("bash", "-c", "awk 'BEGIN {FS = OFS = " "} {s[$6] += 1} END {for (key in s) {print key, s[key]}}' /home/eugenearch/Code/github.com/eugenefoxx/test/readIni/datpatternData > /home/eugenearch/Code/github.com/eugenefoxx/test/readIni/outdatpatternData")
		//countCmd := exec.Command("/bin/sh", "-c", "/home/a20272/Code/github.com/eugenefoxx/readDGS/count.sh")
		countCmd := exec.Command("/bin/sh", "-c", setting.Conf.Bash+"count.sh")
		_, errdcountCmd := countCmd.Output()
		if errdcountCmd != nil {
			panic(errdcountCmd)
		}

		//	[PartsData]
		//	partsData := cfg.Section("PartsData").Body()
		str3 := cfg.Section("PartsData").Body()
		d3 := []byte(str3)

		//err3 := ioutil.WriteFile("/home/a20272/Code/github.com/eugenefoxx/readDGS/files/datpartsData", d3, 0644)
		err3 := ioutil.WriteFile(setting.Conf.Files+"datpartsData", d3, 0644)
		check(err3)
	}

	//recipesap.MakeRecipeSAP("/home/a20272/Code/github.com/eugenefoxx/readDGS/files/outdatpatternData", "/home/a20272/Code/github.com/eugenefoxx/readDGS/files/datpartsData", s)
	recipesap.MakeRecipeSAP(setting.Conf.Files+"outdatpatternData", setting.Conf.Files+"datpartsData", s)

	duration := time.Since(start)
	fmt.Println("Время работы - ", duration)
	//open.RunWith("/home/a20272/Code/github.com/eugenefoxx/readDGS/files/Рецепт для SAP.csv", "soffice")
	open.RunWith(setting.Conf.Files+"recipte.csv", "soffice")

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
