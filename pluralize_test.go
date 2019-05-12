package pluralize

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/gertd/go-pluralize/cmd/pluralize/tflags"
)

type TestEntry struct {
	input    string
	expected string
}

type params struct {
	passLog *bool
	word    *string
	cmd     *string
}

var (
	p params //nolint:gochecknoglobals
)

func TestMain(m *testing.M) {

	p.passLog = flag.Bool("pass", false, "log PASS results")
	p.word = flag.String("word", "", "input value")
	p.cmd = flag.String("cmd", "all", "command name [optional]")

	flag.Parse()

	os.Exit(m.Run())
}

// plog -- PASSED result log
func plog(t *testing.T, format string, a ...interface{}) {
	if *p.passLog {
		t.Logf(format, a...)
	}
}

// slog -- Summary result log
func slog(name string, passed int, failed int, total int) {

	fmt.Fprintf(os.Stdout, ">>> %s PASSED=%d FAILED=%d OF %d\n",
		name, passed, failed, total)
}

func TestCmd(t *testing.T) {
	if p.word == nil || len(*p.word) == 0 {
		t.SkipNow()
		return
	}

	testCmd := tflags.TestCmdString(*p.cmd)

	if testCmd.Has(tflags.TestCmdUnknown) {
		t.Error(fmt.Errorf("unknown -cmd value %s, valid [All|IsPlural|IsSingular|Plural|Singular]", *p.cmd))
		return
	}

	pluralize := NewClient()

	if testCmd.Has(tflags.TestCmdIsPlural) {
		t.Logf("IsPlural(%s)   => %t\n", *p.word, pluralize.IsPlural(*p.word))
	}
	if testCmd.Has(tflags.TestCmdIsSingular) {
		t.Logf("IsSingular(%s) => %t\n", *p.word, pluralize.IsSingular(*p.word))
	}
	if testCmd.Has(tflags.TestCmdPlural) {
		t.Logf("Plural(%s)     => %s\n", *p.word, pluralize.Plural(*p.word))
	}
	if testCmd.Has(tflags.TestCmdSingular) {
		t.Logf("Singular(%s)   => %s\n", *p.word, pluralize.Singular(*p.word))
	}
}

func TestIsPlural(t *testing.T) {

	tests := append(basicTests(), pluralTests()...)
	passed := 0
	failed := 0

	pluralize := NewClient()

	for i, testItem := range tests {
		if actual := pluralize.IsPlural(testItem.expected); actual == true {
			plog(t, "PASS test[%d] func %s(%s) expected %t, actual %t", i, "IsPlural",
				testItem.input, true, actual)
			passed++
		} else {
			t.Errorf("FAIL test[%d] func %s(%s) expected %t, actual %t", i, "IsPlural",
				testItem.input, true, actual)
			failed++
		}
	}

	slog("TestIsPlural", passed, failed, len(tests))
}

func TestIsSingular(t *testing.T) {

	tests := append(basicTests(), singularTests()...)
	passed := 0
	failed := 0

	pluralize := NewClient()

	for i, testItem := range tests {
		if actual := pluralize.IsSingular(testItem.input); actual == true {
			plog(t, "PASS test[%d] func %s(%s) expected %t, actual %t", i, "IsSingular",
				testItem.input, true, actual)
			passed++
		} else {
			t.Errorf("FAIL test[%d] func %s(%s) expected %t, actual %t", i, "IsSingular",
				testItem.input, true, actual)
			failed++
		}
	}

	slog("TestIsSingular", passed, failed, len(tests))
}

func TestPlural(t *testing.T) {

	tests := append(basicTests(), pluralTests()...)
	passed := 0
	failed := 0

	pluralize := NewClient()

	for i, testItem := range tests {

		if actual := pluralize.Plural(testItem.input); actual == testItem.expected {
			plog(t, "PASS test[%d] func %s(%s) expected %s, actual %s", i, "Plural",
				testItem.input, testItem.expected, actual)
			passed++
		} else {
			t.Errorf("FAIL test[%d] func %s(%s) expected %s, actual %s", i, "Plural",
				testItem.input, testItem.expected, actual)
			failed++
		}

	}

	slog("TestPlural", passed, failed, len(tests))
}

