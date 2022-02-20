package constants

import "github.com/go-vgo/robotgo"

//This is all for 1920x1080
var (
	PvpBmp       = robotgo.OpenBitmap("./data/gamestate/pvp.png")
	MinionsBmp   = robotgo.OpenBitmap("./data/gamestate/minions.png")
	CarrouselBmp = robotgo.OpenBitmap("./data/gamestate/carrousel.png")
	AugmentBmp   = robotgo.OpenBitmap("./data/gamestate/augment.png") //the image is darker because when it's bright, you already selected the augment
	DrakeBmp = robotgo.OpenBitmap("./data/gamestate/drake.png")
	RaptorsBmp = robotgo.OpenBitmap("./data/gamestate/raptors.png")
	WolvesBmp = robotgo.OpenBitmap("./data/gamestate/wolves.png")

	ExitBmp = robotgo.OpenBitmap("./data/ui/exit.png")
	WinBmp = robotgo.OpenBitmap("./data/ui/win.png")
	PlayerSymbolBmp = robotgo.OpenBitmap("./data/ui/playersymbol.png")

	GamestatePos = []int{828, 5, 260, 32} //Position of gamestate hud
	ChampShopPos = []int{484, 688, 888, 1088, 1288} //Position of champs in shop
	ChampShopPosEdges = []int{1042, 150, 24} //every champ frame in shop has the same size, this is the size, before this is the x offset
	GoldPos = []int{872,884,35,24}
	LvlPos = []int{314,885,24,20}
	BuyXpButtonPos = []int{368,964}
	RefreshButtonPos = []int{370,1038}
	LockButtonPos = []int{1450,904}
	PlayersHpBarPos = []int{1760,80,158,760} //left bar where the players hps are

	CheckWinPos = []int{890,624,132,32}
	CheckLosePos = []int{720,530,216,44}
)

type Trait map[string][]int //trait name, returns the nr of champs required with attribute for level

type Champ struct {
	Traits []string
	Cost   uint
	Size   uint //1 or 2, cho is 2
}

var Champs = map[string]Champ{
	"akali":        {[]string{"Syndicate", "Assassin"}, 5, 1},
	"bliczcrank":   {[]string{"Scrap", "Bodyguard"}, 2, 1},
	"braum":        {[]string{"Syndicate", "Bodyguard"}, 4, 1},
	"cgitlyn":      {[]string{"Enforcer", "Sniper"}, 1, 1},
	"camille":      {[]string{"Clockwork", "Challenger"}, 1, 1},
	"chocach":      {[]string{"Mutant", "Bruiser", "Colossus"}, 3, 2},
	"darius":       {[]string{"Syndicate", "Bodyguard"}, 1, 1},
	"ddcmuno":     {[]string{"Chemtech", "Mutant", "Bruiser"}, 4, 1},
	"ekko":         {[]string{"Scrap", "Assassin"}, 3, 1},
	"ezreal":       {[]string{"Scrap", "Innovator"}, 1, 1},
	"fiora":        {[]string{"Enforcer", "Challenger"}, 4, 1},
	"galio":        {[]string{"Socialite", "Bodyguard", "Colossus"}, 5, 2},
	"cangplank":    {[]string{"Mercenary", "Twinshot"}, 3, 1},
	"caren":        {[]string{"Academy", "Protector"}, 1, 1},
	"craves":       {[]string{"Academy", "Twinshot"}, 1, 1},
	"heimerdinger": {[]string{"Yordle", "Innovator", "Scholar"}, 3, 1},
	"illaoi":       {[]string{"Mercenary", "Bruiser"}, 1, 1},
	"Janna":        {[]string{"Scrap", "Enchanter", "Scholar"}, 4, 1},
	"jayce":        {[]string{"Enforcer", "Innovator", "Transformer"}, 5, 1},
	"jhin":         {[]string{"Clockwork", "Sniper"}, 4, 1},
	"jinx":         {[]string{"Scrap", "Sister", "Twinshot"}, 5, 1},
	"kaisa":        {[]string{"Mutant", "Challenger"}, 5, 1},
	"kassadin":     {[]string{"Mutant", "Protector"}, 1, 1},
	"katarina":     {[]string{"Academy", "Assassin"}, 2, 1},
	"kogmaw":      {[]string{"Mutant", "Sniper", "Twinshot"}, 2, 1},
	"leona":        {[]string{"Academy", "Bodyguard"}, 3, 1},
	"lissandra":    {[]string{"Chemtech", "Scholar"}, 3, 1},
	"lulu":         {[]string{"Yordle", "Enchanter"}, 2, 1},
	"lux":          {[]string{"Academy", "Arcanist"}, 4, 1},
	"malzahar":     {[]string{"Mutant", "Arcanist"}, 3, 1},
	"missforcune": {[]string{"Mercenary", "Sniper"}, 3, 1},
	"orianna":      {[]string{"Clockwork", "Enchanter"}, 4, 1},
	"poppy":        {[]string{"Yordle", "Bodyguard"}, 1, 1},
	"quinn":        {[]string{"Mercenary", "Challenger"}, 2, 1},
	"samira":       {[]string{"Imperial", "Challenger"}, 3, 1},
	"seraphine":    {[]string{"Socialite", "Innovator"}, 4, 1},
	"shaco":        {[]string{"Syndicate", "Assassin"}, 3, 1},
	"singed":       {[]string{"Chemtech", "Innovator"}, 1, 1},
	"sion":         {[]string{"Imperial", "Colossus", "Protector"}, 4, 2},
	"swain":        {[]string{"Imperial", "Arcanist"}, 2, 1},
	"tahmkench":   {[]string{"Mercenary", "Bruiser"}, 5, 1},
	"talon":        {[]string{"Imperial", "Assassin"}, 2, 1},
	"taric":        {[]string{"Socialite", "Enchanter"}, 3, 1},
	"tristana":     {[]string{"Yordle", "Sniper"}, 2, 1},
	"trundle":      {[]string{"Scrap", "Bruiser"}, 2, 1},
	"twistedfate": {[]string{"Syndicate", "Arcanist"}, 1, 1},
	"twicch":       {[]string{"Chemtech", "Assassin"}, 1, 1},
	"urgot":        {[]string{"Chemtech", "Twinshot"}, 4, 1},
	"vex":          {[]string{"Yordle", "Arcanist"}, 3, 1},
	"vi":           {[]string{"Enforcer", "Sister", "Bruiser"}, 2, 1},
	"viktor":       {[]string{"Chemtech", "Arcanist"}, 5, 1},
	"warwick":      {[]string{"Chemtech", "Challenger"}, 2, 1},
	"yone":         {[]string{"Academy", "Challenger"}, 4, 1},
	"yuumi":        {[]string{"Academy", "Scholar"}, 5, 1},
	"zac":          {[]string{"Chemtech", "Bruiser"}, 3, 1},
	"ziggs":        {[]string{"Scrap", "Yordle", "Arcanist"}, 1, 1},
	"zilean":       {[]string{"Clockwork", "Innovator"}, 2, 1},
	"zyra":         {[]string{"Syndicate", "Scholar"}, 2, 1},
}