func TestSingular(t *testing.T) {

	tests := append(basicTests(), singularTests()...)
	passed := 0
	failed := 0

	pluralize := NewClient()

	for i, testItem := range tests {
		if actual := pluralize.Singular(testItem.expected); actual == testItem.input {
			plog(t, "PASS test[%d] func %s(%s) expected %s, actual %s", i, "Singular",
				testItem.input, testItem.expected, actual)
			passed++
		} else {
			t.Errorf("FAIL test[%d] func %s(%s) expected %s, actual %s", i, "Singular",
				testItem.input, testItem.expected, actual)
			failed++
		}
	}

	slog("TestSingular", passed, failed, len(tests))
}

func TestNewPluralRule(t *testing.T) {

	pluralize := NewClient()

	if pluralize.Plural(`regex`) != `regexes` {
		t.Fail()
	}

	pluralize.AddPluralRule(`(?i)gex$`, `gexii`)

	if pluralize.Plural(`regex`) != `regexii` {
		t.Fail()
	}
}

func TestNewSingularRule(t *testing.T) {

	pluralize := NewClient()

	if pluralize.Singular(`singles`) != `single` {
		t.Fail()
	}

	pluralize.AddSingularRule(`(?i)singles$`, `singular`)

	if pluralize.Singular(`singles`) != `singular` {
		t.Fail()
	}
}

func TestNewIrregularRule(t *testing.T) {

	pluralize := NewClient()

	if pluralize.Plural(`irregular`) != `irregulars` {
		t.Fail()
	}

	pluralize.AddIrregularRule(`irregular`, `regular`)

	if pluralize.Plural(`irregular`) != `regular` {
		t.Fail()
	}
}

func TestNewUncountableRule(t *testing.T) {

	pluralize := NewClient()

	if pluralize.Plural(`paper`) != `papers` {
		t.Fail()
	}

	pluralize.AddUncountableRule(`paper`)

	if pluralize.Plural(`paper`) != `paper` {
		t.Fail()
	}
}

func TestPluralize(t *testing.T) {

	const test = "test"
	const tests = "tests"

	pluralize := NewClient()

	if pluralize.Pluralize(test, 0, false) != tests {
		t.Fail()
	}
	if pluralize.Pluralize(test, 0, false) != tests {
		t.Fail()
	}
	if pluralize.Pluralize(test, 1, false) != test {
		t.Fail()
	}
	if pluralize.Pluralize(test, 5, false) != tests {
		t.Fail()
	}
	if pluralize.Pluralize(test, 1, true) != `1 test` {
		t.Fail()
	}
	if pluralize.Pluralize(test, 5, true) != `5 tests` {
		t.Fail()
	}
	if pluralize.Pluralize(`蘋果`, 2, true) != `2 蘋果` {
		t.Fail()
	}
}

// Basic test cases of singular - plural pairs
func basicTests() []TestEntry {
	return []TestEntry{
		// Uncountables.
		{`firmware`, `firmware`},
		{`fish`, `fish`},
		{`media`, `media`},
		{`moose`, `moose`},
		{`police`, `police`},
		{`sheep`, `sheep`},
		{`series`, `series`},
		{`species`, `species`},
		{`agenda`, `agenda`},
		{`news`, `news`},
		{`reindeer`, `reindeer`},
		{`starfish`, `starfish`},
		{`smallpox`, `smallpox`},
		{`tennis`, `tennis`},
		{`chickenpox`, `chickenpox`},
		{`shambles`, `shambles`},
		{`garbage`, `garbage`},
		{`you`, `you`},
		{`wildlife`, `wildlife`},
		{`Staff`, `Staff`},
		{`STAFF`, `STAFF`},
		{`turquois`, `turquois`},
		{`carnivorous`, `carnivorous`},
		{`only`, `only`},
		{`aircraft`, `aircraft`},
		// Latin.
		{`veniam`, `veniam`},
		// Pluralization.
		{`this`, `these`},
		{`that`, `those`},
		{`is`, `are`},
		{`man`, `men`},
		{`superman`, `supermen`},
		{`ox`, `oxen`},
		{`bus`, `buses`},
		{`airbus`, `airbuses`},
		{`railbus`, `railbuses`},
		{`wife`, `wives`},
		{`guest`, `guests`},
		{`thing`, `things`},
		{`mess`, `messes`},
		{`guess`, `guesses`},
		{`person`, `people`},
		{`meteor`, `meteors`},
		{`chateau`, `chateaus`},
		{`lap`, `laps`},
		{`cough`, `coughs`},
		{`death`, `deaths`},
		{`coach`, `coaches`},
		{`boy`, `boys`},
		{`toy`, `toys`},
		{`guy`, `guys`},
		{`girl`, `girls`},
		{`chair`, `chairs`},
		{`toe`, `toes`},
		{`tiptoe`, `tiptoes`},
		{`tomato`, `tomatoes`},
		{`potato`, `potatoes`},
		{`tornado`, `tornadoes`},
		{`torpedo`, `torpedoes`},
		{`hero`, `heroes`},
		{`superhero`, `superheroes`},
		{`volcano`, `volcanoes`},
		{`canto`, `cantos`},
		{`hetero`, `heteros`},
		{`photo`, `photos`},
		{`portico`, `porticos`},
		{`quarto`, `quartos`},
		{`kimono`, `kimonos`},
		{`albino`, `albinos`},
		{`cherry`, `cherries`},
		{`piano`, `pianos`},
		{`pro`, `pros`},
		{`combo`, `combos`},
		{`turbo`, `turbos`},
		{`bar`, `bars`},
		{`crowbar`, `crowbars`},
		{`van`, `vans`},
		{`tobacco`, `tobaccos`},
		{`afficionado`, `afficionados`},
		{`monkey`, `monkeys`},
		{`neutrino`, `neutrinos`},
		{`rhino`, `rhinos`},
		{`steno`, `stenos`},
		{`latino`, `latinos`},
		{`casino`, `casinos`},
		{`avocado`, `avocados`},
		{`commando`, `commandos`},
		{`tuxedo`, `tuxedos`},
		{`speedo`, `speedos`},
		{`dingo`, `dingoes`},
		{`echo`, `echoes`},
		{`nacho`, `nachos`},
		{`motto`, `mottos`},
		{`psycho`, `psychos`},
		{`poncho`, `ponchos`},
		{`pass`, `passes`},
		{`ghetto`, `ghettos`},
		{`mango`, `mangos`},
		{`lady`, `ladies`},
		{`bath`, `baths`},
		{`professional`, `professionals`},
		{`dwarf`, `dwarves`}, // Proper spelling is "dwarfs".
		{`encyclopedia`, `encyclopedias`},
		{`louse`, `lice`},
		{`roof`, `roofs`},
		{`woman`, `women`},
		{`formula`, `formulas`},
		{`polyhedron`, `polyhedra`},
		{`index`, `indices`}, // Maybe "indexes".
		{`matrix`, `matrices`},
		{`vertex`, `vertices`},
		{`axe`, `axes`}, // Could also be plural of "ax".
		{`pickaxe`, `pickaxes`},
		{`crisis`, `crises`},
		{`criterion`, `criteria`},
		{`phenomenon`, `phenomena`},
		{`addendum`, `addenda`},
		{`datum`, `data`},
		{`forum`, `forums`},
		{`millennium`, `millennia`},
		{`alumnus`, `alumni`},
		{`medium`, `mediums`},
		{`census`, `censuses`},
		{`genus`, `genera`},
		{`dogma`, `dogmata`},
		{`life`, `lives`},
		{`hive`, `hives`},
		{`kiss`, `kisses`},
		{`dish`, `dishes`},
		{`human`, `humans`},
		{`knife`, `knives`},
		{`phase`, `phases`},
		{`judge`, `judges`},
		{`class`, `classes`},
		{`witch`, `witches`},
		{`church`, `churches`},
		{`massage`, `massages`},
		{`prospectus`, `prospectuses`},
		{`syllabus`, `syllabi`},
		{`viscus`, `viscera`},
		{`cactus`, `cacti`},
		{`hippopotamus`, `hippopotamuses`},
		{`octopus`, `octopuses`},
		{`platypus`, `platypuses`},
		{`kangaroo`, `kangaroos`},
		{`atlas`, `atlases`},
		{`stigma`, `stigmata`},
		{`schema`, `schemata`},
		{`phenomenon`, `phenomena`},
		{`diagnosis`, `diagnoses`},
		{`mongoose`, `mongooses`},
		{`mouse`, `mice`},
		{`liturgist`, `liturgists`},
		{`box`, `boxes`},
		{`gas`, `gases`},
		{`self`, `selves`},
		{`chief`, `chiefs`},
		{`quiz`, `quizzes`},
		{`child`, `children`},
		{`shelf`, `shelves`},
		{`fizz`, `fizzes`},
		{`tooth`, `teeth`},
		{`thief`, `thieves`},
		{`day`, `days`},
		{`loaf`, `loaves`},
		{`fix`, `fixes`},
		{`spy`, `spies`},
		{`vertebra`, `vertebrae`},
		{`clock`, `clocks`},
		{`lap`, `laps`},
		{`cuff`, `cuffs`},
		{`leaf`, `leaves`},
		{`calf`, `calves`},
		{`moth`, `moths`},
		{`mouth`, `mouths`},
		{`house`, `houses`},
		{`proof`, `proofs`},
		{`hoof`, `hooves`},
		{`elf`, `elves`},
		{`turf`, `turfs`},
		{`craft`, `crafts`},
		{`die`, `dice`},
		{`penny`, `pennies`},
		{`campus`, `campuses`},
		{`virus`, `viri`},
		{`iris`, `irises`},
		{`bureau`, `bureaus`},
		{`kiwi`, `kiwis`},
		{`wiki`, `wikis`},
		{`igloo`, `igloos`},
		{`ninja`, `ninjas`},
		{`pizza`, `pizzas`},
		{`kayak`, `kayaks`},
		{`canoe`, `canoes`},
		{`tiding`, `tidings`},
		{`pea`, `peas`},
		{`drive`, `drives`},
		{`nose`, `noses`},
		{`movie`, `movies`},
		{`status`, `statuses`},
		{`alias`, `aliases`},
		{`memorandum`, `memorandums`},
		{`language`, `languages`},
		{`plural`, `plurals`},
		{`word`, `words`},
		{`multiple`, `multiples`},
		{`reward`, `rewards`},
		{`sandwich`, `sandwiches`},
		{`subway`, `subways`},
		{`direction`, `directions`},
		{`land`, `lands`},
		{`row`, `rows`},
		{`grow`, `grows`},
		{`flow`, `flows`},
		{`rose`, `roses`},
		{`raise`, `raises`},
		{`friend`, `friends`},
		{`follower`, `followers`},
		{`male`, `males`},
		{`nail`, `nails`},
		{`sex`, `sexes`},
		{`tape`, `tapes`},
		{`ruler`, `rulers`},
		{`king`, `kings`},
		{`queen`, `queens`},
		{`zero`, `zeros`},
		{`quest`, `quests`},
		{`goose`, `geese`},
		{`foot`, `feet`},
		{`ex`, `exes`},
		{`reflex`, `reflexes`},
		{`heat`, `heats`},
		{`train`, `trains`},
		{`test`, `tests`},
		{`pie`, `pies`},
		{`fly`, `flies`},
		{`eye`, `eyes`},
		{`lie`, `lies`},
		{`node`, `nodes`},
		{`trade`, `trades`},
		{`chinese`, `chinese`},
		{`please`, `pleases`},
		{`japanese`, `japanese`},
		{`regex`, `regexes`},
		{`license`, `licenses`},
		{`zebra`, `zebras`},
		{`general`, `generals`},
		{`corps`, `corps`},
		{`pliers`, `pliers`},
		{`flyer`, `flyers`},
		{`scissors`, `scissors`},
		{`fireman`, `firemen`},
		{`chirp`, `chirps`},
		{`harp`, `harps`},
		{`corpse`, `corpses`},
		{`dye`, `dyes`},
		{`move`, `moves`},
		{`zombie`, `zombies`},
		{`variety`, `varieties`},
		{`talkie`, `talkies`},
		{`walkie-talkie`, `walkie-talkies`},
		{`groupie`, `groupies`},
		{`goonie`, `goonies`},
		{`lassie`, `lassies`},
		{`genie`, `genies`},
		{`foodie`, `foodies`},
		{`faerie`, `faeries`},
		{`collie`, `collies`},
		{`obloquy`, `obloquies`},
		{`looey`, `looies`},
		{`osprey`, `ospreys`},
		{`cover`, `covers`},
		{`tie`, `ties`},
		{`groove`, `grooves`},
		{`bee`, `bees`},
		{`ave`, `aves`},
		{`wave`, `waves`},
		{`wolf`, `wolves`},
		{`airwave`, `airwaves`},
		{`archive`, `archives`},
		{`arch`, `arches`},
		{`dive`, `dives`},
		{`aftershave`, `aftershaves`},
		{`cave`, `caves`},
		{`grave`, `graves`},
		{`gift`, `gifts`},
		{`nerve`, `nerves`},
		{`nerd`, `nerds`},
		{`carve`, `carves`},
		{`rave`, `raves`},
		{`scarf`, `scarves`},
		{`sale`, `sales`},
		{`sail`, `sails`},
		{`swerve`, `swerves`},
		{`love`, `loves`},
		{`dove`, `doves`},
		{`glove`, `gloves`},
		{`wharf`, `wharves`},
		{`valve`, `valves`},
		{`werewolf`, `werewolves`},
		{`view`, `views`},
		{`emu`, `emus`},
		{`menu`, `menus`},
		{`wax`, `waxes`},
		{`fax`, `faxes`},
		{`nut`, `nuts`},
		{`crust`, `crusts`},
		{`lemma`, `lemmata`},
		{`anathema`, `anathemata`},
		{`analysis`, `analyses`},
		{`locus`, `loci`},
		{`uterus`, `uteri`},
		{`curriculum`, `curricula`},
		{`quorum`, `quora`},
		{`genie`, `genies`},
		{`genius`, `geniuses`},
		{`flower`, `flowers`},
		{`crash`, `crashes`},
		{`soul`, `souls`},
		{`career`, `careers`},
		{`planet`, `planets`},
		{`son`, `sons`},
		{`sun`, `suns`},
		{`drink`, `drinks`},
		{`diploma`, `diplomas`},
		{`dilemma`, `dilemmas`},
		{`grandma`, `grandmas`},
		{`no`, `nos`},
		{`yes`, `yeses`},
		{`employ`, `employs`},
		{`employee`, `employees`},
		{`history`, `histories`},
		{`story`, `stories`},
		{`purchase`, `purchases`},
		{`order`, `orders`},
		{`key`, `keys`},
		{`bomb`, `bombs`},
		{`city`, `cities`},
		{`sanity`, `sanities`},
		{`ability`, `abilities`},
		{`activity`, `activities`},
		{`cutie`, `cuties`},
		{`validation`, `validations`},
		{`floaty`, `floaties`},
		{`nicety`, `niceties`},
		{`goalie`, `goalies`},
		{`crawly`, `crawlies`},
		{`duty`, `duties`},
		{`scrutiny`, `scrutinies`},
		{`deputy`, `deputies`},
		{`beauty`, `beauties`},
		{`bank`, `banks`},
		{`family`, `families`},
		{`tally`, `tallies`},
		{`ally`, `allies`},
		{`alley`, `alleys`},
		{`valley`, `valleys`},
		{`medley`, `medleys`},
		{`melody`, `melodies`},
		{`trolly`, `trollies`},
		{`thunk`, `thunks`},
		{`koala`, `koalas`},
		{`special`, `specials`},
		{`book`, `books`},
		{`knob`, `knobs`},
		{`crab`, `crabs`},
		{`plough`, `ploughs`},
		{`high`, `highs`},
		{`low`, `lows`},
		{`hiccup`, `hiccups`},
		{`bonus`, `bonuses`},
		{`circus`, `circuses`},
		{`abacus`, `abacuses`},
		{`phobia`, `phobias`},
		{`case`, `cases`},
		{`lace`, `laces`},
		{`trace`, `traces`},
		{`mage`, `mages`},
		{`lotus`, `lotuses`},
		{`motorbus`, `motorbuses`},
		{`cutlas`, `cutlases`},
		{`tequila`, `tequilas`},
		{`liar`, `liars`},
		{`delta`, `deltas`},
		{`visa`, `visas`},
		{`flea`, `fleas`},
		{`favela`, `favelas`},
		{`cobra`, `cobras`},
		{`finish`, `finishes`},
		{`gorilla`, `gorillas`},
		{`mass`, `masses`},
		{`face`, `faces`},
		{`rabbit`, `rabbits`},
		{`adventure`, `adventures`},
		{`breeze`, `breezes`},
		{`brew`, `brews`},
		{`canopy`, `canopies`},
		{`copy`, `copies`},
		{`spy`, `spies`},
		{`cave`, `caves`},
		{`charge`, `charges`},
		{`cinema`, `cinemas`},
		{`coffee`, `coffees`},
		{`favourite`, `favourites`},
		{`themself`, `themselves`},
		{`country`, `countries`},
		{`issue`, `issues`},
		{`authority`, `authorities`},
		{`force`, `forces`},
		{`objective`, `objectives`},
		{`present`, `presents`},
		{`industry`, `industries`},
		{`believe`, `believes`},
		{`century`, `centuries`},
		{`category`, `categories`},
		{`eve`, `eves`},
		{`fee`, `fees`},
		{`gene`, `genes`},
		{`try`, `tries`},
		{`currency`, `currencies`},
		{`pose`, `poses`},
		{`cheese`, `cheeses`},
		{`clue`, `clues`},
		{`cheer`, `cheers`},
		{`litre`, `litres`},
		{`money`, `monies`},
		{`attorney`, `attorneys`},
		{`balcony`, `balconies`},
		{`cockney`, `cockneys`},
		{`donkey`, `donkeys`},
		{`honey`, `honeys`},
		{`smiley`, `smilies`},
		{`survey`, `surveys`},
		{`whiskey`, `whiskies`},
		{`volley`, `volleys`},
		{`tongue`, `tongues`},
		{`suit`, `suits`},
		{`suite`, `suites`},
		{`cruise`, `cruises`},
		{`eave`, `eaves`},
		{`consultancy`, `consultancies`},
		{`pouch`, `pouches`},
		{`wallaby`, `wallabies`},
		{`abyss`, `abysses`},
		{`weekly`, `weeklies`},
		{`whistle`, `whistles`},
		{`utilise`, `utilises`},
		{`utilize`, `utilizes`},
		{`mercy`, `mercies`},
		{`mercenary`, `mercenaries`},
		{`take`, `takes`},
		{`flush`, `flushes`},
		{`gate`, `gates`},
		{`evolve`, `evolves`},
		{`slave`, `slaves`},
		{`native`, `natives`},
		{`revolve`, `revolves`},
		{`twelve`, `twelves`},
		{`sleeve`, `sleeves`},
		{`subjective`, `subjectives`},
		{`stream`, `streams`},
		{`beam`, `beams`},
		{`foam`, `foams`},
		{`callus`, `calluses`},
		{`use`, `uses`},
		{`beau`, `beaus`},
		{`gateau`, `gateaus`},
		{`fetus`, `fetuses`},
		{`luau`, `luaus`},
		{`pilau`, `pilaus`},
		{`shoe`, `shoes`},
		{`sandshoe`, `sandshoes`},
		{`zeus`, `zeuses`},
		{`nucleus`, `nuclei`},
		{`sky`, `skies`},
		{`beach`, `beaches`},
		{`brush`, `brushes`},
		{`hoax`, `hoaxes`},
		{`scratch`, `scratches`},
		{`nanny`, `nannies`},
		{`negro`, `negroes`},
		{`taco`, `tacos`},
		{`cafe`, `cafes`},
		{`cave`, `caves`},
		{`giraffe`, `giraffes`},
		{`goodwife`, `goodwives`},
		{`housewife`, `housewives`},
		{`safe`, `safes`},
		{`save`, `saves`},
		{`pocketknife`, `pocketknives`},
		{`tartufe`, `tartufes`},
		{`tartuffe`, `tartuffes`},
		{`truffle`, `truffles`},
		{`jefe`, `jefes`},
		{`agrafe`, `agrafes`},
		{`agraffe`, `agraffes`},
		{`bouffe`, `bouffes`},
		{`carafe`, `carafes`},
		{`chafe`, `chafes`},
		{`pouffe`, `pouffes`},
		{`pouf`, `poufs`},
		{`piaffe`, `piaffes`},
		{`gaffe`, `gaffes`},
		{`executive`, `executives`},
		{`cove`, `coves`},
		{`dove`, `doves`},
		{`fave`, `faves`},
		{`positive`, `positives`},
		{`solve`, `solves`},
		{`trove`, `troves`},
		{`treasure`, `treasures`},
		{`suave`, `suaves`},
		{`bluff`, `bluffs`},
		{`half`, `halves`},
		{`knockoff`, `knockoffs`},
		{`handkerchief`, `handkerchiefs`},
		{`reed`, `reeds`},
		{`reef`, `reefs`},
		{`yourself`, `yourselves`},
		{`sunroof`, `sunroofs`},
		{`plateau`, `plateaus`},
		{`radius`, `radii`},
		{`stratum`, `strata`},
		{`stratus`, `strati`},
		{`focus`, `foci`},
		{`fungus`, `fungi`},
		{`appendix`, `appendices`},
		{`seraph`, `seraphim`},
		{`cherub`, `cherubim`},
		{`memo`, `memos`},
		{`cello`, `cellos`},
		{`automaton`, `automata`},
		{`button`, `buttons`},
		{`crayon`, `crayons`},
		{`captive`, `captives`},
		{`abrasive`, `abrasives`},
		{`archive`, `archives`},
		{`additive`, `additives`},
		{`hive`, `hives`},
		{`beehive`, `beehives`},
		{`olive`, `olives`},
		{`black olive`, `black olives`},
		{`chive`, `chives`},
		{`adjective`, `adjectives`},
		{`cattle drive`, `cattle drives`},
		{`explosive`, `explosives`},
		{`executive`, `executives`},
		{`negative`, `negatives`},
		{`fugitive`, `fugitives`},
		{`progressive`, `progressives`},
		{`laxative`, `laxatives`},
		{`incentive`, `incentives`},
		{`genesis`, `geneses`},
		{`surprise`, `surprises`},
		{`enterprise`, `enterprises`},
		{`relative`, `relatives`},
		{`positive`, `positives`},
		{`perspective`, `perspectives`},
		{`superlative`, `superlatives`},
		{`afterlife`, `afterlives`},
		{`native`, `natives`},
		{`detective`, `detectives`},
		{`collective`, `collectives`},
		{`lowlife`, `lowlives`},
		{`low-life`, `low-lives`},
		{`strife`, `strifes`},
		{`pony`, `ponies`},
		{`phony`, `phonies`},
		{`felony`, `felonies`},
		{`colony`, `colonies`},
		{`symphony`, `symphonies`},
		{`semicolony`, `semicolonies`},
		{`radiotelephony`, `radiotelephonies`},
		{`company`, `companies`},
		{`ceremony`, `ceremonies`},
		{`carnivore`, `carnivores`},
		{`emphasis`, `emphases`},
		{`abuse`, `abuses`},
		{`ass`, `asses`},
		{`mile`, `miles`},
		{`consensus`, `consensuses`},
		{`coatdress`, `coatdresses`},
		{`courthouse`, `courthouses`},
		{`playhouse`, `playhouses`},
		{`crispness`, `crispnesses`},
		{`racehorse`, `racehorses`},
		{`greatness`, `greatnesses`},
		{`christmas`, `christmases`},
		{`zymase`, `zymases`},
		{`accomplice`, `accomplices`},
		{`amice`, `amices`},
		{`titmouse`, `titmice`},
		{`slice`, `slices`},
		// Prototype inheritance.
		{`constructor`, `constructors`},
		// Non-standard case.
		{`randomWord`, `randomWords`},
		{`camelCase`, `camelCases`},
		{`PascalCase`, `PascalCases`},
		{`Alumnus`, `Alumni`},
		{`CHICKEN`, `CHICKENS`},
		{`日本語`, `日本語`},
		{`한국`, `한국`},
		{`中文`, `中文`},
		{`اللغة العربية`, `اللغة العربية`},
		{`四 chicken`, `四 chickens`},
	}
}

// Odd plural to singular tests.
func singularTests() []TestEntry {
	return []TestEntry{
		{`dingo`, `dingos`},
		{`mango`, `mangoes`},
		{`echo`, `echos`},
		{`ghetto`, `ghettoes`},
		{`nucleus`, `nucleuses`},
		{`bureau`, `bureaux`},
		{`seraph`, `seraphs`},
	}
}

// Odd singular to plural tests.
func pluralTests() []TestEntry {
	return []TestEntry{
		{`whisky`, `whiskies`},
		{`plateaux`, `plateaux`},
		{`axis`, `axes`},
		{`automatum`, `automata`},
		{`thou`, `you`},
		{`passerby`, `passersby`},
	}
}
